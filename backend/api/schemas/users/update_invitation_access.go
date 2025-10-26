package users

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

type updateInvitationAccessBody struct {
	Email  string            `validate:"required,email,max=255" json:"email"`
	Access model.AccessLevel `validate:"required,gte=1,lte=3" json:"access"`
}

func (c *Controller) apiUpdateInvitationAccess(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, schemaIDParam)
	if err != nil {
		return err
	}

	var body updateInvitationAccessBody
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

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = c.updateAlreadyAcceptedUser(requestCtx, schemaID, &body)
	}
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}

func (c *Controller) updateAlreadyAcceptedUser(ctx context.Context, schemaID model.ID, body *updateInvitationAccessBody) error {
	user, err := c.users.GeyByEmail(ctx, body.Email)
	if err != nil {
		return err
	}

	return c.userSchemas.Update(ctx, &repositoryuserschemas.UpdateOptions{
		UserID:   user.ID,
		SchemaID: schemaID,
		Updates:  &model.UserSchema{Access: body.Access},
	})
}
