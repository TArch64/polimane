package model

import "gorm.io/gorm"

type Scope func(*gorm.DB) *gorm.DB
