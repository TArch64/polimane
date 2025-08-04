package awsconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"

	"polimane/backend/env"
)

func Provider(ctx context.Context, environment *env.Environment) (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx, func(options *config.LoadOptions) error {
		return configure(environment, options)
	})

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
