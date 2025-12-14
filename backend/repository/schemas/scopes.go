package schemas

import (
	"polimane/backend/model"
	"polimane/backend/repository"
)

func IncludeUserSchemaScope(userID model.ID) repository.Scope {
	return repository.Join("JOIN user_schemas ON user_schemas.schema_id = schemas.id AND user_schemas.user_id = ?", userID)
}

var (
	FilterScreenshoted = repository.Where("screenshoted_at IS NOT NULL")
)
