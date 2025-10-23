package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

func (c *Controller) apiById(ctx *fiber.Ctx) error {
	schemaId, err := base.GetParamID(ctx, schemaIdParam)
	if err != nil {
		return err
	}

	var schema model.SchemaWithAccess
	err = c.schemas.GetOutByID(ctx.Context(), &repositoryschemas.ByIDOptions{
		User:     auth.GetSessionUser(ctx),
		SchemaID: schemaId,
		Select:   []string{"schemas.*", "user_schemas.access AS access"},
	}, &schema)

	if err != nil {
		return err
	}

	return ctx.JSON(schema)
}
