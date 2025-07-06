package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	repositoryschemas "polimane/backend/repository/schemas"
)

func apiDelete(ctx *fiber.Ctx) error {
	schemaId, err := base.GetParamID(ctx, "schemaId")
	if err != nil {
		return err
	}

	err = repositoryschemas.Delete(&repositoryschemas.DeleteOptions{
		Ctx:      ctx.Context(),
		User:     auth.GetSessionUser(ctx),
		SchemaID: schemaId,
	})

	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
