package workos

import (
	"context"
	"errors"
	"strings"

	"github.com/workos/workos-go/v4/pkg/usermanagement"
)

var (
	SessionEndedErr = errors.New("session has ended")
)

type RefreshAuthOptions struct {
	Token     string
	UserAgent string
}

func (i *Impl) AuthenticateWithRefreshToken(ctx context.Context, options *RefreshAuthOptions) (*usermanagement.RefreshAuthenticationResponse, error) {
	res, err := i.userManagement.AuthenticateWithRefreshToken(ctx, usermanagement.AuthenticateWithRefreshTokenOpts{
		ClientID:     i.env.WorkOS.ClientID,
		RefreshToken: options.Token,
		UserAgent:    options.UserAgent,
	})

	if err != nil {
		if strings.Contains(err.Error(), "invalid_grant") {
			return nil, SessionEndedErr
		}
		return nil, err
	}

	return &res, nil
}
