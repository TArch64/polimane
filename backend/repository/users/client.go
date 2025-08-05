package users

import (
	"context"

	"gorm.io/gorm"

	"polimane/backend/model"
)

type Client interface {
	ByID(ctx context.Context, id model.ID) (*model.User, error)
	CreateIfNeeded(ctx context.Context, workosID string) (*model.User, error)
}

type Impl struct {
	db *gorm.DB
}

var _ Client = (*Impl)(nil)

func Provider(db *gorm.DB) Client {
	return &Impl{db: db}
}
