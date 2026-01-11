package subscriptioncounters

import (
	"context"
	"database/sql"
	"fmt"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/signal"
)

const syncUserCounterSQL = `
UPDATE user_subscriptions AS us
SET counters = json_set(counters, @counter_path, TO_JSONB(c.count))
FROM (%[1]s) AS c
WHERE us.user_id = c.user_id
RETURNING us.user_id AS id, (counters->>@counter_name)::smallint AS count
`

const changeUserCounterSQL = `
UPDATE user_subscriptions
SET counters = json_set(counters, @counter_path, TO_JSONB((counters->>@counter_name)::smallint %[1]c @delta))
WHERE user_id IN @user_ids
RETURNING user_id AS id, (counters->>@counter_name)::smallint AS count
`

type UserCounter struct {
	counterName string
	counterPath string
	syncSQL     string
	addSQL      string
	removeSQL   string
	db          *gorm.DB
	signals     *signal.Container
	local       *accessor[model.UserSubscription]
}

type userCounterDeps struct {
	DB      *gorm.DB
	Signals *signal.Container
}

type userCounterOptions struct {
	Deps     *userCounterDeps
	Name     string
	CountSQL string
	Local    *accessor[model.UserSubscription]
}

func newUserCounter(options *userCounterOptions) *UserCounter {
	return &UserCounter{
		counterName: options.Name,
		counterPath: fmt.Sprintf("{%s}", options.Name),
		syncSQL:     fmt.Sprintf(syncUserCounterSQL, options.CountSQL),
		addSQL:      fmt.Sprintf(changeUserCounterSQL, '+'),
		removeSQL:   fmt.Sprintf(changeUserCounterSQL, '-'),
		db:          options.Deps.DB,
		signals:     options.Deps.Signals,
		local:       options.Local,
	}
}

func (p *UserCounter) Sync(ctx context.Context, userIDs ...model.ID) error {
	var updated []*updatedCounter

	err := gorm.
		G[model.UserSubscription](p.db).
		Raw(p.syncSQL,
			sql.Named("counter_name", p.counterName),
			sql.Named("counter_path", p.counterPath),
			sql.Named("user_ids", userIDs),
		).
		Scan(ctx, &updated)

	if err != nil {
		return err
	}

	p.updateCache(ctx, updated)
	return nil
}

func (p *UserCounter) AddTx(
	ctx context.Context,
	tx *gorm.DB,
	value uint16,
	userIDs ...model.ID,
) error {
	return p.change(ctx, tx, p.addSQL, value, userIDs)
}

func (p *UserCounter) RemoveTx(
	ctx context.Context,
	tx *gorm.DB,
	value uint16,
	userIDs ...model.ID,
) error {
	return p.change(ctx, tx, p.removeSQL, value, userIDs)
}

func (p *UserCounter) change(
	ctx context.Context,
	tx *gorm.DB,
	querySQL string,
	value uint16,
	userIDs []model.ID,
) error {
	var updated []*updatedCounter

	err := gorm.
		G[model.UserSubscription](tx).
		Raw(querySQL,
			sql.Named("counter_name", p.counterName),
			sql.Named("counter_path", p.counterPath),
			sql.Named("delta", value),
			sql.Named("user_ids", userIDs),
		).
		Scan(ctx, &updated)

	if err != nil {
		return err
	}

	p.updateCache(ctx, updated)
	return nil
}

func (p *UserCounter) updateCache(ctx context.Context, updated []*updatedCounter) {
	for _, row := range updated {
		event := signal.NewUpdateUserCacheEvent(row.ID, func(user *model.User) {
			p.local.Set(user.Subscription, row.Count)
		})

		p.signals.UpdateUserCacheSync.Emit(ctx, event)
	}
}
