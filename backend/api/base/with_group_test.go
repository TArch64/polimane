package base

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestWithGroup(t *testing.T) {
	t.Run("creates and configures route group", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		path := "/api/v1"

		// Act
		WithGroup(app, path, func(group fiber.Router) {
			group.Get("/test", func(c *fiber.Ctx) error {
				return c.SendString("success")
			})
		})

		req := httptest.NewRequest("GET", "/api/v1/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("allows nested route registration", func(t *testing.T) {
		// Arrange
		app := fiber.New()

		// Act
		WithGroup(app, "/api", func(apiGroup fiber.Router) {
			apiGroup.Get("/health", func(c *fiber.Ctx) error {
				return c.SendString("healthy")
			})

			WithGroup(apiGroup, "/v1", func(v1Group fiber.Router) {
				v1Group.Get("/users", func(c *fiber.Ctx) error {
					return c.SendString("users")
				})
			})
		})

		// Test root group route
		req1 := httptest.NewRequest("GET", "/api/health", nil)
		resp1, err := app.Test(req1)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp1.StatusCode)

		// Test nested group route
		req2 := httptest.NewRequest("GET", "/api/v1/users", nil)
		resp2, err := app.Test(req2)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp2.StatusCode)
	})

	t.Run("handles empty path", func(t *testing.T) {
		// Arrange
		app := fiber.New()

		// Act
		WithGroup(app, "", func(group fiber.Router) {
			group.Get("/root", func(c *fiber.Ctx) error {
				return c.SendString("root")
			})
		})

		req := httptest.NewRequest("GET", "/root", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("handles root path", func(t *testing.T) {
		// Arrange
		app := fiber.New()

		// Act
		WithGroup(app, "/", func(group fiber.Router) {
			group.Get("/slash", func(c *fiber.Ctx) error {
				return c.SendString("slash")
			})
		})

		req := httptest.NewRequest("GET", "/slash", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("supports multiple HTTP methods", func(t *testing.T) {
		// Arrange
		app := fiber.New()

		// Act
		WithGroup(app, "/api", func(group fiber.Router) {
			group.Get("/resource", func(c *fiber.Ctx) error {
				return c.SendString("GET")
			})
			group.Post("/resource", func(c *fiber.Ctx) error {
				return c.SendString("POST")
			})
			group.Put("/resource", func(c *fiber.Ctx) error {
				return c.SendString("PUT")
			})
			group.Delete("/resource", func(c *fiber.Ctx) error {
				return c.SendString("DELETE")
			})
		})

		// Test GET
		req1 := httptest.NewRequest("GET", "/api/resource", nil)
		resp1, err := app.Test(req1)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp1.StatusCode)

		// Test POST
		req2 := httptest.NewRequest("POST", "/api/resource", nil)
		resp2, err := app.Test(req2)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp2.StatusCode)

		// Test PUT
		req3 := httptest.NewRequest("PUT", "/api/resource", nil)
		resp3, err := app.Test(req3)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp3.StatusCode)

		// Test DELETE
		req4 := httptest.NewRequest("DELETE", "/api/resource", nil)
		resp4, err := app.Test(req4)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp4.StatusCode)
	})

	t.Run("supports middleware in group", func(t *testing.T) {
		// Arrange
		app := fiber.New()

		// Act
		WithGroup(app, "/protected", func(group fiber.Router) {
			// Add middleware to the group
			group.Use(func(c *fiber.Ctx) error {
				c.Set("X-Middleware", "applied")
				return c.Next()
			})

			group.Get("/resource", func(c *fiber.Ctx) error {
				return c.SendString("protected")
			})
		})

		req := httptest.NewRequest("GET", "/protected/resource", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "applied", resp.Header.Get("X-Middleware"))
	})

	t.Run("handles path with trailing slash", func(t *testing.T) {
		// Arrange
		app := fiber.New()

		// Act
		WithGroup(app, "/api/", func(group fiber.Router) {
			group.Get("/test", func(c *fiber.Ctx) error {
				return c.SendString("test")
			})
		})

		req := httptest.NewRequest("GET", "/api/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("works with route parameters", func(t *testing.T) {
		// Arrange
		app := fiber.New()

		// Act
		WithGroup(app, "/users", func(group fiber.Router) {
			group.Get("/:id", func(c *fiber.Ctx) error {
				id := c.Params("id")
				return c.SendString("user_" + id)
			})
		})

		req := httptest.NewRequest("GET", "/users/123", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("can be called with existing group", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		existingGroup := app.Group("/api")

		// Act
		WithGroup(existingGroup, "/v1", func(group fiber.Router) {
			group.Get("/test", func(c *fiber.Ctx) error {
				return c.SendString("nested")
			})
		})

		req := httptest.NewRequest("GET", "/api/v1/test", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	t.Run("handler function receives correct group", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		var receivedGroup fiber.Router

		// Act
		WithGroup(app, "/test", func(group fiber.Router) {
			receivedGroup = group
			// Add a route to verify the group works
			group.Get("/verify", func(c *fiber.Ctx) error {
				return c.SendString("verified")
			})
		})

		// Assert
		assert.NotNil(t, receivedGroup)

		// Verify the group works by testing the route
		req := httptest.NewRequest("GET", "/test/verify", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
}
