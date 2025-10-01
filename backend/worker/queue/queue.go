package queue

import (
	"context"

	"polimane/backend/worker/events"
)

type Interface interface {
	Name() string
	Process(ctx context.Context, message *events.Message) error
}

type Base struct {
	events map[string]EventProcessor
}

func NewBase() *Base {
	return &Base{
		events: make(map[string]EventProcessor),
	}
}
