package auth

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/workos/workos-go/v4/pkg/usermanagement"

	"polimane/backend/model"
)

func TestUserSession(t *testing.T) {
	t.Run("struct initialization", func(t *testing.T) {
		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		user := &model.User{
			Identifiable: &model.Identifiable{ID: testUserID},
		}

		workosUser := &usermanagement.User{
			ID:         "workos-123",
			ExternalID: testUserID.String(),
		}

		session := &UserSession{
			ID:         "session-456",
			User:       user,
			WorkosUser: workosUser,
		}

		assert.Equal(t, "session-456", session.ID)
		assert.Equal(t, user, session.User)
		assert.Equal(t, workosUser, session.WorkosUser)
	})

	t.Run("zero value", func(t *testing.T) {
		var session UserSession

		assert.Equal(t, "", session.ID)
		assert.Nil(t, session.User)
		assert.Nil(t, session.WorkosUser)
	})
}

func TestSetSession(t *testing.T) {
	t.Run("sets session in fiber context", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		req := httptest.NewRequest("GET", "/", nil)

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		user := &model.User{
			Identifiable: &model.Identifiable{ID: testUserID},
		}

		session := &UserSession{
			ID:   "test-session",
			User: user,
			WorkosUser: &usermanagement.User{
				ID:         "workos-456",
				ExternalID: testUserID.String(),
			},
		}

		// Act & Assert using middleware pattern
		app.Get("/", func(c *fiber.Ctx) error {
			SetSession(c, session)

			// Verify session is stored in locals
			storedSession := c.Locals(sessionKey)
			assert.NotNil(t, storedSession)
			assert.Equal(t, session, storedSession)
			return c.SendStatus(200)
		})

		_, err := app.Test(req)
		assert.NoError(t, err)
	})

	t.Run("overwrites existing session", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		req := httptest.NewRequest("GET", "/", nil)

		// Act & Assert using middleware pattern
		app.Get("/", func(c *fiber.Ctx) error {
			// Set initial session
			initialSession := &UserSession{ID: "initial"}
			SetSession(c, initialSession)

			// Create new session
			newSession := &UserSession{ID: "new"}
			SetSession(c, newSession)

			// Assert
			storedSession := c.Locals(sessionKey)
			assert.Equal(t, newSession, storedSession)
			assert.NotEqual(t, initialSession, storedSession)
			return c.SendStatus(200)
		})

		_, err := app.Test(req)
		assert.NoError(t, err)
	})
}

func TestGetSession(t *testing.T) {
	t.Run("retrieves session from fiber context", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		req := httptest.NewRequest("GET", "/", nil)

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		user := &model.User{
			Identifiable: &model.Identifiable{ID: testUserID},
		}

		originalSession := &UserSession{
			ID:   "retrieve-test",
			User: user,
			WorkosUser: &usermanagement.User{
				ID:         "workos-retrieve",
				ExternalID: testUserID.String(),
			},
		}

		// Act & Assert using middleware pattern
		app.Get("/", func(c *fiber.Ctx) error {
			SetSession(c, originalSession)

			// Act
			retrievedSession := GetSession(c)

			// Assert
			assert.NotNil(t, retrievedSession)
			assert.Equal(t, originalSession, retrievedSession)
			assert.Equal(t, "retrieve-test", retrievedSession.ID)
			assert.Equal(t, user, retrievedSession.User)
			assert.Equal(t, "workos-retrieve", retrievedSession.WorkosUser.ID)
			return c.SendStatus(200)
		})

		_, err := app.Test(req)
		assert.NoError(t, err)
	})

	t.Run("returns same instance that was set", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		req := httptest.NewRequest("GET", "/", nil)

		session := &UserSession{ID: "same-instance-test"}

		// Act & Assert using middleware pattern
		app.Get("/", func(c *fiber.Ctx) error {
			SetSession(c, session)

			// Act
			retrievedSession := GetSession(c)

			// Assert
			assert.Same(t, session, retrievedSession)
			return c.SendStatus(200)
		})

		_, err := app.Test(req)
		assert.NoError(t, err)
	})
}

func TestGetSessionUser(t *testing.T) {
	t.Run("retrieves user from session", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		req := httptest.NewRequest("GET", "/", nil)

		testUserID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
		user := &model.User{
			Identifiable: &model.Identifiable{ID: testUserID},
		}

		session := &UserSession{
			ID:   "user-test",
			User: user,
			WorkosUser: &usermanagement.User{
				ID:         "workos-user-test",
				ExternalID: testUserID.String(),
			},
		}

		// Act & Assert using middleware pattern
		app.Get("/", func(c *fiber.Ctx) error {
			SetSession(c, session)

			// Act
			retrievedUser := GetSessionUser(c)

			// Assert
			assert.NotNil(t, retrievedUser)
			assert.Equal(t, user, retrievedUser)
			assert.Equal(t, testUserID, retrievedUser.ID)
			return c.SendStatus(200)
		})

		_, err := app.Test(req)
		assert.NoError(t, err)
	})

	t.Run("returns same user instance from session", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		req := httptest.NewRequest("GET", "/", nil)

		user := &model.User{
			Identifiable: &model.Identifiable{
				ID: model.MustStringToID("550e8400-e29b-41d4-a716-446655440000"),
			},
		}

		session := &UserSession{
			ID:   "user-instance-test",
			User: user,
		}

		// Act & Assert using middleware pattern
		app.Get("/", func(c *fiber.Ctx) error {
			SetSession(c, session)

			// Act
			retrievedUser := GetSessionUser(c)

			// Assert
			assert.Same(t, user, retrievedUser)
			return c.SendStatus(200)
		})

		_, err := app.Test(req)
		assert.NoError(t, err)
	})
}

func TestSessionKey(t *testing.T) {
	t.Run("session key is defined", func(t *testing.T) {
		// The sessionKey should be a zero-value UserSession struct
		// used as a unique key for storing session data in fiber locals
		assert.NotNil(t, sessionKey)
		assert.IsType(t, UserSession{}, sessionKey)
	})

	t.Run("session key is zero value", func(t *testing.T) {
		assert.Equal(t, "", sessionKey.ID)
		assert.Nil(t, sessionKey.User)
		assert.Nil(t, sessionKey.WorkosUser)
	})
}
