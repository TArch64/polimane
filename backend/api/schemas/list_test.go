package schemas

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
)

func TestController_apiList(t *testing.T) {
	t.Run("returns schemas list successfully", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()

		schema1 := &model.Schema{
			Identifiable: &model.Identifiable{
				ID: model.MustStringToID("650e8400-e29b-41d4-a716-446655440001"),
			},
			Name: "Schema 1",
		}
		schema2 := &model.Schema{
			Identifiable: &model.Identifiable{
				ID: model.MustStringToID("650e8400-e29b-41d4-a716-446655440002"),
			},
			Name: "Schema 2",
		}
		schemas := []*model.Schema{schema1, schema2}

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Get("/schemas", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiList(c)
		})

		mockSchemas.On("ByUser", mock.MatchedBy(func(options *repositoryschemas.ByUserOptions) bool {
			return options.User == testUser && len(options.Select) == 3
		})).Return(schemas, nil)

		req := httptest.NewRequest("GET", "/schemas", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var responseSchemas []*model.Schema
		err = json.Unmarshal(body, &responseSchemas)
		assert.NoError(t, err)

		assert.Len(t, responseSchemas, 2)
		assert.Equal(t, schema1.ID, responseSchemas[0].ID)
		assert.Equal(t, schema1.Name, responseSchemas[0].Name)
		assert.Equal(t, schema2.ID, responseSchemas[1].ID)
		assert.Equal(t, schema2.Name, responseSchemas[1].Name)

		mockSchemas.AssertExpectations(t)
	})

	t.Run("returns empty array when no schemas found", func(t *testing.T) {
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

		app.Get("/schemas", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiList(c)
		})

		// Return nil instead of empty slice to test the nil handling
		mockSchemas.On("ByUser", mock.MatchedBy(func(options *repositoryschemas.ByUserOptions) bool {
			return options.User == testUser && len(options.Select) == 3
		})).Return(nil, nil)

		req := httptest.NewRequest("GET", "/schemas", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var responseSchemas []*model.Schema
		err = json.Unmarshal(body, &responseSchemas)
		assert.NoError(t, err)

		// Should return empty array, not null
		assert.NotNil(t, responseSchemas)
		assert.Len(t, responseSchemas, 0)

		mockSchemas.AssertExpectations(t)
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

		app.Get("/schemas", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiList(c)
		})

		mockSchemas.On("ByUser", mock.MatchedBy(func(options *repositoryschemas.ByUserOptions) bool {
			return options.User == testUser && len(options.Select) == 3
		})).Return(nil, assert.AnError)

		req := httptest.NewRequest("GET", "/schemas", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)

		mockSchemas.AssertExpectations(t)
	})

	t.Run("uses correct select fields", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()
		schemas := []*model.Schema{}

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Get("/schemas", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiList(c)
		})

		// Verify the exact select fields are used
		mockSchemas.On("ByUser", mock.MatchedBy(func(options *repositoryschemas.ByUserOptions) bool {
			return options.User == testUser && len(options.Select) == 3
		})).Return(schemas, nil)

		req := httptest.NewRequest("GET", "/schemas", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		mockSchemas.AssertExpectations(t)
	})

	t.Run("returns single schema in array", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()

		schema := &model.Schema{
			Identifiable: &model.Identifiable{
				ID: model.MustStringToID("650e8400-e29b-41d4-a716-446655440001"),
			},
			Name: "Single Schema",
		}
		schemas := []*model.Schema{schema}

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Get("/schemas", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiList(c)
		})

		mockSchemas.On("ByUser", mock.MatchedBy(func(options *repositoryschemas.ByUserOptions) bool {
			return options.User == testUser && len(options.Select) == 3
		})).Return(schemas, nil)

		req := httptest.NewRequest("GET", "/schemas", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var responseSchemas []*model.Schema
		err = json.Unmarshal(body, &responseSchemas)
		assert.NoError(t, err)

		assert.Len(t, responseSchemas, 1)
		assert.Equal(t, schema.ID, responseSchemas[0].ID)
		assert.Equal(t, schema.Name, responseSchemas[0].Name)

		mockSchemas.AssertExpectations(t)
	})
}
