package argon

import "github.com/matthewhartstonge/argon2"

func Hash(value string) ([]byte, error) {
	argon := argon2.DefaultConfig()
	return argon.HashEncoded([]byte(value))
}

func HashString(value string) (string, error) {
	bytes, err := Hash(value)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
