package workos

import (
	"context"
	"errors"
	"strings"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

var (
	AccessTokenExpiredErr = errors.New("access token expired")
)

type AccessTokenClaims struct {
	UserID    string
	SessionID string
}

func (i *Client) AuthenticateWithAccessToken(ctx context.Context, tokenStr string) (*AccessTokenClaims, error) {
	jwksURL, err := i.UserManagement.GetJWKSURL(i.env.WorkOS.ClientID)
	if err != nil {
		return nil, err
	}

	keySet, err := jwk.Fetch(ctx, jwksURL.String())
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(
		[]byte(tokenStr),
		jwt.WithContext(ctx),
		jwt.WithKeySet(keySet),
		jwt.WithValidate(true),
	)

	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, AccessTokenExpiredErr
		}

		return nil, err
	}

	userID, _ := token.Subject()

	var sessionID string
	if err = token.Get("sid", &sessionID); err != nil {
		return nil, err
	}

	return &AccessTokenClaims{
		UserID:    userID,
		SessionID: sessionID,
	}, nil
}
