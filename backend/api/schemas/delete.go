package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	repositoryschemas "polimane/backend/repository/schemas"
)

func (c *Controller) apiDelete(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, schemaIDParam)
	if err != nil {
		return err
	}

	err = c.schemas.Delete(ctx.Context(), &repositoryschemas.DeleteOptions{
		User:     auth.GetSessionUser(ctx),
		SchemaID: schemaID,
	})

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
