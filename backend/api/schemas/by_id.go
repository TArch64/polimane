package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	repositoryschemas "polimane/backend/repository/schemas"
)

func apiById(ctx *fiber.Ctx) error {
	schemaId, err := base.GetParamID(ctx, "schemaId")
	if err != nil {
		return err
	}

	schema, err := repositoryschemas.ByID(&repositoryschemas.ByIDOptions{
		Ctx:      ctx.Context(),
		User:     auth.GetSessionUser(ctx),
		SchemaID: schemaId,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(schema)
}
