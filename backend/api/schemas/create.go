package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

type createBody struct {
	Name string `json:"name" validate:"required"`
}

func (c *Controller) apiCreate(ctx *fiber.Ctx) error {
	var body createBody
	err := base.ParseBody(ctx, &body)
	if err != nil {
		return err
	}

	schema, err := c.schemas.Create(ctx.Context(), &repositoryschemas.CreateOptions{
		User: auth.GetSessionUser(ctx),
		Name: body.Name,
	})

	if err != nil {
		return err
	}

	if err = c.updateScreenshot(ctx.Context(), schema.ID, false); err != nil {
		return err
	}

	return ctx.JSON(newListItem(&model.SchemaWithAccess{
		Schema: *schema,
		Access: model.AccessAdmin,
	}))
}
