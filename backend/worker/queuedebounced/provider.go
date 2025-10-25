package queuedebounced

import (
	"go.uber.org/fx"

	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
	"polimane/backend/worker/queuedebounced/handlerschemascreenshot"
)

type Queue struct {
	*queue.Base
}

type ProviderOptions struct {
	fx.In
	SchemaScreenshotHandler *handlerschemascreenshot.Handler
}

func Provider(options ProviderOptions) queue.Interface {
	q := &Queue{Base: queue.NewBase()}
	q.HandleEvent(events.EventSchemaScreenshot, options.SchemaScreenshotHandler.Handle)
	return q
}

func (q *Queue) Name() string {
	return events.QueueDebounced
}
