package handlerdeleteusers

import (
	"go.uber.org/fx"

	repositoryfolders "polimane/backend/repository/folders"
	repositoryschemas "polimane/backend/repository/schemas"
	repositoryusers "polimane/backend/repository/users"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/logpersistent"
	"polimane/backend/services/schemadelete"
	"polimane/backend/services/workos"
)

type Handler struct {
	users    *repositoryusers.Client
	deleters []Deleter
}

type ProviderOptions struct {
	fx.In
	Users            *repositoryusers.Client
	Folders          *repositoryfolders.Client
	UserSchemas      *repositoryuserschemas.Client
	Schemas          *repositoryschemas.Client
	SchemaDelete     *schemadelete.Service
	Workos           *workos.Client
	PersistentLogger *logpersistent.Logger
}

func Provider(options ProviderOptions) *Handler {
	return &Handler{
		users: options.Users,
		deleters: []Deleter{
			&DeleterFolders{
				Folders:          options.Folders,
				PersistentLogger: options.PersistentLogger,
			},
			&DeleterWorkos{
				Workos:           options.Workos,
				PersistentLogger: options.PersistentLogger,
			},
			&DeleterUser{
				Users:            options.Users,
				PersistentLogger: options.PersistentLogger,
			},
			&DeleterOrphanSchemas{
				Schemas:          options.Schemas,
				UserSchemas:      options.UserSchemas,
				SchemaDelete:     options.SchemaDelete,
				PersistentLogger: options.PersistentLogger,
			},
		},
	}
}
