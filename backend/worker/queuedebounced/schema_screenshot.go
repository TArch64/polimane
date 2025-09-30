package queuedebounced

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"

	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

func (q *Queue) ProcessSchemaScreenshot(_ context.Context, message *types.Message) error {
	var body events.SchemaScreenshotBody
	if err := queue.ParseBody(message, &body); err != nil {
		return err
	}

	j, _ := json.MarshalIndent(body, "", "  ")
	println("DEBOUNCED QUEUE MESSAGE", string(j))
	return nil
}
