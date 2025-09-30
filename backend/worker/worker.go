package worker

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"go.uber.org/fx"

	"polimane/backend/services/awssqs"
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
	messages chan *types.Message,
) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10)

	for message := range messages {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(q queue.Interface, message *types.Message) {
			defer wg.Done()
			defer func() { <-semaphore }()

			if err := q.Process(ctx, message); err != nil {
				c.handleError(err)
				return
			}

			if err := c.sqs.Delete(ctx, q.Name(), *message.ReceiptHandle); err != nil {
				c.handleError(err)
			}
		}(q, message)
	}

	wg.Wait()
}
