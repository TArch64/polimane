package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
)

func (c *Controller) PasswordReset(ctx *fiber.Ctx) error {
	session := auth.GetSession(ctx)

	_, err := c.workos.UserManagement.CreatePasswordReset(ctx.Context(), usermanagement.CreatePasswordResetOpts{
		Email: session.WorkosUser.Email,
	})

	if err != nil {
		return err
	}

	c.signals.InvalidateAuthCache.Emit(ctx.Context(), session.ID)
	return base.NewSuccessResponse(ctx)
}
