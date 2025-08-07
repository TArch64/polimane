package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
)

type loginCompleteQuery struct {
	Code string `query:"code"`
}

func (c *Controller) apiLoginComplete(ctx *fiber.Ctx) error {
	var query loginCompleteQuery
	if err := ctx.QueryParser(&query); err != nil {
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

	user, err := c.users.CreateIfNeeded(reqCtx, data.User.ID)
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

	redirectUrl := c.env.AppURL().JoinPath("auth/complete")
	redirectQuery := redirectUrl.Query()
	redirectQuery.Set("access-token", data.AccessToken)
	redirectQuery.Set("refresh-token", data.RefreshToken)
	redirectUrl.RawQuery = redirectQuery.Encode()
	return ctx.Redirect(redirectUrl.String())
}
