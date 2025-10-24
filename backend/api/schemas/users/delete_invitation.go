package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
)

type deleteInvitationQuery struct {
	Email string `query:"email" validate:"required,email,max=255"`
}

func (c *Controller) apiDeleteInvitation(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, schemaIDParam)
	if err != nil {
		return err
	}

	var query deleteInvitationQuery
	if err = base.ParseQuery(ctx, &query); err != nil {
		return err
	}

	currentUser := auth.GetSessionUser(ctx)
	requestCtx := ctx.Context()
	err = c.userSchemas.HasAccess(requestCtx, currentUser.ID, schemaID, model.AccessAdmin)
	if err != nil {
		return nil
	}

	err = c.schemaInvitations.Delete(requestCtx, &repositoryschemainvitations.DeleteOptions{
		Email:    query.Email,
		SchemaID: schemaID,
	})

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
