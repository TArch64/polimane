package repositorybase

import (
	"gorm.io/gorm"
)

type Client[M any] struct {
	DB *gorm.DB
}

func New[M any](db *gorm.DB) *Client[M] {
	return &Client[M]{DB: db}
}
