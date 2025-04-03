package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"

	"polimane/backend/api/users"
)

type ConfigureApp func(config *fiber.Config)

func New(configFns ...ConfigureApp) *fiber.App {
	config := fiber.Config{
		AppName: "Polymane",
	}

	for _, configure := range configFns {
		configure(&config)
	}

	app := fiber.New(config)

	app.Use(recover2.New(recover2.Config{
		EnableStackTrace: true,
	}))

	app.Use(helmet.New())

	group := app.Group("/api")
	users.Group(group)

	return app
}
