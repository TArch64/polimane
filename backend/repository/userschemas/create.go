package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type CreateOptions struct {
	UserID   model.ID
	SchemaID model.ID
	Access   model.AccessLevel
}

func (c *Client) Create(ctx context.Context, options *CreateOptions) (*model.UserSchema, error) {
	return c.CreateTx(ctx, c.db, options)
}

func (c *Client) CreateTx(ctx context.Context, tx *gorm.DB, options *CreateOptions) (*model.UserSchema, error) {
	userSchema := &model.UserSchema{
		UserID:   options.UserID,
		SchemaID: options.SchemaID,
		Access:   options.Access,
	}

	result := gorm.WithResult()

	err := gorm.
		G[model.UserSchema](tx, result).
		Create(ctx, userSchema)

	if err != nil {
		return nil, err
	}

	return userSchema, nil
}
