package awsconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"go.uber.org/fx"

	"polimane/backend/env"
)

type Options struct {
	fx.In
	Ctx context.Context
	Env *env.Environment
}

func Provider(options Options) (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(options.Ctx, func(loadOptions *config.LoadOptions) error {
		return configure(options.Env, loadOptions)
	})

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
