package base

import (
	"github.com/gofiber/fiber/v2"
)

func ParseQuery[B any](ctx *fiber.Ctx, query *B) (err error) {
	if err = ctx.QueryParser(query); err != nil {
		return err
	}
	if err = Validate(query); err != nil {
		return err
	}
	return nil
}
