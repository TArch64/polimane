package base

import (
	"github.com/gofiber/fiber/v2"
)

func ParseBody[B any](ctx *fiber.Ctx, body *B) (err error) {
	if err = ctx.BodyParser(body); err != nil {
		return err
	}
	if err = Validate(body); err != nil {
		return err
	}
	return nil
}
