package schemas

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/guregu/dynamo/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	repositoryschemas "polimane/backend/repository/schemas"
)

func apiDelete(ctx *fiber.Ctx) error {
	schemaId, err := base.GetRequiredParam(ctx, "schemaId")
	if err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)

	err = repositoryschemas.Delete(ctx.Context(), user, schemaId)
	if errors.Is(err, dynamo.ErrNotFound) {
		return base.NotFoundErr
	}
	if err != nil {
		return err
	}

	return base.NewSuccessResponse(ctx)
}
