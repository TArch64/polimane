package repository

import "gorm.io/gorm"

func Paginate(offset uint16, limit uint8) Scope {
	return func(stmt *gorm.Statement) {
		stmt.
			Offset(int(offset)).
			Limit(int(limit))
	}
}
