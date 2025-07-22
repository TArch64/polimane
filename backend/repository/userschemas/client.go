package userschemas

import "gorm.io/gorm"

type Client struct {
	db *gorm.DB
}

func Provider(db *gorm.DB) *Client {
	return &Client{db: db}
}
