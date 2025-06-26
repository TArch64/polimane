package awsssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"

	"polimane/backend/base"
)

var client *ssm.Client

func newConfig(ctx context.Context) (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	return &cfg, err
}

func Init(ctx context.Context) error {
	cfg, err := newConfig(ctx)
	if err != nil {
		return base.TagError("ssm.config", err)
	}

	client = ssm.NewFromConfig(*cfg)
	return nil
}

func Client() *ssm.Client {
	return client
}
