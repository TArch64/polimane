package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"polimane/backend/api/base"
	repositoryusers "polimane/backend/repository/users"
	"polimane/backend/services/argon"
)

var invalidCredentialsErr = base.NewReasonedError(fiber.StatusForbidden, "InvalidCredentials")

type loginBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func apiLogin(ctx *fiber.Ctx) error {
	var body loginBody
	err := base.ParseBody(ctx, &body)
	if err != nil {
		return err
	}

	user, err := repositoryusers.ByName(ctx.Context(), body.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return invalidCredentialsErr
	}
	if err != nil {
		return err
	}

	if !argon.Compare(body.Password, user.PasswordHash) {
		return invalidCredentialsErr
	}

	expiresAt := newTokenExpiresAt()
	token, err := newAuthToken(user, expiresAt)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}
