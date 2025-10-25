resource "aws_sqs_queue" "debounced" {
  name                  = "polimane-debounced.fifo"
  fifo_queue            = true
  deduplication_scope   = "messageGroup"
  delay_seconds         = 300
  fifo_throughput_limit = "perMessageGroupId"
  tags                  = local.aws_common_tags
}

resource "aws_sqs_queue" "scheduled" {
  name = "polimane-scheduled"
  tags = local.aws_common_tags
}
