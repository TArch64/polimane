locals {
  worker_name = "${local.app_name}-worker"
}

resource "aws_lambda_function" "worker" {
  depends_on       = [null_resource.worker_build]
  filename         = local.worker_build_zip
  function_name    = local.worker_name
  role             = aws_iam_role.lambda_role.arn
  handler          = "lambda"
  runtime          = "provided.al2023"
  architectures    = ["arm64"]
  timeout          = 30
  memory_size      = 128
  source_code_hash = local.worker_sources_hash
  tags             = local.aws_common_tags

  environment {
    variables = local.lambda_environment
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_lambda_event_source_mapping" "sqs_worker_debounced" {
  function_name                      = aws_lambda_function.worker.function_name
  event_source_arn                   = aws_sqs_queue.debounced.arn
  maximum_batching_window_in_seconds = 300
  tags                               = local.aws_common_tags
}

resource "aws_lambda_event_source_mapping" "sqs_worker_scheduled" {
  function_name                      = aws_lambda_function.worker.function_name
  event_source_arn                   = aws_sqs_queue.scheduled.arn
  maximum_batching_window_in_seconds = 300
  tags                               = local.aws_common_tags
}

resource "aws_cloudwatch_log_group" "worker_logs" {
  name              = "/aws/lambda/${local.worker_name}"
  retention_in_days = 1
  tags              = local.aws_common_tags
}

