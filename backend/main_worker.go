package main

import (
	"go.uber.org/fx"

	"polimane/backend/env"
	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	repositoryschemas "polimane/backend/repository/schemas"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/appcontext"
	"polimane/backend/services/awsconfig"
	"polimane/backend/services/awss3"
	"polimane/backend/services/awssqs"
	"polimane/backend/services/bitwarden"
	"polimane/backend/services/db"
	"polimane/backend/services/schemascreenshot"
	"polimane/backend/services/sentry"
	"polimane/backend/signal"
	"polimane/backend/views"
	"polimane/backend/worker"
	"polimane/backend/worker/queue"
	"polimane/backend/worker/queuedebounced"
	"polimane/backend/worker/queuedebounced/handlerschemascreenshot"
	"polimane/backend/worker/queuescheduled"
	"polimane/backend/worker/queuescheduled/handlercleanupinvitations"
)

func Queue(f any) any {
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
			signal.Provider,
			views.Provider,

			// repositories
			repositoryschemas.Provider,
			repositoryuserschemas.Provider,
			repositoryschemainvitations.Provider,

			// queues
			handlerschemascreenshot.Provider,
			Queue(queuedebounced.Provider),

			handlercleanupinvitations.Provider,
			Queue(queuescheduled.Provider),

			worker.Provider,
		),
		fx.Invoke(worker.Start),
	).Run()
}
