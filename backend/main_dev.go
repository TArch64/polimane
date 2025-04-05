//go:build dev

package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"polimane/backend/api"
	"polimane/backend/awsdynamodb"
	"polimane/backend/env"
)

func main() {
	var err error

	err = env.Init()
	if err != nil {
		log.Panic(err)
	}

	err = awsdynamodb.Init(context.Background())
	if err != nil {
		log.Panic(err)
	}

	app := api.New(func(config *fiber.Config) {
		config.EnablePrintRoutes = true
	})

	app.Use(cors.New())

	err = app.Listen(":3000")
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		log.Println("Shutting down application")

		if err = app.Shutdown(); err != nil {
			log.Println(err)
		}
	}()
}
