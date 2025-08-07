package workos

import (
	"context"
	"crypto"
	"net/url"

	"github.com/lestrrat-go/jwx/v3/cert"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/stretchr/testify/mock"
	"github.com/workos/workos-go/v4/pkg/mfa"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
)

// MockUserManagement is a mock implementation of the UserManagement interface
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

// MockMFA is a mock implementation of the MFA interface
type MockMFA struct {
	mock.Mock
}

func (m *MockMFA) DeleteFactor(ctx context.Context, opts mfa.DeleteFactorOpts) error {
	args := m.Called(ctx, opts)
	return args.Error(0)
}

func (m *MockMFA) VerifyChallenge(ctx context.Context, opts mfa.VerifyChallengeOpts) (mfa.VerifyChallengeResponse, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(mfa.VerifyChallengeResponse), args.Error(1)
}

func (m *MockMFA) GetFactor(ctx context.Context, opts mfa.GetFactorOpts) (mfa.Factor, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(mfa.Factor), args.Error(1)
}

// MockJWKClient is a mock implementation of the JWK client
type MockJWKClient struct {
	mock.Mock
}

func (m *MockJWKClient) Fetch(ctx context.Context, urlStr string) (jwk.Set, error) {
	args := m.Called(ctx, urlStr)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(jwk.Set), args.Error(1)
}

// MockJWKSet is a mock implementation of jwk.Set
type MockJWKSet struct {
	mock.Mock
}

// AddKey adds a key to the set
func (m *MockJWKSet) AddKey(key jwk.Key) error {
	args := m.Called(key)
	return args.Error(0)
}

// Clear removes all keys from the set
func (m *MockJWKSet) Clear() error {
	args := m.Called()
	return args.Error(0)
}

// Key returns the key at the specified index
func (m *MockJWKSet) Key(idx int) (jwk.Key, bool) {
	args := m.Called(idx)
	if args.Get(0) == nil {
		return nil, args.Bool(1)
	}
	return args.Get(0).(jwk.Key), args.Bool(1)
}

// Get retrieves a field value from the set
func (m *MockJWKSet) Get(key string, value any) error {
	args := m.Called(key, value)
	return args.Error(0)
}

// Set assigns a field value to the set
func (m *MockJWKSet) Set(key string, value any) error {
	args := m.Called(key, value)
	return args.Error(0)
}

// Remove removes a field from the set
func (m *MockJWKSet) Remove(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

// Index returns the index of the specified key
func (m *MockJWKSet) Index(key jwk.Key) int {
	args := m.Called(key)
	return args.Int(0)
}

// Len returns the number of keys in the set
func (m *MockJWKSet) Len() int {
	args := m.Called()
	return args.Int(0)
}

// LookupKeyID looks up a key by its ID
func (m *MockJWKSet) LookupKeyID(keyID string) (jwk.Key, bool) {
	args := m.Called(keyID)
	if args.Get(0) == nil {
		return nil, args.Bool(1)
	}
	return args.Get(0).(jwk.Key), args.Bool(1)
}

// RemoveKey removes a specific key from the set
func (m *MockJWKSet) RemoveKey(key jwk.Key) error {
	args := m.Called(key)
	return args.Error(0)
}

// Keys returns the list of key field names
func (m *MockJWKSet) Keys() []string {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).([]string)
}

// Clone creates a copy of the set
func (m *MockJWKSet) Clone() (jwk.Set, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(jwk.Set), args.Error(1)
}

// MockKey is a mock implementation of jwk.Key
type MockKey struct {
	mock.Mock
}

func (m *MockKey) Has(key string) bool {
	args := m.Called(key)
	return args.Bool(0)
}

func (m *MockKey) Get(key string, value any) error {
	args := m.Called(key, value)
	return args.Error(0)
}

func (m *MockKey) Set(key string, value any) error {
	args := m.Called(key, value)
	return args.Error(0)
}

func (m *MockKey) Remove(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *MockKey) Validate() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockKey) Thumbprint(hash crypto.Hash) ([]byte, error) {
	args := m.Called(hash)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockKey) Keys() []string {
	args := m.Called()
	return args.Get(0).([]string)
}

func (m *MockKey) Clone() (jwk.Key, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(jwk.Key), args.Error(1)
}

func (m *MockKey) PublicKey() (jwk.Key, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(jwk.Key), args.Error(1)
}

func (m *MockKey) KeyType() jwa.KeyType {
	args := m.Called()
	return args.Get(0).(jwa.KeyType)
}

func (m *MockKey) KeyUsage() (string, bool) {
	args := m.Called()
	return args.String(0), args.Bool(1)
}

func (m *MockKey) KeyOps() (jwk.KeyOperationList, bool) {
	args := m.Called()
	return args.Get(0).(jwk.KeyOperationList), args.Bool(1)
}

func (m *MockKey) Algorithm() (jwa.KeyAlgorithm, bool) {
	args := m.Called()
	return args.Get(0).(jwa.KeyAlgorithm), args.Bool(1)
}

func (m *MockKey) KeyID() (string, bool) {
	args := m.Called()
	return args.String(0), args.Bool(1)
}

func (m *MockKey) X509URL() (string, bool) {
	args := m.Called()
	return args.String(0), args.Bool(1)
}

func (m *MockKey) X509CertChain() (*cert.Chain, bool) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Bool(1)
	}
	return args.Get(0).(*cert.Chain), args.Bool(1)
}

func (m *MockKey) X509CertThumbprint() (string, bool) {
	args := m.Called()
	return args.String(0), args.Bool(1)
}

func (m *MockKey) X509CertThumbprintS256() (string, bool) {
	args := m.Called()
	return args.String(0), args.Bool(1)
}
