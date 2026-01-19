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

type userCounterOptions[CV counterValue, CD counterDelta] struct {
	*userCounterDeps
	name         string
	counterValue *model.Accessor[*model.UserSubscription, CV]
	counterLimit *model.Accessor[*model.UserSubscription, *CV]
}

type UserCounter[CV counterValue, CD counterDelta] struct {
	*userCounterOptions[CV, CD]
}

func newUserCounter[CV counterValue, CD counterDelta](options *userCounterOptions[CV, CD]) *UserCounter[CV, CD] {
	return &UserCounter[CV, CD]{userCounterOptions: options}
}

func (u *UserCounter[CV, CD]) CanAdd(subscription *model.UserSubscription, delta CV) bool {
	limit := u.counterLimit.Get(subscription)
	if limit == nil {
		return true
	}
	value := u.counterValue.Get(subscription) + delta
	return value <= *limit
}

func (u *UserCounter[CV, CD]) ChangeTx(ctx context.Context, tx *gorm.DB, values ChangeSet) error {
	queryValues, args := repository.NamedUpdateValues(values)
	query := fmt.Sprintf(changeUserCounterSQL, queryValues)
	args = append(args, sql.Named("counter_name", u.name))
	var updated []*updatedCounter[CV]

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

func (u *UserCounter[CV, CD]) updateCache(ctx context.Context, updated []*updatedCounter[CV]) {
	for _, row := range updated {
		event := signal.NewUpdateUserCacheEvent(row.ID, func(user *model.User) {
			u.counterValue.Set(user.Subscription, row.Count)
		})

		u.signals.UpdateUserCacheSync.Emit(ctx, event)
	}
}
