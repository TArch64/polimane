package users

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
)

type UpdateInvitationAccessBody struct {
	InvitationBody
	Access model.AccessLevel `validate:"required,gte=1,lte=3" json:"access"`
}

func (c *Controller) UpdateInvitationAccess(ctx *fiber.Ctx) (err error) {
	var body UpdateInvitationAccessBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	if err = c.filterSchemaIDsByAccess(ctx, &body.IDs); err != nil {
		return err
	}

	reqCtx := ctx.Context()
	hasInvitations, err := c.schemaInvitations.Exists(reqCtx, repository.EmailEq(body.Email))
	if err != nil {
		return err
	}

	if hasInvitations {
		err = c.schemaInvitations.UpsertMany(reqCtx, &repositoryschemainvitations.UpsertManyOptions{
			Email:     body.Email,
			SchemaIDs: body.IDs,
			Updates:   &model.SchemaInvitation{Access: body.Access},
		})
	} else {
		err = c.updateAlreadyAcceptedUser(reqCtx, &body)
	}

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}

func (c *Controller) updateAlreadyAcceptedUser(ctx context.Context, body *UpdateInvitationAccessBody) error {
	user, err := c.users.Get(
		ctx,
		repository.Select("id"),
		repository.EmailEq(body.Email),
	)
	if err != nil {
		return err
	}

	return c.updateUserAccess(ctx, user.ID, body.IDs, body.Access)
}
