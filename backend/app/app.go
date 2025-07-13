package app

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api"
	"polimane/backend/env"
	"polimane/backend/services/db"
	"polimane/backend/services/workos"
)

type Config struct {
	ApiOptions *api.Options
}

func New(config *Config) (*fiber.App, error) {
	var err error

	if err = env.Init(); err != nil {
		return nil, err
	}

	if err = db.Init(); err != nil {
		return nil, err
	}

	workos.Init()
	return api.New(config.ApiOptions)
}
