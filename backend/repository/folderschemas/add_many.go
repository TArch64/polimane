package folderschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type AddManyOptions struct {
	FolderID  model.ID
	SchemaIDs []model.ID
}

func (c *Client) AddManyTx(ctx context.Context, tx *gorm.DB, options *AddManyOptions) (err error) {
	folderSchemas := make([]model.FolderSchema, len(options.SchemaIDs))
	for i, schemaID := range options.SchemaIDs {
		folderSchemas[i] = model.FolderSchema{
			FolderID: options.FolderID,
			SchemaID: schemaID,
		}
	}

	if err = c.CreateManyTx(ctx, tx, &folderSchemas); err != nil {
		return err
	}

	// Delete schemas from old folder later
	return nil
}
