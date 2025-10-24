package userschemas

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/model"
)

type ClientOptions struct {
	fx.In
	DB *gorm.DB
}

type Client struct {
	db *gorm.DB
}

func Provider(options ClientOptions) *Client {
	return &Client{db: options.DB}
}

type WithAccessCheck[O any] struct {
	CurrentUser *model.User
	Operation   *O
}
