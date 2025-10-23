package queuedebounced

import (
	"go.uber.org/fx"

	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

type Queue struct {
	*queue.Base
	schemas          *repositoryschemas.Client
	schemaScreenshot schemascreenshot.Interface
}

type ProviderOptions struct {
	fx.In
	Schemas          *repositoryschemas.Client
	SchemaScreenshot schemascreenshot.Interface
}

func Provider(options ProviderOptions) queue.Interface {
	q := &Queue{
		Base:             queue.NewBase(),
		schemas:          options.Schemas,
		schemaScreenshot: options.SchemaScreenshot,
	}

	q.HandleEvent(events.EventSchemaScreenshot, q.ProcessSchemaScreenshot)
	return q
}

func (q *Queue) Name() string {
	return events.QueueDebounced
}
