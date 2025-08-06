package auth

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"polimane/backend/api/base"
	"polimane/backend/env"
	"polimane/backend/services/workos"
	"polimane/backend/signal"
)

func TestControllerProvider(t *testing.T) {
	t.Run("creates controller with dependencies", func(t *testing.T) {
		// Arrange
		workosClient := &workos.Client{}
		environment := &env.Environment{}
		usersClient := &MockUsersClient{}
		signalsContainer := &signal.Container{}

		// Act
		controller := Provider(ControllerOptions{
			WorkosClient: workosClient,
			Env:          environment,
			Users:        usersClient,
			Signals:      signalsContainer,
		})

		// Assert
		assert.NotNil(t, controller)
		assert.IsType(t, &Controller{}, controller)

		// Verify it implements base.Controller interface
		var _ base.Controller = controller
	})

	t.Run("controller has correct internal state", func(t *testing.T) {
		// Arrange
		workosClient := &workos.Client{}
		environment := &env.Environment{}
		usersClient := &MockUsersClient{}
		signalsContainer := &signal.Container{}

		// Act
		controller := Provider(ControllerOptions{
			WorkosClient: workosClient,
			Env:          environment,
			Users:        usersClient,
			Signals:      signalsContainer,
		}).(*Controller)

		// Assert
		assert.Equal(t, workosClient, controller.workosClient)
		assert.Equal(t, environment, controller.env)
		assert.Equal(t, usersClient, controller.users)
		assert.Equal(t, signalsContainer, controller.signals)
	})
}

func TestControllerPublic(t *testing.T) {
	t.Run("registers public routes correctly", func(t *testing.T) {
		// Create test controller
		controller := &Controller{}

		// Create fiber app for testing
		app := fiber.New()
		router := app.Group("/api")

		// Act
		controller.Public(router)

		// Get routes
		routes := app.GetRoutes()

		// Assert that login routes are registered
		var foundLoginRoute, foundLoginCompleteRoute bool
		for _, route := range routes {
			if route.Path == "/api/auth/login" && route.Method == "GET" {
				foundLoginRoute = true
			}
			if route.Path == "/api/auth/login/complete" && route.Method == "GET" {
				foundLoginCompleteRoute = true
			}
		}

		assert.True(t, foundLoginRoute, "Login route should be registered")
		assert.True(t, foundLoginCompleteRoute, "Login complete route should be registered")
	})
}

func TestControllerPrivate(t *testing.T) {
	t.Run("registers private routes correctly", func(t *testing.T) {
		// Create test controller
		controller := &Controller{}

		// Create fiber app for testing
		app := fiber.New()
		router := app.Group("/api")

		// Act
		controller.Private(router)

		// Get routes
		routes := app.GetRoutes()

		// Assert that logout route is registered
		var foundLogoutRoute bool
		for _, route := range routes {
			if route.Path == "/api/auth/logout" && route.Method == "POST" {
				foundLogoutRoute = true
			}
		}

		assert.True(t, foundLogoutRoute, "Logout route should be registered")
	})
}

func TestGroupPrefix(t *testing.T) {
	t.Run("group prefix is correct", func(t *testing.T) {
		assert.Equal(t, "auth", groupPrefix)
	})
}
