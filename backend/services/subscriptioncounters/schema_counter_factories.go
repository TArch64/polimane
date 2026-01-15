package subscriptioncounters

import (
	"polimane/backend/model"
)

func newSchemaBeads(deps *schemaCounterDeps) *SchemaCounter {
	return newSchemaCounter(&schemaCounterOptions{
		deps: deps,
		name: "schemaBeads",

		localSet: func(target *model.UserSchema, value uint16) {
			target.Counters.Data().SchemaBeads = value
		},
	})
}

func newSharedAccess(deps *schemaCounterDeps) *SchemaCounter {
	return newSchemaCounter(&schemaCounterOptions{
		deps: deps,
		name: "sharedAccess",

		localSet: func(target *model.UserSchema, value uint16) {
			target.Counters.Data().SharedAccess = uint8(value)
		},
	})
}
