package userschemas

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type Client interface {
	CreateTx(tx *gorm.DB, userID, schemaID model.ID) error
	DeleteTx(tx *gorm.DB, userID, schemaID model.ID) error
	HasAccess(ctx context.Context, userID, schemaID model.ID) error
}

type Impl struct {
	db *gorm.DB
}

var _ Client = (*Impl)(nil)

func Provider(db *gorm.DB) Client {
	return &Impl{db: db}
}
