package bitwarden

import (
	"os"
	"time"

	"github.com/bitwarden/sdk-go"
	"github.com/stretchr/testify/mock"
)

type MockSecretsManager struct {
	mock.Mock
}

func (m *MockSecretsManager) Create(key, value, note string, organizationID string, projectIDs []string) (*sdk.SecretResponse, error) {
	args := m.Called(key, value, note, organizationID, projectIDs)
	return args.Get(0).(*sdk.SecretResponse), args.Error(1)
}

func (m *MockSecretsManager) List(organizationID string) (*sdk.SecretIdentifiersResponse, error) {
	args := m.Called(organizationID)
	return args.Get(0).(*sdk.SecretIdentifiersResponse), args.Error(1)
}

func (m *MockSecretsManager) Get(secretID string) (*sdk.SecretResponse, error) {
	args := m.Called(secretID)
	return args.Get(0).(*sdk.SecretResponse), args.Error(1)
}

func (m *MockSecretsManager) GetByIDS(ids []string) (*sdk.SecretsResponse, error) {
	args := m.Called(ids)
	return args.Get(0).(*sdk.SecretsResponse), args.Error(1)
}

func (m *MockSecretsManager) Update(secretID string, key, value, note string, organizationID string, projectIDs []string) (*sdk.SecretResponse, error) {
	args := m.Called(secretID, key, value, note, organizationID, projectIDs)
	return args.Get(0).(*sdk.SecretResponse), args.Error(1)
}

func (m *MockSecretsManager) Delete(secretIDs []string) (*sdk.SecretsDeleteResponse, error) {
	args := m.Called(secretIDs)
	return args.Get(0).(*sdk.SecretsDeleteResponse), args.Error(1)
}

func (m *MockSecretsManager) Sync(organizationID string, lastSyncedDate *time.Time) (*sdk.SecretsSyncResponse, error) {
	args := m.Called(organizationID, lastSyncedDate)
	return args.Get(0).(*sdk.SecretsSyncResponse), args.Error(1)
}

type MockBitwardenClient struct {
	mock.Mock
}

func (m *MockBitwardenClient) AccessTokenLogin(accessToken string, stateFile *string) error {
	args := m.Called(accessToken, stateFile)
	return args.Error(0)
}

func (m *MockBitwardenClient) Projects() sdk.ProjectsInterface {
	args := m.Called()
	return args.Get(0).(sdk.ProjectsInterface)
}

func (m *MockBitwardenClient) Secrets() sdk.SecretsInterface {
	args := m.Called()
	return args.Get(0).(sdk.SecretsInterface)
}

func (m *MockBitwardenClient) Generators() sdk.GeneratorsInterface {
	args := m.Called()
	return args.Get(0).(sdk.GeneratorsInterface)
}

func (m *MockBitwardenClient) Close() {
	m.Called()
}

type MockEnv struct {
	mock.Mock
}

func (m *MockEnv) Getenv(key string) string {
	args := m.Called(key)
	return args.String(0)
}

func (m *MockEnv) Setenv(key, value string) error {
	args := m.Called(key, value)
	return args.Error(0)
}

type MockFS struct {
	mock.Mock
}

func (m *MockFS) MkdirAll(path string, perm os.FileMode) error {
	args := m.Called(path, perm)
	return args.Error(0)
}

func (m *MockFS) WriteFile(name string, data []byte, perm os.FileMode) error {
	args := m.Called(name, data, perm)
	return args.Error(0)
}
