package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
)

func (c *Controller) apiLogout(ctx *fiber.Ctx) error {
	url, err := c.workosClient.UserManagement.GetLogoutURL(usermanagement.GetLogoutURLOpts{
		SessionID: GetSessionID(ctx),
		ReturnTo:  c.env.AppURL().JoinPath("auth").String(),
	})

	if err != nil {
		return err
	}

	c.signals.InvalidateAuthCache.Emit(ctx.Context(), GetSessionID(ctx))

	return ctx.JSON(fiber.Map{
		"url": url.String(),
	})
}
