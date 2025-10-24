package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

func (c *Controller) apiDelete(ctx *fiber.Ctx) error {
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

	requestCtx := ctx.Context()
	err = c.userSchemas.HasAccess(requestCtx, currentUser.ID, schemaID, model.AccessAdmin)
	if err != nil {
		return nil
	}

	err = c.userSchemas.Delete(requestCtx, &repositoryuserschemas.DeleteOptions{
		UserID:   userID,
		SchemaID: schemaID,
	})

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
