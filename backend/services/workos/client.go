package workos

import (
	"context"

	"github.com/workos/workos-go/v4/pkg/mfa"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"go.uber.org/fx"

	"polimane/backend/env"
	"polimane/backend/services/jwk"
)

type ClientOptions struct {
	fx.In
	Env *env.Environment
	JWK jwk.Client
}

type Client interface {
	UserManagement() UserManagement
	MFA() MFA
	AuthenticateWithAccessToken(ctx context.Context, tokenStr string) (*AccessTokenClaims, error)
}

type Impl struct {
	userManagement UserManagement
	mfa            MFA
	env            *env.Environment
	jwk            jwk.Client
}

func Provider(options ClientOptions) Client {
	usermanagement.SetAPIKey(options.Env.WorkOS.ApiKey)
	mfa.SetAPIKey(options.Env.WorkOS.ApiKey)

	return &Impl{
		userManagement: usermanagement.DefaultClient,
		mfa:            mfa.DefaultClient,
		env:            options.Env,
		jwk:            options.JWK,
	}
}
