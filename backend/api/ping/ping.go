package ping

import "github.com/gofiber/fiber/v2"

func (c *Controller) apiPing(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "pong"})
}
