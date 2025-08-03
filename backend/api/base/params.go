package base

import (
	"fmt"

	"polimane/backend/model"

	"github.com/gofiber/fiber/v2"
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

func GetParamID(ctx *fiber.Ctx, name string) (model.ID, error) {
	value, err := GetRequiredParam(ctx, name)
	if err != nil {
		return model.ID{}, err
	}

	id, err := model.StringToID(value)
	if err != nil {
		return model.ID{}, newMissingParamErr(name)
	}

	return id, nil
}
