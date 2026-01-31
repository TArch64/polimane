package users

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/model"
	repositorybase "polimane/backend/repository/base"
)

type ClientOptions struct {
	fx.In
	DB *gorm.DB
}

type Client struct {
	*repositorybase.Client[model.User]
}

func Provider(options ClientOptions) *Client {
	return &Client{
		Client: repositorybase.New[model.User](options.DB),
	}
}
