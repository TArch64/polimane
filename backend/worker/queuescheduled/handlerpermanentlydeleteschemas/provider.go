package handlerpermanentlydeleteschemas

import (
	"go.uber.org/fx"

	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/schemadelete"
)

type Handler struct {
	schemas *repositoryschemas.Client
	delete  *schemadelete.Service
}

type ProviderOptions struct {
	fx.In
	Schemas *repositoryschemas.Client
	Delete  *schemadelete.Service
}

func Provider(options ProviderOptions) *Handler {
	return &Handler{
		schemas: options.Schemas,
		delete:  options.Delete,
	}
}
