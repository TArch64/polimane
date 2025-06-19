//go:build !dev

package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"polimane/backend/api"
	"polimane/backend/app"
)

func main() {
	instance, err := app.New(&app.Config{
		ApiOptions: &api.Options{
			Protocol: "https",
			Configure: func(config *fiber.Config) {
				config.DisableStartupMessage = true
			},
		},
	})

	if err != nil {
		log.Panic(err)
	}

	handler := adaptor.FiberApp(instance)

	lambda.Start(func(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
		url := req.RawPath
		if req.RawQueryString != "" {
			url += "?" + req.RawQueryString
		}

		httpReq, _ := http.NewRequestWithContext(ctx, req.RequestContext.HTTP.Method, url, strings.NewReader(req.Body))

		for key, value := range req.Headers {
			httpReq.Header.Set(key, value)
		}

		httpReq.RequestURI = url

		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, httpReq)

		headers := make(map[string]string)
		for key, values := range recorder.Header() {
			if len(values) > 0 {
				headers[key] = strings.Join(values, ", ")
			}
		}

		return events.LambdaFunctionURLResponse{
			StatusCode: recorder.Code,
			Headers:    headers,
			Body:       recorder.Body.String(),
		}, nil
	})
}
