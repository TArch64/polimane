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
UPDATE user_subscriptions
SET counters = json_set(counters, '{%s}', TO_JSONB(computed.count))
FROM (%s) AS computed
WHERE user_subscriptions.user_id = computed.user_id
`

const changeUserCounterSQL = `
UPDATE user_subscriptions
SET counters = json_set(counters, '{%s}', TO_JSONB((counters->>'%s')::smallint %c @value))
WHERE user_id IN @user_ids
`

type PerUser struct {
	syncSQL   string
	addSQL    string
	removeSQL string
	db        *gorm.DB
	signals   *signal.Container
}

type perUserDeps struct {
	DB      *gorm.DB
	Signals *signal.Container
}

type perUserOptions struct {
	Deps     *perUserDeps
	Name     string
	CountSQL string
}

func newPerUser(options *perUserOptions) *PerUser {
	return &PerUser{
		syncSQL:   fmt.Sprintf(syncUserCounterSQL, options.Name, options.CountSQL),
		addSQL:    fmt.Sprintf(changeUserCounterSQL, options.Name, options.Name, '+'),
		removeSQL: fmt.Sprintf(changeUserCounterSQL, options.Name, options.Name, '-'),
		db:        options.Deps.DB,
		signals:   options.Deps.Signals,
	}
}

func (p *PerUser) Sync(ctx context.Context, userIDs ...model.ID) error {
	err := gorm.
		G[model.UserSubscription](p.db).
		Exec(ctx, p.syncSQL, sql.Named("user_ids", userIDs))

	if err != nil {
		return err
	}

	p.invalidateCache(ctx, userIDs)
	return nil
}

func (p *PerUser) Add(
	ctx context.Context,
	value uint16,
	userIDs ...model.ID,
) error {
	return p.change(ctx, p.addSQL, value, userIDs)
}

func (p *PerUser) Remove(
	ctx context.Context,
	value uint16,
	userIDs ...model.ID,
) error {
	return p.change(ctx, p.removeSQL, value, userIDs)
}

func (p *PerUser) change(
	ctx context.Context,
	querySQL string,
	value uint16,
	userIDs []model.ID,
) error {
	err := gorm.
		G[model.UserSubscription](p.db).
		Exec(ctx, querySQL,
			sql.Named("value", value),
			sql.Named("user_ids", userIDs),
		)

	if err != nil {
		return err
	}

	p.invalidateCache(ctx, userIDs)
	return nil
}

func (p *PerUser) invalidateCache(ctx context.Context, userIDs []model.ID) {
	for _, userID := range userIDs {
		p.signals.InvalidateUserCache.Emit(ctx, userID)
	}
}
