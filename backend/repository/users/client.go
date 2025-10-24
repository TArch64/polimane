package users

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	repositoryuserschemas "polimane/backend/repository/userschemas"
)

type ClientOptions struct {
	fx.In
	DB                *gorm.DB
	SchemaInvitations *repositoryschemainvitations.Client
	UserSchemas       *repositoryuserschemas.Client
}

type Client struct {
	db                *gorm.DB
	schemaInvitations *repositoryschemainvitations.Client
	userSchemas       *repositoryuserschemas.Client
}

func Provider(options ClientOptions) *Client {
	return &Client{
		db:                options.DB,
		schemaInvitations: options.SchemaInvitations,
		userSchemas:       options.UserSchemas,
	}
}
