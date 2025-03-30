package api

import "github.com/gofiber/fiber/v2"

type ConfigureApp func(config *fiber.Config)

func New(configFns ...ConfigureApp) *fiber.App {
	config := fiber.Config{
		AppName: "Polymane",
	}

	for _, configure := range configFns {
		configure(&config)
	}

	app := fiber.New(config)

	group := app.Group("/api")

	group.Get("", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}
