package env

import (
	"github.com/stretchr/testify/mock"

	"polimane/backend/services/bitwarden"
)

// MockBitwardenClient implements bitwarden.Client for testing
type MockBitwardenClient struct {
	mock.Mock
}

func (m *MockBitwardenClient) DownloadCerts(certs []*bitwarden.DownloadingCert) error {
	args := m.Called(certs)
	return args.Error(0)
}

func (m *MockBitwardenClient) Load(sids []string) (map[string]string, error) {
	args := m.Called(sids)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(map[string]string), args.Error(1)
}

func (m *MockBitwardenClient) LoadToEnviron(names []string) error {
	args := m.Called(names)
	return args.Error(0)
}
