package folderschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func (c *Client) DeleteTx(ctx context.Context, tx *gorm.DB, scopes ...repository.Scope) error {
	_, err := gorm.
		G[model.FolderSchema](tx).
		Scopes(scopes...).
		Delete(ctx)

	return err
}
