package base

import (
	"github.com/gofiber/fiber/v2"

	"polimane/backend/model"
)

type Controller interface {
	Private(group fiber.Router)
	Public(group fiber.Router)
}

type BulkOperationBody struct {
	IDs []model.ID `json:"ids" validate:"required"`
}
