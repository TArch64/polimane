//go:build !dev

package awsconfig

import (
	"github.com/aws/aws-sdk-go-v2/config"

	"polimane/backend/env"
)

func configure(_ *env.Environment, _ *config.LoadOptions) {}
