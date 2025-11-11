package schemas

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/signal"
)

type ClientOptions struct {
	fx.In
	DB          *gorm.DB
	UserSchemas *repositoryuserschemas.Client
	Signals     *signal.Container
}

type Client struct {
	db          *gorm.DB
	userSchemas *repositoryuserschemas.Client
	signals     *signal.Container
}

func Provider(options ClientOptions) *Client {
	return &Client{
		db:          options.DB,
		userSchemas: options.UserSchemas,
		signals:     options.Signals,
	}
}
