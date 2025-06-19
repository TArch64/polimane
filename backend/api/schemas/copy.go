package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	repositoryschemas "polimane/backend/repository/schemas"
)

func apiCopy(ctx *fiber.Ctx) error {
	schemaId, err := base.GetRequiredParam(ctx, "schemaId")
	if err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)

	schema, err := repositoryschemas.Copy(ctx.Context(), &repositoryschemas.CopyOptions{
		User:     user,
		SchemaID: schemaId,
	})

	if err != nil {
		return err
	}

	schema.Content = nil
	return ctx.JSON(schema)
}
