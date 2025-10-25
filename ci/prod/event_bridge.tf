resource "aws_scheduler_schedule" "cleanup_invitations" {
  name                         = "${local.app_name}-cleanup-invitations"
  schedule_expression          = "cron(0 2 * * ? *)"
  schedule_expression_timezone = "Europe/Kyiv"

  flexible_time_window {
    mode                      = "FLEXIBLE"
    maximum_window_in_minutes = 15
  }

  target {
    arn      = aws_sqs_queue.scheduled.arn
    role_arn = aws_iam_role.scheduler_sqs.arn

    input = jsonencode({
      eventType = "cleanup-invitations"
      payload   = {}
    })
  }
}

resource "aws_iam_role" "scheduler_sqs" {
  name = "${local.app_name}-scheduler-sqs"
  tags = local.aws_common_tags

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action    = "sts:AssumeRole"
        Effect    = "Allow"
        Principal = { Service = "scheduler.amazonaws.com" }
      }
    ]
  })
}

resource "aws_iam_role_policy" "scheduler_sqs" {
  name = "${local.app_name}-scheduler-sqs"
  role = aws_iam_role.scheduler_sqs.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = ["sqs:SendMessage"]
        Resource = aws_sqs_queue.scheduled.arn
      }
    ]
  })
}
