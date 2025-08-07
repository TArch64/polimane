package signal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProvider(t *testing.T) {
	t.Run("creates container with all signals initialized", func(t *testing.T) {
		container := Provider()

		require.NotNil(t, container)
		assert.NotNil(t, container.InvalidateUserCache)
		assert.NotNil(t, container.InvalidateWorkosUserCache)
		assert.NotNil(t, container.InvalidateAuthCache)
	})

	t.Run("returns new instance on each call", func(t *testing.T) {
		container1 := Provider()
		container2 := Provider()

		assert.NotSame(t, container1, container2)
		assert.NotSame(t, container1.InvalidateUserCache, container2.InvalidateUserCache)
		assert.NotSame(t, container1.InvalidateWorkosUserCache, container2.InvalidateWorkosUserCache)
		assert.NotSame(t, container1.InvalidateAuthCache, container2.InvalidateAuthCache)
	})
}

func TestContainer(t *testing.T) {
	t.Run("zero value container has nil signals", func(t *testing.T) {
		var container Container

		assert.Nil(t, container.InvalidateUserCache)
		assert.Nil(t, container.InvalidateWorkosUserCache)
		assert.Nil(t, container.InvalidateAuthCache)
	})
}
