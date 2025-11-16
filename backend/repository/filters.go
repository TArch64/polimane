package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
)

func AddWhere(stmt *gorm.Statement, expr ...clause.Expression) {
	stmt.AddClause(clause.Where{Exprs: expr})
}

func IDEq(id model.ID) Scope {
	return func(stmt *gorm.Statement) {
		AddWhere(stmt, gorm.Expr("id = ?", id))
	}
}

func UserIDEq(id model.ID) Scope {
	return func(stmt *gorm.Statement) {
		AddWhere(stmt, gorm.Expr("user_id = ?", id))
	}
}

func EmailEq(email string) Scope {
	return func(stmt *gorm.Statement) {
		AddWhere(stmt, gorm.Expr("email = ?", email))
	}
}

func IDsIn(IDs []model.ID) Scope {
	return func(stmt *gorm.Statement) {
		AddWhere(stmt, gorm.Expr("id IN (?)", IDs))
	}
}

func SchemaIDsIn(IDs []model.ID) Scope {
	return func(stmt *gorm.Statement) {
		AddWhere(stmt, gorm.Expr("schema_id IN (?)", IDs))
	}
}
