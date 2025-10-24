locals {
  cdn_domain = "cdn.${local.domain}"
}

data "aws_cloudfront_cache_policy" "caching_optimized" {
  name = "Managed-CachingOptimized"
}

data "aws_cloudfront_cache_policy" "images" {
  name = "polimane-images"
}

resource "aws_cloudfront_origin_access_control" "cdn" {
  name                              = "${local.app_name}-cdn-oac"
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

resource "aws_cloudfront_distribution" "cdn" {
  http_version        = "http2and3"
  price_class         = "PriceClass_100"
  aliases             = [local.cdn_domain]
  enabled             = true
  wait_for_deployment = true
  comment             = local.app_name

  tags = merge(local.aws_common_tags, {
    Name = local.app_name
  })

  viewer_certificate {
    acm_certificate_arn      = aws_acm_certificate.cloudfront.arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1.2_2021"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  origin {
    origin_id                = aws_s3_bucket.bucket.id
    domain_name              = aws_s3_bucket.bucket.bucket_regional_domain_name
    origin_access_control_id = aws_cloudfront_origin_access_control.cdn.id
    origin_path              = "/data"
  }

  default_cache_behavior {
    allowed_methods        = ["HEAD", "GET", "OPTIONS"]
    cached_methods         = ["HEAD", "GET"]
    target_origin_id       = aws_s3_bucket.bucket.id
    cache_policy_id        = data.aws_cloudfront_cache_policy.caching_optimized.id
    viewer_protocol_policy = "https-only"
  }

  ordered_cache_behavior {
    allowed_methods            = ["HEAD", "GET", "OPTIONS"]
    cached_methods             = ["HEAD", "GET"]
    path_pattern               = "/images/*"
    target_origin_id           = aws_s3_bucket.bucket.id
    cache_policy_id            = data.aws_cloudfront_cache_policy.images.id
    response_headers_policy_id = aws_cloudfront_response_headers_policy.images.id
    viewer_protocol_policy     = "https-only"
    compress                   = false
  }
}

resource "aws_s3_bucket_policy" "cdn" {
  bucket = aws_s3_bucket.bucket.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowCloudFrontServicePrincipal"
        Effect = "Allow"
        Principal = {
          Service = "cloudfront.amazonaws.com"
        }
        Action   = "s3:GetObject"
        Resource = "${aws_s3_bucket.bucket.arn}/*"
        Condition = {
          StringEquals = {
            "AWS:SourceArn" = aws_cloudfront_distribution.cdn.arn
          }
        }
      }
    ]
  })
}

resource "aws_cloudfront_response_headers_policy" "images" {
  name = "${local.app_name}-images"

  cors_config {
    origin_override                  = false
    access_control_max_age_sec       = 600
    access_control_allow_credentials = false

    access_control_allow_headers {
      items = ["*"]
    }

    access_control_allow_methods {
      items = ["GET", "OPTIONS"]
    }

    access_control_allow_origins {
      items = ["https://${local.webapp_domain}"]
    }
  }
}

