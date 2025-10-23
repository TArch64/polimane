package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

type userListItem struct {
	ID        string            `json:"id"`
	Email     string            `json:"email"`
	FirstName string            `json:"firstName"`
	LastName  string            `json:"lastName"`
	Access    model.AccessLevel `json:"access"`
}

func (c *Controller) apiList(ctx *fiber.Ctx) error {
	schemaId, err := base.GetParamID(ctx, schemaIdParam)
	if err != nil {
		return err
	}

	var users []*userListItem
	err = c.userSchemas.ListBySchemaOut(ctx.Context(), &repositoryuserschemas.ListBySchemaOptions{
		SchemaID: schemaId,

		Select: []string{
			"user_id AS id",
			"access",
			"email",
			"first_name",
			"last_name",
		},
	}, &users)

	return ctx.JSON(users)
}
