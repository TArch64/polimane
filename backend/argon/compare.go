package argon

import "github.com/matthewhartstonge/argon2"

func Compare(raw, hash string) bool {
	equal, err := argon2.VerifyEncoded([]byte(raw), []byte(hash))
	return err == nil && equal
}
