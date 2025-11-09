package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
)

type listQuery struct {
	IDs []string `query:"ids"`
}

func (c *Controller) apiList(ctx *fiber.Ctx) error {
	var err error
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

	var response listResponse
	eg, egCtx := errgroup.WithContext(ctx.Context())
	_, _ = eg, egCtx

	eg.Go(func() error {
		return c.listUsers(egCtx, IDs, &response)
	})

	eg.Go(func() error {
		return c.listInvitations(egCtx, IDs, &response)
	})

	if err = eg.Wait(); err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (c *Controller) listUsers(ctx context.Context, schemaIDs []model.ID, res *listResponse) error {
	return c.userSchemas.ListUserSchemaAccessOut(ctx, schemaIDs, &res.Users)
}

func (c *Controller) listInvitations(ctx context.Context, schemaIDs []model.ID, res *listResponse) error {
	return c.schemaInvitations.ListBySchemaIDsOut(ctx, &repositoryschemainvitations.ListBySchemaIDsOptions{
		SchemaIDs: schemaIDs,

		Select: []string{
			"DISTINCT ON (email) email",
			"access",
		},

		Order: []string{"email"},
	}, &res.Invitations)
}
