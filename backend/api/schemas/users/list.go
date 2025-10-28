package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"

	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

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

type listInvitation struct {
	Email  string            `json:"email"`
	Access model.AccessLevel `json:"access"`
}

func newUserListItem(user *model.User, schemaUser *model.UserSchema) *listUser {
	return &listUser{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Access:    schemaUser.Access,
	}
}

func (c *Controller) apiList(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, schemaIDParam)
	if err != nil {
		return err
	}

	var response listResponse
	eg, egCtx := errgroup.WithContext(ctx.Context())

	eg.Go(func() error {
		return c.listUsers(egCtx, schemaID, &response)
	})

	eg.Go(func() error {
		return c.listInvitations(egCtx, schemaID, &response)
	})

	if err = eg.Wait(); err != nil {
		return err
	}

	return ctx.JSON(response)
}

func (c *Controller) listUsers(ctx context.Context, schemaID model.ID, res *listResponse) error {
	return c.userSchemas.ListBySchemaIDOut(ctx, &repositoryuserschemas.ListBySchemaIDOptions{
		SchemaID: schemaID,

		Select: []string{
			"user_id AS id",
			"access",
			"email",
			"first_name",
			"last_name",
		},
	}, &res.Users)
}

func (c *Controller) listInvitations(ctx context.Context, schemaID model.ID, res *listResponse) error {
	return c.schemaInvitations.ListBySchemaIDOut(ctx, &repositoryschemainvitations.ListBySchemaIDOptions{
		SchemaID: schemaID,
		Select:   []string{"email", "access"},
	}, &res.Invitations)
}
