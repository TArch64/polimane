package subscriptioncounters

import (
	"polimane/backend/model"
)

func newSchemasCreated(deps *userCounterDeps) *UserCounter {
	return newUserCounter(&userCounterOptions{
		deps: deps,
		name: "schemasCreated",

		localSet: func(subscription *model.UserSubscription, value uint16) {
			subscription.Counters.Data().SchemasCreated = value
		},
	})
}
