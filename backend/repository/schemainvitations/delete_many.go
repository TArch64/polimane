package schemainvitations

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

type DeleteManyOptions struct {
	Email     string
	SchemaIDs []model.ID
}

func (c *Client) DeleteMany(ctx context.Context, scopes ...repository.Scope) error {
	return c.DeleteManyTx(ctx, c.db, scopes...)
}

func (c *Client) DeleteManyTx(ctx context.Context, tx *gorm.DB, scopes ...repository.Scope) error {
	_, err := gorm.
		G[model.SchemaInvitation](tx).
		Scopes(scopes...).
		Delete(ctx)

	return err
}
