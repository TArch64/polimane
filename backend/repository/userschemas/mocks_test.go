package userschemas

import (
	"context"

	tmock "github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"polimane/backend/model/modelbase"
)

type MockUserSchemasClient struct {
	tmock.Mock
}

func (m *MockUserSchemasClient) CreateTx(tx *gorm.DB, userID, schemaID modelbase.ID) error {
	args := m.Called(tx, userID, schemaID)
	return args.Error(0)
}

func (m *MockUserSchemasClient) DeleteTx(tx *gorm.DB, userID, schemaID modelbase.ID) error {
	args := m.Called(tx, userID, schemaID)
	return args.Error(0)
}

func (m *MockUserSchemasClient) HasAccess(ctx context.Context, userID, schemaID modelbase.ID) error {
	args := m.Called(ctx, userID, schemaID)
	return args.Error(0)
}
