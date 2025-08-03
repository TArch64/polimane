package users

import (
	"context"

	"github.com/stretchr/testify/mock"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
)

type MockUsersClient struct {
	mock.Mock
}

func (m *MockUsersClient) ByID(ctx context.Context, id modelbase.ID) (*model.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUsersClient) CreateIfNeeded(ctx context.Context, workosID string) (*model.User, error) {
	args := m.Called(ctx, workosID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}
