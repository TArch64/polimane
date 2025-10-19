package queue

import (
	"context"

	"polimane/backend/worker/events"
)

type Interface interface {
	Name() string
	Process(ctx context.Context, message *events.Message) error
	GetEventHandlers() EventHandlers
}

type Base struct {
	EventHandlers EventHandlers
}

func NewBase() *Base {
	return &Base{
		EventHandlers: make(EventHandlers),
	}
}
