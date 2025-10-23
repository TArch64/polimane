package schemascreenshot

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/fx"

	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/views"
)

type Service struct {
	renderer *views.Renderer
	s3       *s3.Client
	schemas  *repositoryschemas.Client
}

type ProviderOptions struct {
	fx.In
	Renderer *views.Renderer
	S3       *s3.Client
	Schemas  *repositoryschemas.Client
}

func Provider(options ProviderOptions) *Service {
	return &Service{
		renderer: options.Renderer,
		s3:       options.S3,
		schemas:  options.Schemas,
	}
}
