package dberror

import (
	"errors"

	"gorm.io/gorm"
)

var UniqueConstraintErr = errors.New("unique constraint violation")

func (p *Plugin) handleUniqueConstraintError(db *gorm.DB) bool {
	if p.matchesErr(db, "SQLSTATE 23505") {
		db.Error = UniqueConstraintErr
		return true
	}
	return false
}
