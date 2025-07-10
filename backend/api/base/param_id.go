package base

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/model/modelbase"
)

func GetParamID(ctx *fiber.Ctx, name string) (modelbase.ID, error) {
	id, err := modelbase.StringToID(ctx.Params(name))
	if err != nil {
		return modelbase.ID{}, NewReasonedError(fiber.StatusBadRequest, fmt.Sprintf("MissingRequiredParam[%s]", name))
	}

	return id, nil
}
