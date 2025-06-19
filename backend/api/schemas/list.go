package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

func apiList(ctx *fiber.Ctx) error {
	user := auth.GetSessionUser(ctx)

	schemas, err := repositoryschemas.ByUser(ctx.Context(), user, []string{"Name"})
	if err != nil {
		return err
	}

	if schemas == nil {
		schemas = []*model.Schema{}
	}

	return ctx.JSON(schemas)
}
