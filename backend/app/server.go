package app

import "github.com/gofiber/fiber/v2"

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "Polymane",
		Prefork: true,
	})

	group := app.Group("/api")

	group.Get("", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}
