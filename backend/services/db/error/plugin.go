package dberror

import (
	"strings"

	"gorm.io/gorm"
)

type Plugin struct{}

func New() gorm.Plugin {
	return &Plugin{}
}

func (p *Plugin) Name() string {
	return "error_handler"
}

func (p *Plugin) Initialize(db *gorm.DB) (err error) {
	err = db.
		Callback().
		Create().
		After("gorm:after_create").
		Register(p.Name()+":after_create", p.handleCreateErr)

	if err != nil {
		return err
	}

	return db.
		Callback().
		Update().
		After("gorm:after_update").
		Register(p.Name()+":after_update", p.handleUpdateErr)
}

func (p *Plugin) matchesErr(db *gorm.DB, text string) bool {
	return db.Error != nil && strings.Contains(db.Error.Error(), text)
}
