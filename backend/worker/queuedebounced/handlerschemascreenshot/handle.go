package handlerschemascreenshot

import (
	"context"
	"errors"

	"gorm.io/gorm"

	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

func (h *Handler) Handle(ctx context.Context, message *events.Message) error {
	var body events.SchemaScreenshotBody
	if err := queue.ParseBody(message, &body); err != nil {
		return err
	}

	schema, err := h.schemas.GetByID(ctx, &repositoryschemas.ByIDOptions{
		SchemaID: body.SchemaID,
	})

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
