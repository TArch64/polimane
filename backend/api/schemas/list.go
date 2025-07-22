package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

func (c *Controller) apiList(ctx *fiber.Ctx) error {
	schemas, err := c.schemas.ByUser(&repositoryschemas.ByUserOptions{
		Ctx:    ctx.Context(),
		User:   auth.GetSessionUser(ctx),
		Select: []string{"id", "name"},
	})

	if err != nil {
		return err
	}

	if schemas == nil {
		schemas = []*model.Schema{}
	}

	return ctx.JSON(schemas)
}
