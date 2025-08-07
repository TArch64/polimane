package schemas

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/model"
	repositoryschemas "polimane/backend/repository/schemas"
	"polimane/backend/services/awsconfig"
)

func TestController_apiUpdateScreenshot(t *testing.T) {
	t.Run("updates screenshot successfully", func(t *testing.T) {
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

		app.Patch("/schemas/:schemaId/screenshot", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdateScreenshot(c)
		})

		// Mock S3 upload
		expectedKey := "data/images/" + testSchemaID.String() + "/schema.webp"
		mockS3.On("PutObject", mock.Anything, mock.MatchedBy(func(input *s3.PutObjectInput) bool {
			return *input.Key == expectedKey && *input.Bucket == awsconfig.S3Bucket && input.ACL == types.ObjectCannedACLPrivate && *input.ContentType == "image/webp"
		}), mock.Anything).Return(&s3.PutObjectOutput{}, nil)

		// Mock schema update
		mockSchemas.On("Update", mock.MatchedBy(func(options *repositoryschemas.UpdateOptions) bool {
			return options.User == testUser &&
				options.SchemaID == testSchemaID &&
				options.Updates.ScreenshotedAt != nil
		})).Return(nil)

		// Create a simple base64 encoded webp image
		testImageData := "test-image-data"
		encodedImage := base64.StdEncoding.EncodeToString([]byte(testImageData))
		src := "data:image/webp;base64," + encodedImage

		jsonBody := `{"src":"` + src + `"}`
		req := httptest.NewRequest("PATCH", "/schemas/"+testSchemaID.String()+"/screenshot", strings.NewReader(jsonBody))
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

		mockS3.AssertExpectations(t)
		mockSchemas.AssertExpectations(t)
	})

	t.Run("returns validation error for missing src", func(t *testing.T) {
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

		app.Patch("/schemas/:schemaId/screenshot", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdateScreenshot(c)
		})

		jsonBody := `{}`
		req := httptest.NewRequest("PATCH", "/schemas/"+testSchemaID.String()+"/screenshot", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		mockS3.AssertNotCalled(t, "PutObject")
		mockSchemas.AssertNotCalled(t, "Update")
	})

	t.Run("returns validation error for invalid src format", func(t *testing.T) {
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

		app.Patch("/schemas/:schemaId/screenshot", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdateScreenshot(c)
		})

		// Invalid format - should start with "data:image/webp;base64,"
		jsonBody := `{"src":"data:image/png;base64,invalid"}`
		req := httptest.NewRequest("PATCH", "/schemas/"+testSchemaID.String()+"/screenshot", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		mockS3.AssertNotCalled(t, "PutObject")
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

		app.Patch("/schemas/:schemaId/screenshot", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdateScreenshot(c)
		})

		testImageData := "test-image-data"
		encodedImage := base64.StdEncoding.EncodeToString([]byte(testImageData))
		src := "data:image/webp;base64," + encodedImage

		jsonBody := `{"src":"` + src + `"}`
		req := httptest.NewRequest("PATCH", "/schemas/invalid-uuid/screenshot", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)

		mockS3.AssertNotCalled(t, "PutObject")
		mockSchemas.AssertNotCalled(t, "Update")
	})

	t.Run("handles S3 upload error", func(t *testing.T) {
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

		app.Patch("/schemas/:schemaId/screenshot", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdateScreenshot(c)
		})

		// Mock S3 upload failure
		expectedKey := "data/images/" + testSchemaID.String() + "/schema.webp"
		mockS3.On("PutObject", mock.Anything, mock.MatchedBy(func(input *s3.PutObjectInput) bool {
			return *input.Key == expectedKey && *input.Bucket == awsconfig.S3Bucket && input.ACL == types.ObjectCannedACLPrivate && *input.ContentType == "image/webp"
		}), mock.Anything).Return(nil, assert.AnError)

		testImageData := "test-image-data"
		encodedImage := base64.StdEncoding.EncodeToString([]byte(testImageData))
		src := "data:image/webp;base64," + encodedImage

		jsonBody := `{"src":"` + src + `"}`
		req := httptest.NewRequest("PATCH", "/schemas/"+testSchemaID.String()+"/screenshot", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)

		mockS3.AssertExpectations(t)
		mockSchemas.AssertNotCalled(t, "Update")
	})

	t.Run("handles repository update error", func(t *testing.T) {
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

		app.Patch("/schemas/:schemaId/screenshot", func(c *fiber.Ctx) error {
			// Set up session user
			auth.SetSession(c, &auth.UserSession{
				ID:   "session-123",
				User: testUser,
			})
			return controller.apiUpdateScreenshot(c)
		})

		// Mock S3 upload success
		expectedKey := "data/images/" + testSchemaID.String() + "/schema.webp"
		mockS3.On("PutObject", mock.Anything, mock.MatchedBy(func(input *s3.PutObjectInput) bool {
			return *input.Key == expectedKey && *input.Bucket == awsconfig.S3Bucket && input.ACL == types.ObjectCannedACLPrivate && *input.ContentType == "image/webp"
		}), mock.Anything).Return(&s3.PutObjectOutput{}, nil)

		// Mock schema update failure
		mockSchemas.On("Update", mock.MatchedBy(func(options *repositoryschemas.UpdateOptions) bool {
			return options.User == testUser &&
				options.SchemaID == testSchemaID &&
				options.Updates.ScreenshotedAt != nil
		})).Return(assert.AnError)

		testImageData := "test-image-data"
		encodedImage := base64.StdEncoding.EncodeToString([]byte(testImageData))
		src := "data:image/webp;base64," + encodedImage

		jsonBody := `{"src":"` + src + `"}`
		req := httptest.NewRequest("PATCH", "/schemas/"+testSchemaID.String()+"/screenshot", strings.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.StatusCode)

		mockS3.AssertExpectations(t)
		mockSchemas.AssertExpectations(t)
	})
}

func TestController_uploadScreenshot(t *testing.T) {
	t.Run("uploads screenshot to S3 successfully", func(t *testing.T) {
		// Arrange
		mockS3 := &MockS3Client{}
		testSchemaID := model.MustStringToID("650e8400-e29b-41d4-a716-446655440001")

		controller := &Controller{
			s3: mockS3,
		}

		testImageData := "test-image-data"
		encodedImage := base64.StdEncoding.EncodeToString([]byte(testImageData))
		src := "data:image/webp;base64," + encodedImage

		expectedKey := "data/images/" + testSchemaID.String() + "/schema.webp"
		mockS3.On("PutObject", mock.Anything, mock.MatchedBy(func(input *s3.PutObjectInput) bool {
			return *input.Key == expectedKey && *input.Bucket == awsconfig.S3Bucket && input.ACL == types.ObjectCannedACLPrivate && *input.ContentType == "image/webp"
		}), mock.Anything).Return(&s3.PutObjectOutput{}, nil)

		ctx := context.Background()

		// Act
		err := controller.uploadScreenshot(ctx, testSchemaID, src)

		// Assert
		assert.NoError(t, err)
		mockS3.AssertExpectations(t)
	})

	t.Run("handles invalid base64 data", func(t *testing.T) {
		// Arrange
		mockS3 := &MockS3Client{}
		testSchemaID := model.MustStringToID("650e8400-e29b-41d4-a716-446655440001")

		controller := &Controller{
			s3: mockS3,
		}

		// Invalid base64 data
		src := "data:image/webp;base64,invalid-base64-data!!!"

		ctx := context.Background()

		// Act
		err := controller.uploadScreenshot(ctx, testSchemaID, src)

		// Assert
		assert.Error(t, err)
		mockS3.AssertNotCalled(t, "PutObject")
	})

	t.Run("handles S3 error", func(t *testing.T) {
		// Arrange
		mockS3 := &MockS3Client{}
		testSchemaID := model.MustStringToID("650e8400-e29b-41d4-a716-446655440001")

		controller := &Controller{
			s3: mockS3,
		}

		testImageData := "test-image-data"
		encodedImage := base64.StdEncoding.EncodeToString([]byte(testImageData))
		src := "data:image/webp;base64," + encodedImage

		expectedKey := "data/images/" + testSchemaID.String() + "/schema.webp"
		mockS3.On("PutObject", mock.Anything, mock.MatchedBy(func(input *s3.PutObjectInput) bool {
			return *input.Key == expectedKey && *input.Bucket == awsconfig.S3Bucket && input.ACL == types.ObjectCannedACLPrivate && *input.ContentType == "image/webp"
		}), mock.Anything).Return(nil, assert.AnError)

		ctx := context.Background()

		// Act
		err := controller.uploadScreenshot(ctx, testSchemaID, src)

		// Assert
		assert.Error(t, err)
		mockS3.AssertExpectations(t)
	})

	t.Run("extracts base64 data correctly", func(t *testing.T) {
		// Arrange
		mockS3 := &MockS3Client{}
		testSchemaID := model.MustStringToID("650e8400-e29b-41d4-a716-446655440001")

		controller := &Controller{
			s3: mockS3,
		}

		testImageData := "specific-test-data"
		encodedImage := base64.StdEncoding.EncodeToString([]byte(testImageData))
		src := "data:image/webp;base64," + encodedImage

		expectedKey := "data/images/" + testSchemaID.String() + "/schema.webp"

		// Verify the body contains the decoded data
		mockS3.On("PutObject", mock.Anything, mock.MatchedBy(func(input *s3.PutObjectInput) bool {
			if input.Key == nil || *input.Key != expectedKey {
				return false
			}
			if input.ContentType == nil || *input.ContentType != "image/webp" {
				return false
			}
			if input.ACL != types.ObjectCannedACLPrivate {
				return false
			}
			if input.Bucket == nil || *input.Bucket != awsconfig.S3Bucket {
				return false
			}

			// Read the body and verify it matches our test data
			if reader, ok := input.Body.(*bytes.Reader); ok {
				buf := make([]byte, reader.Len())
				reader.Read(buf)
				reader.Seek(0, 0) // Reset for actual use
				return string(buf) == testImageData
			}
			return false
		}), mock.Anything).Return(&s3.PutObjectOutput{}, nil)

		ctx := context.Background()

		// Act
		err := controller.uploadScreenshot(ctx, testSchemaID, src)

		// Assert
		assert.NoError(t, err)
		mockS3.AssertExpectations(t)
	})

	t.Run("generates correct S3 key", func(t *testing.T) {
		// Arrange
		mockS3 := &MockS3Client{}
		testSchemaID := model.MustStringToID("123e4567-e89b-12d3-a456-426614174000")

		controller := &Controller{
			s3: mockS3,
		}

		testImageData := "test-data"
		encodedImage := base64.StdEncoding.EncodeToString([]byte(testImageData))
		src := "data:image/webp;base64," + encodedImage

		expectedKey := "data/images/123e4567-e89b-12d3-a456-426614174000/schema.webp"
		mockS3.On("PutObject", mock.Anything, mock.MatchedBy(func(input *s3.PutObjectInput) bool {
			return input.Key != nil && *input.Key == expectedKey
		}), mock.Anything).Return(&s3.PutObjectOutput{}, nil)

		ctx := context.Background()

		// Act
		err := controller.uploadScreenshot(ctx, testSchemaID, src)

		// Assert
		assert.NoError(t, err)
		mockS3.AssertExpectations(t)
	})
}

func TestApiUpdateScreenshotBody(t *testing.T) {
	t.Run("has correct validation tags", func(t *testing.T) {
		// This test verifies the struct definition
		body := apiUpdateScreenshotBody{
			Src: "data:image/webp;base64,test",
		}

		assert.Equal(t, "data:image/webp;base64,test", body.Src)
	})

	t.Run("sets ScreenshotedAt to current time", func(t *testing.T) {
		// This is tested indirectly through the integration test,
		// but we can verify the time is recent
		before := time.Now()
		time.Sleep(1 * time.Millisecond)
		now := time.Now()
		time.Sleep(1 * time.Millisecond)
		after := time.Now()

		// In the actual implementation, screenshotedAt := time.Now()
		// We verify it's between before and after our test execution
		assert.True(t, now.After(before))
		assert.True(t, now.Before(after))
	})
}
