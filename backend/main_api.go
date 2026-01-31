package main

import (
	"go.uber.org/fx"

	"polimane/backend/api"
	apiauth "polimane/backend/api/auth"
	"polimane/backend/api/base"
	apifolders "polimane/backend/api/folders"
	apiping "polimane/backend/api/ping"
	apischemas "polimane/backend/api/schemas"
	apischemasusers "polimane/backend/api/schemas/users"
	apiusers "polimane/backend/api/users"
	apiauthfactors "polimane/backend/api/users/authfactors"
	"polimane/backend/env"
	repositoryfolders "polimane/backend/repository/folders"
	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	repositoryschemas "polimane/backend/repository/schemas"
	repositoryusers "polimane/backend/repository/users"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/appcontext"
	"polimane/backend/services/awscloudwatch"
	"polimane/backend/services/awsconfig"
	"polimane/backend/services/awss3"
	"polimane/backend/services/awssqs"
	"polimane/backend/services/bitwarden"
	"polimane/backend/services/db"
	"polimane/backend/services/fxlogger"
	"polimane/backend/services/logpersistent"
	"polimane/backend/services/logstdout"
	"polimane/backend/services/schemadelete"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/services/sentry"
	"polimane/backend/services/usercreate"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
	"polimane/backend/views"
)

func AsController(f any) any {
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

			// repositories
			repositoryuserschemas.Provider,
			repositoryusers.Provider,
			repositoryschemas.Provider,
			repositoryschemainvitations.Provider,
			repositoryfolders.Provider,

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
			awscloudwatch.Provider,
			views.Provider,
			schemascreenshot.Provider,
			schemadelete.Provider,
			logstdout.Provider,
			logpersistent.Provider,
			usercreate.Provider,

			// api
			apiauth.MiddlewareProvider,
			AsController(apiping.Provider),
			AsController(apiauth.Provider),
			AsController(apiusers.Provider),
			apiauthfactors.Provider, // users child controller
			AsController(apischemas.Provider),
			apischemasusers.Provider, // schemas child controller
			AsController(apifolders.Provider),
			api.OptionsProvider,
			api.Provider,
		),
		fxlogger.Provider,
		fx.Invoke(api.Start),
	).Run()
}
