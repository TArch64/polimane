//go:build !dev

package env

import (
	"github.com/Netflix/go-env"

	"polimane/backend/base"
	"polimane/backend/services/bitwarden"
)

func loadEnvs() error {
	err := bitwarden.Init()
	if err != nil {
		return err
	}

	err = bitwarden.LoadToEnviron([]string{
		"BACKEND_DEFAULT_USER",
		"BACKEND_DEFAULT_PASSWORD",
		"BACKEND_SECRET_KEY",
		"BACKEND_SENTRY_DSN",
	})

	if err != nil {
		return base.TagError("env.load.bitwarden", err)
	}

	_, err = env.UnmarshalFromEnviron(environment)
	return err
}
