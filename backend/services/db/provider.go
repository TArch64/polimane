package db

import (
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"

	"polimane/backend/base"
	"polimane/backend/env"
	"polimane/backend/services/db/autoanalyze"
	dberror "polimane/backend/services/db/error"
	"polimane/backend/services/logstdout"
	"polimane/backend/services/sentry"
)

type Options struct {
	fx.In
	Env    *env.Environment
	Sentry *sentry.Container
	Stdout *logstdout.Logger
}

func Provider(options Options) (*gorm.DB, error) {
	dialect := postgres.New(postgres.Config{
		DSN:                  options.Env.Database.URL,
		PreferSimpleProtocol: true,
	})

	instance, err := gorm.Open(dialect, &gorm.Config{
		Logger:      newLogger(),
		PrepareStmt: false,
	})

	if err != nil {
		return nil, base.TagError("db.open", err)
	}

	if err = instance.Use(dberror.New()); err != nil {
		return nil, base.TagError("db.error_handler", err)
	}

	if env.IsDev {
		err = instance.Use(autoanalyze.New(&autoanalyze.PluginOptions{
			Stdout: options.Stdout,
		}))
		if err != nil {
			return nil, base.TagError("db.auto_analyze", err)
		}
	} else {
		if err = instance.Use(tracing.NewPlugin()); err != nil {
			return nil, base.TagError("db.tracing", err)
		}
	}

	return instance, nil
}
