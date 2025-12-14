package schemas

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/model"
	repositorybase "polimane/backend/repository/base"
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
	*repositorybase.Client[model.Schema]
	userSchemas *repositoryuserschemas.Client
	signals     *signal.Container
}

func Provider(options ClientOptions) *Client {
	return &Client{
		Client:      repositorybase.New[model.Schema](options.DB),
		userSchemas: options.UserSchemas,
		signals:     options.Signals,
	}
}
