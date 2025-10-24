package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
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

	err = c.userSchemas.DeleteWithAccessCheck(ctx.Context(), &repositoryuserschemas.DeleteWithAccessCheckOptions{
		CurrentUser: currentUser,

		Operation: &repositoryuserschemas.DeleteOptions{
			UserID:   userID,
			SchemaID: schemaID,
		},
	})

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
