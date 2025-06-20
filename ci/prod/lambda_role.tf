resource "aws_iam_role" "lambda_role" {
  name = "${local.lambda_name}-role"
  tags = local.aws_common_tags
  assume_role_policy = file("${path.module}/lambda_assume_role.json")
}

resource "aws_iam_role_policy_attachment" "lambda_basic_execution" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy" "lambda_permissions" {
  role = aws_iam_role.lambda_role.name
  name = "X-${local.lambda_name}-Permissions"
  policy = file("${path.module}/lambda_permissions_policy.json")
}
