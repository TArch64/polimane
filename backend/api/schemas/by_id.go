package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	repositoryschemas "polimane/backend/repository/schemas"
)

func (c *Controller) apiById(ctx *fiber.Ctx) error {
	schemaId, err := base.GetParamID(ctx, schemaIdParam)
	if err != nil {
		return err
	}

	schema, err := c.schemas.ByID(ctx.Context(), &repositoryschemas.ByIDOptions{
		User:     auth.GetSessionUser(ctx),
		SchemaID: schemaId,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(schema)
}
