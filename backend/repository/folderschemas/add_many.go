package folderschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

type AddManyOptions struct {
	SchemaIDs   []model.ID
	FolderID    model.ID
	OldFolderID *model.ID
}

func (c *Client) AddMany(ctx context.Context, options *AddManyOptions) (err error) {
	return c.db.Transaction(func(tx *gorm.DB) error {
		return c.AddManyTx(ctx, tx, options)
	})
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

	if options.OldFolderID == nil {
		return nil
	}

	return c.DeleteTx(ctx, tx,
		FolderIDEq(*options.OldFolderID),
		repository.SchemaIDsIn(options.SchemaIDs),
	)
}
