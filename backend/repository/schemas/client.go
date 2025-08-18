package schemas

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	"polimane/backend/model"
	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/services/awss3"
	"polimane/backend/signal"
)

type Client interface {
	ByID(options *ByIDOptions) (*model.Schema, error)
	ByUser(options *ByUserOptions) ([]*model.Schema, error)
	Copy(options *CopyOptions) (*model.Schema, error)
	Create(options *CreateOptions) (schema *model.Schema, err error)
	Delete(options *DeleteOptions) (err error)
	Update(options *UpdateOptions) (err error)
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
