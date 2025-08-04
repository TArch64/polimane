package auth

import (
	"context"
	"net/url"

	"github.com/maniartech/signals"
	"github.com/stretchr/testify/mock"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
	"polimane/backend/services/workos"
)

// MockUsersClient implements repositoryusers.Client for testing
type MockUsersClient struct {
	mock.Mock
}

func (m *MockUsersClient) CreateIfNeeded(ctx context.Context, userID string) (*model.User, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUsersClient) ByID(ctx context.Context, id modelbase.ID) (*model.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

// MockUserManagement implements workos.UserManagement interface for testing
type MockUserManagement struct {
	mock.Mock
}

func (m *MockUserManagement) GetAuthorizationURL(opts usermanagement.GetAuthorizationURLOpts) (*url.URL, error) {
	args := m.Called(opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*url.URL), args.Error(1)
}

func (m *MockUserManagement) AuthenticateWithCode(ctx context.Context, opts usermanagement.AuthenticateWithCodeOpts) (usermanagement.AuthenticateResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.AuthenticateResponse), args.Error(1)
}

func (m *MockUserManagement) UpdateUser(ctx context.Context, opts usermanagement.UpdateUserOpts) (usermanagement.User, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.User), args.Error(1)
}

func (m *MockUserManagement) RevokeSession(ctx context.Context, opts usermanagement.RevokeSessionOpts) error {
	args := m.Called(ctx, opts)
	return args.Error(0)
}

func (m *MockUserManagement) AuthenticateWithRefreshToken(ctx context.Context, opts usermanagement.AuthenticateWithRefreshTokenOpts) (usermanagement.RefreshAuthenticationResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.RefreshAuthenticationResponse), args.Error(1)
}

func (m *MockUserManagement) GetUser(ctx context.Context, opts usermanagement.GetUserOpts) (usermanagement.User, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.User), args.Error(1)
}

func (m *MockUserManagement) EnrollAuthFactor(ctx context.Context, opts usermanagement.EnrollAuthFactorOpts) (usermanagement.EnrollAuthFactorResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.EnrollAuthFactorResponse), args.Error(1)
}

func (m *MockUserManagement) ListAuthFactors(ctx context.Context, opts usermanagement.ListAuthFactorsOpts) (usermanagement.ListAuthFactorsResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.ListAuthFactorsResponse), args.Error(1)
}

func (m *MockUserManagement) VerifyEmail(ctx context.Context, opts usermanagement.VerifyEmailOpts) (usermanagement.UserResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.UserResponse), args.Error(1)
}

func (m *MockUserManagement) SendVerificationEmail(ctx context.Context, opts usermanagement.SendVerificationEmailOpts) (usermanagement.UserResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.UserResponse), args.Error(1)
}

func (m *MockUserManagement) CreatePasswordReset(ctx context.Context, opts usermanagement.CreatePasswordResetOpts) (usermanagement.PasswordReset, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.PasswordReset), args.Error(1)
}

func (m *MockUserManagement) GetJWKSURL(clientID string) (*url.URL, error) {
	args := m.Called(clientID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*url.URL), args.Error(1)
}

// MockWorkosClient for testing workos client functionality
type MockWorkosClient struct {
	mock.Mock
	UserManagement *MockUserManagement
}

func (m *MockWorkosClient) AuthenticateWithAccessToken(ctx context.Context, token string) (*workos.AccessTokenClaims, error) {
	args := m.Called(ctx, token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*workos.AccessTokenClaims), args.Error(1)
}

// MockSignal implements signals.Signal for testing
type MockSignal[T any] struct {
	mock.Mock
}

var _ signals.Signal[string] = (*MockSignal[string])(nil)

func (m *MockSignal[T]) Emit(ctx context.Context, payload T) {
	m.Called(ctx, payload)
}

func (m *MockSignal[T]) AddListener(handler signals.SignalListener[T], key ...string) int {
	args := m.Called(handler, key)
	return args.Int(0)
}

func (m *MockSignal[T]) RemoveListener(key string) int {
	args := m.Called(key)
	return args.Int(0)
}

func (m *MockSignal[T]) Reset() {
	m.Called()
}

func (m *MockSignal[T]) Len() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockSignal[T]) IsEmpty() bool {
	args := m.Called()
	return args.Bool(0)
}
