//go:build !dev

package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"polimane/backend/api"
	"polimane/backend/awsdynamodb"
	"polimane/backend/env"
)

func main() {
	var err error

	err = env.Init()
	if err != nil {
		log.Panic(err)
	}

	err = awsdynamodb.Init(context.Background())
	if err != nil {
		log.Panic(err)
	}

	app := api.New(func(config *fiber.Config) {
		config.Prefork = true
		config.DisableStartupMessage = true
	})

	app.Use(cors.New())

	fiberLambda := fiberadapter.New(app)

	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return fiberLambda.ProxyWithContext(ctx, request)
	})
}
