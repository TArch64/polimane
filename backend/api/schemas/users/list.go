package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

type listUser struct {
	ID        model.ID          `json:"id"`
	Email     string            `json:"email"`
	FirstName string            `json:"firstName"`
	LastName  string            `json:"lastName"`
	Access    model.AccessLevel `json:"access"`
}

type listInvitation struct {
	Email  string            `json:"email"`
	Access model.AccessLevel `json:"access"`
}

func newUserListItem(user *model.User, schemaUser *model.UserSchema) *listUser {
	return &listUser{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Access:    schemaUser.Access,
	}
}

func (c *Controller) apiList(ctx *fiber.Ctx) error {
	schemaID, err := base.GetParamID(ctx, schemaIDParam)
	if err != nil {
		return err
	}

	var users []*listUser
	err = c.userSchemas.ListBySchemaOut(ctx.Context(), &repositoryuserschemas.ListBySchemaOptions{
		SchemaID: schemaID,

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
