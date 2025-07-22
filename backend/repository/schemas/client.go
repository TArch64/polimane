package schemas

import (
	"gorm.io/gorm"

	repositoryuserschemas "polimane/backend/repository/userschemas"
	"polimane/backend/signal"
)

type Client struct {
	db          *gorm.DB
	userSchemas *repositoryuserschemas.Client
	signals     *signal.Container
}

func Provider(
	db *gorm.DB,
	userSchemas *repositoryuserschemas.Client,
	signals *signal.Container,
) *Client {
	return &Client{
		db:          db,
		userSchemas: userSchemas,
		signals:     signals,
	}
}
