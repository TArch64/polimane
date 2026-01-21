package schemas

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	"polimane/backend/repository"
	repositoryschemas "polimane/backend/repository/schemas"
)

type SchemaDetails struct {
	model.Schema
	Access model.AccessLevel `json:"access"`
}

func (c *Controller) ByID(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, ParamSchemaID)
	if err != nil {
		return err
	}

	var schema SchemaDetails
	user := auth.GetSessionUser(ctx)
	err = c.schemas.GetOut(ctx.Context(), &schema,
		repository.IDEq(schemaID),
		repository.Select("schemas.*", "access"),
		repositoryschemas.IncludeUserSchemaScope(user.ID),
	)

	if err != nil {
		return err
	}

	return ctx.JSON(schema)
}
