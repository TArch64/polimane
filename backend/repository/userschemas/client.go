package userschemas

import (
	"context"

	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/model"
)

type Client interface {
	CreateTx(tx *gorm.DB, userID, schemaID model.ID) error
	DeleteTx(tx *gorm.DB, userID, schemaID model.ID) error
	HasAccess(ctx context.Context, userID, schemaID model.ID) error
}

type ClientOptions struct {
	fx.In
	DB *gorm.DB
}

type Impl struct {
	db *gorm.DB
}

var _ Client = (*Impl)(nil)

func Provider(options ClientOptions) Client {
	return &Impl{db: options.DB}
}
