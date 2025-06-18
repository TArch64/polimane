package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"polimane/backend/env"
	"polimane/backend/model"
)

const (
	cookieName = "pa"
)

func newTokenExpiresAt() time.Time {
	return time.Now().Add(time.Hour * 24 * 30)
}

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID model.ID `json:"userId"`
}

func newAuthToken(user *model.User, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		UserID: user.ID,
	})

	return token.SignedString([]byte(env.Env().SecretKey))
}
