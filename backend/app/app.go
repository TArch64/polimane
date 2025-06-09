package app

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/api"
	"polimane/backend/env"
	awsdynamodb "polimane/backend/services/dynamodb"
	awsssm "polimane/backend/services/ssm"
)

type Config struct {
	ApiConfig api.Config
}

func New(config *Config) (*fiber.App, error) {
	var err error

	err = env.Init()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	err = awsssm.Init(ctx)
	if err != nil {
		return nil, err
	}

	err = awsdynamodb.Init(ctx)
	if err != nil {
		return nil, err
	}

	return api.New(config.ApiConfig), nil
}
