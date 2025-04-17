//go:build dev

package awsdynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/smithy-go/middleware"
)

func newConfig(ctx context.Context) (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	return &cfg, err
}

func configureClient(options *dynamodb.Options) {
	options.BaseEndpoint = aws.String("http://dynamodb:8000")

	options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
		return stack.Initialize.Add(&queryLoggerMiddleware{}, middleware.After)
	})
}
