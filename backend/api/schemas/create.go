package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

type createBody struct {
	Name    string              `json:"name" validate:"required"`
	Content model.SchemaContent `json:"content"`
}

func apiCreate(ctx *fiber.Ctx) error {
	var body createBody
	err := base.ParseBody(ctx, &body)
	if err != nil {
		return err
	}

	user := auth.GetSessionUser(ctx)

	schema, err := repositoryschemas.Create(ctx.Context(), &repositoryschemas.CreateOptions{
		User:    user,
		Name:    body.Name,
		Content: body.Content,
	})

	if err != nil {
		return err
	}

	schema.Content = nil
	return ctx.JSON(schema)
}
