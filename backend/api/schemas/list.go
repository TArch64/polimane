package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/model"
	"polimane/backend/repositoryschemas"
)

func apiList(ctx *fiber.Ctx) error {
	user := auth.GetSessionUser(ctx)

	schemas, err := repositoryschemas.ByUser(ctx.Context(), user)
	if err != nil {
		return err
	}

	if schemas == nil {
		schemas = []*model.Schema{}
	}

	return ctx.JSON(schemas)
}
