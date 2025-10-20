resource "aws_apigatewayv2_api" "lambda_api" {
  name          = "${local.lambda_name}-api"
  protocol_type = "HTTP"
  tags          = local.aws_common_tags
  route_selection_expression = "$request.method $request.path"

  cors_configuration {
    allow_origins = ["https://${local.webapp_domain}"]
    allow_methods = ["*"]
    allow_headers = [
      "Origin",
      "Content-Type",
      "Accept",
      "Authorization",
      "X-Refresh-Token",
      "X-Requested-With",
      "X-CSRF-Token",
      "Cookie"
    ]
    allow_credentials = true
    expose_headers = ["Set-Cookie"]
    max_age = 300
  }
}

resource "aws_apigatewayv2_integration" "lambda_integration" {
  api_id                 = aws_apigatewayv2_api.lambda_api.id
  integration_type       = "AWS_PROXY"
  integration_uri        = aws_lambda_function.lambda.invoke_arn
  payload_format_version = "2.0"
  integration_method     = "POST"
}

resource "aws_apigatewayv2_route" "lambda_route" {
  api_id    = aws_apigatewayv2_api.lambda_api.id
  route_key = "ANY /{proxy+}"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_integration.id}"
}

resource "aws_apigatewayv2_stage" "lambda_stage" {
  api_id      = aws_apigatewayv2_api.lambda_api.id
  name        = "$default"
  auto_deploy = true
  tags        = local.aws_common_tags

  default_route_settings {
    throttling_burst_limit = 10
    throttling_rate_limit  = 2
  }
}

resource "aws_apigatewayv2_domain_name" "cloudflare" {
  domain_name = local.api_domain
  tags        = local.aws_common_tags

  domain_name_configuration {
    certificate_arn = aws_acm_certificate.cloudflare.arn
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }
}

resource "aws_apigatewayv2_api_mapping" "example" {
  api_id      = aws_apigatewayv2_api.lambda_api.id
  domain_name = aws_apigatewayv2_domain_name.cloudflare.id
  stage       = aws_apigatewayv2_stage.lambda_stage.id
}

resource "aws_lambda_permission" "api_gateway_invoke" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.lambda_api.execution_arn}/*/*/*"
}
