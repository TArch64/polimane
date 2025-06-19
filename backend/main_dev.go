//go:build dev

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/api"
	"polimane/backend/app"
)

func main() {
	instance, err := app.New(&app.Config{
		ApiOptions: &api.Options{
			Protocol: "http",
			Configure: func(config *fiber.Config) {
				config.EnablePrintRoutes = true
			},
		},
	})

	if err != nil {
		log.Panic(err)
	}

	defer func() {
		log.Println("Shutting down application")

		if err = instance.Shutdown(); err != nil {
			log.Println(err)
		}
	}()

	err = instance.Listen(":3000")
	if err != nil {
		log.Panic(err)
	}
}
