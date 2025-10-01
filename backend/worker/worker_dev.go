//go:build dev

package worker

import (
	"context"
	"log"
	"time"

	"go.uber.org/fx"

	"polimane/backend/services/awssqs"
	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

func (c *Controller) handleError(err error) {
	log.Println(err)
}

type StartOptions struct {
	fx.In
	SQS        awssqs.Client
	Controller *Controller
}

func Start(options StartOptions) {
	ctx := context.Background()

	log.Println("starting worker...")

	for _, q := range options.Controller.queues {
		go watchQueue(
			ctx,
			q,
			options.Controller,
			options.SQS,
		)
	}
}

func watchQueue(
	ctx context.Context,
	q queue.Interface,
	controller *Controller,
	client awssqs.Client,
) {
	log.Println("registered queue:", q.Name())

	messagesChan := make(chan *events.Message, 100)
	go controller.Process(ctx, q, messagesChan)

	for {
		select {
		case <-ctx.Done():
			close(messagesChan)
			log.Println("context cancelled, stopping queue watcher:", q.Name())
			return
		default:
		}

		time.Sleep(1 * time.Second)
		messages, err := client.Receive(ctx, q.Name())

		if err != nil {
			log.Println("error receiving message:", err)
			continue
		}

		if len(messages) > 0 {
			log.Printf("received %d messages from %s\n", len(messages), q.Name())

			for _, message := range messages {
				messagesChan <- &events.Message{
					Body:          *message.Body,
					ReceiptHandle: *message.ReceiptHandle,
					EventType:     *message.MessageAttributes["EventType"].StringValue,
					OnEnd:         func() {},
				}
			}
		}
	}
}
