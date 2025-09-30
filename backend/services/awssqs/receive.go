package awssqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func (i *Impl) Receive(ctx context.Context, queue string) ([]types.Message, error) {
	output, err := i.sqs.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:              i.buildQueueUrl(queue),
		MessageAttributeNames: []string{"All"},
	})

	if err != nil {
		return nil, err
	}

	return output.Messages, nil
}
