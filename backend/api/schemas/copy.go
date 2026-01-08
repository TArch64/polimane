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

	reqCtx := ctx.Context()
	user := auth.GetSessionUser(ctx)
	schema, err := c.schemas.Copy(reqCtx, &repositoryschemas.CopyOptions{
		User:     user,
		SchemaID: schemaID,
	})

	if err != nil {
		return err
	}

	err = c.subscriptionCounters.SchemasCreated.Add(reqCtx, 1, user.ID)
	if err != nil {
		return err
	}

	if err = c.updateScreenshot(reqCtx, schema.ID, false); err != nil {
		return err
	}

	return ctx.JSON(NewListSchema(schema, model.AccessAdmin))
}
