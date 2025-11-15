package folderschemas

import (
	"polimane/backend/model"
	"polimane/backend/repository"

	"gorm.io/gorm"
)

func FolderIDEq(id model.ID) repository.Scope {
	return func(db *gorm.Statement) {
		repository.AddWhere(db, gorm.Expr("folder_id = ?", id))
	}
}
