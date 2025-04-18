package env

import "github.com/Netflix/go-env"

type Environment struct {
	SecretKey      string `env:"BACKEND_SECRET_KEY"`
	FrontendOrigin string `env:"BACKEND_FRONTEND_ORIGIN"`

	DefaultUser struct {
		User     string `env:"BACKEND_DEFAULT_USER"`
		Password string `env:"BACKEND_DEFAULT_PASSWORD"`
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
