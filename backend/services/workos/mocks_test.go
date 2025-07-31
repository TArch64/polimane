package workos

import (
	"context"
	"net/url"

	"github.com/stretchr/testify/mock"
)

// Mock implementations for WorkOS service testing
// These mocks are kept minimal since most WorkOS functionality
// requires integration testing with real JWT tokens and API calls

type MockJWKInterface struct {
	mock.Mock
}

func (m *MockJWKInterface) Fetch(ctx context.Context, u string) (interface{}, error) {
	args := m.Called(ctx, u)
	return args.Get(0), args.Error(1)
}

type MockUserManagementInterface struct {
	mock.Mock
}

func (m *MockUserManagementInterface) GetJWKSURL(clientID string) (*url.URL, error) {
	args := m.Called(clientID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*url.URL), args.Error(1)
}
