package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

func apiList(ctx *fiber.Ctx) error {
	schemas, err := repositoryschemas.ByUser(&repositoryschemas.ByUserOptions{
		Ctx:    ctx.Context(),
		User:   auth.GetSessionUser(ctx),
		Select: []string{"Name"},
	})

	if err != nil {
		return err
	}

	if schemas == nil {
		schemas = []*model.Schema{}
	}

	return ctx.JSON(schemas)
}
