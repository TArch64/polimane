package base

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetRequiredParam(ctx *fiber.Ctx, name string) (string, error) {
	value := ctx.Params(name)
	if len(value) == 0 {
		return "", NewReasonedError(fiber.StatusBadRequest, fmt.Sprintf("MissingRequiredParam[%s]", name))
	}
	return value, nil
}
