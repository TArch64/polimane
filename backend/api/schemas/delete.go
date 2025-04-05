package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/repositoryschemas"
)

func apiDelete(ctx *fiber.Ctx) error {
	schemaId, err := base.GetRequiredParam(ctx, "schemaId")
	if err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)

	err = repositoryschemas.Delete(ctx.Context(), user, schemaId)
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
