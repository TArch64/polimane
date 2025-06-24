data "bitwarden_secret" "backend_default_user" {
  key = "backend_default_user"
}

data "bitwarden_secret" "backend_default_password" {
  key = "backend_default_password"
}

data "bitwarden_secret" "backend_secret_key" {
  key = "backend_secret_key"
}

data "bitwarden_secret" "backend_sentry_dsn" {
  key = "backend_sentry_dsn"
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
      BACKEND_APP_DOMAIN       = local.domain
      BACKEND_SENTRY_RELEASE   = local.lambda_sources_hash,
      BACKEND_BITWARDEN_TOKEN      = var.bitwarden_token
      BACKEND_DEFAULT_USER_SID     = data.bitwarden_secret.backend_default_user.id
      BACKEND_DEFAULT_PASSWORD_SID = data.bitwarden_secret.backend_default_password.id
      BACKEND_SECRET_KEY_SID       = data.bitwarden_secret.backend_secret_key.id
      BACKEND_SENTRY_DSN_SID       = data.bitwarden_secret.backend_sentry_dsn.id
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
