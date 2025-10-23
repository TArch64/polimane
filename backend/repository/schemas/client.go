package schemas

import (
	"context"

	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/model"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/awss3"
	"polimane/backend/signal"
)

type Client interface {
	ByID(ctx context.Context, options *ByIDOptions) (*model.Schema, error)
	ByUser(ctx context.Context, options *ByUserOptions) ([]*model.Schema, error)
	Copy(ctx context.Context, options *CopyOptions) (*model.Schema, error)
	Create(ctx context.Context, options *CreateOptions) (schema *model.Schema, err error)
	Delete(ctx context.Context, options *DeleteOptions) (err error)
	Update(ctx context.Context, options *UpdateOptions) (err error)
}

type ClientOptions struct {
	fx.In
	DB          *gorm.DB
	UserSchemas repositoryuserschemas.Client
	Signals     *signal.Container
	S3          awss3.Client
}

type Impl struct {
	db          *gorm.DB
	userSchemas repositoryuserschemas.Client
	signals     *signal.Container
	s3          awss3.Client
}

var _ Client = (*Impl)(nil)

func Provider(options ClientOptions) Client {
	return &Impl{
		db:          options.DB,
		userSchemas: options.UserSchemas,
		signals:     options.Signals,
		s3:          options.S3,
	}
}
