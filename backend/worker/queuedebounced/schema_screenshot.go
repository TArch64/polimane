package queuedebounced

import (
	"context"

	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

func (q *Queue) ProcessSchemaScreenshot(ctx context.Context, message *events.Message) error {
	var body events.SchemaScreenshotBody
	if err := queue.ParseBody(message, &body); err != nil {
		return err
	}

	schema, err := q.schemas.ByID(ctx, &repositoryschemas.ByIDOptions{
		SchemaID: body.SchemaID,
	})

	if err != nil {
		return err
	}

	return q.schemaScreenshot.Screenshot(ctx, &schemascreenshot.ScreenshotOptions{
		Schema: schema,
	})
}
