//go:build !dev

package db

import "gorm.io/gorm/logger"

func newLogger() logger.Interface {
	return logger.Discard
}
