//go:build dev

package awsdynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"os"
)

func newConfig(ctx context.Context) (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithDefaultRegion(os.Getenv("AWS_DEFAULT_REGION")),
	)
	return &cfg, err
}

func configureClient(options *dynamodb.Options) {
	options.BaseEndpoint = aws.String("http://dynamodb:8000")
}
