package subscriptioncounters

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/signal"
)

type ChangeSet map[model.ID]int16

type Service struct {
	SchemasCreated *UserCounter
	SchemaBeads    *SchemaCounter
	SharedAccess   *SchemaCounter
}

type ProviderOptions struct {
	fx.In
	DB      *gorm.DB
	Signals *signal.Container
}

func Provider(options ProviderOptions) *Service {
	userDeps := &userCounterDeps{
		db:      options.DB,
		signals: options.Signals,
	}

	schemaDeps := &schemaCounterDeps{
		db: options.DB,
	}

	return &Service{
		SchemasCreated: newSchemasCreated(userDeps),
		SchemaBeads:    newSchemaBeads(schemaDeps),
		SharedAccess:   newSharedAccess(schemaDeps),
	}
}

type updatedCounter struct {
	ID    model.ID
	Count uint16
}
