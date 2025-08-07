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

func TestController_apiById(t *testing.T) {
	t.Run("returns schema when found", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()
		testSchema := createTestSchema()

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Get("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiById(c)
		})

		mockSchemas.On("ByID", mock.MatchedBy(func(options *repositoryschemas.ByIDOptions) bool {
			return options.User == testUser && options.SchemaID == testSchema.ID
		})).Return(testSchema, nil)

		req := httptest.NewRequest("GET", "/schemas/"+testSchema.ID.String(), nil)

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

		app.Get("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiById(c)
		})

		mockSchemas.On("ByID", mock.MatchedBy(func(options *repositoryschemas.ByIDOptions) bool {
			return options.User == testUser && options.SchemaID == testSchemaID
		})).Return(nil, gorm.ErrRecordNotFound)

		req := httptest.NewRequest("GET", "/schemas/"+testSchemaID.String(), nil)

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

		app.Get("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiById(c)
		})

		req := httptest.NewRequest("GET", "/schemas/invalid-uuid", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		mockSchemas.AssertNotCalled(t, "ByID")
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

		app.Get("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiById(c)
		})

		mockSchemas.On("ByID", mock.MatchedBy(func(options *repositoryschemas.ByIDOptions) bool {
			return options.User == testUser && options.SchemaID == testSchemaID
		})).Return(nil, assert.AnError)

		req := httptest.NewRequest("GET", "/schemas/"+testSchemaID.String(), nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)

		mockSchemas.AssertExpectations(t)
	})
}
