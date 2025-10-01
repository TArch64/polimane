package schemascreenshot

import (
	"context"

	"go.uber.org/fx"

	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/awss3"
	"polimane/backend/views"
)

type Interface interface {
	Screenshot(ctx context.Context, options *ScreenshotOptions) error
}

type Impl struct {
	renderer views.Renderer
	s3       awss3.Client
	schemas  repositoryschemas.Client
}

type ProviderOptions struct {
	fx.In
	Renderer views.Renderer
	S3       awss3.Client
	Schemas  repositoryschemas.Client
}

func Provider(options ProviderOptions) Interface {
	return &Impl{
		renderer: options.Renderer,
		s3:       options.S3,
		schemas:  options.Schemas,
	}
}
