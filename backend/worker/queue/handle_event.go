package queue

import (
	"context"

	"polimane/backend/worker/events"
)

type EventProcessor = func(ctx context.Context, message *events.Message) error

func (b *Base) HandleEvent(group string, processor EventProcessor) {
	b.events[group] = processor
}
