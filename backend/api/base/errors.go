package base

import (
	"github.com/gofiber/fiber/v2"
)

var (
	SchemasCreatedLimitReachedErr = NewReasonedError(fiber.StatusBadRequest, "LimitReached[SchemasCreated]")
)
