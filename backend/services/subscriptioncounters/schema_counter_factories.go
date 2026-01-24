package subscriptioncounters

import (
	"polimane/backend/model"
)

func newSchemaBeads(deps *schemaCounterDeps) *SchemaCounter[uint16, int16] {
	return newSchemaCounter[uint16, int16](&schemaCounterOptions[uint16, int16]{
		name:              "schemaBeads",
		schemaCounterDeps: deps,

		counterValue: model.NewAccessor[*model.UserSchema, uint16](
			func(target *model.UserSchema) uint16 {
				return target.Counters.Data().SchemaBeads
			},
			func(target *model.UserSchema, value uint16) {
				target.Counters.Data().SchemaBeads = value
			},
		),

		counterLimit: model.NewAccessor[*model.UserSubscription, *uint16](
			func(target *model.UserSubscription) *uint16 {
				return target.Plan().Limits.SchemaBeads
			}, nil,
		),
	})
}

func newSharedAccess(deps *schemaCounterDeps) *SchemaCounter[uint8, int8] {
	return newSchemaCounter[uint8, int8](&schemaCounterOptions[uint8, int8]{
		name:              "sharedAccess",
		schemaCounterDeps: deps,

		counterValue: model.NewAccessor[*model.UserSchema, uint8](
			func(target *model.UserSchema) uint8 {
				return target.Counters.Data().SharedAccess
			},
			func(target *model.UserSchema, value uint8) {
				target.Counters.Data().SharedAccess = value
			},
		),

		counterLimit: model.NewAccessor[*model.UserSubscription, *uint8](
			func(target *model.UserSubscription) *uint8 {
				return target.Plan().Limits.SharedAccess
			}, nil,
		),
	})
}
