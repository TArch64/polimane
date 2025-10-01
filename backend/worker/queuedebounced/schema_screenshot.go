package queuedebounced

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"

	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

func (q *Queue) ProcessSchemaScreenshot(ctx context.Context, message *types.Message) error {
	var body events.SchemaScreenshotBody
	if err := queue.ParseBody(message, &body); err != nil {
		return err
	}

	schema, err := q.schemas.ByID(&repositoryschemas.ByIDOptions{
		Ctx:      ctx,
		SchemaID: body.SchemaID,
	})

	if err != nil {
		return err
	}

	err = q.schemaScreenshot.Screenshot(ctx, &schemascreenshot.ScreenshotOptions{
		Schema: schema,
	})

	return nil
}
