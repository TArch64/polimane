package subscriptioncounters

import (
	"polimane/backend/model"
)

const countSchemasScopedSQL = `
	SELECT user_schemas.user_id, COUNT(user_schemas.schema_id) AS count
	FROM user_schemas
	WHERE user_schemas.user_id IN @user_ids AND user_schemas.deleted_at IS NULL
	GROUP BY user_schemas.user_id
`

func newSchemasCreated(deps *userCounterDeps) *UserCounter {
	return newUserCounter(&userCounterOptions{
		Deps:     deps,
		Name:     "schemasCreated",
		CountSQL: countSchemasScopedSQL,

		Local: &accessor[model.UserSubscription]{
			Get: func(subscription *model.UserSubscription) uint16 {
				return subscription.Counters.Data().SchemasCreated
			},

			Set: func(subscription *model.UserSubscription, value uint16) {
				subscription.Counters.Data().SchemasCreated = value
			},
		},
	})
}
