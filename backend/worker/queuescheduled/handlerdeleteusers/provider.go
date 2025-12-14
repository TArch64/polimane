package handlerdeleteusers

import (
	"go.uber.org/fx"

	repositoryusers "polimane/backend/repository/users"
)

type Handler struct {
	users *repositoryusers.Client
}

type ProviderOptions struct {
	fx.In
	Users *repositoryusers.Client
}

func Provider(options ProviderOptions) *Handler {
	return &Handler{
		users: options.Users,
	}
}
