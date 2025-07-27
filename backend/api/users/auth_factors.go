package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/auth"
)

func (c *Controller) apiListAuthFactors(ctx *fiber.Ctx) error {
	user := auth.GetSessionUser(ctx)

	factors, err := c.workosClient.UserManagement.ListAuthFactors(ctx.Context(), usermanagement.ListAuthFactorsOpts{
		User: user.WorkosID,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(factors.Data)
}
