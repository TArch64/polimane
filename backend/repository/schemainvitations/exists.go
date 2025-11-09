package schemainvitations

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) Exists(ctx context.Context, filters ...repository.Filter) (bool, error) {
	var exists bool

	err := gorm.
		G[model.SchemaInvitation](c.db).
		Select("1 AS exists").
		Scopes(filters...).
		Scan(ctx, &exists)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return exists, nil
}
