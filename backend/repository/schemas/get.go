package schemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) GetOut(ctx context.Context, out interface{}, scopes ...repository.Scope) error {
	return gorm.
		G[model.Schema](c.db).
		Scopes(scopes...).
		Scan(ctx, out)
}

func (c *Client) Get(ctx context.Context, scopes ...repository.Scope) (*model.Schema, error) {
	return gorm.
		G[*model.Schema](c.db).
		Scopes(scopes...).
		Take(ctx)
}
