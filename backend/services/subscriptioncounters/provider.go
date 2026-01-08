package subscriptioncounters

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/signal"
)

type Service struct {
	db             *gorm.DB
	signals        *signal.Container
	SchemasCreated *PerUser
}

type ProviderOptions struct {
	fx.In
	DB      *gorm.DB
	Signals *signal.Container
}

func Provider(options ProviderOptions) *Service {
	userDeps := &perUserDeps{
		DB:      options.DB,
		Signals: options.Signals,
	}

	return &Service{
		signals:        options.Signals,
		SchemasCreated: newSchemasCreated(userDeps),
	}
}
