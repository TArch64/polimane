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
  hostnames = ["*.${local.domain}", local.domain]
  request_type       = "origin-rsa"
  requested_validity = 5475
}

resource "aws_acm_certificate" "cloudflare" {
  private_key      = tls_private_key.aws_origin.private_key_pem
  certificate_body = cloudflare_origin_ca_certificate.aws_origin.certificate
  tags             = local.aws_common_tags
}

resource "cloudflare_dns_record" "webapp" {
  name    = local.domain
  type    = "CNAME"
  proxied = true
  ttl     = 1
  zone_id = local.cloudflare_zone_id
  content = "${cloudflare_pages_project.webapp.name}.pages.dev"
}
