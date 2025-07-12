package auth

import (
	"github.com/golang-jwt/jwt/v5"

	"polimane/backend/env"
)

type loginStateClaims struct {
	jwt.RegisteredClaims
	ReturnTo string `json:"returnTo"`
}

func newLoginState(query *loginQueryParams) (string, error) {
	claims := &loginStateClaims{
		RegisteredClaims: jwt.RegisteredClaims{},
		ReturnTo:         query.ReturnTo,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(env.Instance.SecretKey))
}

func parseLoginState(state string) (*loginStateClaims, error) {
	var claims loginStateClaims

	_, err := jwt.ParseWithClaims(state, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.Instance.SecretKey), nil
	})

	return &claims, err
}
