//go:build !dev

package worker

import (
	"context"
	"strings"
	"sync"

	lambdaEvents "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/getsentry/sentry-go"
	"go.uber.org/fx"

	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

func (c *Controller) handleError(ctx context.Context, err error) {
	hub := sentry.GetHubFromContext(ctx)
	client, scope := hub.Client(), hub.Scope()

	client.CaptureException(
		err,
		&sentry.EventHint{Context: ctx},
		scope,
	)
}

type StartOptions struct {
	fx.In
	Ctx        context.Context
	Controller *Controller
}

type QueueSubscription struct {
	Queue queue.Interface
	Chan  chan *events.Message
}

func Start(options StartOptions) {
	subscriptions := make(map[string]*QueueSubscription)

	lambda.Start(func(ctx context.Context, sqsEvent lambdaEvents.SQSEvent) error {
		wg := sync.WaitGroup{}

		for _, message := range sqsEvent.Records {
			var exists bool
			var subscription *QueueSubscription

			if subscription, exists = subscriptions[message.EventSourceARN]; !exists {
				for _, q := range options.Controller.queues {
					if strings.HasSuffix(message.EventSourceARN, q.Name()) {
						subscription = &QueueSubscription{
							Queue: q,
							Chan:  make(chan *events.Message, 100),
						}
						subscriptions[message.EventSourceARN] = subscription
						go options.Controller.Process(options.Ctx, q, subscription.Chan)
						break
					}
				}

				if subscription == nil {
					continue
				}
			}

			wg.Add(1)

			subscription.Chan <- &events.Message{
				Body:          message.Body,
				ReceiptHandle: message.ReceiptHandle,
				EventType:     *message.MessageAttributes["EventType"].StringValue,
				OnEnd: func() {
					wg.Done()
				},
			}
		}

		wg.Wait()
		return nil
	})
}
