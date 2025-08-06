package users

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"polimane/backend/api/base"
	"polimane/backend/signal"
)

func TestProvider(t *testing.T) {
	t.Run("creates controller with dependencies", func(t *testing.T) {
		// Arrange
		mockWorkosClient := NewMockWorkosClient()
		signalsContainer := &signal.Container{}

		// Act
		controller := Provider(ControllerOptions{
			WorkosClient: mockWorkosClient,
			Signals:      signalsContainer,
		})

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
		controller := &Controller{}

		// Act & Assert
		// Public method should do nothing - this just verifies it doesn't panic
		controller.Public(nil)
	})
}

func TestFactorIdParam(t *testing.T) {
	t.Run("has correct factor ID parameter name", func(t *testing.T) {
		// Act & Assert
		assert.Equal(t, "factorId", factorIdParam)
	})
}
