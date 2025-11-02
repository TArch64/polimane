package model

import "gorm.io/gorm"

type Scope = func(*gorm.DB) *gorm.DB

type Pagination struct {
	Offset uint16
	Limit  uint8
}

func PaginationScope(pagination *Pagination) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Offset(int(pagination.Offset)).
			Limit(int(pagination.Limit))
	}
}
