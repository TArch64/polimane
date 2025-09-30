package awssqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"go.uber.org/fx"

	"polimane/backend/env"
)

type Option = func(*sqs.Options)

type ClientOptions struct {
	fx.In
	Config *aws.Config
	Env    *env.Environment
}

type Client interface {
	Receive(ctx context.Context, queue string) ([]types.Message, error)
	Send(ctx context.Context, options *SendOptions) error
	Delete(ctx context.Context, queue string, handle string) error
}

type Impl struct {
	sqs *sqs.Client
	env *env.Environment
}

func (i *Impl) buildQueueUrl(queue string) *string {
	url := i.env.AWS.SQSBaseURL + "/" + queue
	return &url
}

func Provider(options ClientOptions) Client {
	return &Impl{
		sqs: sqs.NewFromConfig(*options.Config, configure(options.Env)),
		env: options.Env,
	}
}
