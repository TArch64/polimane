package userschemas

import (
	"gorm.io/gorm"
)

func IncludeUsersScope(db *gorm.DB) *gorm.DB {
	return db.Joins("JOIN users ON users.id = user_schemas.user_id")
}
