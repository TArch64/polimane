# AI Test Generation Instructions

## 🚨 CRITICAL RULES - READ FIRST

### Test Execution Command

```bash
docker compose run --rm backend make test test_pattern="./path/..."
```

**MANDATORY:** Always use this exact command format. Replace `./path/...` with actual package path.

Examples:
- `./base/...` for base package
- `./signal/...` for signal package
- `./env/...` for env package
- `./model/...` for model package

**NEVER use bare `go test` commands - always use the docker compose wrapper.**

### Build Tag Rules

- ❌ **SKIP** files with `//go:build dev` (development-only)
- ✅ **TEST** files with `//go:build !dev` (production implementations)
- ✅ **TEST** files without build tags (run in all environments)
- When you see dev/prod file pairs, only test the production version

### Testing Scope

- ✅ Test code YOU wrote in the project
- ❌ Don't test third-party library internals (signals.AddListener, GORM methods, etc.)
- ✅ Test integration points where your code interfaces with libraries
- ✅ Test factory functions (Provider()) for initialization behavior
- ✅ Test struct types for zero-value behavior and field initialization

---

## FOUNDATION: Required Setup

### Frameworks and Imports

```go
import (
"testing"
"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/mock"
"github.com/DATA-DOG/go-sqlmock" // for database tests
)
```

### File Structure Rules

- **Test files:** Same package, `{source_file}_test.go` suffix
- **Mock files:** `mocks_test.go` in same package
- **One test file per source file**

### Required Test Cases

For every function/method, include these test cases when applicable:

1. **✅ Success case** - Happy path
2. **❌ Error cases** - What can go wrong?
3. **🔍 Edge cases** - Empty inputs, nil values, zero values
4. **🗄️ Database errors** (if applicable) - Connection failures, not found, etc.

---

## CORE PATTERNS: Basic Test Structure

### Simple Functions (like TagError, Provider functions)

```go
func TestFunctionName(t *testing.T) {
t.Run("describes what this test does", func (t *testing.T) {
// Arrange
input := "test input"

// Act
result := FunctionName(input)

// Assert
assert.Equal(t, expectedValue, result)
})
}
```

### Mock Implementation Template

```go
type MockClient struct {
mock.Mock
}

func (m *MockClient) Method(ctx context.Context, param Type) (*Model, error) {
args := m.Called(ctx, param)
if args.Get(0) == nil {
return nil, args.Error(1)
}
return args.Get(0).(*Model), args.Error(1)
}
```

---

## DEPENDENCY INJECTION: Options Struct Pattern

### 🚨 Critical Pattern: Provider Functions

This codebase uses `fx.In` dependency injection through options structs. **Always use the options struct format:**

```go
// ✅ Correct - Use options struct
client := Provider(ClientOptions{
DB:          gormDB,
UserSchemas: userSchemas,
Signals:     signals,
})

// ❌ Wrong - Individual parameters (old pattern)
client := Provider(gormDB, userSchemas, signals)
```

### Common Provider Patterns

**Repository Clients:**

```go
client := Provider(ClientOptions{
DB: gormDB,
})
```

**Service Clients:**

```go
client := Provider(ClientOptions{
Env: environment,
JWK: jwkClient,
})
```

**API Controllers:**

```go
controller := Provider(ControllerOptions{
WorkosClient: workosClient,
Env:          environment,
Users:        usersClient,
Signals:      signalsContainer,
})
```

**API Server:**

```go
app, err := Provider(ServerOptions{
Controllers:    controllers,
Options:        options,
Sentry:         sentryContainer,
Env:            environment,
AuthMiddleware: authMiddleware,
})
```

---

## MOCKING: Advanced Mock Patterns

### Context Handling in Mocks

**⚠️ Common Problem:** S3 and other external service mocks fail due to context type mismatches:

```go
// ❌ This fails - runtime context differs from expected
mockS3.On("PutObject", mock.AnythingOfType("*context.timerCtx"), ...)

// ✅ This works - flexible context matching
mockS3.On("PutObject", mock.Anything, ...)
```

**Rule:** When mocking functions that accept `context.Context`, use `mock.Anything` for context parameters unless
specifically testing context behavior.

### Complex Parameter Matching with mock.MatchedBy

**Pattern for Repository Calls:**

```go
// ✅ Use matcher functions instead of exact struct matching
mockSchemas.On("Update", mock.MatchedBy(func(options *repositoryschemas.UpdateOptions) bool {
return options.User == testUser &&
options.SchemaID == testSchemaID &&
options.Updates != nil
})).Return(nil)

// ❌ Avoid exact struct matching (context type issues)
mockSchemas.On("Update", &repositoryschemas.UpdateOptions{
Ctx:      testUser, // Type error!
User:     testUser,
SchemaID: testSchemaID,
})
```

### Interface vs Implementation

**🚨 Critical Rule:** When creating test fixtures that need access to unexported fields, always instantiate the concrete
implementation:

```go
// ✅ Correct - Use concrete implementation
client := &Impl{
api: mockAPI,
fs:  mockFS,
env: mockEnv,
}

// ❌ Wrong - Cannot instantiate interface
client := &Client{
api: mockAPI, // Compilation error
}
```

**Pattern:** Look for the concrete type (usually `Impl`) that implements the interface (`Client`).

---

## DATABASE TESTING: GORM & SQL Patterns

### Database Test Setup Pattern

```go
func setupTest(t *testing.T) (*Impl, sqlmock.Sqlmock, func ()) {
db, mock, err := sqlmock.New()
require.NoError(t, err)

gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
require.NoError(t, err)

client := &Impl{db: gormDB}
cleanup := func () { db.Close() }

return client, mock, cleanup
}

func TestMethodName(t *testing.T) {
client, mock, cleanup := setupTest(t)
defer cleanup()

t.Run("success case", func (t *testing.T) {
// Setup expectations
mock.ExpectQuery("SELECT").WillReturnRows(...)

// Call method
result, err := client.Method()

// Verify
assert.NoError(t, err)
assert.NoError(t, mock.ExpectationsWereMet())
})
}
```

### SQL Operation Mapping

- `INSERT/UPDATE/DELETE` → Use `mock.ExpectExec()`
- `SELECT` → Use `mock.ExpectQuery()`

### GORM Behaviors

- `Take()`, `Find()` return empty struct + error (NOT nil) when not found
- `Create()`, `Delete()` automatically wrapped in transactions
- GORM adds `created_at`/`updated_at` to INSERT - use `sqlmock.AnyArg()`

### Common SQL Patterns

```go
// CREATE with timestamps
mock.ExpectBegin()
mock.ExpectExec(`INSERT INTO "table_name"`).
WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), userArg, schemaArg).
WillReturnResult(sqlmock.NewResult(1, 1))
mock.ExpectCommit()

// SELECT with WHERE
mock.ExpectQuery(`SELECT \* FROM "table_name" WHERE col = \$1`).
WithArgs(arg).
WillReturnRows(sqlmock.NewRows([]string{"col"}).AddRow(value))
```

---

## HTTP & FIBER TESTING: Web Application Patterns

### HTTP Testing Rules

- ❌ **NEVER send real HTTP requests in tests**
- ✅ Mock all external HTTP calls using `testify/mock` or `httptest`
- ✅ Use `httptest.NewServer()` for integration-style testing
- ✅ Always verify mock expectations with `AssertExpectations(t)`

### 🚨 CRITICAL: Fiber Custom Error Handler Configuration

When testing Fiber applications that use custom error types (like `base.CustomError`), **always configure the test app
with the same error handler as production:**

```go
import "polimane/backend/api/base"

// ✅ Correct - Configure custom error handler
app := fiber.New(fiber.Config{
ErrorHandler: base.ErrorHandler,
})

// ❌ Wrong - Default error handler doesn't handle custom errors
app := fiber.New()
```

**Why this matters:**
- Custom errors return 500 instead of intended status codes (e.g., 401) without proper handling
- Fiber's default error handler treats custom errors as generic errors
- Test behavior must match production behavior
- Especially critical when `env.IsDev` is false in test environments

### Fiber Test Pattern

```go
func TestMiddlewareHandler(t *testing.T) {
t.Run("returns unauthorized when access token missing", func (t *testing.T) {
// Arrange
middleware := &Middleware{...}

// Create fiber app with custom error handler
app := fiber.New(fiber.Config{
ErrorHandler: base.ErrorHandler,
})
app.Use(middleware.Handler)
app.Get("/test", func (c *fiber.Ctx) error {
return c.SendString("success")
})

req := httptest.NewRequest("GET", "/test", nil)

// Act
resp, err := app.Test(req)

// Assert
assert.NoError(t, err)
assert.Equal(t, 401, resp.StatusCode) // Now works correctly!
})
}
```

---

## SPECIAL CASES: Global State & Cache Testing

### Global State and Test Isolation

**Rule:** Tests that use validation must initialize the validator in each test:

```go
func TestController_apiCreate(t *testing.T) {
t.Run("validates input", func (t *testing.T) {
// Initialize validator for this test
base.InitValidator()

// Test code...
})
}
```

**Why:** Tests should not depend on execution order or shared global state.

### Cache Testing Guidelines

**Pattern:** Use real cache implementations for core functionality, mocks for external dependencies:

```go
// ✅ Correct - Real cache for core functionality
cacheOptions := []localcache.Option{
localcache.WithDefaultExpiration(10 * time.Minute),
localcache.WithCleanupInterval(5 * time.Minute),
}

middleware := &Middleware{
userCache:       localcache.New[*model.User](cacheOptions...),
workosUserCache: localcache.New[*usermanagement.User](cacheOptions...),
workosClient:    &workos.Client{}, // Mock external service
env:             &env.Environment{}, // Mock configuration
users:           &MockUsersClient{}, // Mock repository
}
```

**Why:**
- Cache behavior is integral to the functionality being tested
- Real cache implementations provide more accurate integration testing
- External services should be mocked to avoid network dependencies

---

## DEBUGGING & TROUBLESHOOTING

### Systematic Test Fixing Strategy

When encountering widespread test failures:

1. **Categorize by error type:**
    - Provider function signature mismatches
    - Mock type errors
    - Context type mismatches
    - Interface vs implementation issues

2. **Fix one category at a time:**
    - Use pattern matching tools (`grep`, `sed`) for bulk fixes
    - Don't jump between different error types

3. **Verify incrementally:**
    - Run tests after each category of fixes
    - Use `docker compose run --rm backend make test` to verify progress

4. **Update patterns:**
    - Document new patterns discovered during fixes
    - Add to this instruction file for future reference

### Common Debugging Steps

1. **Run the test** and capture the exact error message
2. **Check operation type** - Are you using `ExpectExec` vs `ExpectQuery` correctly?
3. **Extract SQL pattern** from error log and match exactly
4. **Verify argument count** - Include timestamps for GORM operations
5. **Check transactions** - GORM auto-wraps Create/Delete operations
6. **Update this file** with new patterns you discover

---

## ANTI-PATTERNS: Common Pitfalls to Avoid

### Don't Test Language Mechanics

**❌ Avoid testing interface definitions:**

```go
// BAD - This test is pointless
func TestController(t *testing.T) {
t.Run("implements Controller interface", func (t *testing.T) {
mock := &mockController{}
var controller Controller = mock // Compiler already checks this
assert.NotNil(t, controller)
})
}
```

**✅ Instead, test concrete implementations and behavior:**

```go
// GOOD - Test actual functionality
func TestUserController_CreateUser(t *testing.T) {
// Test the actual business logic of CreateUser method
}
```

**Why:** The Go compiler already enforces interface compliance. Focus tests on business logic and behavior.

### Fiber Context Testing Issues

**⚠️ Common Issue:** Cannot directly create Fiber contexts for unit testing

```go
// ❌ This doesn't work
ctx := app.AcquireCtx(&httptest.Request{}) // Wrong type
```

**✅ Use integration-style testing with Fiber's Test method:**

```go
// ✅ Correct approach
app := fiber.New()
app.Get("/test", handlerFunction)
req := httptest.NewRequest("GET", "/test", nil)
resp, err := app.Test(req)
```

### Empty Request Body Behavior

**Learning:** Empty request bodies cause JSON parsing errors (500), not validation errors (400)

```go
// Empty body with application/json content-type
req := httptest.NewRequest("POST", "/test", strings.NewReader(""))
req.Header.Set("Content-Type", "application/json")
// Results in 500 (JSON parse error), not 400 (validation error)
```

### JSON Marshaling of Empty Maps

**Learning:** Empty `CustomErrorData{}` maps serialize differently than `nil`:

```go
// CustomErrorData{} → serializes as {} → unmarshals as non-nil empty map  
// nil → omitted from JSON (due to omitempty tag) → unmarshals as nil
```

---

## CONTINUOUS IMPROVEMENT

### Auto-Update Enabled

When you encounter new patterns, GORM behaviors, or debugging solutions, immediately add them to this file in the
appropriate sections.

**Learning Categories to Update:**

- New GORM operation behaviors
- SQL pattern variations
- Transaction handling edge cases
- Mock setup requirements
- Fiber error handler configuration
- Cache implementation patterns
- Common pitfalls and resolutions
- Provider function signature updates
- Context handling solutions
- Interface vs implementation issues