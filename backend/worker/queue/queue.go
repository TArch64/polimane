package queue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type Interface interface {
	Name() string
	Process(ctx context.Context, message *types.Message) error
}

type Base struct {
	events map[string]EventProcessor
}
