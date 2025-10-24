import {
  id = "polimane-prod"
  to = aws_s3_bucket.bucket
}

resource "aws_s3_bucket" "bucket" {
  bucket = local.app_name
  tags   = local.aws_common_tags
}

resource "aws_s3_bucket_ownership_controls" "bucket" {
  bucket = aws_s3_bucket.bucket.id

  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_acl" "bucket" {
  depends_on = [aws_s3_bucket_ownership_controls.bucket]

  bucket = aws_s3_bucket.bucket.id
  acl    = "private"
}
