package users

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

type addUserBody struct {
	Email string `validate:"required,email" json:"email"`
}

type addUserResponse struct {
	Invited bool          `json:"invited"`
	User    *userListItem `json:"user"`
}

func (c *Controller) apiAdd(ctx *fiber.Ctx) error {
	schemaId, err := base.GetParamID(ctx, schemaIdParam)
	if err != nil {
		return err
	}

	var body addUserBody
	if err = base.ParseBody(ctx, &body); err != nil {
		return err
	}

	requestCtx := ctx.Context()
	user, err := c.users.GeyByEmail(requestCtx, body.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.inviteUser(requestCtx, body.Email)
	}
	if err != nil {
		return err
	}

	currentUser := auth.GetSessionUser(ctx)
	if currentUser.ID == user.ID {
		return base.InvalidRequestErr
	}

	userSchema, err := c.userSchemas.CreateWithAccessCheck(requestCtx, &repositoryuserschemas.CreateWithAccessCheckOptions{
		CurrentUser: currentUser,

		Operation: &repositoryuserschemas.CreateOptions{
			UserID:   user.ID,
			SchemaID: schemaId,
			Access:   model.AccessRead,
		},
	})

	if err != nil {
		return err
	}

	return ctx.JSON(&addUserResponse{
		Invited: false,
		User:    newUserListItem(user, userSchema),
	})
}

func (c *Controller) inviteUser(ctx context.Context, email string) error {
	return nil
}
