package workos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImpl_MFA(t *testing.T) {
	t.Run("returns MFA client", func(t *testing.T) {
		mockMFA := &MockMFA{}
		client := &Impl{
			mfa: mockMFA,
		}

		result := client.MFA()

		assert.Equal(t, mockMFA, result)
		assert.NotNil(t, result)
	})

	t.Run("returns nil when MFA is nil", func(t *testing.T) {
		client := &Impl{
			mfa: nil,
		}

		result := client.MFA()

		assert.Nil(t, result)
	})
}
