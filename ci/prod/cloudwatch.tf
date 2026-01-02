locals {
  persistent_streams = [
    "schema_deletion",
    "user_deletion",
    "cleanup_expired_invitations",
  ]
}

resource "aws_cloudwatch_log_group" "persistent" {
  name              = "/${local.app_name}/persistent"
  retention_in_days = 60
  tags              = local.aws_common_tags
}

resource "aws_cloudwatch_log_stream" "persistent_streams" {
  for_each       = toset(local.persistent_streams)
  name           = each.value
  log_group_name = aws_cloudwatch_log_group.persistent.name
}

resource "aws_cloudwatch_log_group" "lambda_logs" {
  name              = "/${local.app_name}/lambda/${local.lambda_name}"
  retention_in_days = 1
  tags              = local.aws_common_tags
}

resource "aws_cloudwatch_log_group" "worker_logs" {
  name              = "/${local.app_name}/lambda/${local.worker_name}"
  retention_in_days = 1
  tags              = local.aws_common_tags
}
