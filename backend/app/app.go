package app

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/api"
	"polimane/backend/env"
	"polimane/backend/services/bitwarden"
	awsdynamodb "polimane/backend/services/dynamodb"
	awsssm "polimane/backend/services/ssm"
)

type Config struct {
	ApiOptions *api.Options
}

func New(config *Config) (*fiber.App, error) {
	var err error

	err = bitwarden.Init()
	if err != nil {
		return nil, err
	}

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

	return api.New(config.ApiOptions)
}
