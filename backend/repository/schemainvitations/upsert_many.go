package schemainvitations

import (
	"context"

	"gorm.io/gorm/clause"

	"polimane/backend/model"
)

type UpsertManyOptions struct {
	Email     string
	SchemaIDs []model.ID
	Updates   *model.SchemaInvitation
}

func (c *Client) UpsertMany(ctx context.Context, options *UpsertManyOptions) error {
	schemaInvitations := make([]model.SchemaInvitation, len(options.SchemaIDs))
	for idx, schemaID := range options.SchemaIDs {
		schemaInvitations[idx] = *options.Updates
		schemaInvitations[idx].Email = options.Email
		schemaInvitations[idx].SchemaID = schemaID
	}

	onConflict := &clause.OnConflict{
		Columns: []clause.Column{
			{Name: "email"},
			{Name: "schema_id"},
		},

		DoUpdates: clause.AssignmentColumns([]string{
			"access",
		}),
	}

	return c.InsertMany(ctx, &schemaInvitations, onConflict)
}
