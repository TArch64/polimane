//go:build dev

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/app"
)

func main() {
	api, err := app.New(&app.Config{
		ApiConfig: func(config *fiber.Config) {
			config.EnablePrintRoutes = true
		},
	})

	if err != nil {
		log.Panic(err)
	}

	err = api.Listen(":3000")
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		log.Println("Shutting down application")

		if err = api.Shutdown(); err != nil {
			log.Println(err)
		}
	}()
}
