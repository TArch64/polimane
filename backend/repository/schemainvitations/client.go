package schemainvitations

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
	*repositorybase.Client[model.SchemaInvitation]
}

func Provider(options ClientOptions) *Client {
	return &Client{
		Client: repositorybase.New[model.SchemaInvitation](options.DB),
	}
}
