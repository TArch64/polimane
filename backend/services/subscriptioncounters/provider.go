package subscriptioncounters

import (
	"go.uber.org/fx"

	repositoryuserschemas "polimane/backend/repository/userschemas"
	repositoryusersubscriptions "polimane/backend/repository/usersubscriptions"
	"polimane/backend/signal"
)

type Service struct {
	userSchemas       *repositoryuserschemas.Client
	userSubscriptions *repositoryusersubscriptions.Client
	signals           *signal.Container
}

type ProviderOptions struct {
	fx.In
	UserSchemas       *repositoryuserschemas.Client
	UserSubscriptions *repositoryusersubscriptions.Client
	Signals           *signal.Container
}

func Provider(options ProviderOptions) *Service {
	return &Service{
		userSchemas:       options.UserSchemas,
		userSubscriptions: options.UserSubscriptions,
		signals:           options.Signals,
	}
}
