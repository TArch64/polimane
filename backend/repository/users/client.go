package users

import (
	"context"

	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/model"
)

type Client interface {
	ByID(ctx context.Context, id model.ID) (*model.User, error)
	CreateIfNeeded(ctx context.Context, workosID string) (*model.User, error)
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
