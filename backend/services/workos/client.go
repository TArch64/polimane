package workos

import (
	"context"
	"net/url"

	"github.com/workos/workos-go/v4/pkg/mfa"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"go.uber.org/fx"

	"polimane/backend/env"
	"polimane/backend/services/jwk"
)

type UserManagement interface {
	GetAuthorizationURL(opts usermanagement.GetAuthorizationURLOpts) (*url.URL, error)
	AuthenticateWithCode(ctx context.Context, opts usermanagement.AuthenticateWithCodeOpts) (usermanagement.AuthenticateResponse, error)
	UpdateUser(ctx context.Context, opts usermanagement.UpdateUserOpts) (usermanagement.User, error)
	RevokeSession(ctx context.Context, opts usermanagement.RevokeSessionOpts) error
	AuthenticateWithRefreshToken(ctx context.Context, opts usermanagement.AuthenticateWithRefreshTokenOpts) (usermanagement.RefreshAuthenticationResponse, error)
	GetUser(ctx context.Context, opts usermanagement.GetUserOpts) (usermanagement.User, error)
	EnrollAuthFactor(ctx context.Context, opts usermanagement.EnrollAuthFactorOpts) (usermanagement.EnrollAuthFactorResponse, error)
	ListAuthFactors(ctx context.Context, opts usermanagement.ListAuthFactorsOpts) (usermanagement.ListAuthFactorsResponse, error)
	VerifyEmail(ctx context.Context, opts usermanagement.VerifyEmailOpts) (usermanagement.UserResponse, error)
	SendVerificationEmail(ctx context.Context, opts usermanagement.SendVerificationEmailOpts) (usermanagement.UserResponse, error)
	CreatePasswordReset(ctx context.Context, opts usermanagement.CreatePasswordResetOpts) (usermanagement.PasswordReset, error)
	GetJWKSURL(clientID string) (*url.URL, error)
}

type MFA interface {
	DeleteFactor(ctx context.Context, opts mfa.DeleteFactorOpts) error
	VerifyChallenge(ctx context.Context, opts mfa.VerifyChallengeOpts) (mfa.VerifyChallengeResponse, error)
	GetFactor(ctx context.Context, opts mfa.GetFactorOpts) (mfa.Factor, error)
}

type ClientOptions struct {
	fx.In
	Env *env.Environment
	JWK jwk.Client
}

type Client struct {
	UserManagement UserManagement
	MFA            MFA
	env            *env.Environment
	jwk            jwk.Client
}

func Provider(options ClientOptions) *Client {
	usermanagement.SetAPIKey(options.Env.WorkOS.ApiKey)
	mfa.SetAPIKey(options.Env.WorkOS.ApiKey)

	return &Client{
		UserManagement: usermanagement.DefaultClient,
		MFA:            mfa.DefaultClient,
		env:            options.Env,
		jwk:            options.JWK,
	}
}
