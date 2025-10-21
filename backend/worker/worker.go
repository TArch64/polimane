package worker

import (
	"context"
	"sync"

	"go.uber.org/fx"

	"polimane/backend/services/awssqs"
	"polimane/backend/services/sentry"
	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

type Controller struct {
	queues []queue.Interface
	sqs    awssqs.Client
}

type ProviderOptions struct {
	fx.In
	Queues []queue.Interface `group:"queues"`
	SQS    awssqs.Client
	Sentry *sentry.Container
}

func Provider(options ProviderOptions) *Controller {
	return &Controller{
		queues: options.Queues,
		sqs:    options.SQS,
	}
}

func (c *Controller) Process(
	ctx context.Context,
	q queue.Interface,
	messages chan *events.Message,
) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10)

	for message := range messages {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(q queue.Interface, message *events.Message) {
			defer wg.Done()
			defer func() { <-semaphore }()

			if err := q.Process(ctx, message); err != nil {
				c.handleError(ctx, err)
				message.OnEnd()
				return
			}

			if err := c.sqs.Delete(ctx, q.Name(), message.ReceiptHandle); err != nil {
				c.handleError(ctx, err)
			}

			message.OnEnd()
		}(q, message)
	}

	wg.Wait()
}
