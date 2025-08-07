package env

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"polimane/backend/services/bitwarden"
)

func TestLoadEnvs(t *testing.T) {
	t.Run("successfully loads environment variables", func(t *testing.T) {
		// Arrange
		mockClient := &MockBitwardenClient{}
		instance := &Environment{}

		// Mock successful calls
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

		mockClient.On("DownloadCerts", mock.MatchedBy(func(certs []*bitwarden.DownloadingCert) bool {
			return len(certs) == 1 &&
				certs[0].SID == "BACKEND_DATABASE_CERT_SID" &&
				certs[0].Dest == "/tmp/postgres/ca-cert.pem"
		})).Return(nil)

		// Act
		err := loadEnvs(instance, mockClient)

		// Assert
		assert.NoError(t, err)
		mockClient.AssertExpectations(t)
	})

	t.Run("returns error when LoadToEnviron fails", func(t *testing.T) {
		// Arrange
		mockClient := &MockBitwardenClient{}
		instance := &Environment{}
		expectedError := errors.New("bitwarden load failed")

		mockClient.On("LoadToEnviron", mock.AnythingOfType("[]string")).Return(expectedError)

		// Act
		err := loadEnvs(instance, mockClient)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "env.load.bitwarden.envs")
		assert.Contains(t, err.Error(), "bitwarden load failed")
		mockClient.AssertExpectations(t)
	})

	t.Run("returns error when DownloadCerts fails", func(t *testing.T) {
		// Arrange
		mockClient := &MockBitwardenClient{}
		instance := &Environment{}
		expectedError := errors.New("cert download failed")

		// Mock successful LoadToEnviron but failed DownloadCerts
		mockClient.On("LoadToEnviron", mock.AnythingOfType("[]string")).Return(nil)
		mockClient.On("DownloadCerts", mock.AnythingOfType("[]*bitwarden.DownloadingCert")).Return(expectedError)

		// Act
		err := loadEnvs(instance, mockClient)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "env.load.bitwarden.certs")
		assert.Contains(t, err.Error(), "cert download failed")
		mockClient.AssertExpectations(t)
	})
}
