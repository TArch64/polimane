package repository

import "gorm.io/gorm"

type Scope = func(db *gorm.Statement)
