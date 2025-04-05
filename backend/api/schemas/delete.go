package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/repositoryschemas"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
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

	return base.NewSuccessResponse().AsJSON(ctx)
}
