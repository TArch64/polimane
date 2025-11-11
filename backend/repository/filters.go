package repository

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
)

type Scope = model.Scope

func Select(columns ...string) Scope {
	return func(db *gorm.Statement) {
		db.AddClause(clause.Select{
			Expression: gorm.Expr(strings.Join(columns, ", ")),
		})
	}
}

func AddWhere(db *gorm.Statement, expr ...clause.Expression) {
	db.AddClause(clause.Where{Exprs: expr})
}

func IDEq(id model.ID) Scope {
	return func(db *gorm.Statement) {
		AddWhere(db, gorm.Expr("id = ?", id))
	}
}

func UserIDEq(id model.ID) Scope {
	return func(db *gorm.Statement) {
		AddWhere(db, gorm.Expr("user_id = ?", id))
	}
}

func EmailEq(email string) Scope {
	return func(db *gorm.Statement) {
		AddWhere(db, gorm.Expr("email = ?", email))
	}
}

func IDsIn(IDs []model.ID) Scope {
	return func(db *gorm.Statement) {
		AddWhere(db, gorm.Expr("id IN (?)", IDs))
	}
}

func SchemaIDsIn(IDs []model.ID) Scope {
	return func(db *gorm.Statement) {
		AddWhere(db, gorm.Expr("schema_id IN (?)", IDs))
	}
}
