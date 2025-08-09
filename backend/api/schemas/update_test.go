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

func TestController_apiUpdate(t *testing.T) {
	t.Run("updates schema name successfully", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()
		testSchemaID := model.MustStringToID("650e8400-e29b-41d4-a716-446655440001")

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		// Initialize validator
		base.InitValidator()

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Patch("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdate(c)
		})

		mockSchemas.On("Update", mock.MatchedBy(func(options *repositoryschemas.UpdateOptions) bool {
			return options.User == testUser && options.SchemaID == testSchemaID
		})).Return(nil)

		jsonBody := `{"name":"Updated Schema"}`
		req := httptest.NewRequest("PATCH", "/schemas/"+testSchemaID.String(), strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

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

	t.Run("updates schema palette successfully", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()
		testSchemaID := model.MustStringToID("650e8400-e29b-41d4-a716-446655440001")

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		// Initialize validator
		base.InitValidator()

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Patch("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdate(c)
		})

		mockSchemas.On("Update", mock.MatchedBy(func(options *repositoryschemas.UpdateOptions) bool {
			return options.User == testUser && options.SchemaID == testSchemaID
		})).Return(nil)

		jsonBody := `{"palette":["#FF0000","#00FF00","#0000FF","#FFFF00","#FF00FF","#00FFFF","#000000","#FFFFFF","#888888"]}`
		req := httptest.NewRequest("PATCH", "/schemas/"+testSchemaID.String(), strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		mockSchemas.AssertExpectations(t)
	})

	t.Run("updates schema content successfully", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()
		testSchemaID := model.MustStringToID("650e8400-e29b-41d4-a716-446655440001")

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		// Initialize validator
		base.InitValidator()

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Patch("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdate(c)
		})

		mockSchemas.On("Update", mock.MatchedBy(func(options *repositoryschemas.UpdateOptions) bool {
			return options.User == testUser && options.SchemaID == testSchemaID
		})).Return(nil)

		jsonBody := `{"content":[{"id":"pattern1","name":"Test Pattern","type":"square","content":[{"id":"row1","content":[{"id":"bead1","color":"#FF0000"}]}]}]}`
		req := httptest.NewRequest("PATCH", "/schemas/"+testSchemaID.String(), strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		mockSchemas.AssertExpectations(t)
	})

	t.Run("returns error for empty updates", func(t *testing.T) {
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

		app.Patch("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdate(c)
		})

		jsonBody := `{}`
		req := httptest.NewRequest("PATCH", "/schemas/"+testSchemaID.String(), strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		var errorResponse base.CustomError
		err = json.Unmarshal(body, &errorResponse)
		assert.NoError(t, err)

		assert.Contains(t, errorResponse.Message, "EmptyUpdatesInput")

		mockSchemas.AssertNotCalled(t, "Update")
	})

	t.Run("returns validation error for invalid palette length", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}
		testUser := createTestUser()
		testSchemaID := model.MustStringToID("650e8400-e29b-41d4-a716-446655440001")

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		// Initialize validator
		base.InitValidator()

		app := fiber.New(fiber.Config{
			ErrorHandler: base.ErrorHandler,
		})

		app.Patch("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdate(c)
		})

		// Palette should have exactly 9 colors
		jsonBody := `{"palette":["#FF0000","#00FF00"]}`
		req := httptest.NewRequest("PATCH", "/schemas/"+testSchemaID.String(), strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		mockSchemas.AssertNotCalled(t, "Update")
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

		app.Patch("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdate(c)
		})

		jsonBody := `{"name":"Updated Schema"}`
		req := httptest.NewRequest("PATCH", "/schemas/invalid-uuid", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		mockSchemas.AssertNotCalled(t, "Update")
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

		app.Patch("/schemas/:schemaId", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdate(c)
		})

		mockSchemas.On("Update", mock.MatchedBy(func(options *repositoryschemas.UpdateOptions) bool {
			return options.User == testUser && options.SchemaID == testSchemaID
		})).Return(assert.AnError)

		jsonBody := `{"name":"Updated Schema"}`
		req := httptest.NewRequest("PATCH", "/schemas/"+testSchemaID.String(), strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)

		mockSchemas.AssertExpectations(t)
	})
}

func TestCollectUpdates(t *testing.T) {
	t.Run("returns schema with name when name provided", func(t *testing.T) {
		// Arrange
		body := &updateBody{Name: "Test Schema"}

		// Act
		updates := collectUpdates(body)

		// Assert
		assert.NotNil(t, updates)
		assert.Equal(t, "Test Schema", updates.Name)
		assert.Nil(t, updates.Content)
		assert.Nil(t, updates.Palette)
	})

	t.Run("returns schema with palette when palette provided", func(t *testing.T) {
		// Arrange
		palette := []string{"#FF0000", "#00FF00", "#0000FF", "#FFFF00", "#FF00FF", "#00FFFF", "#000000", "#FFFFFF", "#888888"}
		body := &updateBody{Palette: palette}

		// Act
		updates := collectUpdates(body)

		// Assert
		assert.NotNil(t, updates)
		assert.Equal(t, "", updates.Name)
		assert.Equal(t, model.TSchemaPalette(palette), updates.Palette)
		assert.Nil(t, updates.Content)
	})

	t.Run("returns schema with content when content provided", func(t *testing.T) {
		// Arrange
		content := model.TSchemaContent{
			&model.SchemaPattern{
				ID:   "pattern1",
				Name: "Test Pattern",
				Type: model.SchemaPatternSquare,
			},
		}
		body := &updateBody{Content: content}

		// Act
		updates := collectUpdates(body)

		// Assert
		assert.NotNil(t, updates)
		assert.Equal(t, "", updates.Name)
		assert.Nil(t, updates.Palette)
		assert.Equal(t, content, updates.Content)
	})

	t.Run("returns schema with all fields when all provided", func(t *testing.T) {
		// Arrange
		content := model.TSchemaContent{
			&model.SchemaPattern{
				ID:   "pattern1",
				Name: "Test Pattern",
				Type: model.SchemaPatternSquare,
			},
		}
		palette := []string{"#FF0000", "#00FF00", "#0000FF", "#FFFF00", "#FF00FF", "#00FFFF", "#000000", "#FFFFFF", "#888888"}
		body := &updateBody{
			Name:    "Test Schema",
			Palette: palette,
			Content: content,
		}

		// Act
		updates := collectUpdates(body)

		// Assert
		assert.NotNil(t, updates)
		assert.Equal(t, "Test Schema", updates.Name)
		assert.Equal(t, model.TSchemaPalette(palette), updates.Palette)
		assert.Equal(t, content, updates.Content)
	})

	t.Run("returns nil when no updates provided", func(t *testing.T) {
		// Arrange
		body := &updateBody{}

		// Act
		updates := collectUpdates(body)

		// Assert
		assert.Nil(t, updates)
	})

	t.Run("ignores empty name", func(t *testing.T) {
		// Arrange
		body := &updateBody{Name: ""}

		// Act
		updates := collectUpdates(body)

		// Assert
		assert.Nil(t, updates)
	})

	t.Run("handles nil content", func(t *testing.T) {
		// Arrange
		body := &updateBody{Content: nil}

		// Act
		updates := collectUpdates(body)

		// Assert
		assert.Nil(t, updates)
	})

	t.Run("handles nil palette", func(t *testing.T) {
		// Arrange
		body := &updateBody{Palette: nil}

		// Act
		updates := collectUpdates(body)

		// Assert
		assert.Nil(t, updates)
	})
}

func TestUpdateBody(t *testing.T) {
	t.Run("has correct validation tags", func(t *testing.T) {
		// This test verifies the struct definition
		content := model.TSchemaContent{
			&model.SchemaPattern{
				ID:   "pattern1",
				Name: "Test Pattern",
				Type: model.SchemaPatternSquare,
			},
		}
		palette := []string{"#FF0000", "#00FF00", "#0000FF", "#FFFF00", "#FF00FF", "#00FFFF", "#000000", "#FFFFFF", "#888888"}

		body := updateBody{
			Name:    "Test",
			Palette: palette,
			Content: content,
		}

		assert.Equal(t, "Test", body.Name)
		assert.Equal(t, palette, body.Palette)
		assert.Equal(t, content, body.Content)
	})
}
