locals {
  cloudflare_zone_id    = "a42f20d4ce852205c537bb0bb8eda260"
  cloudflare_account_id = "b9b2371ca5c6bb7fe6d42ed2a37b04ad"
  api_domain            = "api.${local.domain}"
}

resource "cloudflare_dns_record" "api" {
  name    = local.api_domain
  type    = "CNAME"
  proxied = true
  ttl     = 1
  zone_id = local.cloudflare_zone_id
  content = aws_apigatewayv2_domain_name.cloudflare.domain_name_configuration[0].target_domain_name
}

resource "tls_private_key" "aws_origin" {
  algorithm = "RSA"
}

resource "tls_cert_request" "aws_origin" {
  private_key_pem = tls_private_key.aws_origin.private_key_pem

  subject {
    common_name  = local.domain
    organization = "Taras Turchenko"
  }
}

resource "cloudflare_origin_ca_certificate" "aws_origin" {
  csr                = tls_cert_request.aws_origin.cert_request_pem
  hostnames          = ["*.${local.domain}", local.domain]
  request_type       = "origin-rsa"
  requested_validity = 5475
}

resource "aws_acm_certificate" "cloudflare" {
  private_key      = tls_private_key.aws_origin.private_key_pem
  certificate_body = cloudflare_origin_ca_certificate.aws_origin.certificate
  tags             = local.aws_common_tags

  lifecycle {
    create_before_destroy = true
  }
}

# only used to create a redirect from old room domain to webapp subdomain
resource "cloudflare_dns_record" "root" {
  name    = local.domain
  type    = "A"
  proxied = true
  ttl     = 1
  zone_id = local.cloudflare_zone_id
  content = "192.0.2.1" # Dummy IP (TEST-NET-1). Cloudflare redirect before origin is used.
}

resource "aws_acm_certificate" "cloudfront" {
  provider          = aws.us_east_1
  domain_name       = "*.${local.domain}"
  validation_method = "DNS"
  tags              = local.aws_common_tags

  lifecycle {
    create_before_destroy = true
  }
}

resource "cloudflare_dns_record" "cloudfront_validation" {
  for_each = {
    for dvo in aws_acm_certificate.cloudfront.domain_validation_options : dvo.domain_name => {
      name   = trimsuffix(dvo.resource_record_name, ".")
      record = trimsuffix(dvo.resource_record_value, ".")
      type   = dvo.resource_record_type
    }
  }

  zone_id = local.cloudflare_zone_id
  name    = each.value.name
  content = each.value.record
  type    = each.value.type
  ttl     = 60
}

resource "aws_acm_certificate_validation" "cloudfront" {
  provider        = aws.us_east_1
  certificate_arn = aws_acm_certificate.cloudfront.arn
  validation_record_fqdns = [
    for record in cloudflare_dns_record.cloudfront_validation :
    "${record.name}."
  ]
}

resource "cloudflare_dns_record" "cdn" {
  zone_id = local.cloudflare_zone_id
  name    = local.cdn_domain
  content = aws_cloudfront_distribution.cdn.domain_name
  type    = "CNAME"
  ttl     = 1
  proxied = false
}
