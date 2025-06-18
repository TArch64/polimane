package schemas

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/guregu/dynamo/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	repositoryschemas "polimane/backend/repository/schemas"
)

func apiById(ctx *fiber.Ctx) error {
	schemaId, err := base.GetRequiredParam(ctx, "schemaId")
	if err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)

	schema, err := repositoryschemas.ById(ctx.Context(), user, schemaId)
	if errors.Is(err, dynamo.ErrNotFound) {
		return base.NotFoundErr
	}
	if err != nil {
		return err
	}

	return ctx.JSON(schema)
}
