package schemas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/mock"

	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

// MockSchemasClient mocks the repository schemas client
type MockSchemasClient struct {
	mock.Mock
}

func (m *MockSchemasClient) ByID(options *repositoryschemas.ByIDOptions) (*model.Schema, error) {
	args := m.Called(options)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Schema), args.Error(1)
}

func (m *MockSchemasClient) ByUser(options *repositoryschemas.ByUserOptions) ([]*model.Schema, error) {
	args := m.Called(options)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Schema), args.Error(1)
}

func (m *MockSchemasClient) Copy(options *repositoryschemas.CopyOptions) (*model.Schema, error) {
	args := m.Called(options)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Schema), args.Error(1)
}

func (m *MockSchemasClient) Create(options *repositoryschemas.CreateOptions) (*model.Schema, error) {
	args := m.Called(options)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Schema), args.Error(1)
}

func (m *MockSchemasClient) Delete(options *repositoryschemas.DeleteOptions) error {
	args := m.Called(options)
	return args.Error(0)
}

func (m *MockSchemasClient) Update(options *repositoryschemas.UpdateOptions) error {
	args := m.Called(options)
	return args.Error(0)
}

// MockS3Client mocks AWS S3 client
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

// Helper function to create test user
func createTestUser() *model.User {
	return &model.User{
		Identifiable: &model.Identifiable{
			ID: model.MustStringToID("550e8400-e29b-41d4-a716-446655440000"),
		},
	}
}

// Helper function to create test schema
func createTestSchema() *model.Schema {
	return &model.Schema{
		Identifiable: &model.Identifiable{
			ID: model.MustStringToID("650e8400-e29b-41d4-a716-446655440001"),
		},
		Name:    "Test Schema",
		Palette: model.TSchemaPalette{"#FF0000", "#00FF00", "#0000FF", "#FFFF00", "#FF00FF", "#00FFFF", "#000000", "#FFFFFF", "#888888"},
		Content: model.TSchemaContent{
			&model.SchemaPattern{
				ID:   "pattern1",
				Name: "Test Pattern",
				Type: model.SchemaPatternSquare,
				Content: []*model.SchemaRow{
					{
						ID: "row1",
						Content: []model.SchemaBead{
							{ID: "bead1", Color: "#FF0000"},
						},
					},
				},
			},
		},
	}
}
