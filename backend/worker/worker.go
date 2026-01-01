package worker

import (
	"context"

	"go.uber.org/fx"

	"polimane/backend/services/awssqs"
	"polimane/backend/services/logstdout"
	"polimane/backend/services/sentry"
	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

type Controller struct {
	queues []queue.Interface
	sqs    *awssqs.Client
	stdout *logstdout.Logger
}

type ProviderOptions struct {
	fx.In
	Queues []queue.Interface `group:"queues"`
	SQS    *awssqs.Client
	Sentry *sentry.Container
	Stdout *logstdout.Logger
}

func Provider(options ProviderOptions) *Controller {
	return &Controller{
		queues: options.Queues,
		sqs:    options.SQS,
		stdout: options.Stdout,
	}
}

func (c *Controller) Process(
	ctx context.Context,
	q queue.Interface,
	message *events.Message,
) {
	if err := q.Process(ctx, message); err != nil {
		c.handleError(ctx, err, map[string]string{
			"Queue": q.Name(),
		})
		return
	}

	if err := c.sqs.Delete(ctx, q.Name(), message.ReceiptHandle); err != nil {
		c.handleError(ctx, err, map[string]string{
			"Queue":     q.Name(),
			"EventType": message.EventType,
		})
	}
}
