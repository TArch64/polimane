package subscriptioncounters

import (
	"polimane/backend/model"
)

func newSchemasCreated(deps *userCounterDeps) *UserCounter[uint16, int16] {
	return newUserCounter[uint16, int16](&userCounterOptions[uint16, int16]{
		name:            "schemasCreated",
		userCounterDeps: deps,

		counterValue: model.NewAccessor[*model.UserSubscription, uint16](
			func(target *model.UserSubscription) uint16 {
				return target.Counters.Data().SchemasCreated
			},
			func(target *model.UserSubscription, value uint16) {
				target.Counters.Data().SchemasCreated = value
			},
		),

		counterLimit: model.NewAccessor[*model.UserSubscription, *uint16](
			func(target *model.UserSubscription) *uint16 {
				return target.Limits().SchemasCreated
			}, nil,
		),
	})
}
