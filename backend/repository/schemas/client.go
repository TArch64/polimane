package schemas

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
	S3          *s3.Client
}

type Client struct {
	db          *gorm.DB
	userSchemas *repositoryuserschemas.Client
	signals     *signal.Container
	s3          *s3.Client
}

func Provider(options ClientOptions) *Client {
	return &Client{
		db:          options.DB,
		userSchemas: options.UserSchemas,
		signals:     options.Signals,
		s3:          options.S3,
	}
}
