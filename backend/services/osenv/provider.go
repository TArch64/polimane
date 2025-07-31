package osenv

import "os"

type Env interface {
	Getenv(key string) string
	Setenv(key, value string) error
}

type envImpl struct{}

var _ Env = (*envImpl)(nil)

func (e *envImpl) Getenv(key string) string {
	return os.Getenv(key)
}

func (e *envImpl) Setenv(key, value string) error {
	return os.Setenv(key, value)
}

func Provider() Env {
	return &envImpl{}
}
