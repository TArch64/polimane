package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/env"
	"polimane/backend/services/workos"
)

type loginQueryParams struct {
	ReturnTo string `query:"return-to"`
}

func apiLogin(ctx *fiber.Ctx) error {
	var query loginQueryParams
	if err := ctx.QueryParser(&query); err != nil {
		return err
	}

	state, err := newLoginState(&query)
	if err != nil {
		return err
	}

	url, err := workos.UserManagement.GetAuthorizationURL(usermanagement.GetAuthorizationURLOpts{
		ClientID:    env.Instance.WorkOS.ClientID,
		RedirectURI: env.Instance.ApiURL().JoinPath("api/auth/login/complete").String(),
		State:       state,
		Provider:    "authkit",
	})

	if err != nil {
		return err
	}

	return ctx.Redirect(url.String(), fiber.StatusFound)
}
