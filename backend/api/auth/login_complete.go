package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/env"
	repositoryusers "polimane/backend/repository/users"
	"polimane/backend/services/workos"
)

type loginCompleteQuery struct {
	Code  string `query:"code"`
	State string `query:"state"`
}

func apiLoginComplete(ctx *fiber.Ctx) error {
	var query loginCompleteQuery
	if err := ctx.QueryParser(&query); err != nil {
		return err
	}

	reqCtx := ctx.Context()
	data, err := workos.UserManagement.AuthenticateWithCode(reqCtx, usermanagement.AuthenticateWithCodeOpts{
		ClientID:  env.Instance.WorkOS.ClientID,
		Code:      query.Code,
		UserAgent: ctx.Get("User-Agent"),
	})

	if err != nil {
		return err
	}

	user, err := repositoryusers.CreateIfNeeded(reqCtx, data.User.ID)
	if err != nil {
		return err
	}

	_, err = workos.UserManagement.UpdateUser(reqCtx, usermanagement.UpdateUserOpts{
		User:       data.User.ID,
		ExternalID: user.ID.String(),
	})

	if err != nil {
		return err
	}

	state, err := parseLoginState(query.State)
	if err != nil {
		return err
	}

	redirectUrl := env.Instance.AppURL().JoinPath(state.ReturnTo)
	redirectQuery := redirectUrl.Query()
	redirectQuery.Set("access-token", data.AccessToken)
	redirectQuery.Set("refresh-token", data.RefreshToken)
	redirectUrl.RawQuery = redirectQuery.Encode()
	return ctx.Redirect(redirectUrl.String())
}
