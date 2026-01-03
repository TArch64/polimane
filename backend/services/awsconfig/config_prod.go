//go:build !dev

package awsconfig

import (
	"github.com/aws/aws-sdk-go-v2/config"
)

func configure(_ *Options, _ *config.LoadOptions) {}
