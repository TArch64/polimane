package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"polimane/backend/base"
	"polimane/backend/env"
)

var client *gorm.DB

func Init() (err error) {
	dialect := postgres.Open(env.Env().Database.URL)

	client, err = gorm.Open(dialect, &gorm.Config{
		Logger: newLogger(),
	})

	return base.TagError("db.open", err)
}

func Client() *gorm.DB {
	return client
}
