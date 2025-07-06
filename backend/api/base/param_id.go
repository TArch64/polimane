package base

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/model/modelbase"
)

func GetParamID(ctx *fiber.Ctx, name string) (modelbase.ID, error) {
	value, err := ctx.ParamsInt(name)

	if err != nil {
		return 0, NewReasonedError(fiber.StatusBadRequest, fmt.Sprintf("MissingRequiredParam[%s]", name))
	}

	return modelbase.ID(value), nil
}
