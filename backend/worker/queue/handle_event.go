package queue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type EventProcessor = func(ctx context.Context, message *types.Message) error

func (b *Base) HandleEvent(group string, processor EventProcessor) {
	if b.events == nil {
		b.events = make(map[string]EventProcessor)
	}

	b.events[group] = processor
}
