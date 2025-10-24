package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

type updateAccessBody struct {
	Access model.AccessLevel `validate:"required,gte=1,lte=3" json:"access"`
}

func (c *Controller) apiUpdateAccess(ctx *fiber.Ctx) error {
	userID, err := base.GetParamID(ctx, userIDParam)
	if err != nil {
		return err
	}

	currentUser := auth.GetSessionUser(ctx)
	if currentUser.ID == userID {
		return base.InvalidRequestErr
	}

	schemaID, err := base.GetParamID(ctx, schemaIDParam)
	if err != nil {
		return err
	}

	var body updateAccessBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	requestCtx := ctx.Context()
	err = c.userSchemas.HasAccess(requestCtx, currentUser.ID, schemaID, model.AccessAdmin)
	if err != nil {
		return nil
	}

	err = c.userSchemas.Update(requestCtx, &repositoryuserschemas.UpdateOptions{
		UserID:   userID,
		SchemaID: schemaID,
		Updates:  &model.UserSchema{Access: body.Access},
	})

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
