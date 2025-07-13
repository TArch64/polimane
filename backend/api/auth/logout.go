package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/env"
	"polimane/backend/services/workos"
)

func apiLogout(ctx *fiber.Ctx) error {
	url, err := workos.UserManagement.GetLogoutURL(usermanagement.GetLogoutURLOpts{
		SessionID: GetSessionID(ctx),
		ReturnTo:  env.Instance.AppURL().JoinPath("auth").String(),
	})

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"url": url.String(),
	})
}
