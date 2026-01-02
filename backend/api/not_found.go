package api

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"

	"polimane/backend/services/logstdout"
)

func NotFound(ctx *fiber.Ctx, stdout *logstdout.Logger) error {
	stdout.InfoContext(ctx.Context(), "unhandled route",
		slog.String("method", ctx.Method()),
		slog.String("path", ctx.Path()),
	)

	return ctx.
		Status(404).
		JSON(fiber.Map{"error": "Not Found"})
}
