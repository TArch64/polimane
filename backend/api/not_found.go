package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	log.Println("Unhandled route:", c.Path())

	return c.
		Status(404).
		JSON(fiber.Map{"error": "Not Found"})
}
