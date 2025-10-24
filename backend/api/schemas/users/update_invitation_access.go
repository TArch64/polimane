package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
)

type updateInvitationAccess struct {
	Email  string            `validate:"required,email,max=255" json:"email"`
	Access model.AccessLevel `validate:"required,gte=1,lte=3" json:"access"`
}

func (c *Controller) apiUpdateInvitationAccess(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, schemaIDParam)
	if err != nil {
		return err
	}

	var body updateInvitationAccess
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	currentUser := auth.GetSessionUser(ctx)
	requestCtx := ctx.Context()
	err = c.userSchemas.HasAccess(requestCtx, currentUser.ID, schemaID, model.AccessAdmin)
	if err != nil {
		return nil
	}

	err = c.schemaInvitations.Update(requestCtx, &repositoryschemainvitations.UpdateOptions{
		Email:    body.Email,
		SchemaID: schemaID,
		Updates:  &model.SchemaInvitation{Access: body.Access},
	})

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
