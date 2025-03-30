package awsdynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var client *dynamodb.Client

func Client() *dynamodb.Client {
	return client
}

func Init(ctx context.Context) error {
	cfg, err := newConfig(ctx)
	if err != nil {
		return err
	}

	client = dynamodb.NewFromConfig(*cfg, configureClient)
	return nil
}
