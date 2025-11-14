package model

import "gorm.io/gorm"

type LegacyScope = func(*gorm.DB) *gorm.DB
type Scope = func(db *gorm.Statement)
