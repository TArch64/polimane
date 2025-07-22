//go:build dev

package api

import "github.com/gofiber/fiber/v2"

func OptionsProvider() *Options {
	return &Options{
		Protocol: "http",
		Configure: func(config *fiber.Config) {
			config.EnablePrintRoutes = true
		},
	}
}

func Start(app *fiber.App) error {
	defer func(app *fiber.App) {
		_ = app.Shutdown()
	}(app)

	return app.Listen(":3000")
}
