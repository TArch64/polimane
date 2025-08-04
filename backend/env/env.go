package env

import (
	"net/url"

	"polimane/backend/base"
	"polimane/backend/services/bitwarden"
)

type Environment struct {
	SecretKey   string `env:"BACKEND_SECRET_KEY,required=true"`
	AppDomain   string `env:"BACKEND_APP_DOMAIN,required=true"`
	AppProtocol string `env:"BACKEND_APP_PROTOCOL,required=true"`

	Database struct {
		URL string `env:"BACKEND_DATABASE_URL,required=true"`
	}

	Sentry struct {
		Dsn     string `env:"BACKEND_SENTRY_DSN"`
		Release string `env:"BACKEND_SENTRY_RELEASE"`
	}

	WorkOS struct {
		ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
		ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
	}

	AWS struct {
		// Used only in dev env since in prod we connect using IAM roles
		Region          string `env:"BACKEND_AWS_DEFAULT_REGION"`
		AccessKeyID     string `env:"BACKEND_AWS_ACCESS_KEY_ID"`
		SecretAccessKey string `env:"BACKEND_AWS_SECRET_ACCESS_KEY"`
	}
}

func (e *Environment) AppURL() *url.URL {
	return &url.URL{
		Scheme: e.AppProtocol,
		Host:   e.AppDomain,
	}
}

func (e *Environment) ApiURL() *url.URL {
	return &url.URL{
		Scheme: e.AppProtocol,
		Host:   "api." + e.AppDomain,
	}
}

func Provider(bitwardenClient *bitwarden.Client) (*Environment, error) {
	instance := &Environment{}
	if err := loadEnvs(instance, bitwardenClient); err != nil {
		return nil, base.TagError("env.load", err)
	}

	return instance, nil
}
