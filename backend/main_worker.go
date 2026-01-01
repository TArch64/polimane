package main

import (
	"go.uber.org/fx"

	"polimane/backend/env"
	repositoryfolders "polimane/backend/repository/folders"
	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	repositoryschemas "polimane/backend/repository/schemas"
	repositoryusers "polimane/backend/repository/users"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/appcontext"
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
	"polimane/backend/services/workos"
	"polimane/backend/signal"
	"polimane/backend/views"
	"polimane/backend/worker"
	"polimane/backend/worker/queue"
	"polimane/backend/worker/queuedebounced"
	"polimane/backend/worker/queuedebounced/handlerschemascreenshot"
	"polimane/backend/worker/queuescheduled"
	"polimane/backend/worker/queuescheduled/handlercleanupinvitations"
	"polimane/backend/worker/queuescheduled/handlerdeleteusers"
	"polimane/backend/worker/queuescheduled/handlerpermanentlydeleteschemas"
)

func AsQueue(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(queue.Interface)),
		fx.ResultTags(`group:"queues"`),
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
			sentry.Provider,
			awsconfig.Provider,
			awss3.Provider,
			awssqs.Provider,
			schemascreenshot.Provider,
			schemadelete.Provider,
			signal.Provider,
			views.Provider,
			workos.Provider,
			logstdout.Provider,
			logpersistent.Provider,

			// repositories
			repositoryusers.Provider,
			repositoryschemas.Provider,
			repositoryuserschemas.Provider,
			repositoryschemainvitations.Provider,
			repositoryfolders.Provider,

			// queues
			handlerschemascreenshot.Provider,
			AsQueue(queuedebounced.Provider),

			handlercleanupinvitations.Provider,
			handlerdeleteusers.Provider,
			handlerpermanentlydeleteschemas.Provider,
			AsQueue(queuescheduled.Provider),

			worker.Provider,
		),
		fx.WithLogger(fxlogger.Provider),
		fx.Invoke(worker.Start),
	).Run()
}
