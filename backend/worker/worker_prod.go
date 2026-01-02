//go:build !dev

package worker

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"strings"

	lambdaEvents "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/getsentry/sentry-go"
	"go.uber.org/fx"

	"polimane/backend/services/appcontext"
	"polimane/backend/services/awssqs"
	"polimane/backend/services/logstdout"
	"polimane/backend/worker/events"
	"polimane/backend/worker/queue"
)

var (
	queueNotFoundErr = errors.New("queue not found for message")
)

func (c *Controller) handleError(ctx context.Context, err error, attrs map[string]string) {
	hub := sentry.GetHubFromContext(ctx)
	client, scope := hub.Client(), hub.Scope()

	client.CaptureException(
		err,
		&sentry.EventHint{
			Context: ctx,
			Data:    attrs,
		},
		scope,
	)
}

type StartOptions struct {
	fx.In
	Ctx        *appcontext.Ctx
	Controller *Controller
	Stdout     *logstdout.Logger
}

func Start(options StartOptions) {
	lambda.Start(func(ctx context.Context, sqsEvent lambdaEvents.SQSEvent) error {
		for _, message := range sqsEvent.Records {
			q, err := getQueue(options, &message)
			if err != nil {
				options.Stdout.ErrorContext(options.Ctx, "error getting queue for message",
					slog.String("EventSourceARN", message.EventSourceARN),
				)
				continue
			}

			var body awssqs.QueueEvent
			err = json.Unmarshal([]byte(message.Body), &body)
			if err != nil {
				options.Stdout.ErrorContext(options.Ctx, "error unmarshaling message body",
					slog.String("Queue", q.Name()),
					slog.String("Error", err.Error()),
				)
				continue
			}

			options.Controller.Process(options.Ctx, q, &events.Message{
				Body:          string(body.Payload),
				ReceiptHandle: message.ReceiptHandle,
				EventType:     body.EventType,
			})
		}

		return nil
	})
}

func getQueue(options StartOptions, message *lambdaEvents.SQSMessage) (queue.Interface, error) {
	for _, q := range options.Controller.queues {
		if strings.HasSuffix(message.EventSourceARN, q.Name()) {
			return q, nil
		}
	}

	return nil, queueNotFoundErr
}
