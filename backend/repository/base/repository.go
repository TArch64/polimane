package repositorybase

import (
	"gorm.io/gorm"
)

type Customizer[M any] func(chain gorm.ChainInterface[M]) gorm.ChainInterface[M]

type Client[M any] struct {
	DB *gorm.DB
}

func New[M any](db *gorm.DB) *Client[M] {
	return &Client[M]{DB: db}
}
