package main

import (
	"go.uber.org/fx"

	"polimane/backend/env"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/appcontext"
	"polimane/backend/services/awsconfig"
	"polimane/backend/services/awssqs"
	"polimane/backend/services/bitwarden"
	"polimane/backend/services/db"
	"polimane/backend/services/osenv"
	"polimane/backend/services/osfs"
	"polimane/backend/worker"
	"polimane/backend/worker/queue"
	"polimane/backend/worker/queuedebounced"
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
			osfs.Provider,
			osenv.Provider,

			// services
			bitwarden.Provider,
			env.Provider,
			db.Provider,
			// sentry.Provider,
			awsconfig.Provider,
			awssqs.Provider,

			// repositories
			repositoryschemas.Provider,

			// queues
			Queue(queuedebounced.Provider),
			worker.Provider,
		),
		fx.Invoke(worker.Start),
	).Run()
}
