package schemas

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/awss3"
	"polimane/backend/signal"
)

type ClientOptions struct {
	fx.In
	DB          *gorm.DB
	UserSchemas *repositoryuserschemas.Client
	Signals     *signal.Container
	S3          awss3.Client
}

type Client struct {
	db          *gorm.DB
	userSchemas *repositoryuserschemas.Client
	signals     *signal.Container
	s3          awss3.Client
}

func Provider(options ClientOptions) *Client {
	return &Client{
		db:          options.DB,
		userSchemas: options.UserSchemas,
		signals:     options.Signals,
		s3:          options.S3,
	}
}
