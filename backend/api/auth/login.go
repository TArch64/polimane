package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/api/base"
)

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

	log.Println(body)

	return nil
}
