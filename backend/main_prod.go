//go:build !dev

package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"

	"polimane/backend/app"
)

func main() {
	api, err := app.New(&app.Config{
		ApiConfig: func(config *fiber.Config) {
			config.Prefork = true
			config.DisableStartupMessage = true
		},
	})

	if err != nil {
		log.Panic(err)
	}

	fiberLambda := fiberadapter.New(api)

	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return fiberLambda.ProxyWithContext(ctx, request)
	})
}
