//go:build dev

package api

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func OptionsProvider() *Options {
	return &Options{
		Protocol: "http",
		Configure: func(config *fiber.Config) {
			config.EnablePrintRoutes = true
		},
	}
}

func getErrorHandlerMiddleware() fiber.Handler {
	return recover.New()
}

func Start(app *fiber.App) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		_ = app.Shutdown()
	}()

	return app.Listen(":3000")
}
