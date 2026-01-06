package subscriptioncounters

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"polimane/backend/model"
)

const syncSchemasCreatedSQL = `
UPDATE user_subscriptions
SET counters = JSON_SET(counters, '{schemasCreated}', TO_JSONB(computed.count))
FROM (
	SELECT user_schemas.user_id, COUNT(user_schemas.schema_id) AS count
	FROM user_schemas
	WHERE user_schemas.user_id IN @user_ids AND user_schemas.deleted_at IS NULL
	GROUP BY user_schemas.user_id
) AS computed
WHERE user_subscriptions.user_id = computed.user_id
`

func (s *Service) SyncSchemasCreated(ctx context.Context, userIDs ...model.ID) error {
	err := gorm.
		G[model.UserSubscription](s.userSubscriptions.DB).
		Exec(ctx, syncSchemasCreatedSQL,
			sql.Named("user_ids", userIDs),
		)

	if err != nil {
		return err
	}

	return nil
}
