package schemadelete

import (
	"go.uber.org/fx"

	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/schemascreenshot"
)

type Service struct {
	schemas    *repositoryschemas.Client
	screenshot *schemascreenshot.Service
}

type ProviderOptions struct {
	fx.In
	Schemas    *repositoryschemas.Client
	Screenshot *schemascreenshot.Service
}

func Provider(options ProviderOptions) *Service {
	return &Service{
		schemas:    options.Schemas,
		screenshot: options.Screenshot,
	}
}
