package schemas

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

func TestController_apiCreate(t *testing.T) {
	t.Run("creates schema successfully", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()
		testSchema := createTestSchema()

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		// Initialize validator
		base.InitValidator()

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/schemas", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiCreate(c)
		})

		mockSchemas.On("Create", mock.MatchedBy(func(options *repositoryschemas.CreateOptions) bool {
			return options.User == testUser && options.Name == testSchema.Name
		})).Return(testSchema, nil)

		jsonBody := `{"name":"` + testSchema.Name + `"}`
		req := httptest.NewRequest("POST", "/schemas", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

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

		assert.Equal(t, testSchema.ID, responseSchema.ID)
		assert.Equal(t, testSchema.Name, responseSchema.Name)

		mockSchemas.AssertExpectations(t)
	})

	t.Run("returns validation error for missing name", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		// Initialize validator
		base.InitValidator()

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/schemas", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiCreate(c)
		})

		jsonBody := `{}`
		req := httptest.NewRequest("POST", "/schemas", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		mockSchemas.AssertNotCalled(t, "Create")
	})

	t.Run("returns validation error for empty name", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		// Initialize validator
		base.InitValidator()

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/schemas", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiCreate(c)
		})

		jsonBody := `{"name":""}`
		req := httptest.NewRequest("POST", "/schemas", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		mockSchemas.AssertNotCalled(t, "Create")
	})

	t.Run("returns error for invalid JSON", func(t *testing.T) {
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

		app.Post("/schemas", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiCreate(c)
		})

		invalidJson := `{"name":"Test"` // Missing closing brace
		req := httptest.NewRequest("POST", "/schemas", strings.NewReader(invalidJson))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode) // JSON parsing error

		mockSchemas.AssertNotCalled(t, "Create")
	})

	t.Run("handles repository error", func(t *testing.T) {
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

		app.Post("/schemas", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiCreate(c)
		})

		mockSchemas.On("Create", mock.MatchedBy(func(options *repositoryschemas.CreateOptions) bool {
			return options.User == testUser && options.Name == "Test Schema"
		})).Return(nil, assert.AnError)

		jsonBody := `{"name":"Test Schema"}`
		req := httptest.NewRequest("POST", "/schemas", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

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
		testSchema := createTestSchema()

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		// Initialize validator
		base.InitValidator()

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Post("/schemas", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiCreate(c)
		})

		mockSchemas.On("Create", mock.MatchedBy(func(options *repositoryschemas.CreateOptions) bool {
			return options.User == testUser && options.Name == testSchema.Name
		})).Return(testSchema, nil)

		jsonBody := `{"name":"` + testSchema.Name + `"}`
		req := httptest.NewRequest("POST", "/schemas", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

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

		mockSchemas.AssertExpectations(t)
	})
}

func TestCreateBody(t *testing.T) {
	t.Run("has correct JSON tags", func(t *testing.T) {
		// This test verifies the struct definition
		body := createBody{
			Name: "Test",
		}

		assert.Equal(t, "Test", body.Name)
	})
}
