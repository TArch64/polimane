package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
)

func (c *Controller) apiCurrent(ctx *fiber.Ctx) error {
	return ctx.JSON(auth.GetSession(ctx))
}
