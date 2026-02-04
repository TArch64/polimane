package users

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) GetWithSubscription(ctx context.Context, scopes ...repository.Scope) (*model.User, error) {
	return c.GetCustomizable(ctx,
		func(chain gorm.ChainInterface[*model.User]) gorm.ChainInterface[*model.User] {
			join := clause.InnerJoin.Association("Subscription").As("us")
			return chain.Joins(join, nil)
		},
		scopes...,
	)
}
