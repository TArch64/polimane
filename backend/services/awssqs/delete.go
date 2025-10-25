package awssqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func (c *Client) Delete(ctx context.Context, queue string, receipt string) error {
	_, err := c.sqs.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      c.buildQueueUrl(queue),
		ReceiptHandle: &receipt,
	})

	return err
}
