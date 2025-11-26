package schemas

import (
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/repository"
)

func IncludeUserSchemaScope(userID model.ID) repository.Scope {
	expr := gorm.Expr("JOIN user_schemas ON user_schemas.schema_id = schemas.id AND user_schemas.user_id = ?", userID)
	return repository.Join(expr)
}

var (
	FilterScreenshoted = repository.Where(gorm.Expr("screenshoted_at IS NOT NULL"))
)
