package base

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/model/modelbase"
)

func newMissingParamErr(name string) error {
	return NewReasonedError(fiber.StatusBadRequest, fmt.Sprintf("MissingRequiredParam[%s]", name))
}

func GetRequiredParam(ctx *fiber.Ctx, name string) (string, error) {
	value := ctx.Params(name)
	if value == "" {
		return "", newMissingParamErr(name)
	}
	return value, nil
}

func GetParamID(ctx *fiber.Ctx, name string) (modelbase.ID, error) {
	value, err := GetRequiredParam(ctx, name)
	if err != nil {
		return modelbase.ID{}, err
	}

	id, err := modelbase.StringToID(value)
	if err != nil {
		return modelbase.ID{}, newMissingParamErr(name)
	}

	return id, nil
}
