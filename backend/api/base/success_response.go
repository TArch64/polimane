package base

import "github.com/gofiber/fiber/v2"

func NewSuccessResponse(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"success": true,
	})
}
