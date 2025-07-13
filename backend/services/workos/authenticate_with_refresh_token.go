package workos

import (
	"context"
	"fmt"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/env"
)

func AuthenticateWithRefreshToken(ctx context.Context, tokenStr string) (*usermanagement.User, error) {
	jwksURL := fmt.Sprintf("https://api.workos.com/sso/jwks/%s", env.Instance.WorkOS.ClientID)
	keySet, err := jwk.Fetch(ctx, jwksURL)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(
		[]byte(tokenStr),
		jwt.WithKeySet(keySet),
		jwt.WithValidate(true),
	)

	if err != nil {
		return nil, err
	}

	userID, _ := token.Subject()
	user, err := UserManagement.GetUser(ctx, usermanagement.GetUserOpts{
		User: userID,
	})

	return &user, err
}
