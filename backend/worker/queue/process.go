package queue

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func (b *Base) Process(ctx context.Context, message *types.Message) error {
	eventType := message.MessageAttributes["EventType"].StringValue
	processor := b.events[*eventType]

	if processor == nil {
		return errors.New("no processor found for event type: " + *eventType)
	}

	return processor(ctx, message)
}
