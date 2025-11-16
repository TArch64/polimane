package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) Update(ctx context.Context, updates model.UserSchema, scopes ...repository.Scope) error {
	return c.UpdateTx(ctx, c.db, updates, scopes...)
}

func (c *Client) UpdateTx(ctx context.Context, tx *gorm.DB, updates model.UserSchema, scopes ...repository.Scope) error {
	_, err := gorm.
		G[model.UserSchema](tx).
		Scopes(scopes...).
		Updates(ctx, updates)

	return err
}
