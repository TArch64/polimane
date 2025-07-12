//go:build dev

package env

import "github.com/Netflix/go-env"

func loadEnvs() error {
	_, err := env.UnmarshalFromEnviron(Instance)
	return err
}
