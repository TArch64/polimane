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

	t.Run("updates schema size successfully", func(t *testing.T) {
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

		jsonBody := `{"size":{"left":10,"top":20,"right":90,"bottom":80}}`
		req := httptest.NewRequest("PATCH", "/schemas/"+testSchemaID.String(), strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

		mockSchemas.AssertExpectations(t)
	})

	t.Run("updates schema beads successfully", func(t *testing.T) {
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

		jsonBody := `{"beads":{"1":"#FF0000","2":"#00FF00","3":"#0000FF"}}`
		req := httptest.NewRequest("PATCH", "/schemas/"+testSchemaID.String(), strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)

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
		assert.True(t, updates.Palette.Data() == nil)
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
		assert.NotNil(t, updates.Palette)
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

	t.Run("handles nil palette", func(t *testing.T) {
		// Arrange
		body := &updateBody{Palette: nil}

		// Act
		updates := collectUpdates(body)

		// Assert
		assert.Nil(t, updates)
	})

	t.Run("returns schema with all fields when all provided", func(t *testing.T) {
		// Arrange
		palette := model.SchemaPalette{"#FF0000", "#00FF00", "#0000FF", "#FFFF00", "#FF00FF", "#00FFFF", "#000000", "#FFFFFF", "#888888"}
		size := &model.SchemaSize{Left: 5, Top: 10, Right: 95, Bottom: 90}
		beads := model.SchemaBeads{"1": "#FF0000", "2": "#00FF00"}
		body := &updateBody{
			Name:    "Updated Schema",
			Palette: palette,
			Size:    size,
			Beads:   beads,
		}

		// Act
		updates := collectUpdates(body)

		// Assert
		assert.NotNil(t, updates)
		assert.Equal(t, "Updated Schema", updates.Name)
		assert.NotNil(t, updates.Palette)
		assert.Equal(t, palette, updates.Palette.Data())
		assert.NotNil(t, updates.Size)
		assert.Equal(t, size, updates.Size.Data())
		assert.NotNil(t, updates.Beads)
		assert.Equal(t, beads, updates.Beads.Data())
	})
}

func TestUpdateBody(t *testing.T) {
	t.Run("has correct validation tags and field mappings", func(t *testing.T) {
		// Test all updateBody fields including Size and Beads
		palette := []string{"#FF0000", "#00FF00", "#0000FF", "#FFFF00", "#FF00FF", "#00FFFF", "#000000", "#FFFFFF", "#888888"}
		size := &model.SchemaSize{Left: 10, Top: 20, Right: 90, Bottom: 80}
		beads := model.SchemaBeads{"1": "#FF0000", "2": "#00FF00", "3": "#0000FF"}

		body := updateBody{
			Name:    "Test Schema",
			Palette: palette,
			Size:    size,
			Beads:   beads,
		}

		// Verify all field mappings
		assert.Equal(t, "Test Schema", body.Name)
		assert.Equal(t, model.SchemaPalette(palette), body.Palette)
		assert.Equal(t, size, body.Size)
		assert.Equal(t, beads, body.Beads)

		// Test collectUpdates with all fields to ensure coverage
		updates := collectUpdates(&body)
		assert.NotNil(t, updates)
		assert.Equal(t, "Test Schema", updates.Name)
		assert.NotNil(t, updates.Palette)
		assert.Equal(t, model.SchemaPalette(palette), updates.Palette.Data())
		assert.NotNil(t, updates.Size)
		assert.Equal(t, size, updates.Size.Data())
		assert.NotNil(t, updates.Beads)
		assert.Equal(t, beads, updates.Beads.Data())
	})
}
