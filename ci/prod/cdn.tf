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
  enabled             = true
  http_version        = "http2and3"
  price_class         = "PriceClass_100"
  wait_for_deployment = true
  comment             = local.app_name

  tags = merge(local.aws_common_tags, {
    Name = local.app_name
  })

  viewer_certificate {
    cloudfront_default_certificate = true
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
    allowed_methods = ["HEAD", "GET", "OPTIONS"]
    cached_methods = ["HEAD", "GET"]
    target_origin_id       = aws_s3_bucket.bucket.id
    cache_policy_id        = data.aws_cloudfront_cache_policy.caching_optimized.id
    viewer_protocol_policy = "https-only"
  }

  ordered_cache_behavior {
    allowed_methods = ["HEAD", "GET", "OPTIONS"]
    cached_methods = ["HEAD", "GET"]
    path_pattern           = "/images/*"
    target_origin_id       = aws_s3_bucket.bucket.id
    cache_policy_id        = data.aws_cloudfront_cache_policy.images.id
    viewer_protocol_policy = "https-only"
    compress               = false
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
