package awssqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func (c *Client) Receive(ctx context.Context, queue string) ([]types.Message, error) {
	output, err := c.sqs.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:              c.buildQueueUrl(queue),
		MessageAttributeNames: []string{"All"},
		WaitTimeSeconds:       20,
	}, SkipLogging)

	if err != nil {
		return nil, err
	}

	return output.Messages, nil
}
