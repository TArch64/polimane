package dberror

import "gorm.io/gorm"

func (p *Plugin) handleUpdateErr(db *gorm.DB) {
	if p.handleUniqueConstraintError(db) {
		return
	}
}
