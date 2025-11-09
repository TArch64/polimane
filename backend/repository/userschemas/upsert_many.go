package userschemas

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
)

type UpsertManyOptions struct {
	UserID    model.ID
	SchemaIDs []model.ID
	Updates   *model.UserSchema
}

func (c *Client) UpsertMany(ctx context.Context, options *UpsertManyOptions) error {
	userSchemas := make([]model.UserSchema, len(options.SchemaIDs))
	for idx, schemaID := range options.SchemaIDs {
		userSchemas[idx] = *options.Updates
		userSchemas[idx].UserID = options.UserID
		userSchemas[idx].SchemaID = schemaID
	}

	onConflict := &clause.OnConflict{
		Columns: []clause.Column{
			{Name: "user_id"},
			{Name: "schema_id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"access",
			"updated_at",
		}),
	}

	return gorm.
		G[model.UserSchema](c.db, onConflict).
		CreateInBatches(ctx, &userSchemas, model.DefaultBatch)
}
