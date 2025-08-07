package schemas

import (
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"polimane/backend/model"
)

type MockUserSchemas struct {
	mock.Mock
}

func (m *MockUserSchemas) HasAccess(ctx context.Context, userID, schemaID model.ID) error {
	args := m.Called(ctx, userID, schemaID)
	return args.Error(0)
}

func (m *MockUserSchemas) CreateTx(tx *gorm.DB, userID, schemaID model.ID) error {
	args := m.Called(tx, userID, schemaID)
	return args.Error(0)
}

func (m *MockUserSchemas) DeleteTx(tx *gorm.DB, userID, schemaID model.ID) error {
	args := m.Called(tx, userID, schemaID)
	return args.Error(0)
}
