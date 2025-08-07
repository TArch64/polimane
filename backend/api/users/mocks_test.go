package users

import (
	"context"
	"net/url"

	"github.com/stretchr/testify/mock"
	"github.com/workos/workos-go/v4/pkg/mfa"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/model"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
)

// MockWorkosUserManagement mocks the WorkOS UserManagement interface
type MockWorkosUserManagement struct {
	mock.Mock
}

func (m *MockWorkosUserManagement) GetAuthorizationURL(opts usermanagement.GetAuthorizationURLOpts) (*url.URL, error) {
	args := m.Called(opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*url.URL), args.Error(1)
}

func (m *MockWorkosUserManagement) AuthenticateWithCode(ctx context.Context, opts usermanagement.AuthenticateWithCodeOpts) (usermanagement.AuthenticateResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.AuthenticateResponse), args.Error(1)
}

func (m *MockWorkosUserManagement) UpdateUser(ctx context.Context, opts usermanagement.UpdateUserOpts) (usermanagement.User, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.User), args.Error(1)
}

func (m *MockWorkosUserManagement) RevokeSession(ctx context.Context, opts usermanagement.RevokeSessionOpts) error {
	args := m.Called(ctx, opts)
	return args.Error(0)
}

func (m *MockWorkosUserManagement) AuthenticateWithRefreshToken(ctx context.Context, opts usermanagement.AuthenticateWithRefreshTokenOpts) (usermanagement.RefreshAuthenticationResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.RefreshAuthenticationResponse), args.Error(1)
}

func (m *MockWorkosUserManagement) GetUser(ctx context.Context, opts usermanagement.GetUserOpts) (usermanagement.User, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.User), args.Error(1)
}

func (m *MockWorkosUserManagement) EnrollAuthFactor(ctx context.Context, opts usermanagement.EnrollAuthFactorOpts) (usermanagement.EnrollAuthFactorResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.EnrollAuthFactorResponse), args.Error(1)
}

func (m *MockWorkosUserManagement) ListAuthFactors(ctx context.Context, opts usermanagement.ListAuthFactorsOpts) (usermanagement.ListAuthFactorsResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.ListAuthFactorsResponse), args.Error(1)
}

func (m *MockWorkosUserManagement) VerifyEmail(ctx context.Context, opts usermanagement.VerifyEmailOpts) (usermanagement.UserResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.UserResponse), args.Error(1)
}

func (m *MockWorkosUserManagement) SendVerificationEmail(ctx context.Context, opts usermanagement.SendVerificationEmailOpts) (usermanagement.UserResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.UserResponse), args.Error(1)
}

func (m *MockWorkosUserManagement) CreatePasswordReset(ctx context.Context, opts usermanagement.CreatePasswordResetOpts) (usermanagement.PasswordReset, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(usermanagement.PasswordReset), args.Error(1)
}

func (m *MockWorkosUserManagement) GetJWKSURL(clientID string) (*url.URL, error) {
	args := m.Called(clientID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*url.URL), args.Error(1)
}

// MockWorkosMFA mocks the WorkOS MFA interface
type MockWorkosMFA struct {
	mock.Mock
}

func (m *MockWorkosMFA) DeleteFactor(ctx context.Context, opts mfa.DeleteFactorOpts) error {
	args := m.Called(ctx, opts)
	return args.Error(0)
}

func (m *MockWorkosMFA) VerifyChallenge(ctx context.Context, opts mfa.VerifyChallengeOpts) (mfa.VerifyChallengeResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(mfa.VerifyChallengeResponse), args.Error(1)
}

func (m *MockWorkosMFA) GetFactor(ctx context.Context, opts mfa.GetFactorOpts) (mfa.Factor, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(mfa.Factor), args.Error(1)
}

// MockWorkosClient mocks the WorkOS client
type MockWorkosClient struct {
	mock.Mock
	userManagement workos.UserManagement
	mfa            workos.MFA
}

func (m *MockWorkosClient) UserManagement() workos.UserManagement {
	return m.userManagement
}

func (m *MockWorkosClient) MFA() workos.MFA {
	return m.mfa
}

func (m *MockWorkosClient) AuthenticateWithAccessToken(_ context.Context, _ string) (*workos.AccessTokenClaims, error) {
	return nil, nil
}

func NewMockWorkosClient() workos.Client {
	return &MockWorkosClient{
		userManagement: &MockWorkosUserManagement{},
		mfa:            &MockWorkosMFA{},
	}
}

// NewMockSignalsContainer creates a mock signals container for testing
func NewMockSignalsContainer() *signal.Container {
	return signal.Provider()
}

// Helper function to create test user
func createTestUser() *model.User {
	return &model.User{
		Identifiable: &model.Identifiable{
			ID: model.MustStringToID("550e8400-e29b-41d4-a716-446655440000"),
		},
		WorkosID: "workos_user_123",
	}
}

// Helper function to create test WorkOS user
func createTestWorkosUser() *usermanagement.User {
	return &usermanagement.User{
		ID:                "workos_user_123",
		Email:             "test@example.com",
		FirstName:         "Test",
		LastName:          "User",
		EmailVerified:     true,
		ProfilePictureURL: "https://example.com/profile.jpg",
	}
}
