package users

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	repositoryusers "polimane/backend/repository/users"
)

type updateInvitationAccessBody struct {
	bulkOperationBody
	Email  string            `validate:"required,email,max=255" json:"email"`
	Access model.AccessLevel `validate:"required,gte=1,lte=3" json:"access"`
}

func (c *Controller) apiUpdateInvitationAccess(ctx *fiber.Ctx) error {
	var err error
	var body updateInvitationAccessBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	currentUser := auth.GetSessionUser(ctx)
	requestCtx := ctx.Context()
	err = c.userSchemas.FilterByAccess(requestCtx, currentUser, &body.IDs, model.AccessAdmin)
	if err != nil {
		return nil
	}

	hasInvitations, err := c.schemaInvitations.Exists(requestCtx, repository.EqEmail(body.Email))
	if err != nil {
		return err
	}

	if hasInvitations {
		err = c.schemaInvitations.UpsertMany(requestCtx, &repositoryschemainvitations.UpsertManyOptions{
			Email:     body.Email,
			SchemaIDs: body.IDs,
			Updates:   &model.SchemaInvitation{Access: body.Access},
		})
	} else {
		err = c.updateAlreadyAcceptedUser(requestCtx, &body)
	}

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}

func (c *Controller) updateAlreadyAcceptedUser(ctx context.Context, body *updateInvitationAccessBody) error {
	user, err := c.users.GeyByEmail(ctx, &repositoryusers.GetByEmailOptions{
		Email: body.Email,
	})

	if err != nil {
		return err
	}

	return c.updateUserAccess(ctx, user.ID, body.IDs, body.Access)
}
