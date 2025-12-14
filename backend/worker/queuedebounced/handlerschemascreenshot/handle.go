package handlerschemascreenshot

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"polimane/backend/repository"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

func (h *Handler) Handle(ctx context.Context, message *events.Message) error {
	var body events.SchemaScreenshotBody
	if err := queue.ParseBody(message, &body); err != nil {
		return err
	}

	schema, err := h.schemas.Get(ctx, repository.IDEq(body.SchemaID))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Exit on removed schema
		return nil
	}
	if err != nil {
		return err
	}

	return h.schemaScreenshot.Screenshot(ctx, &schemascreenshot.ScreenshotOptions{
		Schema: schema,
	})
}
