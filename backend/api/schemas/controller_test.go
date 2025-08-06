package schemas

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"polimane/backend/api/base"
)

func TestProvider(t *testing.T) {
	t.Run("creates controller with dependencies", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}

		options := ControllerOptions{
			Schemas: mockSchemas,
			S3:      mockS3,
		}

		// Act
		controller := Provider(options)

		// Assert
		assert.NotNil(t, controller)
		assert.IsType(t, &Controller{}, controller)

		// Verify it implements base.Controller interface
		var baseController base.Controller = controller
		assert.NotNil(t, baseController)
	})
}

func TestController_Public(t *testing.T) {
	t.Run("does not register public routes", func(t *testing.T) {
		// Arrange
		mockSchemas := &MockSchemasClient{}
		mockS3 := &MockS3Client{}

		controller := &Controller{
			schemas: mockSchemas,
			s3:      mockS3,
		}

		// Act & Assert
		// Public method should do nothing - this just verifies it doesn't panic
		controller.Public(nil)
	})
}

func TestSchemaIdParam(t *testing.T) {
	t.Run("has correct schema ID parameter name", func(t *testing.T) {
		// Act & Assert
		assert.Equal(t, "schemaId", schemaIdParam)
	})
}
