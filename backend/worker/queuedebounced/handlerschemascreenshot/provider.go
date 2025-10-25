package handlerschemascreenshot

import (
	"go.uber.org/fx"

	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/schemascreenshot"
)

type Handler struct {
	schemas          *repositoryschemas.Client
	schemaScreenshot *schemascreenshot.Service
}

type ProviderOptions struct {
	fx.In
	Schemas          *repositoryschemas.Client
	SchemaScreenshot *schemascreenshot.Service
}

func Provider(options ProviderOptions) *Handler {
	return &Handler{
		schemas:          options.Schemas,
		schemaScreenshot: options.SchemaScreenshot,
	}
}
