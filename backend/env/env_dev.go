//go:build dev

package env

import (
	"github.com/Netflix/go-env"

	"polimane/backend/services/bitwarden"
)

const IsDev = true

func loadEnvs(instance *Environment, _ *bitwarden.Client) error {
	_, err := env.UnmarshalFromEnviron(instance)
	return err
}
