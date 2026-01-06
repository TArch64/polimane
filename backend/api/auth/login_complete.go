package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/api/base"
)

type LoginCompleteQuery struct {
	Code string `query:"code" validate:"required"`
}

func (c *Controller) LoginComplete(ctx *fiber.Ctx) error {
	var query LoginCompleteQuery
	if err := base.ParseQuery(ctx, &query); err != nil {
		return err
	}

	reqCtx := ctx.Context()
	data, err := c.workos.UserManagement.AuthenticateWithCode(reqCtx, usermanagement.AuthenticateWithCodeOpts{
		ClientID:  c.env.WorkOS.ClientID,
		Code:      query.Code,
		UserAgent: ctx.Get("User-Agent"),
	})

	if err != nil {
		return err
	}

	user, flags, err := c.users.CreateIfNeeded(reqCtx, &data.User)
	if err != nil {
		return err
	}
	if user.SoftDeletable != nil && user.DeletedAt.Valid {
		return c.completeRedirect(ctx, map[string]string{
			"deleted": "1",
		})
	}

	if flags.NeedSyncSchemaCreatedCounter {
		err = c.subscriptionCounters.SyncSchemasCreated(reqCtx, user.ID)
		if err != nil {
			return err
		}
	}

	_, err = c.workos.UserManagement.UpdateUser(reqCtx, usermanagement.UpdateUserOpts{
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

	return c.completeRedirect(ctx, nil)
}

func (c *Controller) completeRedirect(ctx *fiber.Ctx, query map[string]string) error {
	redirectUrl := c.env.AppURL.JoinPath("auth/complete")

	if len(query) > 0 {
		redirectQuery := redirectUrl.Query()
		for key, value := range query {
			redirectQuery.Set(key, value)
		}
		redirectUrl.RawQuery = redirectQuery.Encode()
	}

	return ctx.Redirect(redirectUrl.String())
}
