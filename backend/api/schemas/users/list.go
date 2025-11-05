package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

type listQuery struct {
	IDs []string `query:"ids"`
}

type listResponse struct {
	Users       []*listUser       `json:"users"`
	Invitations []*listInvitation `json:"invitations"`
}

type listUser struct {
	ID        model.ID          `json:"id"`
	Email     string            `json:"email"`
	FirstName string            `json:"firstName"`
	LastName  string            `json:"lastName"`
	Access    model.AccessLevel `json:"access"`
}

func newUserListItem(user *model.User, access model.AccessLevel) *listUser {
	return &listUser{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Access:    access,
	}
}

type listInvitation struct {
	Email  string            `json:"email"`
	Access model.AccessLevel `json:"access"`
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
	return c.userSchemas.ListBySchemaIDsOut(ctx, &repositoryuserschemas.ListBySchemaIDsOptions{
		SchemaIDs: schemaIDs,

		Scopes: []model.Scope{
			repositoryuserschemas.IncludeUsersScope,
		},

		Select: []string{
			"DISTINCT ON (user_id) user_id AS id",
			"access",
			"email",
			"first_name",
			"last_name",
		},

		Order: []string{
			"user_id",
			"user_schemas.created_at ASC",
		},
	}, &res.Users)
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
