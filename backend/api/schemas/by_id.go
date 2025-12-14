package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
	repositoryschemas "polimane/backend/repository/schemas"
)

type schemaDetails struct {
	model.Schema
	Access model.AccessLevel `json:"access"`
}

func (c *Controller) apiByID(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, schemaIDParam)
	if err != nil {
		return err
	}

	var schema schemaDetails
	user := auth.GetSessionUser(ctx)
	err = c.schemas.GetOut(ctx.Context(), &schema,
		repository.IDEq(schemaID),
		repository.Select("schemas.*", "user_schemas.access AS access"),
		repositoryschemas.IncludeUserSchemaScope(user.ID),
	)

	if err != nil {
		return err
	}

	return ctx.JSON(schema)
}
