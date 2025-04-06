package app

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/api"
	"polimane/backend/awsdynamodb"
	"polimane/backend/env"
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

	err = awsdynamodb.Init(context.Background())
	if err != nil {
		return nil, err
	}

	return api.New(config.ApiConfig), nil
}
