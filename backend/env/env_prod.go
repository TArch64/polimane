//go:build !dev

package env

import (
	"github.com/Netflix/go-env"

	"polimane/backend/base"
	"polimane/backend/services/bitwarden"
)

const IsDev = false

func loadEnvs(instance *Environment, bitwardenClient bitwarden.Client) error {
	err := bitwardenClient.LoadToEnviron([]string{
		"BACKEND_SECRET_KEY",
		"BACKEND_SENTRY_DSN",
		"BACKEND_DATABASE_URL",
		"BACKEND_WORKOS_CLIENT_ID",
		"BACKEND_WORKOS_API_KEY",
		"BACKEND_SQS_BASE_URL",
	})

	if err != nil {
		return base.TagError("env.load.bitwarden.envs", err)
	}

	err = bitwardenClient.DownloadCerts([]*bitwarden.DownloadingCert{
		{
			SID:  "BACKEND_DATABASE_CERT_SID",
			Dest: "/tmp/postgres/ca-cert.pem",
		},
	})

	if err != nil {
		return base.TagError("env.load.bitwarden.certs", err)
	}

	_, err = env.UnmarshalFromEnviron(instance)
	return err
}
