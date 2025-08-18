package schemas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
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

// MockS3Client mocks AWS S3 client for repository tests
type MockS3Client struct {
	mock.Mock
}

func (m *MockS3Client) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	args := m.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.PutObjectOutput), args.Error(1)
}

func (m *MockS3Client) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	args := m.Called(ctx, params, optFns)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*s3.DeleteObjectOutput), args.Error(1)
}
