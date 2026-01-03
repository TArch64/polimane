package users

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/model"
	repositorybase "polimane/backend/repository/base"
	repositoryschemainvitations "polimane/backend/repository/schemainvitations"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	repositoryusersubscriptions "polimane/backend/repository/usersubscriptions"
)

type ClientOptions struct {
	fx.In
	DB                *gorm.DB
	SchemaInvitations *repositoryschemainvitations.Client
	UserSchemas       *repositoryuserschemas.Client
	UserSubscriptions *repositoryusersubscriptions.Client
}

type Client struct {
	*repositorybase.Client[model.User]
	schemaInvitations *repositoryschemainvitations.Client
	userSchemas       *repositoryuserschemas.Client
	userSubscriptions *repositoryusersubscriptions.Client
}

func Provider(options ClientOptions) *Client {
	return &Client{
		Client:            repositorybase.New[model.User](options.DB),
		schemaInvitations: options.SchemaInvitations,
		userSchemas:       options.UserSchemas,
		userSubscriptions: options.UserSubscriptions,
	}
}
