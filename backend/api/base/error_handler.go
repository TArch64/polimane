package base

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    CustomErrorData `json:"data,omitempty"`
}

func getErrorStatus(err error) int {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.StatusNotFound
	}

	switch v := err.(type) {
	case *CustomError:
		return v.Code
	case *fiber.Error:
		return v.Code
	default:
		return fiber.StatusInternalServerError
	}
}

func getErrorData(err error) CustomErrorData {
	switch v := err.(type) {
	case *CustomError:
		return v.Data
	default:
		return nil
	}
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	return ctx.Status(getErrorStatus(err)).JSON(ErrorResponse{
		Success: false,
		Message: err.Error(),
		Data:    getErrorData(err),
	})
}
