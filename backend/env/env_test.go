package env

import (
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEnvironment(t *testing.T) {
	t.Run("AppURL constructs correct URL", func(t *testing.T) {
		env := &Environment{
			AppProtocol: "https",
			AppDomain:   "example.com",
		}

		expectedURL := &url.URL{
			Scheme: "https",
			Host:   "example.com",
		}

		result := env.AppURL()
		assert.Equal(t, expectedURL, result)
	})

	t.Run("ApiURL constructs correct API URL", func(t *testing.T) {
		env := &Environment{
			AppProtocol: "https",
			AppDomain:   "example.com",
		}

		expectedURL := &url.URL{
			Scheme: "https",
			Host:   "api.example.com",
		}

		result := env.ApiURL()
		assert.Equal(t, expectedURL, result)
	})

	t.Run("AppURL with http protocol", func(t *testing.T) {
		env := &Environment{
			AppProtocol: "http",
			AppDomain:   "localhost:3000",
		}

		expectedURL := &url.URL{
			Scheme: "http",
			Host:   "localhost:3000",
		}

		result := env.AppURL()
		assert.Equal(t, expectedURL, result)
	})

	t.Run("ApiURL with subdomain", func(t *testing.T) {
		env := &Environment{
			AppProtocol: "https",
			AppDomain:   "staging.example.com",
		}

		expectedURL := &url.URL{
			Scheme: "https",
			Host:   "api.staging.example.com",
		}

		result := env.ApiURL()
		assert.Equal(t, expectedURL, result)
	})

	t.Run("zero value environment", func(t *testing.T) {
		var env Environment

		appURL := env.AppURL()
		apiURL := env.ApiURL()

		assert.Equal(t, "", appURL.Scheme)
		assert.Equal(t, "", appURL.Host)
		assert.Equal(t, "", apiURL.Scheme)
		assert.Equal(t, "api.", apiURL.Host)
	})

	t.Run("struct fields have correct types", func(t *testing.T) {
		env := Environment{
			SecretKey:   "test-secret",
			AppDomain:   "test.com",
			AppProtocol: "https",
		}

		env.Database.URL = "postgres://test"
		env.Sentry.Dsn = "https://sentry.test"
		env.Sentry.Release = "v1.0.0"
		env.WorkOS.ClientID = "client-123"
		env.WorkOS.ApiKey = "api-key-456"

		assert.Equal(t, "test-secret", env.SecretKey)
		assert.Equal(t, "test.com", env.AppDomain)
		assert.Equal(t, "https", env.AppProtocol)
		assert.Equal(t, "postgres://test", env.Database.URL)
		assert.Equal(t, "https://sentry.test", env.Sentry.Dsn)
		assert.Equal(t, "v1.0.0", env.Sentry.Release)
		assert.Equal(t, "client-123", env.WorkOS.ClientID)
		assert.Equal(t, "api-key-456", env.WorkOS.ApiKey)
	})
}

func TestProvider(t *testing.T) {
	t.Run("successfully creates environment when loadEnvs succeeds", func(t *testing.T) {
		// Arrange
		mockClient := &MockBitwardenClient{}

		// Mock successful loadEnvs call
		mockClient.On("LoadToEnviron", mock.MatchedBy(func(names []string) bool {
			expectedNames := []string{
				"BACKEND_SECRET_KEY",
				"BACKEND_SENTRY_DSN",
				"BACKEND_DATABASE_URL",
				"BACKEND_WORKOS_CLIENT_ID",
				"BACKEND_WORKOS_API_KEY",
			}
			return len(names) == len(expectedNames)
		})).Return(nil)

		mockClient.On("DownloadCerts", mock.AnythingOfType("[]*bitwarden.DownloadingCert")).Return(nil)

		options := Options{
			BitwardenClient: mockClient,
		}

		// Act
		env, err := Provider(options)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, env)
		assert.IsType(t, &Environment{}, env)
		mockClient.AssertExpectations(t)
	})

	t.Run("returns error when loadEnvs fails", func(t *testing.T) {
		// Arrange
		mockClient := &MockBitwardenClient{}
		expectedError := errors.New("bitwarden load failed")

		// Mock failed loadEnvs call
		mockClient.On("LoadToEnviron", mock.AnythingOfType("[]string")).Return(expectedError)

		options := Options{
			BitwardenClient: mockClient,
		}

		// Act
		env, err := Provider(options)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, env)
		assert.Contains(t, err.Error(), "env.load")
		assert.Contains(t, err.Error(), "bitwarden load failed")
		mockClient.AssertExpectations(t)
	})

	t.Run("returns error when DownloadCerts fails", func(t *testing.T) {
		// Arrange
		mockClient := &MockBitwardenClient{}
		expectedError := errors.New("cert download failed")

		// Mock successful LoadToEnviron but failed DownloadCerts
		mockClient.On("LoadToEnviron", mock.AnythingOfType("[]string")).Return(nil)
		mockClient.On("DownloadCerts", mock.AnythingOfType("[]*bitwarden.DownloadingCert")).Return(expectedError)

		options := Options{
			BitwardenClient: mockClient,
		}

		// Act
		env, err := Provider(options)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, env)
		assert.Contains(t, err.Error(), "env.load.bitwarden.certs")
		assert.Contains(t, err.Error(), "cert download failed")
		mockClient.AssertExpectations(t)
	})
}
