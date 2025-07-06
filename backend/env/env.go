package env

import "polimane/backend/base"

type Environment struct {
	SecretKey string `env:"BACKEND_SECRET_KEY,required=true"`
	AppDomain string `env:"BACKEND_APP_DOMAIN,required=true"`

	Database struct {
		URL string `env:"BACKEND_DATABASE_URL,required=true"`
	}

	Sentry struct {
		Dsn     string `env:"BACKEND_SENTRY_DSN"`
		Release string `env:"BACKEND_SENTRY_RELEASE"`
	}
}

var environment *Environment

func Init() error {
	environment = &Environment{}
	return base.TagError("env.load", loadEnvs())
}

func Env() *Environment {
	return environment
}
