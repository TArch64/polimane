package schemainvitations

import (
	"context"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
)

type CreateManyOptions struct {
	Email     string
	SchemaIDs []model.ID
	Access    model.AccessLevel
	ExpiresAt time.Time
}

func (c *Client) CreateMany(ctx context.Context, options *CreateManyOptions) error {
	invitations := make([]model.SchemaInvitation, len(options.SchemaIDs))
	for idx, schemaID := range options.SchemaIDs {
		invitations[idx] = model.SchemaInvitation{
			Email:     options.Email,
			SchemaID:  schemaID,
			Access:    options.Access,
			ExpiresAt: options.ExpiresAt,
		}
	}

	return gorm.
		G[model.SchemaInvitation](c.db, clause.OnConflict{DoNothing: true}).
		CreateInBatches(ctx, &invitations, 100)
}
