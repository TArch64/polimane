package schemainvitations

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) Delete(ctx context.Context, scopes ...repository.Scope) error {
	return c.DeleteTx(ctx, c.db, scopes...)
}

func (c *Client) DeleteTx(ctx context.Context, tx *gorm.DB, scopes ...repository.Scope) error {
	_, err := gorm.
		G[model.SchemaInvitation](tx).
		Scopes(scopes...).
		Delete(ctx)

	return err
}
