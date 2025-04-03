//go:build !dev

package awsdynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func newConfig(ctx context.Context) (*aws.Config, error) {
	panic("not implemented")
}

func configureClient(options *dynamodb.Options) {
	panic("not implemented")
}
