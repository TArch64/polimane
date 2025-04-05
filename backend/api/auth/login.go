package auth

import (
	"context"
	"errors"

	"github.com/guregu/dynamo/v2"

	"polimane/backend/api/base"
	"polimane/backend/awsdynamodb"
	"polimane/backend/model"

	"github.com/gofiber/fiber/v2"
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
			return base.NewReasonedError(fiber.StatusForbidden, "invalid_credentials")
		}
		return err
	}

	return ctx.JSON(user)
}
