package workos

import (
	"context"
	"net/url"

	"github.com/workos/workos-go/v4/pkg/usermanagement"
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

func (i *Impl) UserManagement() UserManagement {
	return i.userManagement
}
