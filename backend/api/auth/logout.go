package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/base"
)

func (c *Controller) apiLogout(ctx *fiber.Ctx) error {
	session := GetSession(ctx)

	err := c.workosClient.UserManagement.RevokeSession(ctx.Context(), usermanagement.RevokeSessionOpts{
		SessionID: session.ID,
	})

	if err != nil {
		return err
	}

	c.signals.InvalidateAuthCache.Emit(ctx.Context(), session.ID)

	return base.NewSuccessResponse(ctx)
}
