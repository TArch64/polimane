package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/env"
	"polimane/backend/services/workos"
)

func apiLogin(ctx *fiber.Ctx) error {
	url, err := workos.UserManagement.GetAuthorizationURL(usermanagement.GetAuthorizationURLOpts{
		ClientID:    env.Instance.WorkOS.ClientID,
		RedirectURI: env.Instance.ApiURL().JoinPath("api/auth/login/complete").String(),
		Provider:    "authkit",
	})

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"url": url.String(),
	})
}
