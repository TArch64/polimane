package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"polimane/backend/base"
	"polimane/backend/env"
)

var Instance *gorm.DB

func Init() (err error) {
	dialect := postgres.Open(env.Instance.Database.URL)

	Instance, err = gorm.Open(dialect, &gorm.Config{
		Logger: newLogger(),
	})

	return base.TagError("db.open", err)
}
