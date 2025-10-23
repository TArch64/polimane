package awssqs

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SendOptions struct {
	Queue           string
	Event           string
	DeduplicationId string
	Body            interface{}
}

func (i *Client) Send(ctx context.Context, options *SendOptions) (err error) {
	var body string
	var deduplicationId *string

	if options.Body != nil {
		var bodyJson []byte
		bodyJson, err = json.Marshal(options.Body)
		if err != nil {
			return err
		}

		body = string(bodyJson)
	} else {
		body = "{}"
	}

	if options.DeduplicationId != "" {
		deduplicationId = aws.String(options.Event + "-" + options.DeduplicationId)
	}

	_, err = i.sqs.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:               i.buildQueueUrl(options.Queue),
		MessageBody:            &body,
		MessageGroupId:         &options.Event,
		MessageDeduplicationId: deduplicationId,

		MessageAttributes: map[string]types.MessageAttributeValue{
			"EventType": {
				DataType:    aws.String("String"),
				StringValue: &options.Event,
			},
		},
	})

	return err
}
