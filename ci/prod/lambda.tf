variable "backend_default_user" {
  type      = string
  sensitive = true
  nullable  = false
}

variable "backend_default_password" {
  type      = string
  sensitive = true
  nullable  = false
}

variable "backend_secret_key" {
  type      = string
  sensitive = true
  nullable  = false
}

variable "backend_sentry_dsn" {
  type      = string
  sensitive = true
  nullable  = false
}

locals {
  lambda_name = local.app_name
}

resource "aws_lambda_function" "lambda" {
  depends_on = [null_resource.lambda_build]
  filename = local.lambda_build_zip
  function_name    = local.lambda_name
  role             = aws_iam_role.lambda_role.arn
  handler          = "lambda"
  runtime          = "provided.al2023"
  timeout          = 30
  memory_size      = 128
  source_code_hash = local.lambda_sources_hash
  tags             = local.aws_common_tags

  environment {
    variables = {
      BACKEND_DEFAULT_USER     = var.backend_default_user
      BACKEND_DEFAULT_PASSWORD = var.backend_default_password
      BACKEND_SECRET_KEY       = var.backend_secret_key
      BACKEND_APP_DOMAIN       = local.domain
      BACKEND_SENTRY_DSN       = var.backend_sentry_dsn
      BACKEND_SENTRY_RELEASE   = local.lambda_sources_hash,
    }
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_cloudwatch_log_group" "lambda_logs" {
  name              = "/aws/lambda/${local.lambda_name}"
  retention_in_days = 1
  tags              = local.aws_common_tags
}
