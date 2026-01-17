package subscriptioncounters

import (
	"polimane/backend/model"
)

func newSchemaBeads(deps *schemaCounterDeps) *SchemaCounter {
	return newSchemaCounter(&schemaCounterOptions{
		name:              "schemaBeads",
		schemaCounterDeps: deps,

		localSet: func(target *model.UserSchema, value uint16) {
			target.Counters.Data().SchemaBeads = value
		},
	})
}

func newSharedAccess(deps *schemaCounterDeps) *SchemaCounter {
	return newSchemaCounter(&schemaCounterOptions{
		name:              "sharedAccess",
		schemaCounterDeps: deps,

		localSet: func(target *model.UserSchema, value uint16) {
			target.Counters.Data().SharedAccess = uint8(value)
		},
	})
}
