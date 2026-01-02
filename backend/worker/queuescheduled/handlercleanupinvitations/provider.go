package handlercleanupinvitations

import (
	"go.uber.org/fx"

	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	"polimane/backend/services/logpersistent"
)

type Handler struct {
	schemaInvitations *repositoryschemainvitations.Client
	persistentLogger  *logpersistent.Logger
}

type ProviderOptions struct {
	fx.In
	SchemaInvitations *repositoryschemainvitations.Client
	PersistentLogger  *logpersistent.Logger
}

func Provider(options ProviderOptions) *Handler {
	return &Handler{
		schemaInvitations: options.SchemaInvitations,
		persistentLogger:  options.PersistentLogger,
	}
}
