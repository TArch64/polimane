package schemas

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

func TestController_apiCopy(t *testing.T) {
	t.Run("copies schema successfully", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()
		originalSchema := createTestSchema()
		copiedSchema := createTestSchema()
		copiedSchema.ID = model.MustStringToID("750e8400-e29b-41d4-a716-446655440002")
		copiedSchema.Name = "Test Schema (Copy)"

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/schemas/:schemaId/copy", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiCopy(c)
		})

		mockSchemas.On("Copy", mock.MatchedBy(func(options *repositoryschemas.CopyOptions) bool {
			return options.User == testUser && options.SchemaID == originalSchema.ID
		})).Return(copiedSchema, nil)

		req := httptest.NewRequest("POST", "/schemas/"+originalSchema.ID.String()+"/copy", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var responseSchema model.Schema
		err = json.Unmarshal(body, &responseSchema)
		assert.NoError(t, err)

		assert.Equal(t, copiedSchema.ID, responseSchema.ID)
		assert.Equal(t, copiedSchema.Name, responseSchema.Name)
		// Content should be nil in response (as per the implementation)
		assert.Nil(t, responseSchema.Content)

		mockSchemas.AssertExpectations(t)
	})

	t.Run("returns error when schema not found", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()
		testSchemaID := model.MustStringToID("650e8400-e29b-41d4-a716-446655440001")

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/schemas/:schemaId/copy", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiCopy(c)
		})

		mockSchemas.On("Copy", mock.MatchedBy(func(options *repositoryschemas.CopyOptions) bool {
			return options.User == testUser && options.SchemaID == testSchemaID
		})).Return(nil, gorm.ErrRecordNotFound)

		req := httptest.NewRequest("POST", "/schemas/"+testSchemaID.String()+"/copy", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)

		mockSchemas.AssertExpectations(t)
	})

	t.Run("returns error for invalid schema ID", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/schemas/:schemaId/copy", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiCopy(c)
		})

		req := httptest.NewRequest("POST", "/schemas/invalid-uuid/copy", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		mockSchemas.AssertNotCalled(t, "Copy")
	})

	t.Run("handles repository error", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()
		testSchemaID := model.MustStringToID("650e8400-e29b-41d4-a716-446655440001")

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/schemas/:schemaId/copy", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiCopy(c)
		})

		mockSchemas.On("Copy", mock.MatchedBy(func(options *repositoryschemas.CopyOptions) bool {
			return options.User == testUser && options.SchemaID == testSchemaID
		})).Return(nil, assert.AnError)

		req := httptest.NewRequest("POST", "/schemas/"+testSchemaID.String()+"/copy", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)

		mockSchemas.AssertExpectations(t)
	})

	t.Run("sets content to nil in response", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()
		originalSchema := createTestSchema()
		copiedSchema := createTestSchema()
		copiedSchema.ID = model.MustStringToID("750e8400-e29b-41d4-a716-446655440002")

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/schemas/:schemaId/copy", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiCopy(c)
		})

		mockSchemas.On("Copy", mock.MatchedBy(func(options *repositoryschemas.CopyOptions) bool {
			return options.User == testUser && options.SchemaID == originalSchema.ID
		})).Return(copiedSchema, nil)

		req := httptest.NewRequest("POST", "/schemas/"+originalSchema.ID.String()+"/copy", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var responseSchema model.Schema
		err = json.Unmarshal(body, &responseSchema)
		assert.NoError(t, err)

		// Verify content is explicitly set to nil
		assert.Nil(t, responseSchema.Content)

		mockSchemas.AssertExpectations(t)
	})
}
