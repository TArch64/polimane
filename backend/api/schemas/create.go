package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repositoryschemas"
)

type createBody struct {
	Name string `json:"name" validate:"required"`
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
		Content: model.SchemaContent{},
	})

	if err != nil {
		return err
	}

	return ctx.JSON(schema)
}
