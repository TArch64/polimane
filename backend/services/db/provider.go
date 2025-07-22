package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"polimane/backend/base"
	"polimane/backend/env"
)

func Provider(environment *env.Environment) (*gorm.DB, error) {
	dialect := postgres.Open(environment.Database.URL)

	instance, err := gorm.Open(dialect, &gorm.Config{
		Logger: newLogger(),
	})

	if err != nil {
		return nil, base.TagError("db.open", err)
	}

	return instance, nil
}
