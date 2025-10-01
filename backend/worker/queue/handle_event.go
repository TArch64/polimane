package queue

import (
	"context"

	"polimane/backend/worker/events"
)

type EventProcessor = func(ctx context.Context, message *events.Message) error

func (b *Base) HandleEvent(group string, processor EventProcessor) {
	if b.events == nil {
		b.events = make(map[string]EventProcessor)
	}

	b.events[group] = processor
}
