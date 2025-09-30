package awssqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func (i *Impl) Delete(ctx context.Context, queue string, receipt string) error {
	_, err := i.sqs.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      i.buildQueueUrl(queue),
		ReceiptHandle: &receipt,
	})

	return err
}
