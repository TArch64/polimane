package workos

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/env"
)

var (
	AccessTokenExpired = errors.New("access token expired")
)

func AuthenticateWithAccessToken(ctx context.Context, tokenStr string) (*usermanagement.User, error) {
	jwksURL := fmt.Sprintf("https://api.workos.com/sso/jwks/%s", env.Instance.WorkOS.ClientID)
	keySet, err := jwk.Fetch(ctx, jwksURL)
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

	user, err := UserManagement.GetUser(ctx, usermanagement.GetUserOpts{
		User: userID,
	})

	if err != nil {
		return nil, err
	}

	var sessionID string
	if err = token.Get("sid", &sessionID); err == nil {
		user.Metadata["SessionID"] = sessionID
	}

	return &user, nil
}
