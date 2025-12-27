package repository

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"polimane/backend/model"
)

func Where(expr string, args ...interface{}) Scope {
	return func(stmt *gorm.Statement) {
		stmt.AddClause(clause.Where{
			Exprs: []clause.Expression{gorm.Expr(expr, args...)},
		})
	}
}

func IDEq(id model.ID) Scope {
	return Where("id = ?", id)
}

func UserIDEq(id model.ID) Scope {
	return Where("user_id = ?", id)
}

func EmailEq(email string) Scope {
	return Where("email = ?", email)
}

func IDsIn(IDs []model.ID) Scope {
	return Where("id IN (?)", IDs)
}

func SchemaIDsIn(IDs []model.ID) Scope {
	return Where("schema_id IN (?)", IDs)
}

func IncludeSoftDeleted(stmt *gorm.Statement) {
	stmt.Unscoped = true
}

func SoftDeletedOnly(table ...string) Scope {
	column := "deleted_at"
	if len(table) > 0 {
		column = fmt.Sprintf("%s.%s", table[0], column)
	}

	return func(stmt *gorm.Statement) {
		IncludeSoftDeleted(stmt)
		Where(fmt.Sprintf("%s IS NOT NULL", column))(stmt)
	}
}

func SoftDeletedDaysAgo(days uint8, table ...string) Scope {
	return func(stmt *gorm.Statement) {
		SoftDeletedOnly(table...)(stmt)
		interval := fmt.Sprintf("%d days", days)
		Where("deleted_at <= NOW() - ?::INTERVAL", interval)(stmt)
	}
}
