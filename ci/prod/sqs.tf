resource "aws_sqs_queue" "debounced" {
  name                      = "polimane-debounced.fifo"
  deduplication_scope       = "messageGroup"
  fifo_throughput_limit     = "perMessageGroupId"
  fifo_queue                = true
  delay_seconds             = 300
  receive_wait_time_seconds = 20
  tags                      = local.aws_common_tags
}

resource "aws_sqs_queue" "scheduled" {
  name                      = "polimane-scheduled"
  receive_wait_time_seconds = 20
  tags                      = local.aws_common_tags
}
