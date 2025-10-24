package env

import (
	"net/url"

	"go.uber.org/fx"

	"polimane/backend/base"
	"polimane/backend/services/bitwarden"
)

type Environment struct {
	SecretKey   string `env:"BACKEND_SECRET_KEY,required=true"`
	AppDomain   string `env:"BACKEND_APP_DOMAIN,required=true"`
	AppProtocol string `env:"BACKEND_APP_PROTOCOL,required=true"`
	AppURL      *url.URL
	ApiURL      *url.URL

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
		SQSBaseURL      string `env:"BACKEND_SQS_BASE_URL,required=true"`
	}
}

type Options struct {
	fx.In
	BitWardenClient *bitwarden.Client
}

func Provider(options Options) (*Environment, error) {
	env := &Environment{}
	if err := loadEnvs(env, options.BitWardenClient); err != nil {
		return nil, base.TagError("env.load", err)
	}

	env.AppURL = &url.URL{
		Scheme: env.AppProtocol,
		Host:   "app." + env.AppDomain,
	}

	env.ApiURL = &url.URL{
		Scheme: env.AppProtocol,
		Host:   "api." + env.AppDomain,
	}

	return env, nil
}
