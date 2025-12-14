package userschemas

import (
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func IncludeUsersLegacyScope(db *gorm.DB) *gorm.DB {
	return db.Joins("JOIN users ON users.id = user_schemas.user_id")
}

func IncludeSchemasScope() repository.Scope {
	return repository.Join("JOIN schemas ON schemas.id = user_schemas.schema_id")
}

var (
	FilterWithoutFolder = repository.Where("folder_id IS NULL")
)

func FolderIDEq(id *model.ID) repository.Scope {
	if id == nil {
		return FilterWithoutFolder
	} else {
		return repository.Where("folder_id = ?", *id)
	}
}
