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

func TestController_apiDelete(t *testing.T) {
	t.Run("deletes schema successfully", func(t *testing.T) {
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

		app.Delete("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiDelete(c)
		})

		mockSchemas.On("Delete", mock.MatchedBy(func(options *repositoryschemas.DeleteOptions) bool {
			return options.User == testUser && options.SchemaID == testSchemaID
		})).Return(nil)

		req := httptest.NewRequest("DELETE", "/schemas/"+testSchemaID.String(), nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var response map[string]interface{}
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		assert.Equal(t, true, response["success"])

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

		app.Delete("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiDelete(c)
		})

		mockSchemas.On("Delete", mock.MatchedBy(func(options *repositoryschemas.DeleteOptions) bool {
			return options.User == testUser && options.SchemaID == testSchemaID
		})).Return(gorm.ErrRecordNotFound)

		req := httptest.NewRequest("DELETE", "/schemas/"+testSchemaID.String(), nil)

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

		app.Delete("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiDelete(c)
		})

		req := httptest.NewRequest("DELETE", "/schemas/invalid-uuid", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		mockSchemas.AssertNotCalled(t, "Delete")
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

		app.Delete("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiDelete(c)
		})

		mockSchemas.On("Delete", mock.MatchedBy(func(options *repositoryschemas.DeleteOptions) bool {
			return options.User == testUser && options.SchemaID == testSchemaID
		})).Return(assert.AnError)

		req := httptest.NewRequest("DELETE", "/schemas/"+testSchemaID.String(), nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)

		mockSchemas.AssertExpectations(t)
	})

	t.Run("returns success response format", func(t *testing.T) {
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

		app.Delete("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiDelete(c)
		})

		mockSchemas.On("Delete", mock.MatchedBy(func(options *repositoryschemas.DeleteOptions) bool {
			return options.User == testUser && options.SchemaID == testSchemaID
		})).Return(nil)

		req := httptest.NewRequest("DELETE", "/schemas/"+testSchemaID.String(), nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var response fiber.Map
		err = json.Unmarshal(body, &response)
		assert.NoError(t, err)

		// Verify it matches the expected success response format
		expected := fiber.Map{"success": true}
		assert.Equal(t, expected, response)

		mockSchemas.AssertExpectations(t)
	})
}
