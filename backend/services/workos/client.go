package workos

import (
	"context"

	"github.com/workos/workos-go/v4/pkg/mfa"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"go.uber.org/fx"

	"polimane/backend/env"
)

type ClientOptions struct {
	fx.In
	Env *env.Environment
}

type Client interface {
	UserManagement() UserManagement
	MFA() MFA
	AuthenticateWithAccessToken(ctx context.Context, tokenStr string) (*AccessTokenClaims, error)
	AuthenticateWithRefreshToken(ctx context.Context, options *RefreshAuthOptions) (*usermanagement.RefreshAuthenticationResponse, error)
}

type Impl struct {
	userManagement UserManagement
	mfa            MFA
	env            *env.Environment
}

func Provider(options ClientOptions) Client {
	usermanagement.SetAPIKey(options.Env.WorkOS.ApiKey)
	mfa.SetAPIKey(options.Env.WorkOS.ApiKey)

	return &Impl{
		userManagement: usermanagement.DefaultClient,
		mfa:            mfa.DefaultClient,
		env:            options.Env,
	}
}
