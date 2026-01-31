package usercreate

import (
	"go.uber.org/fx"

	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	repositoryusers "polimane/backend/repository/users"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

type Service struct {
	users             *repositoryusers.Client
	userSchemas       *repositoryuserschemas.Client
	schemaInvitations *repositoryschemainvitations.Client
}

type ProviderOptions struct {
	fx.In
	Users             *repositoryusers.Client
	UserSchemas       *repositoryuserschemas.Client
	SchemaInvitations *repositoryschemainvitations.Client
}

func Provider(options ProviderOptions) *Service {
	return &Service{
		users:             options.Users,
		userSchemas:       options.UserSchemas,
		schemaInvitations: options.SchemaInvitations,
	}
}
