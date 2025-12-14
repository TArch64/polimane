package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

func (c *Controller) Copy(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, ParamSchemaID)
	if err != nil {
		return err
	}

	schema, err := c.schemas.Copy(ctx.Context(), &repositoryschemas.CopyOptions{
		User:     auth.GetSessionUser(ctx),
		SchemaID: schemaID,
	})

	if err != nil {
		return err
	}

	if err = c.updateScreenshot(ctx.Context(), schema.ID, false); err != nil {
		return err
	}

	return ctx.JSON(NewListSchema(schema, model.AccessAdmin))
}
