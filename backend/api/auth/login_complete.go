package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/base"
)

type loginCompleteQuery struct {
	Code string `query:"code" validate:"required"`
}

func (c *Controller) apiLoginComplete(ctx *fiber.Ctx) error {
	var query loginCompleteQuery
	if err := base.ParseQuery(ctx, &query); err != nil {
		return err
	}

	reqCtx := ctx.Context()
	data, err := c.workosClient.UserManagement().AuthenticateWithCode(reqCtx, usermanagement.AuthenticateWithCodeOpts{
		ClientID:  c.env.WorkOS.ClientID,
		Code:      query.Code,
		UserAgent: ctx.Get("User-Agent"),
	})

	if err != nil {
		return err
	}

	user, err := c.users.CreateIfNeeded(reqCtx, data.User)
	if err != nil {
		return err
	}

	_, err = c.workosClient.UserManagement().UpdateUser(reqCtx, usermanagement.UpdateUserOpts{
		User:       data.User.ID,
		ExternalID: user.ID.String(),
	})

	if err != nil {
		return err
	}

	setCookies(ctx, c.env, &authCookies{
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
	})

	redirectUrl := c.env.AppURL.JoinPath("auth/complete")
	return ctx.Redirect(redirectUrl.String())
}
