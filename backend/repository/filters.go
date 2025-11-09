package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
)

type Filter = model.Scope

func AddWhere(db *gorm.Statement, expr ...clause.Expression) {
	db.AddClause(clause.Where{Exprs: expr})
}

func EqEmail(email string) Filter {
	return func(db *gorm.Statement) {
		AddWhere(db, gorm.Expr("email = ?", email))
	}
}

func InSchemaIDs(IDs []model.ID) Filter {
	return func(db *gorm.Statement) {
		AddWhere(db, gorm.Expr("schema_id IN (?)", IDs))
	}
}
