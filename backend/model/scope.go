package model

import "gorm.io/gorm"

type LegacyScope = func(*gorm.DB) *gorm.DB
type Scope = func(db *gorm.Statement)

type Pagination struct {
	Offset uint16
	Limit  uint8
}

func PaginationScope(pagination *Pagination) LegacyScope {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Offset(int(pagination.Offset)).
			Limit(int(pagination.Limit))
	}
}
