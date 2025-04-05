package schemas

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/guregu/dynamo/v2"

	"polimane/backend/repositoryschemas"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
)

func apiById(ctx *fiber.Ctx) error {
	schemaId, err := base.GetRequiredParam(ctx, "schemaId")
	if err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)

	schema, err := repositoryschemas.ById(ctx.Context(), user, schemaId)
	if err != nil {
		if errors.Is(err, dynamo.ErrNotFound) {
			return base.NotFoundErr
		}
		return err
	}

	return ctx.JSON(schema)
}
