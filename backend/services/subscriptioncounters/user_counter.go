package subscriptioncounters

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/signal"
)

const changeUserCounterSQL = `
UPDATE user_subscriptions
SET counters = jsonb_increment(counters, @counter_name, @delta)
WHERE user_id IN @user_ids
RETURNING user_id AS id, (counters->>@counter_name)::smallint AS count
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

func (u *UserCounter) AddTx(ctx context.Context, tx *gorm.DB, value int, userIDs ...model.ID) error {
	return u.change(ctx, tx, value, userIDs)
}

func (u *UserCounter) RemoveTx(ctx context.Context, tx *gorm.DB, value int, userIDs ...model.ID) error {
	return u.change(ctx, tx, -value, userIDs)
}

func (u *UserCounter) change(ctx context.Context, tx *gorm.DB, value int, userIDs []model.ID) error {
	var updated []*updatedCounter

	err := gorm.
		G[model.UserSubscription](tx).
		Raw(changeUserCounterSQL,
			sql.Named("counter_name", u.name),
			sql.Named("delta", value),
			sql.Named("user_ids", userIDs),
		).
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
