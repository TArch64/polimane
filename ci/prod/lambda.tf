locals {
  migrations_hash = sha1(join("", [
    for f in fileset("${local.lambda_sources_dir}/migrations", "**/*.sql")
    : filesha1("${local.lambda_sources_dir}/migrations/${f}")
  ]))

  lambda_name = local.app_name

  lambda_environment = {
    BACKEND_APP_DOMAIN           = local.domain
    BACKEND_APP_PROTOCOL         = "https"
    BACKEND_SENTRY_RELEASE       = local.lambda_sources_hash,
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

resource "aws_lambda_function" "lambda" {
  depends_on       = [null_resource.lambda_build]
  filename         = local.lambda_build_zip
  function_name    = local.lambda_name
  role             = aws_iam_role.lambda_role.arn
  handler          = "lambda"
  runtime          = "provided.al2023"
  architectures    = ["arm64"]
  timeout          = 30
  memory_size      = 128
  source_code_hash = local.lambda_sources_hash
  tags             = local.aws_common_tags

  environment {
    variables = local.lambda_environment
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

resource "null_resource" "lambda_migrations" {
  triggers   = { sources_hash = local.migrations_hash }
  depends_on = [aws_lambda_function.lambda]

  provisioner "local-exec" {
    command = "bash ${path.module}/job/run.sh"

    environment = {
      JOB_ID           = local.migrations_hash
      BUILD_IMAGE      = "polimane-prod-backend-migrations"
      BUILD_DOCKERFILE = abspath("${path.root}/job/backend.Dockerfile")
      BUILD_CONTEXT    = local.lambda_sources_dir

      BUILD_SECRET = jsonencode(["BACKEND_DATABASE_URL", "BACKEND_DATABASE_CERT"])
      # nonsensitive is safe here because values passed using build secrets
      BACKEND_DATABASE_URL  = nonsensitive(bitwarden_secret.backend_database_url.value)
      BACKEND_DATABASE_CERT = nonsensitive(bitwarden_secret.backend_database_cert.value)
    }
  }
}

