package ping

import "github.com/gofiber/fiber/v2"

func apiPing(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "pong"})
}
