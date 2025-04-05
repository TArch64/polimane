package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"

	"polimane/backend/api/users"
	"polimane/backend/env"
)

type ConfigureApp func(config *fiber.Config)

func New(configFns ...ConfigureApp) *fiber.App {
	config := fiber.Config{
		AppName:      "Polymane",
		ErrorHandler: base.ErrorHandler,
	}

	for _, configure := range configFns {
		configure(&config)
	}

	app := fiber.New(config)

	app.Use(recover2.New(recover2.Config{
		EnableStackTrace: true,
	}))

	app.Use(helmet.New())

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: env.Env().SecretKey,
	}))

	base.InitValidator()

	group := app.Group("/api")
	auth.Group(group)
	users.Group(group)

	return app
}
