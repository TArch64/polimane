//go:build dev

package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"runtime"
	"text/tabwriter"
	"time"

	"go.uber.org/fx"

	"polimane/backend/base"
	"polimane/backend/services/awssqs"
	"polimane/backend/services/logstdout"
	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

func (c *Controller) handleError(_ context.Context, err error, attrs map[string]string) {
	var args []any
	for name, value := range attrs {
		args = append(args, slog.String(name, value))
	}

	c.stdout.Error(err.Error(), args...)
}

type StartOptions struct {
	fx.In
	Ctx        context.Context
	SQS        *awssqs.Client
	Controller *Controller
	Stdout     *logstdout.Logger
}

func Start(options StartOptions) {
	options.Stdout.InfoContext(options.Ctx, "starting worker...")

	for _, q := range options.Controller.queues {
		go watchQueue(q, options)
	}

	go printStartupMessage(options.Controller)
}

func watchQueue(q queue.Interface, options StartOptions) {
	messagesChan := make(chan *events.Message, 100)
	go options.Controller.Process(options.Ctx, q, messagesChan)

	for {
		select {
		case <-options.Ctx.Done():
			close(messagesChan)
			options.Stdout.InfoContext(options.Ctx, "stopping queue watcher",
				slog.String("queue", q.Name()),
			)
			return
		default:
		}

		time.Sleep(1 * time.Second)
		messages, err := options.SQS.Receive(options.Ctx, q.Name())

		if err != nil {
			options.Stdout.ErrorContext(options.Ctx, "error receiving message",
				slog.String("queue", q.Name()),
				slog.String("err", err.Error()),
			)
			continue
		}

		if len(messages) > 0 {
			for _, message := range messages {
				var body awssqs.QueueEvent
				err = json.Unmarshal([]byte(*message.Body), &body)
				if err != nil {
					options.Stdout.ErrorContext(options.Ctx, "error unmarshaling message body",
						slog.String("queue", q.Name()),
						slog.String("err", err.Error()),
					)
					continue
				}

				messagesChan <- &events.Message{
					Body:          string(body.Payload),
					ReceiptHandle: *message.ReceiptHandle,
					EventType:     body.EventType,
					OnEnd:         func() {},
				}

				options.Stdout.InfoContext(options.Ctx, "Processing actions",
					slog.String("queue", q.Name()),
					slog.String("event_type", body.EventType),
					slog.Int("count", len(messages)),
				)
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
		controller.stdout.Error("error printing startup message", slog.String("err", err.Error()))
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
