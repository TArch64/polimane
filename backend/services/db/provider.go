package db

import (
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"polimane/backend/base"
	"polimane/backend/env"
	"polimane/backend/services/sentry"
)

type Options struct {
	fx.In
	Env    *env.Environment
	Sentry *sentry.Container
}

func Provider(options Options) (*gorm.DB, error) {
	dialect := postgres.Open(options.Env.Database.URL)

	instance, err := gorm.Open(dialect, &gorm.Config{
		Logger: newLogger(),
	})

	if err != nil {
		return nil, base.TagError("db.open", err)
	}

	if tracing := newTracingPlugin(); tracing != nil {
		if err = instance.Use(tracing); err != nil {
			return nil, base.TagError("db.tracing", err)
		}
	}

	return instance, nil
}
