package users

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/awsdynamodb"
	"polimane/backend/model"
)

func apiList(ctx *fiber.Ctx) error {
	var users []model.User

	err := awsdynamodb.
		Table().
		Scan().
		Filter(model.TypeFilter(model.PKUser)).
		All(ctx.Context(), &users)

	if err != nil {
		return err
	}

	return ctx.JSON(users)
}
