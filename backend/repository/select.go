package repository

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Select(columns ...string) Scope {
	return func(stmt *gorm.Statement) {
		stmt.AddClause(clause.Select{
			Expression: gorm.Expr(strings.Join(columns, ", ")),
		})
	}
}
