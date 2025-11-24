# Security Review Report - Polimane

**Date:** 2025-11-24
**Reviewer:** Claude (Automated Security Review)
**Application:** Polimane - Schema/Template Management SaaS Platform

## Executive Summary

This security review assessed the Polimane full-stack application consisting of a Vue 3 frontend and Go backend deployed on AWS infrastructure. The application demonstrates good security practices with proper secrets management via Bitwarden, secure authentication through WorkOS, and appropriate use of GORM to prevent SQL injection.

**Overall Security Posture:** Good with some areas requiring attention

**Critical Issues:** 2
**High Priority Issues:** 3
**Medium Priority Issues:** 4
**Low Priority Issues:** 2

---

## 1. Critical Issues

### 1.1 Overly Permissive Content Security Policy (CSP)

**Location:** `backend/api/server.go:44`

**Issue:**
```go
ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval'; ..."
```

The CSP allows `'unsafe-inline'` and `'unsafe-eval'` for scripts, which significantly weakens XSS protection. These directives allow inline JavaScript and eval(), making the application vulnerable to XSS attacks if any user input is improperly sanitized.

**Impact:** High - Reduces effectiveness of CSP as a defense-in-depth mechanism against XSS attacks.

**Recommendation:**
1. Remove `'unsafe-inline'` and `'unsafe-eval'` from script-src
2. Use nonces or hashes for legitimate inline scripts
3. Replace any eval() usage with safer alternatives
4. For Vue 3, configure Vite to generate CSP-compatible builds

### 1.2 Information Disclosure in Development Mode

**Location:** `backend/api/auth/middleware.go:174-176`

**Issue:**
```go
if env.IsDev {
    extra = append(extra, base.CustomErrorData{"internalError": err.Error()})
    return unauthorizedErr.AddCustomData(extra...)
}
```

Internal error details are exposed in development mode. If this code is accidentally deployed to production or if the `env.IsDev` flag is misconfigured, sensitive error information could be leaked to attackers.

**Impact:** High - Stack traces, database errors, or internal paths could be exposed to attackers.

**Recommendation:**
1. Add runtime checks to ensure development error details are NEVER exposed in production
2. Implement environment validation on startup
3. Use structured logging for internal errors instead of returning them to clients
4. Consider adding a separate staging environment with production-like error handling

---

## 2. High Priority Issues

### 2.1 No CSRF Protection Implementation

**Location:** Entire application (backend/api/)

**Issue:**
While the CORS configuration includes `"X-CSRF-Token"` header in the allowed headers list (`lambda_gateway.tf:17`), there's no actual CSRF token generation, validation, or middleware implementation in the backend code.

**Impact:** The application relies solely on SameSite cookie attributes for CSRF protection. While SameSite="None" with Secure flag is used, this is insufficient for a comprehensive CSRF defense strategy.

**Recommendation:**
1. Implement double-submit cookie pattern or synchronizer token pattern
2. Add CSRF middleware to validate tokens on state-changing requests (POST, PUT, PATCH, DELETE)
3. Generate and send CSRF tokens from the backend
4. Validate tokens on the frontend before making API calls
5. Consider using the `gorilla/csrf` package or implementing custom CSRF middleware

### 2.2 Unsafe v-html Usage

**Location:** `frontend/src/modules/schemas/editor/components/modals/schemaExport/SchemaExportPreview.vue:5`

**Issue:**
```vue
<div v-html="source" />
```

The component uses `v-html` to render SVG content from the API. While the SVG comes from a trusted backend endpoint, it processes user-controlled schema data (colors, beads, etc.). If schema data isn't properly sanitized on the backend, this could lead to XSS vulnerabilities.

**Impact:** High - Potential XSS if malicious SVG content is rendered.

**Recommendation:**
1. Validate and sanitize all schema data (especially palette colors and beads data) on the backend before rendering SVG
2. Consider using DOMPurify library to sanitize SVG content before rendering
3. Implement strict SVG schema validation
4. Use Vue's template syntax instead of v-html where possible
5. Review `backend/views/templates/schema_preview.go` to ensure proper output encoding

### 2.3 No Rate Limiting on Critical Endpoints

**Location:** Backend API endpoints (notably auth endpoints)

**Issue:**
While AWS API Gateway has basic throttling configured (burst: 10, rate: 2 req/sec in `lambda_gateway.tf:47-48`), there's no application-level rate limiting for critical endpoints like:
- `/api/auth/login`
- `/api/auth/login/complete`
- `/api/users/password-reset`
- `/api/users/email-verify`

**Impact:** Vulnerable to brute force attacks, credential stuffing, and enumeration attacks.

**Recommendation:**
1. Implement application-level rate limiting using middleware
2. Use different limits for different endpoint categories:
   - Auth endpoints: 5 attempts per 15 minutes per IP
   - Password reset: 3 attempts per hour per email
   - Email verification: 5 attempts per hour per user
3. Consider using Redis for distributed rate limiting
4. Implement progressive delays or temporary account lockouts after repeated failures
5. Add monitoring and alerting for rate limit violations

---

## 3. Medium Priority Issues

### 3.1 Long Cookie Expiration Time

**Location:** `backend/api/auth/cookie.go:17`

**Issue:**
```go
cookieMaxAge = int((time.Hour * 24 * 60).Seconds()) // 60 days
```

Cookies expire after 60 days, which is quite long. If a cookie is stolen, an attacker would have extended access.

**Impact:** Medium - Increases the window of opportunity for session hijacking attacks.

**Recommendation:**
1. Reduce cookie max age to 7-14 days for better security
2. Implement a "remember me" feature with explicit user consent for longer sessions
3. Add session activity monitoring to detect anomalous behavior
4. Consider implementing device fingerprinting for session validation

### 3.2 User Cache TTL May Be Too Long

**Location:** `backend/api/auth/middleware.go:47`

**Issue:**
```go
localcache.WithDefaultExpiration(10 * time.Minute),
```

User and WorkOS user data are cached for 10 minutes. If a user's permissions are revoked or their account is disabled, changes won't take effect for up to 10 minutes.

**Impact:** Medium - Delayed enforcement of security-critical changes (account suspension, permission revocation).

**Recommendation:**
1. Reduce cache TTL to 1-2 minutes for security-sensitive data
2. Implement cache invalidation on security-critical events:
   - Account suspension
   - Permission changes
   - Password changes
3. Add force-logout capability that invalidates all user sessions
4. Consider using Redis with pub/sub for immediate cache invalidation across instances

### 3.3 Missing Security Headers Validation

**Location:** `backend/api/server.go:41-57`

**Issue:**
While Helmet middleware is configured with good defaults, some important headers are missing:
- `Permissions-Policy` (formerly Feature-Policy)
- `X-Content-Type-Options` is set but should be validated

**Recommendation:**
1. Add Permissions-Policy header:
   ```go
   PermissionsPolicy: "geolocation=(), microphone=(), camera=()"
   ```
2. Ensure X-Content-Type-Options: "nosniff" is properly set (currently configured)
3. Consider adding `X-Permitted-Cross-Domain-Policies: none`
4. Review and tighten CSP policy as mentioned in Critical Issues

### 3.4 No Logging of Security Events

**Location:** Throughout the application

**Issue:**
There's no evidence of security event logging for:
- Failed authentication attempts
- Authorization failures
- Suspicious activity patterns
- Account changes (email, password, MFA)

**Impact:** Medium - Difficulty detecting and responding to security incidents.

**Recommendation:**
1. Implement structured security event logging
2. Log the following events:
   - All authentication attempts (success and failure)
   - Authorization failures
   - Account modifications
   - Password changes/resets
   - MFA setup/removal
   - Unusual access patterns
3. Send security logs to centralized logging (CloudWatch, Sentry already integrated)
4. Set up alerts for suspicious patterns
5. Implement audit trail for schema access and modifications

---

## 4. Low Priority Issues

### 4.1 Missing Request ID Tracking

**Location:** Backend API

**Issue:**
No request ID middleware is implemented for request tracing and debugging.

**Recommendation:**
1. Add request ID middleware to generate unique IDs for each request
2. Include request IDs in all logs
3. Return request IDs in error responses for customer support
4. Use OpenTelemetry (already imported) for distributed tracing

### 4.2 Cookie Domain Configuration

**Location:** `backend/api/auth/cookie.go:34`

**Issue:**
```go
domain := "." + environment.AppDomain
```

The cookie domain is prefixed with a dot, making it accessible to all subdomains. This is intentional for the app/api subdomain setup but should be documented.

**Impact:** Low - Minimal risk if all subdomains are controlled by the application.

**Recommendation:**
1. Document the subdomain cookie sharing design decision
2. Ensure no untrusted subdomains are added in the future
3. Consider using path restrictions if possible
4. Review subdomain isolation strategy

---

## 5. Positive Security Findings

The following security controls are properly implemented:

### 5.1 Authentication & Authorization ✅
- WorkOS integration for secure authentication
- JWT token validation with JWKS
- Proper token refresh mechanism
- Access level enforcement (Read, Write, Admin)
- User-schema relationship validation before operations

### 5.2 Secrets Management ✅
- Bitwarden integration for all sensitive secrets
- No hardcoded credentials found in source code
- Secrets referenced by ID in Terraform (SID pattern)
- Environment-based secret loading

### 5.3 Database Security ✅
- GORM used throughout (no raw SQL found)
- Parameterized queries prevent SQL injection
- CockroachDB with TLS encryption
- Database migrations managed via Goose

### 5.4 Input Validation ✅
- go-playground/validator used for request validation
- Validation tags on all input structures
- Custom validators (e.g., iscolor for hex colors)
- Query parameter validation

### 5.5 Transport Security ✅
- TLS 1.2 enforced (lambda_gateway.tf:59)
- HTTPS-only cookies (Secure flag)
- HTTP-only cookies prevent XSS cookie theft
- HSTS headers configured with preload

### 5.6 CORS Configuration ✅
- Specific origin allowlist (no wildcards)
- Credentials enabled with specific origin
- Appropriate methods and headers whitelisted
- Max-age set appropriately

### 5.7 Infrastructure Security ✅
- AWS IAM roles for Lambda (no hardcoded AWS credentials in production)
- CloudWatch logging enabled
- API Gateway throttling configured
- CloudFront CDN for DDoS protection
- Cloudflare security rules

### 5.8 Cookie Security ✅
- HTTPOnly flag prevents JavaScript access
- Secure flag ensures HTTPS-only transmission
- SameSite attribute (though "None" for cross-origin)
- Cookie encryption via encryptcookie middleware

### 5.9 Error Handling ✅
- Custom error handler with appropriate status codes
- GORM record not found errors mapped to 404
- Generic error messages (except dev mode issue noted)

---

## 6. Dependency Analysis

### Backend Dependencies (Go)
**Major Dependencies:**
- `gofiber/fiber/v2` v2.52.9 - Web framework
- `gorm.io/gorm` v1.31.1 - ORM (latest)
- `workos/workos-go/v4` v4.46.1 - Authentication
- `aws/aws-sdk-go-v2` v1.39.6 - AWS integration
- `getsentry/sentry-go` v0.36.2 - Error tracking

**Recommendations:**
1. All dependencies appear to be recent versions
2. Consider running `go audit` or `govulncheck` regularly
3. Enable Dependabot or Renovate for automated dependency updates
4. Review `jackc/pgx/v5` and `gorm` security advisories periodically

### Frontend Dependencies (NPM)
**Major Dependencies:**
- `vue` v3.5.24 - Latest stable
- `@sentry/vue` v10.25.0 - Error tracking
- `vite` v7.2.2 - Build tool
- `pinia` v3.0.4 - State management

**Recommendations:**
1. Run `npm audit` to check for known vulnerabilities
2. Update dependencies with `npm-check-updates` (already in scripts)
3. Consider using `npm audit fix` for automated patching
4. Review `jspdf` and `svg2pdf.js` for any security advisories

---

## 7. Recommendations by Priority

### Immediate Actions (Critical)
1. ✅ Fix CSP to remove unsafe-inline and unsafe-eval
2. ✅ Add production safeguards for error message exposure
3. ✅ Implement CSRF protection

### Short-term (High - within 1 sprint)
1. ✅ Sanitize SVG rendering or add DOMPurify
2. ✅ Implement application-level rate limiting
3. ✅ Add security event logging

### Medium-term (Medium - within 1-2 months)
1. ✅ Reduce cookie expiration time
2. ✅ Reduce cache TTL or implement cache invalidation
3. ✅ Add comprehensive security logging and monitoring
4. ✅ Add missing security headers

### Long-term (Low - ongoing)
1. ✅ Implement request ID tracking
2. ✅ Set up automated dependency scanning
3. ✅ Regular security audits and penetration testing
4. ✅ Security awareness training for team

---

## 8. Testing Recommendations

To validate these findings and ensure ongoing security:

1. **Automated Security Testing:**
   - Set up OWASP ZAP or Burp Suite for regular scanning
   - Integrate security tests into CI/CD pipeline
   - Use gosec for Go static analysis
   - Use npm audit in CI pipeline

2. **Manual Testing:**
   - Test authentication flows for bypass vulnerabilities
   - Verify authorization checks on all endpoints
   - Test file upload/export functionality for injection
   - Validate rate limiting effectiveness

3. **Penetration Testing:**
   - Conduct annual penetration tests
   - Test for OWASP Top 10 vulnerabilities
   - Test API security specifically
   - Test infrastructure security

---

## 9. Compliance Considerations

If applicable to your use case:

1. **GDPR Compliance:**
   - User data deletion capabilities
   - Data export functionality
   - Consent management
   - Data retention policies

2. **SOC 2 Considerations:**
   - Access logging and monitoring
   - Encryption at rest and in transit
   - Change management procedures
   - Incident response plan

---

## 10. Conclusion

The Polimane application demonstrates a strong security foundation with proper use of industry-standard security practices. The use of Bitwarden for secrets management, WorkOS for authentication, and GORM for database access shows good security awareness.

However, the identified issues, particularly the permissive CSP and lack of CSRF protection, should be addressed promptly to maintain a robust security posture. Implementing the recommendations in this report will significantly enhance the application's security.

**Next Steps:**
1. Review and prioritize the identified issues
2. Create tickets for each issue in your issue tracker
3. Assign ownership and timelines
4. Implement fixes following the recommendations
5. Re-test after fixes are deployed
6. Schedule regular security reviews (quarterly recommended)

---

**Report Prepared By:** Claude Code
**Review Methodology:** Static code analysis, dependency review, configuration audit
**Scope:** Full-stack application including infrastructure configuration
