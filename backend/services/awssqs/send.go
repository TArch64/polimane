package awssqs

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SendOptions struct {
	Queue           string
	Event           string
	DeduplicationID string
	Body            interface{}
}

type QueueEvent struct {
	EventType string          `json:"eventType"`
	Payload   json.RawMessage `json:"payload"`
}

func (c *Client) Send(ctx context.Context, options *SendOptions) (err error) {
	var deduplicationID *string
	if options.DeduplicationID != "" {
		deduplicationID = aws.String(options.Event + "-" + options.DeduplicationID)
	}

	payload, err := c.serializeEventPayload(options.Body)
	if err != nil {
		return err
	}

	bodyJson, err := c.serializeEventBody(options.Event, payload)
	if err != nil {
		return err
	}

	_, err = c.sqs.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:               c.buildQueueUrl(options.Queue),
		MessageBody:            &bodyJson,
		MessageGroupId:         &options.Event,
		MessageDeduplicationId: deduplicationID,
	})

	return err
}

func (c *Client) serializeEventPayload(body interface{}) (json.RawMessage, error) {
	if body != nil {
		return json.Marshal(body)
	} else {
		return []byte("{}"), nil
	}
}

func (c *Client) serializeEventBody(eventType string, payload json.RawMessage) (string, error) {
	body := QueueEvent{
		EventType: eventType,
		Payload:   payload,
	}

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	return string(bodyJson), nil
}
