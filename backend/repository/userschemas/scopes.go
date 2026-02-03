package userschemas

import (
	"polimane/backend/model"
	"polimane/backend/repository"
)

var (
	IncludeUsersScope   = repository.Join("JOIN users ON users.id = user_schemas.user_id")
	IncludeSchemasScope = repository.Join("JOIN schemas ON schemas.id = user_schemas.schema_id")

	FilterWithoutFolder = repository.Where("folder_id IS NULL")
)

func FolderIDEq(id *model.ID) repository.Scope {
	if id == nil {
		return FilterWithoutFolder
	}
	return repository.Where("folder_id = ?", *id)
}
