package queue

import (
	"context"

	"polimane/backend/worker/events"
)

type EventHandlers = map[string]EventProcessor
type EventProcessor = func(ctx context.Context, message *events.Message) error

func (b *Base) HandleEvent(group string, processor EventProcessor) {
	b.EventHandlers[group] = processor
}

func (b *Base) GetEventHandlers() EventHandlers {
	return b.EventHandlers
}
