locals {
  lambda_name = "polimane-prod"
}

resource "aws_lambda_function" "lambda" {
  depends_on = [data.external.lambda_build]
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
    variables = sensitive(yamldecode(file("${path.module}/.env-lambda.yaml")))
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
