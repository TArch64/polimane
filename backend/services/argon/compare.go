package argon

import (
	"crypto/subtle"

	"github.com/matthewhartstonge/argon2"
)

func Compare(raw, hash string) bool {
	hashBytes := []byte(hash)
	equal, err := argon2.VerifyEncoded([]byte(raw), hashBytes)

	if err != nil {
		return subtle.ConstantTimeCompare(make([]byte, len(hash)), hashBytes) == 1
	}

	return equal
}
