package workos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImpl_UserManagement(t *testing.T) {
	t.Run("returns UserManagement client", func(t *testing.T) {
		mockUserMgmt := &MockUserManagement{}
		client := &Impl{
			userManagement: mockUserMgmt,
		}

		result := client.UserManagement()

		assert.Equal(t, mockUserMgmt, result)
		assert.NotNil(t, result)
	})

	t.Run("returns nil when UserManagement is nil", func(t *testing.T) {
		client := &Impl{
			userManagement: nil,
		}

		result := client.UserManagement()

		assert.Nil(t, result)
	})
}
