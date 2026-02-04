package usercreate

import (
	"go.uber.org/fx"

	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	repositoryusers "polimane/backend/repository/users"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	repositoryusersubscriptions "polimane/backend/repository/usersubscriptions"
	"polimane/backend/services/subscriptioncounters"
)

type Service struct {
	users                *repositoryusers.Client
	userSubscriptions    *repositoryusersubscriptions.Client
	userSchemas          *repositoryuserschemas.Client
	schemaInvitations    *repositoryschemainvitations.Client
	subscriptionCounters *subscriptioncounters.Service
}

type ProviderOptions struct {
	fx.In
	Users                *repositoryusers.Client
	UserSubscriptions    *repositoryusersubscriptions.Client
	UserSchemas          *repositoryuserschemas.Client
	SchemaInvitations    *repositoryschemainvitations.Client
	SubscriptionCounters *subscriptioncounters.Service
}

func Provider(options ProviderOptions) *Service {
	return &Service{
		users:                options.Users,
		userSubscriptions:    options.UserSubscriptions,
		userSchemas:          options.UserSchemas,
		schemaInvitations:    options.SchemaInvitations,
		subscriptionCounters: options.SubscriptionCounters,
	}
}
