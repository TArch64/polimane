package dberror

import "gorm.io/gorm"

func (p *Plugin) handleCreateErr(db *gorm.DB) {
	if p.handleUniqueConstraintError(db) {
		return
	}
}
