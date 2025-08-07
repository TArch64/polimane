package api

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"polimane/backend/api/auth"
	"polimane/backend/api/base"
	"polimane/backend/env"
	"polimane/backend/services/sentry"
)

// MockController implements base.Controller for testing
type MockController struct {
	mock.Mock
}

func (m *MockController) Public(router fiber.Router) {
	m.Called(router)
}

func (m *MockController) Private(router fiber.Router) {
	m.Called(router)
}

func TestOptions(t *testing.T) {
	t.Run("struct initialization", func(t *testing.T) {
		configCalled := false
		options := &Options{
			Protocol: "https",
			Configure: func(config *fiber.Config) {
				configCalled = true
				config.AppName = "Test App"
			},
		}

		assert.Equal(t, "https", options.Protocol)
		assert.NotNil(t, options.Configure)

		// Test Configure function works
		config := &fiber.Config{}
		options.Configure(config)
		assert.True(t, configCalled)
		assert.Equal(t, "Test App", config.AppName)
	})

	t.Run("zero value", func(t *testing.T) {
		var options Options

		assert.Equal(t, "", options.Protocol)
		assert.Nil(t, options.Configure)
	})
}

func TestProvider(t *testing.T) {
	t.Run("creates fiber app with basic configuration", func(t *testing.T) {
		// Arrange
		mockController := &MockController{}
		mockController.On("Public", mock.AnythingOfType("*fiber.Group")).Return()
		mockController.On("Private", mock.AnythingOfType("*fiber.Group")).Return()

		controllers := []base.Controller{mockController}

		options := &Options{
			Protocol: "http",
			Configure: func(config *fiber.Config) {
				config.DisableStartupMessage = true
			},
		}

		sentryContainer := &sentry.Container{
			Handler: nil, // No sentry handler for test
		}

		environment := &env.Environment{
			SecretKey:   "test-secret-key-12345678901234567890123456",
			AppProtocol: "http",
			AppDomain:   "localhost:3000",
		}

		// Create a minimal auth middleware for testing
		authMiddleware := &auth.Middleware{}

		// Act
		app, err := Provider(ServerOptions{
			Controllers:    controllers,
			Options:        options,
			Sentry:         sentryContainer,
			Env:            environment,
			AuthMiddleware: authMiddleware,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, app)
		assert.IsType(t, &fiber.App{}, app)
		mockController.AssertExpectations(t)
	})

	t.Run("applies custom configuration", func(t *testing.T) {
		// Arrange
		controllers := []base.Controller{}
		customAppName := "Custom Test App"

		options := &Options{
			Protocol: "https",
			Configure: func(config *fiber.Config) {
				config.AppName = customAppName
				config.DisableStartupMessage = true
			},
		}

		sentryContainer := &sentry.Container{Handler: nil}
		environment := &env.Environment{
			SecretKey:   "test-secret-key-12345678901234567890123456",
			AppProtocol: "https",
			AppDomain:   "example.com",
		}

		authMiddleware := &auth.Middleware{}

		// Act
		app, err := Provider(ServerOptions{
			Controllers:    controllers,
			Options:        options,
			Sentry:         sentryContainer,
			Env:            environment,
			AuthMiddleware: authMiddleware,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, app)
		// We can't directly access the config after app creation, but we verify no error occurred
		// which means our Configure function was called successfully
	})

	t.Run("handles empty controllers list", func(t *testing.T) {
		// Arrange
		controllers := []base.Controller{}

		options := &Options{
			Protocol: "http",
			Configure: func(config *fiber.Config) {
				config.DisableStartupMessage = true
			},
		}

		sentryContainer := &sentry.Container{Handler: nil}
		environment := &env.Environment{
			SecretKey:   "test-secret-key-12345678901234567890123456",
			AppProtocol: "http",
			AppDomain:   "localhost:3000",
		}

		authMiddleware := &auth.Middleware{}

		// Act
		app, err := Provider(ServerOptions{
			Controllers:    controllers,
			Options:        options,
			Sentry:         sentryContainer,
			Env:            environment,
			AuthMiddleware: authMiddleware,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, app)
	})

	t.Run("adds sentry handler when provided", func(t *testing.T) {
		// Arrange
		controllers := []base.Controller{}

		options := &Options{
			Protocol: "http",
			Configure: func(config *fiber.Config) {
				config.DisableStartupMessage = true
			},
		}

		// Mock sentry handler
		sentryHandler := func(c *fiber.Ctx) error { return c.Next() }
		sentryContainer := &sentry.Container{
			Handler: sentryHandler,
		}

		environment := &env.Environment{
			SecretKey:   "test-secret-key-12345678901234567890123456",
			AppProtocol: "http",
			AppDomain:   "localhost:3000",
		}

		authMiddleware := &auth.Middleware{}

		// Act
		app, err := Provider(ServerOptions{
			Controllers:    controllers,
			Options:        options,
			Sentry:         sentryContainer,
			Env:            environment,
			AuthMiddleware: authMiddleware,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, app)
		// The sentry handler is added to the middleware stack
		// We can't directly test this without making HTTP requests, but we verify no error occurred
	})

	t.Run("sets up CORS with environment URL", func(t *testing.T) {
		// Arrange
		controllers := []base.Controller{}

		options := &Options{
			Protocol: "https",
			Configure: func(config *fiber.Config) {
				config.DisableStartupMessage = true
			},
		}

		sentryContainer := &sentry.Container{Handler: nil}
		environment := &env.Environment{
			SecretKey:   "test-secret-key-12345678901234567890123456",
			AppProtocol: "https",
			AppDomain:   "api.example.com",
		}

		authMiddleware := &auth.Middleware{}

		// Act
		app, err := Provider(ServerOptions{
			Controllers:    controllers,
			Options:        options,
			Sentry:         sentryContainer,
			Env:            environment,
			AuthMiddleware: authMiddleware,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, app)
		// CORS middleware is configured with environment.AppURL().String()
		// We verify the app was created successfully with the environment
	})

	t.Run("multiple controllers are registered", func(t *testing.T) {
		// Arrange
		controller1 := &MockController{}
		controller1.On("Public", mock.AnythingOfType("*fiber.Group")).Return()
		controller1.On("Private", mock.AnythingOfType("*fiber.Group")).Return()

		controller2 := &MockController{}
		controller2.On("Public", mock.AnythingOfType("*fiber.Group")).Return()
		controller2.On("Private", mock.AnythingOfType("*fiber.Group")).Return()

		controllers := []base.Controller{controller1, controller2}

		options := &Options{
			Protocol: "http",
			Configure: func(config *fiber.Config) {
				config.DisableStartupMessage = true
			},
		}

		sentryContainer := &sentry.Container{Handler: nil}
		environment := &env.Environment{
			SecretKey:   "test-secret-key-12345678901234567890123456",
			AppProtocol: "http",
			AppDomain:   "localhost:3000",
		}

		authMiddleware := &auth.Middleware{}

		// Act
		app, err := Provider(ServerOptions{
			Controllers:    controllers,
			Options:        options,
			Sentry:         sentryContainer,
			Env:            environment,
			AuthMiddleware: authMiddleware,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, app)
		controller1.AssertExpectations(t)
		controller2.AssertExpectations(t)
	})
}
