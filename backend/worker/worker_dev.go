//go:build dev

package worker

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"text/tabwriter"
	"time"

	"polimane/backend/base"

	"go.uber.org/fx"

	"polimane/backend/services/awssqs"
	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

func (c *Controller) handleError(_ context.Context, err error) {
	log.Println(err)
}

type StartOptions struct {
	fx.In
	Ctx        context.Context
	SQS        *awssqs.Client
	Controller *Controller
}

func Start(options StartOptions) {
	log.Println("starting worker...")

	for _, q := range options.Controller.queues {
		go watchQueue(
			options.Ctx,
			q,
			options.Controller,
			options.SQS,
		)
	}

	go printStartupMessage(options.Controller)
}

func watchQueue(
	ctx context.Context,
	q queue.Interface,
	controller *Controller,
	client *awssqs.Client,
) {
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
			for _, message := range messages {
				actionType := *message.MessageAttributes["EventType"].StringValue

				messagesChan <- &events.Message{
					Body:          *message.Body,
					ReceiptHandle: *message.ReceiptHandle,
					EventType:     actionType,
					OnEnd:         func() {},
				}

				log.Printf("Processing %s action from %s queue\n", actionType, q.Name())
			}
		}
	}
}

func printStartupMessage(controller *Controller) {
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	printStartupRow(writer, "Queue", "Event", "Handler")
	printStartupRow(writer, "-----", "-----", "-------")

	for _, q := range controller.queues {
		queueName := q.Name()
		for event, handler := range q.GetEventHandlers() {
			handlerPointer := runtime.FuncForPC(reflect.ValueOf(handler).Pointer())
			printStartupRow(writer, queueName, event, handlerPointer.Name())
		}
	}

	err := writer.Flush()
	if err != nil {
		log.Println("error printing startup message:", err)
	}
}

func printStartupRow(writer *tabwriter.Writer, columns ...string) {
	_, _ = fmt.Fprintf(
		writer,
		"%s\t|\t%s\t|\t%s\n",
		base.Colored(columns[0], base.AnsiBlue),
		base.Colored(columns[1], base.AnsiGreen),
		base.Colored(columns[2], base.AnsiYellow),
	)
}
