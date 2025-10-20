package queue

import (
	"context"
	"errors"

	"polimane/backend/worker/events"
)

func (b *Base) Process(ctx context.Context, message *events.Message) error {
	processor := b.EventHandlers[message.EventType]

	if processor == nil {
		return errors.New("no processor found for event type: " + message.EventType)
	}

	return processor(ctx, message)
}
