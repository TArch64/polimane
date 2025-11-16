package folderschemas

import (
	"polimane/backend/model"
	"polimane/backend/repository"

	"gorm.io/gorm"
)

func FolderIDEq(id model.ID) repository.Scope {
	return func(stmt *gorm.Statement) {
		repository.AddWhere(stmt, gorm.Expr("folder_id = ?", id))
	}
}
