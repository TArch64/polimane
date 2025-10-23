package schemas

import (
	"polimane/backend/model"

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
		Select:   append(model.GetColumns(model.Schema{}), "user_schemas.access AS access"),
	})

	if err != nil {
		return err
	}

	return ctx.JSON(schema)
}
