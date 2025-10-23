package awssqs

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"go.uber.org/fx"

	"polimane/backend/env"
)

type Option = func(*sqs.Options)

type ClientOptions struct {
	fx.In
	Config *aws.Config
	Env    *env.Environment
}

type Client struct {
	sqs *sqs.Client
	env *env.Environment
}

func (i *Client) buildQueueUrl(queue string) *string {
	url := i.env.AWS.SQSBaseURL + "/" + queue
	return &url
}

func Provider(options ClientOptions) *Client {
	return &Client{
		sqs: sqs.NewFromConfig(*options.Config, configure(options.Env)),
		env: options.Env,
	}
}
