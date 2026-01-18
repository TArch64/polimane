package subscriptioncounters

import (
	"context"
	"database/sql"
	"fmt"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
	"polimane/backend/signal"
)

const changeUserCounterSQL = `
UPDATE user_subscriptions
SET counters = jsonb_increment(counters, @counter_name, change_set.delta)
FROM (VALUES %s) as change_set(user_id, delta)
WHERE user_subscriptions.user_id = change_set.user_id
RETURNING user_subscriptions.user_id AS id, (counters->>@counter_name)::smallint AS count
`

type userCounterDeps struct {
	db      *gorm.DB
	signals *signal.Container
}

type userCounterOptions struct {
	*userCounterDeps
	name     string
	localSet model.Set[*model.UserSubscription, uint16]
}

type UserCounter struct {
	*userCounterOptions
}

func newUserCounter(options *userCounterOptions) *UserCounter {
	return &UserCounter{userCounterOptions: options}
}

func (u *UserCounter) ChangeTx(ctx context.Context, tx *gorm.DB, values ChangeSet) error {
	queryValues, args := repository.NamedUpdateValues(values)
	query := fmt.Sprintf(changeUserCounterSQL, queryValues)
	args = append(args, sql.Named("counter_name", u.name))
	var updated []*updatedCounter

	err := gorm.
		G[model.UserSubscription](tx).
		Raw(query, args...).
		Scan(ctx, &updated)

	if err != nil {
		return err
	}

	u.updateCache(ctx, updated)
	return nil
}

func (u *UserCounter) updateCache(ctx context.Context, updated []*updatedCounter) {
	for _, row := range updated {
		event := signal.NewUpdateUserCacheEvent(row.ID, func(user *model.User) {
			u.localSet(user.Subscription, row.Count)
		})

		u.signals.UpdateUserCacheSync.Emit(ctx, event)
	}
}
