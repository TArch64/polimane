package auth

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/guregu/dynamo/v2"

	"polimane/backend/api/base"
	"polimane/backend/argon"
	"polimane/backend/awsdynamodb"
	"polimane/backend/model"
)

var (
	invalidCredentialsErr = base.NewReasonedError(fiber.StatusForbidden, "InvalidCredentials")
)

type loginBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func fetchUserByName(ctx context.Context, username string) (*model.User, error) {
	var user model.User

	err := awsdynamodb.Table().
		Get("SK", model.NewKey(model.SKUser, username)).
		Index(model.IndexUserName).
		One(ctx, &user)

	return &user, err
}

func apiLogin(ctx *fiber.Ctx) error {
	var body loginBody
	err := base.ParseBody(ctx, &body)
	if err != nil {
		return err
	}

	user, err := fetchUserByName(ctx.Context(), body.Username)
	if err != nil {
		if errors.Is(err, dynamo.ErrNotFound) {
			return invalidCredentialsErr
		}
		return err
	}

	if !argon.Compare(body.Password, user.PasswordHash) {
		return invalidCredentialsErr
	}

	expiresAt := newTokenExpiresAt()
	token, err := newToken(user, expiresAt)

	ctx.Cookie(&fiber.Cookie{
		Name:     "pa",
		Value:    token,
		HTTPOnly: true,
		Expires:  expiresAt,
	})

	return ctx.JSON(user)
}
