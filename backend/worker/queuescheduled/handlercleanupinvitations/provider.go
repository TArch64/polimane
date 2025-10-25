package handlercleanupinvitations

import (
	"go.uber.org/fx"
)

type Handler struct{}

type ProviderOptions struct {
	fx.In
}

func Provider(options ProviderOptions) *Handler {
	return &Handler{}
}
