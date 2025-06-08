import {
  id = "polimane-prod"
  to = aws_s3_bucket.bucket
}

resource "aws_s3_bucket" "bucket" {
  bucket = "polimane-prod"
  tags = local.aws_common_tags
}
