package argon

import "github.com/matthewhartstonge/argon2"

func Hash(value string) (string, error) {
	argon := argon2.DefaultConfig()

	bytes, err := argon.HashEncoded([]byte(value))
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
