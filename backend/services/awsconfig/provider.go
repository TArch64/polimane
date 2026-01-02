package awsconfig

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"go.uber.org/fx"

	"polimane/backend/env"
	"polimane/backend/services/appcontext"
)

type Options struct {
	fx.In
	Ctx *appcontext.Ctx
	Env *env.Environment
}

func Provider(options Options) (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(options.Ctx, func(loadOptions *config.LoadOptions) error {
		configure(options.Env, loadOptions)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
