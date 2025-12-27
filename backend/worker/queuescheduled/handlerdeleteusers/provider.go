package handlerdeleteusers

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	repositoryfolders "polimane/backend/repository/folders"
	repositoryschemas "polimane/backend/repository/schemas"
	repositoryusers "polimane/backend/repository/users"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/schemadelete"
	"polimane/backend/services/workos"
)

type Handler struct {
	db           *gorm.DB
	users        *repositoryusers.Client
	folders      *repositoryfolders.Client
	userSchemas  *repositoryuserschemas.Client
	schemas      *repositoryschemas.Client
	schemaDelete *schemadelete.Service
	workos       *workos.Client
}

type ProviderOptions struct {
	fx.In
	DB           *gorm.DB
	Users        *repositoryusers.Client
	Folders      *repositoryfolders.Client
	UserSchemas  *repositoryuserschemas.Client
	Schemas      *repositoryschemas.Client
	SchemaDelete *schemadelete.Service
	Workos       *workos.Client
}

func Provider(options ProviderOptions) *Handler {
	return &Handler{
		db:           options.DB,
		users:        options.Users,
		folders:      options.Folders,
		userSchemas:  options.UserSchemas,
		schemas:      options.Schemas,
		schemaDelete: options.SchemaDelete,
		workos:       options.Workos,
	}
}
