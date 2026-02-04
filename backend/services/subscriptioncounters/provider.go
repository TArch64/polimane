package subscriptioncounters

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/signal"
)

type ChangeSet map[model.ID]int16

type counterValue interface {
	~uint8 | ~uint16
}

type counterDelta interface {
	~int8 | ~int16
}

type updatedCounter[CV counterValue] struct {
	ID    model.ID
	Count CV
}

type Service struct {
	SchemasCreated *UserCounter[uint16, int16]
	SchemaBeads    *SchemaCounter[uint16, int16]
	SharedAccess   *SchemaCounter[uint8, int8]
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
