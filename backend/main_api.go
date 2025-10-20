package main

import (
	"go.uber.org/fx"

	"polimane/backend/api"
	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/api/ping"
	"polimane/backend/api/schemas"
	"polimane/backend/api/users"
	"polimane/backend/env"
	repositoryschemas "polimane/backend/repository/schemas"
	repositoryusers "polimane/backend/repository/users"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/appcontext"
	"polimane/backend/services/awsconfig"
	"polimane/backend/services/awss3"
	"polimane/backend/services/awssqs"
	"polimane/backend/services/bitwarden"
	"polimane/backend/services/db"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/services/sentry"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
	"polimane/backend/views"
)

func Controller(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(base.Controller)),
		fx.ResultTags(`group:"controllers"`),
	)
}

func main() {
	fx.New(
		fx.Provide(
			// external
			appcontext.Provider,

			// services
			bitwarden.Provider,
			env.Provider,
			db.Provider,
			workos.Provider,
			sentry.Provider,
			signal.Provider,
			awsconfig.Provider,
			awss3.Provider,
			awssqs.Provider,
			views.Provider,
			schemascreenshot.Provider,

			// repositories
			repositoryuserschemas.Provider,
			repositoryusers.Provider,
			repositoryschemas.Provider,

			// api
			auth.MiddlewareProvider,
			Controller(ping.Provider),
			Controller(auth.Provider),
			Controller(users.Provider),
			Controller(schemas.Provider),
			api.OptionsProvider,
			api.Provider,
		),
		fx.Invoke(api.Start),
	).Run()
}
