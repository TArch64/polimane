package repository

import "gorm.io/gorm"

func Paginate(offset uint16, limit uint8) Scope {
	return func(db *gorm.Statement) {
		db.
			Offset(int(offset)).
			Limit(int(limit))
	}
}
