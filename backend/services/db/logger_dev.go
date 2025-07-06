//go:build dev

package db

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

func newLogger() logger.Interface {
	return logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      true,
	})
}
