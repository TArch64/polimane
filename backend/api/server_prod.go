//go:build !dev

package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	sentryfiber "github.com/getsentry/sentry-go/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func OptionsProvider() *Options {
	return &Options{
		Protocol: "https",
		Configure: func(config *fiber.Config) {
			config.DisableStartupMessage = true
		},
	}
}

func getErrorHandlerMiddleware() fiber.Handler {
	return sentryfiber.New(sentryfiber.Options{
		Repanic:         true,
		WaitForDelivery: true,
	})
}

func convertIncomingRequest(ctx context.Context, req events.APIGatewayV2HTTPRequest) *http.Request {
	url := req.RawPath
	if req.RawQueryString != "" {
		url += "?" + req.RawQueryString
	}

	httpReq, _ := http.NewRequestWithContext(
		ctx,
		req.RequestContext.HTTP.Method,
		url,
		strings.NewReader(req.Body),
	)

	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	httpReq.RequestURI = url
	return httpReq
}

func convertOutcomingResponse(recorder *httptest.ResponseRecorder) *events.APIGatewayV2HTTPResponse {
	headers := make(map[string]string)
	var cookies []string

	for key, values := range recorder.Header() {
		if strings.ToLower(key) == "set-cookie" {
			cookies = append(cookies, values...)
		} else if len(values) > 0 {
			headers[key] = strings.Join(values, ", ")
		}
	}

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: recorder.Code,
		Headers:    headers,
		Cookies:    cookies,
		Body:       recorder.Body.String(),
	}
}

func runHandler(
	ctx context.Context,
	handler http.Handler,
	req events.APIGatewayV2HTTPRequest,
) (events.APIGatewayV2HTTPResponse, error) {
	httpReq := convertIncomingRequest(ctx, req)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, httpReq)
	res := convertOutcomingResponse(recorder)
	return *res, nil
}

func Start(app *fiber.App) error {
	handler := adaptor.FiberApp(app)

	lambda.Start(func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		return runHandler(ctx, handler, req)
	})

	return nil
}
