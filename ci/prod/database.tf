data "aws_dynamodb_table" "database" {
  depends_on = [aws_lambda_function.lambda]
  name = "polimane-prod"
}
