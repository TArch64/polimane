//go:build !dev

package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

func newLogger() logger.Interface {
	return logger.Discard
}

func newTracingPlugin() gorm.Plugin {
	return tracing.NewPlugin()
}
