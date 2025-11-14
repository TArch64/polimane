package users

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
)

type listQuery struct {
	IDs []string `query:"ids"`
}

func (c *Controller) apiList(ctx *fiber.Ctx) (err error) {
	var query listQuery
	if err = base.ParseQuery(ctx, &query); err != nil {
		return err
	}

	IDs, err := model.StringsToIDs(query.IDs)
	if err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)
	err = c.userSchemas.FilterByAccess(ctx.Context(), user, &IDs, model.AccessAdmin)
	if err != nil {
		return err
	}
	if len(IDs) == 0 {
		return fiber.ErrBadRequest
	}

	var response listResponse
	eg, egCtx := errgroup.WithContext(ctx.Context())
	_, _ = eg, egCtx

	eg.Go(func() error {
		return c.userSchemas.ListSchemasAccessOut(egCtx, IDs, &response.Users)
	})

	eg.Go(func() error {
		return c.schemaInvitations.ListSchemasAccessOut(egCtx, IDs, &response.Invitations)
	})

	if err = eg.Wait(); err != nil {
		return err
	}

	return ctx.JSON(response)
}
