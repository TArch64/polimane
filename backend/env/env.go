package env

import "github.com/Netflix/go-env"

type Environment struct {
	SecretKey string `env:"BACKEND_SECRET_KEY,required=true"`
	AppDomain string `env:"BACKEND_APP_DOMAIN,required=true"`

	DefaultUser struct {
		User     string `env:"BACKEND_DEFAULT_USER,required=true"`
		Password string `env:"BACKEND_DEFAULT_PASSWORD,required=true"`
	}

	Sentry struct {
		Dsn string `env:"BACKEND_SENTRY_DSN"`
	}
}

var environment *Environment

func Init() error {
	environment = &Environment{}
	_, err := env.UnmarshalFromEnviron(environment)
	return err
}

func Env() *Environment {
	return environment
}
