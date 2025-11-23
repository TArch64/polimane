package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
)

func Where(expr ...clause.Expression) Scope {
	return func(stmt *gorm.Statement) {
		stmt.AddClause(clause.Where{Exprs: expr})
	}
}

func IDEq(id model.ID) Scope {
	return Where(gorm.Expr("id = ?", id))
}

func UserIDEq(id model.ID) Scope {
	return Where(gorm.Expr("user_id = ?", id))
}

func EmailEq(email string) Scope {
	return Where(gorm.Expr("email = ?", email))
}

func IDsIn(IDs []model.ID) Scope {
	return Where(gorm.Expr("id IN (?)", IDs))
}

func SchemaIDsIn(IDs []model.ID) Scope {
	return Where(gorm.Expr("schema_id IN (?)", IDs))
}
