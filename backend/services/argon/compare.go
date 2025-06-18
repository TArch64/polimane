package argon

import (
	"crypto/subtle"
)

func Compare(raw, hash string) bool {
	checkingBytes, err := Hash(raw)
	if err != nil {
		checkingBytes = make([]byte, len(hash))
	}

	return subtle.ConstantTimeCompare(checkingBytes, []byte(hash)) == 1
}
