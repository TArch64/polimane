package workos

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"net/url"
	"testing"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"polimane/backend/env"
)

// Test the AccessTokenClaims struct
func TestAccessTokenClaims(t *testing.T) {
	claims := &AccessTokenClaims{
		UserID:    "user_123",
		SessionID: "session_456",
	}

	assert.Equal(t, "user_123", claims.UserID)
	assert.Equal(t, "session_456", claims.SessionID)
}

// Test the AccessTokenExpired error
func TestAccessTokenExpiredError(t *testing.T) {
	assert.Equal(t, "access token expired", AccessTokenExpired.Error())
	assert.NotNil(t, AccessTokenExpired)
}

// Test JWT parsing error scenarios with proper mocks
func TestClient_AuthenticateWithAccessToken_JWTParsingErrors(t *testing.T) {
	ctx := context.Background()

	t.Run("JWT parse error - generic error", func(t *testing.T) {
		mockUserMgmt := &MockUserManagement{}
		mockJWK := &MockJWKClient{}

		client := &Impl{
			userManagement: mockUserMgmt,
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client-id",
					ApiKey:   "test-api-key",
				},
			},
			jwk: mockJWK,
		}

		jwksURL, err := url.Parse("https://api.workos.com/sso/jwks/test-client-id")
		require.NoError(t, err)

		mockUserMgmt.On("GetJWKSURL", "test-client-id").Return(jwksURL, nil)
		mockJWK.On("Fetch", ctx, jwksURL.String()).Return(&MockJWKSet{}, nil)

		// This will cause jwt.Parse to fail due to invalid token format
		tokenStr := "invalid.jwt.token"

		claims, err := client.AuthenticateWithAccessToken(ctx, tokenStr)

		assert.Error(t, err)
		assert.Nil(t, claims)
		// Should return the JWT parsing error, not AccessTokenExpired
		assert.NotEqual(t, AccessTokenExpired, err)
		mockUserMgmt.AssertExpectations(t)
		mockJWK.AssertExpectations(t)
	})

	t.Run("malformed JWT token", func(t *testing.T) {
		mockUserMgmt := &MockUserManagement{}
		mockJWK := &MockJWKClient{}

		client := &Impl{
			userManagement: mockUserMgmt,
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client-id",
					ApiKey:   "test-api-key",
				},
			},
			jwk: mockJWK,
		}

		jwksURL, err := url.Parse("https://api.workos.com/sso/jwks/test-client-id")
		require.NoError(t, err)

		mockUserMgmt.On("GetJWKSURL", "test-client-id").Return(jwksURL, nil)
		mockJWK.On("Fetch", ctx, jwksURL.String()).Return(&MockJWKSet{}, nil)

		// This will cause jwt.Parse to fail due to malformed JWT
		tokenStr := "not.a.jwt"

		claims, err := client.AuthenticateWithAccessToken(ctx, tokenStr)

		assert.Error(t, err)
		assert.Nil(t, claims)
		assert.NotEqual(t, AccessTokenExpired, err)
		mockUserMgmt.AssertExpectations(t)
		mockJWK.AssertExpectations(t)
	})

	t.Run("JWT parse error - token expired", func(t *testing.T) {
		mockUserMgmt := &MockUserManagement{}
		mockJWK := &MockJWKClient{}

		client := &Impl{
			userManagement: mockUserMgmt,
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client-id",
					ApiKey:   "test-api-key",
				},
			},
			jwk: mockJWK,
		}

		jwksURL, err := url.Parse("https://api.workos.com/sso/jwks/test-client-id")
		require.NoError(t, err)

		// Create a real RSA key for signing
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		require.NoError(t, err)

		// Create a JWK from the RSA key
		realKey, err := jwk.Import(privateKey)
		require.NoError(t, err)
		require.NoError(t, realKey.Set(jwk.KeyIDKey, "test-key-id"))
		require.NoError(t, realKey.Set(jwk.AlgorithmKey, jwa.RS256()))

		// Create a JWK set containing our key
		realKeySet := jwk.NewSet()
		require.NoError(t, realKeySet.AddKey(realKey))

		mockUserMgmt.On("GetJWKSURL", "test-client-id").Return(jwksURL, nil)
		mockJWK.On("Fetch", ctx, jwksURL.String()).Return(realKeySet, nil)

		// Create an expired token using the JWT library
		token := jwt.New()
		require.NoError(t, token.Set(jwt.SubjectKey, "user_123"))
		require.NoError(t, token.Set("sid", "session_456"))
		require.NoError(t, token.Set(jwt.ExpirationKey, time.Now().Add(-1*time.Hour))) // Expired 1 hour ago
		require.NoError(t, token.Set(jwt.IssuedAtKey, time.Now().Add(-2*time.Hour)))

		// Sign the token with our private key
		tokenBytes, err := jwt.Sign(token, jwt.WithKey(jwa.RS256(), realKey))
		require.NoError(t, err)

		claims, err := client.AuthenticateWithAccessToken(ctx, string(tokenBytes))

		assert.Error(t, err)
		assert.Nil(t, claims)
		// This tests the specific error handling path in lines 127-129
		// The jwt.Parse should detect the expired token and trigger the "token is expired" check
		assert.Equal(t, AccessTokenExpired, err)
		mockUserMgmt.AssertExpectations(t)
		mockJWK.AssertExpectations(t)
	})
}

// Test successful authentication path
func TestClient_AuthenticateWithAccessToken_Success(t *testing.T) {
	ctx := context.Background()
	mockUserMgmt := &MockUserManagement{}
	mockJWK := &MockJWKClient{}

	client := &Impl{
		userManagement: mockUserMgmt,
		env: &env.Environment{
			WorkOS: struct {
				ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
				ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
			}{
				ClientID: "test-client-id",
				ApiKey:   "test-api-key",
			},
		},
		jwk: mockJWK,
	}

	jwksURL, err := url.Parse("https://api.workos.com/sso/jwks/test-client-id")
	require.NoError(t, err)

	// Create a real RSA key for signing
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	// Create a JWK from the RSA key
	realKey, err := jwk.Import(privateKey)
	require.NoError(t, err)
	require.NoError(t, realKey.Set(jwk.KeyIDKey, "test-key-id"))
	require.NoError(t, realKey.Set(jwk.AlgorithmKey, jwa.RS256()))

	// Create a JWK set containing our key
	realKeySet := jwk.NewSet()
	require.NoError(t, realKeySet.AddKey(realKey))

	mockUserMgmt.On("GetJWKSURL", "test-client-id").Return(jwksURL, nil)
	mockJWK.On("Fetch", ctx, jwksURL.String()).Return(realKeySet, nil)

	// Create a valid token with proper claims
	token := jwt.New()
	require.NoError(t, token.Set(jwt.SubjectKey, "user_123"))
	require.NoError(t, token.Set("sid", "session_456"))
	require.NoError(t, token.Set(jwt.ExpirationKey, time.Now().Add(1*time.Hour))) // Valid for 1 hour
	require.NoError(t, token.Set(jwt.IssuedAtKey, time.Now()))

	// Sign the token with our private key
	tokenBytes, err := jwt.Sign(token, jwt.WithKey(jwa.RS256(), realKey))
	require.NoError(t, err)

	claims, err := client.AuthenticateWithAccessToken(ctx, string(tokenBytes))

	// This tests the success path in lines 134-144
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, "user_123", claims.UserID)
	assert.Equal(t, "session_456", claims.SessionID)
	mockUserMgmt.AssertExpectations(t)
	mockJWK.AssertExpectations(t)
}

// Test session ID extraction error
func TestClient_AuthenticateWithAccessToken_SessionIDError(t *testing.T) {
	ctx := context.Background()
	mockUserMgmt := &MockUserManagement{}
	mockJWK := &MockJWKClient{}

	client := &Impl{
		userManagement: mockUserMgmt,
		env: &env.Environment{
			WorkOS: struct {
				ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
				ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
			}{
				ClientID: "test-client-id",
				ApiKey:   "test-api-key",
			},
		},
		jwk: mockJWK,
	}

	jwksURL, err := url.Parse("https://api.workos.com/sso/jwks/test-client-id")
	require.NoError(t, err)

	// Create a real RSA key for signing
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	// Create a JWK from the RSA key
	realKey, err := jwk.Import(privateKey)
	require.NoError(t, err)
	require.NoError(t, realKey.Set(jwk.KeyIDKey, "test-key-id"))
	require.NoError(t, realKey.Set(jwk.AlgorithmKey, jwa.RS256()))

	// Create a JWK set containing our key
	realKeySet := jwk.NewSet()
	require.NoError(t, realKeySet.AddKey(realKey))

	mockUserMgmt.On("GetJWKSURL", "test-client-id").Return(jwksURL, nil)
	mockJWK.On("Fetch", ctx, jwksURL.String()).Return(realKeySet, nil)

	// Create a valid token WITHOUT the "sid" claim to trigger the session ID extraction error
	token := jwt.New()
	require.NoError(t, token.Set(jwt.SubjectKey, "user_123"))
	// Intentionally NOT setting the "sid" claim to test the error path
	require.NoError(t, token.Set(jwt.ExpirationKey, time.Now().Add(1*time.Hour)))
	require.NoError(t, token.Set(jwt.IssuedAtKey, time.Now()))

	// Sign the token with our private key
	tokenBytes, err := jwt.Sign(token, jwt.WithKey(jwa.RS256(), realKey))
	require.NoError(t, err)

	claims, err := client.AuthenticateWithAccessToken(ctx, string(tokenBytes))

	// This tests the session ID extraction error path in lines 137-139
	assert.Error(t, err)
	assert.Nil(t, claims)
	// The error should occur when trying to extract the "sid" claim
	assert.Contains(t, err.Error(), "sid")
	mockUserMgmt.AssertExpectations(t)
	mockJWK.AssertExpectations(t)
}

// Test empty token string handling
func TestClient_AuthenticateWithAccessToken_EmptyToken(t *testing.T) {
	ctx := context.Background()
	mockUserMgmt := &MockUserManagement{}
	mockJWK := &MockJWKClient{}

	client := &Impl{
		userManagement: mockUserMgmt,
		env: &env.Environment{
			WorkOS: struct {
				ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
				ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
			}{
				ClientID: "test-client-id",
				ApiKey:   "test-api-key",
			},
		},
		jwk: mockJWK,
	}

	jwksURL, err := url.Parse("https://api.workos.com/sso/jwks/test-client-id")
	require.NoError(t, err)

	mockUserMgmt.On("GetJWKSURL", "test-client-id").Return(jwksURL, nil)
	mockJWK.On("Fetch", ctx, jwksURL.String()).Return(&MockJWKSet{}, nil)

	claims, err := client.AuthenticateWithAccessToken(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, claims)
	mockUserMgmt.AssertExpectations(t)
	mockJWK.AssertExpectations(t)
}

// Test error handling scenarios with mocked dependencies
func TestClient_AuthenticateWithAccessToken_ErrorHandling(t *testing.T) {
	ctx := context.Background()
	tokenStr := "test.jwt.token"

	t.Run("GetJWKSURL error", func(t *testing.T) {
		mockUserMgmt := &MockUserManagement{}
		mockJWK := &MockJWKClient{}

		client := &Impl{
			userManagement: mockUserMgmt,
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client-id",
					ApiKey:   "test-api-key",
				},
			},
			jwk: mockJWK,
		}

		mockUserMgmt.On("GetJWKSURL", "test-client-id").Return(nil, assert.AnError)

		claims, err := client.AuthenticateWithAccessToken(ctx, tokenStr)

		assert.Error(t, err)
		assert.Nil(t, claims)
		assert.Equal(t, assert.AnError, err)
		mockUserMgmt.AssertExpectations(t)
	})

	t.Run("JWK fetch error", func(t *testing.T) {
		mockUserMgmt := &MockUserManagement{}
		mockJWK := &MockJWKClient{}

		client := &Impl{
			userManagement: mockUserMgmt,
			env: &env.Environment{
				WorkOS: struct {
					ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
					ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
				}{
					ClientID: "test-client-id",
					ApiKey:   "test-api-key",
				},
			},
			jwk: mockJWK,
		}

		jwksURL, err := url.Parse("https://api.workos.com/sso/jwks/test-client-id")
		require.NoError(t, err)

		mockUserMgmt.On("GetJWKSURL", "test-client-id").Return(jwksURL, nil)
		mockJWK.On("Fetch", ctx, jwksURL.String()).Return(nil, assert.AnError)

		claims, err := client.AuthenticateWithAccessToken(ctx, tokenStr)

		assert.Error(t, err)
		assert.Nil(t, claims)
		assert.Equal(t, assert.AnError, err)
		mockUserMgmt.AssertExpectations(t)
		mockJWK.AssertExpectations(t)
	})
}

// Test client configuration requirements
func TestClient_AuthenticateWithAccessToken_Requirements(t *testing.T) {
	tests := []struct {
		name        string
		client      *Impl
		expectPanic bool
	}{
		{
			name: "client with nil UserManagement",
			client: &Impl{
				userManagement: nil,
				env: &env.Environment{
					WorkOS: struct {
						ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
						ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
					}{
						ClientID: "test-client-id",
						ApiKey:   "test-api-key",
					},
				},
				jwk: &MockJWKClient{},
			},
			expectPanic: true,
		},
		{
			name: "client with nil environment",
			client: &Impl{
				env: nil,
				jwk: &MockJWKClient{},
			},
			expectPanic: true,
		},
		{
			name: "client with nil JWK interface",
			client: &Impl{
				userManagement: &MockUserManagement{},
				env: &env.Environment{
					WorkOS: struct {
						ClientID string `env:"BACKEND_WORKOS_CLIENT_ID,required=true"`
						ApiKey   string `env:"BACKEND_WORKOS_API_KEY,required=true"`
					}{
						ClientID: "test-client-id",
						ApiKey:   "test-api-key",
					},
				},
				jwk: nil,
			},
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			tokenStr := "test.jwt.token"

			if tt.expectPanic {
				assert.Panics(t, func() {
					_, _ = tt.client.AuthenticateWithAccessToken(ctx, tokenStr)
				})
			} else {
				// Should not panic, but may error due to invalid token/config
				_, err := tt.client.AuthenticateWithAccessToken(ctx, tokenStr)
				// We expect an error since these are not real valid configs
				assert.Error(t, err)
			}
		})
	}
}
