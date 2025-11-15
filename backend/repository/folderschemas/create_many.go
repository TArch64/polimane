package folderschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

func (c *Client) CreateManyTx(ctx context.Context, tx *gorm.DB, folderSchemas *[]model.FolderSchema) error {
	return gorm.
		G[model.FolderSchema](tx, gorm.WithResult()).
		CreateInBatches(ctx, folderSchemas, model.DefaultBatch)
}
