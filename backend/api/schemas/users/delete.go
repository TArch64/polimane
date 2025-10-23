package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

var (
	InvalidRequestErr = base.NewReasonedError(fiber.StatusBadRequest, "InvalidRequest")
)

func (c *Controller) apiDelete(ctx *fiber.Ctx) error {
	userId, err := base.GetParamID(ctx, userIdParam)
	if err != nil {
		return err
	}

	currentUser := auth.GetSessionUser(ctx)
	if currentUser.ID == userId {
		return InvalidRequestErr
	}

	schemaId, err := base.GetParamID(ctx, schemaIdParam)
	if err != nil {
		return err
	}

	err = c.userSchemas.DeleteWithAccessCheck(ctx.Context(), &repositoryuserschemas.DeleteWithAccessCheckOptions{
		User:     currentUser,
		UserID:   userId,
		SchemaID: schemaId,
	})

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
