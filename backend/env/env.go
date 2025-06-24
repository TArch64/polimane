package env

type Environment struct {
	SecretKey string `env:"BACKEND_SECRET_KEY,required=true"`
	AppDomain string `env:"BACKEND_APP_DOMAIN,required=true"`

	DefaultUser struct {
		User     string `env:"BACKEND_DEFAULT_USER,required=true"`
		Password string `env:"BACKEND_DEFAULT_PASSWORD,required=true"`
	}

	Sentry struct {
		Dsn     string `env:"BACKEND_SENTRY_DSN"`
		Release string `env:"BACKEND_SENTRY_RELEASE"`
	}
}

var environment *Environment

func Init() error {
	environment = &Environment{}
	return loadEnvs()
}

func Env() *Environment {
	return environment
}
