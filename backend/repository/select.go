package repository

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Select(columns ...string) Scope {
	return func(db *gorm.Statement) {
		db.AddClause(clause.Select{
			Expression: gorm.Expr(strings.Join(columns, ", ")),
		})
	}
}
