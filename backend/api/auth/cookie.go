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

func newCookieToken(user *model.User, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		UserID: user.ID,
	})

	return token.SignedString([]byte(env.Env().SecretKey))
}

func parseCookieToken(token string) (*tokenClaims, error) {
	claims := &tokenClaims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.Env().SecretKey), nil
	})

	return claims, err
}
