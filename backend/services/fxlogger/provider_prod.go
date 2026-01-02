//go:build !dev

package fxlogger

import (
	"go.uber.org/fx"
)

var Provider = fx.NopLogger
