//go:build !dev

package env

import (
	"sync"

	"github.com/Netflix/go-env"

	"polimane/backend/base"
	"polimane/backend/services/bitwarden"
)

var loadEnvMutex = sync.Mutex{}
var loaded = false

func loadEnvs() error {
	loadEnvMutex.Lock()
	defer loadEnvMutex.Unlock()

	if loaded {
		return nil
	}

	loaded = true

	err := bitwarden.Init()
	if err != nil {
		return err
	}

	err = bitwarden.LoadToEnviron([]string{
		"BACKEND_SECRET_KEY",
		"BACKEND_SENTRY_DSN",
		"BACKEND_DATABASE_URL",
		"BACKEND_WORKOS_CLIENT_ID",
		"BACKEND_WORKOS_API_KEY",
	})

	if err != nil {
		return base.TagError("env.load.bitwarden.envs", err)
	}

	err = bitwarden.DownloadCerts([]*bitwarden.DownloadingCert{
		{
			SID:  "BACKEND_DATABASE_CERT_SID",
			Dest: "/tmp/postgres/ca-cert.pem",
		},
	})

	if err != nil {
		return base.TagError("env.load.bitwarden.certs", err)
	}

	_, err = env.UnmarshalFromEnviron(Instance)
	return err
}
