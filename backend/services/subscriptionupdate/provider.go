package subscriptionupdate

import (
	"go.uber.org/fx"

	repositoryusersubscriptions "polimane/backend/repository/usersubscriptions"
	"polimane/backend/signal"
)

type Service struct {
	userSubscriptions *repositoryusersubscriptions.Client
	signals           *signal.Container
}

type ProviderOptions struct {
	fx.In
	UserSubscriptions *repositoryusersubscriptions.Client
	Signals           *signal.Container
}

func Provider(options ProviderOptions) *Service {
	return &Service{
		userSubscriptions: options.UserSubscriptions,
		signals:           options.Signals,
	}
}
