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
	"polimane/backend/services/bitwarden"
	"polimane/backend/services/db"
	"polimane/backend/services/jwk"
	"polimane/backend/services/sentry"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
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
			bitwarden.Provider,
			env.Provider,
			db.Provider,
			jwk.Provider,
			workos.Provider,
			sentry.Provider,
			signal.Provider,
			repositoryuserschemas.Provider,
			repositoryusers.Provider,
			repositoryschemas.Provider,
			auth.MiddlewareProvider,
			Controller(ping.Provider),
			Controller(auth.Provider),
			Controller(users.Provider),
			Controller(schemas.Provider),
			api.OptionsProvider,
			fx.Annotate(
				api.Provider,
				fx.ParamTags(`group:"controllers"`),
			),
		),
		fx.Invoke(api.Start),
	).Run()
}
