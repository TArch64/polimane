package schemainvitations

import (
	"context"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
)

type CreateOptions struct {
	Email     string
	SchemaID  model.ID
	Access    model.AccessLevel
	ExpiresAt time.Time
}

func (c *Client) Create(ctx context.Context, options *CreateOptions) error {
	schemaInvitation := &model.SchemaInvitation{
		Email:     options.Email,
		SchemaID:  options.SchemaID,
		Access:    options.Access,
		ExpiresAt: options.ExpiresAt,
	}

	return gorm.
		G[model.SchemaInvitation](c.db, clause.OnConflict{DoNothing: true}).
		Create(ctx, schemaInvitation)
}
