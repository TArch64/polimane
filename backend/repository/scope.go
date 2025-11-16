package repository

import "gorm.io/gorm"

type Scope = func(stmt *gorm.Statement)
