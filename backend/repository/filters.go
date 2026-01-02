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

func IDsIn(IDs []model.ID, table ...string) Scope {
	column := Column("id", table...)
	return Where(column+" IN (?)", IDs)
}

func SchemaIDsIn(IDs []model.ID) Scope {
	return Where("schema_id IN (?)", IDs)
}

func SoftDeletedOnly(table ...string) Scope {
	return func(stmt *gorm.Statement) {
		column := Column("deleted_at", table...)
		IncludeSoftDeleted(stmt)
		Where(column + " IS NOT NULL")(stmt)
	}
}

func SoftDeletedDaysAgo(days uint8, table ...string) Scope {
	return func(stmt *gorm.Statement) {
		SoftDeletedOnly(table...)(stmt)
		interval := fmt.Sprintf("%d days", days)
		Where("deleted_at <= NOW() - ?::INTERVAL", interval)(stmt)
	}
}

func Column(name string, table ...string) string {
	if len(table) > 0 {
		return fmt.Sprintf("%s.%s", table[0], name)
	}
	return name
}
