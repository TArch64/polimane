package base

import (
	"maps"

	"github.com/gofiber/fiber/v2"
)

type CustomErrorData map[string]interface{}

type CustomError struct {
	Code    int
	Message string
	Data    CustomErrorData
}

var _ error = (*CustomError)(nil)

func NewCustomError(code int, message string, data CustomErrorData) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) AddCustomData(extra ...CustomErrorData) *CustomError {
	err := new(CustomError)
	*err = *e

	err.Data = make(CustomErrorData)
	maps.Copy(err.Data, e.Data)
	for _, data := range extra {
		maps.Copy(err.Data, data)
	}

	return err
}

func NewReasonedError(code int, reason string) *CustomError {
	return NewCustomError(code, reason, CustomErrorData{"reason": reason})
}

var (
	NotFoundErr       = NewReasonedError(fiber.StatusNotFound, "NotFound")
	InvalidRequestErr = NewReasonedError(fiber.StatusBadRequest, "InvalidRequest")
)
