package handlercleanupinvitations

import (
	"go.uber.org/fx"

	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
)

type Handler struct {
	schemaInvitations *repositoryschemainvitations.Client
}

type ProviderOptions struct {
	fx.In
	SchemaInvitations *repositoryschemainvitations.Client
}

func Provider(options ProviderOptions) *Handler {
	return &Handler{
		schemaInvitations: options.SchemaInvitations,
	}
}
