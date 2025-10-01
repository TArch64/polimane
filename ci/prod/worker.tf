locals {
  worker_name = "${local.app_name}-worker"
}

resource "aws_lambda_function" "worker" {
  depends_on = [null_resource.worker_build]
  filename         = local.worker_build_zip
  function_name    = local.worker_name
  role             = aws_iam_role.lambda_role.arn
  handler          = "lambda"
  runtime          = "provided.al2023"
  architectures = ["arm64"]
  timeout          = 30
  memory_size      = 128
  source_code_hash = local.worker_sources_hash
  tags             = local.aws_common_tags

  environment {
    variables = {
      BACKEND_APP_DOMAIN           = local.domain
      BACKEND_APP_PROTOCOL         = "https"
      BACKEND_SENTRY_RELEASE       = local.worker_sources_hash,
      BACKEND_BITWARDEN_TOKEN      = var.bitwarden_token
      BACKEND_SECRET_KEY_SID       = data.bitwarden_secret.backend_secret_key.id
      BACKEND_SENTRY_DSN_SID       = data.bitwarden_secret.backend_sentry_dsn.id,
      BACKEND_DATABASE_URL_SID     = bitwarden_secret.backend_database_url.id,
      BACKEND_DATABASE_CERT_SID    = bitwarden_secret.backend_database_cert.id,
      BACKEND_WORKOS_CLIENT_ID_SID = data.bitwarden_secret.backend_workos_client_id.id,
      BACKEND_WORKOS_API_KEY_SID   = data.bitwarden_secret.backend_workos_api_key.id,
      BACKEND_SQS_BASE_URL_SID     = data.bitwarden_secret.backend_sqs_base_url.id,
    }
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_lambda_event_source_mapping" "sqs_worker_debounced" {
  function_name    = aws_lambda_function.worker.function_name
  event_source_arn = aws_sqs_queue.debounced.arn
  tags             = local.aws_common_tags
}

resource "aws_cloudwatch_log_group" "worker_logs" {
  name              = "/aws/lambda/${local.worker_name}"
  retention_in_days = 1
  tags              = local.aws_common_tags
}

