package schemas

import (
	"gorm.io/gorm"

	"polimane/backend/model"
	repositoryuserschemas "polimane/backend/repository/userschemas"
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

type Impl struct {
	db          *gorm.DB
	userSchemas repositoryuserschemas.Client
	signals     *signal.Container
}

var _ Client = (*Impl)(nil)

func Provider(
	db *gorm.DB,
	userSchemas repositoryuserschemas.Client,
	signals *signal.Container,
) Client {
	return &Impl{
		db:          db,
		userSchemas: userSchemas,
		signals:     signals,
	}
}
