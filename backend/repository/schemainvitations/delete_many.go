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

func (c *Client) DeleteMany(ctx context.Context, filters ...repository.Filter) error {
	return c.DeleteManyTx(ctx, c.db, filters...)
}

func (c *Client) DeleteManyTx(ctx context.Context, tx *gorm.DB, filters ...repository.Filter) error {
	_, err := gorm.
		G[model.SchemaInvitation](tx).
		Scopes(filters...).
		Delete(ctx)

	return err
}
