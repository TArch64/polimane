package schemas

import (
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
)

type MockUserSchemas struct {
	mock.Mock
}

func (m *MockUserSchemas) HasAccess(ctx context.Context, userID, schemaID modelbase.ID) error {
	args := m.Called(ctx, userID, schemaID)
	return args.Error(0)
}

func (m *MockUserSchemas) CreateTx(tx *gorm.DB, userID, schemaID modelbase.ID) error {
	args := m.Called(tx, userID, schemaID)
	return args.Error(0)
}

func (m *MockUserSchemas) DeleteTx(tx *gorm.DB, userID, schemaID modelbase.ID) error {
	args := m.Called(tx, userID, schemaID)
	return args.Error(0)
}

type MockSchemasClient struct {
	mock.Mock
}

func (m *MockSchemasClient) ByID(options *ByIDOptions) (*model.Schema, error) {
	args := m.Called(options)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Schema), args.Error(1)
}

func (m *MockSchemasClient) ByUser(options *ByUserOptions) ([]*model.Schema, error) {
	args := m.Called(options)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Schema), args.Error(1)
}

func (m *MockSchemasClient) Copy(options *CopyOptions) (*model.Schema, error) {
	args := m.Called(options)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Schema), args.Error(1)
}

func (m *MockSchemasClient) Create(options *CreateOptions) (*model.Schema, error) {
	args := m.Called(options)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Schema), args.Error(1)
}

func (m *MockSchemasClient) Delete(options *DeleteOptions) error {
	args := m.Called(options)
	return args.Error(0)
}

func (m *MockSchemasClient) Update(options *UpdateOptions) error {
	args := m.Called(options)
	return args.Error(0)
}
