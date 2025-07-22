package workos

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/lestrrat-go/jwx/v3/jwt"
)

var (
	AccessTokenExpired = errors.New("access token expired")
)

type AccessTokenClaims struct {
	UserID    string
	SessionID string
}

func (c *Client) AuthenticateWithAccessToken(ctx context.Context, tokenStr string) (*AccessTokenClaims, error) {
	jwksURL := fmt.Sprintf("https://api.workos.com/sso/jwks/%s", c.env.WorkOS.ClientID)
	keySet, err := c.jwk.Fetch(ctx, jwksURL)
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
			return nil, AccessTokenExpired
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
