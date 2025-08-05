package modelbase

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringToID(t *testing.T) {
	t.Run("valid UUID string", func(t *testing.T) {
		validUUID := "550e8400-e29b-41d4-a716-446655440000"

		id, err := StringToID(validUUID)

		require.NoError(t, err)
		assert.True(t, id.Valid)
		assert.Equal(t, validUUID, id.String())
	})

	t.Run("empty string", func(t *testing.T) {
		id, err := StringToID("")

		assert.Error(t, err)
		assert.False(t, id.Valid)
	})

	t.Run("invalid UUID string", func(t *testing.T) {
		invalidUUID := "invalid-uuid-string"

		id, err := StringToID(invalidUUID)

		assert.Error(t, err)
		assert.False(t, id.Valid)
	})

	t.Run("UUID without hyphens", func(t *testing.T) {
		uuidWithoutHyphens := "550e8400e29b41d4a716446655440000"

		id, err := StringToID(uuidWithoutHyphens)

		require.NoError(t, err)
		assert.True(t, id.Valid)
		assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", id.String())
	})

	t.Run("uppercase UUID", func(t *testing.T) {
		uppercaseUUID := "550E8400-E29B-41D4-A716-446655440000"

		id, err := StringToID(uppercaseUUID)

		require.NoError(t, err)
		assert.True(t, id.Valid)
		assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", id.String())
	})

	t.Run("malformed UUID", func(t *testing.T) {
		malformedUUID := "550e8400-e29b-41d4-a716"

		id, err := StringToID(malformedUUID)

		assert.Error(t, err)
		assert.False(t, id.Valid)
	})
}

func TestMustStringToID(t *testing.T) {
	t.Run("valid UUID string", func(t *testing.T) {
		validUUID := "550e8400-e29b-41d4-a716-446655440000"

		id := MustStringToID(validUUID)

		assert.True(t, id.Valid)
		assert.Equal(t, validUUID, id.String())
	})

	t.Run("empty string panics", func(t *testing.T) {
		assert.Panics(t, func() {
			MustStringToID("")
		})
	})

	t.Run("invalid UUID string panics", func(t *testing.T) {
		invalidUUID := "invalid-uuid-string"

		assert.Panics(t, func() {
			MustStringToID(invalidUUID)
		})
	})

	t.Run("malformed UUID panics", func(t *testing.T) {
		malformedUUID := "550e8400-e29b-41d4-a716"

		assert.Panics(t, func() {
			MustStringToID(malformedUUID)
		})
	})
}

func TestIdentifiable(t *testing.T) {
	t.Run("struct initialization", func(t *testing.T) {
		testID := MustStringToID("550e8400-e29b-41d4-a716-446655440000")

		identifiable := Identifiable{
			ID: testID,
		}

		assert.Equal(t, testID, identifiable.ID)
		assert.True(t, identifiable.ID.Valid)
		assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", identifiable.ID.String())
	})

	t.Run("zero value", func(t *testing.T) {
		var identifiable Identifiable

		assert.False(t, identifiable.ID.Valid)
		assert.Equal(t, pgtype.UUID{}, identifiable.ID)
	})
}

func TestIDType(t *testing.T) {
	t.Run("ID is pgtype.UUID", func(t *testing.T) {
		var id ID
		var pgUUID pgtype.UUID

		// Verify ID is the same type as pgtype.UUID
		assert.IsType(t, pgUUID, id)
	})

	t.Run("ID assignment from pgtype.UUID", func(t *testing.T) {
		pgUUID := pgtype.UUID{}
		err := pgUUID.Scan("550e8400-e29b-41d4-a716-446655440000")
		require.NoError(t, err)

		var id ID = pgUUID

		assert.True(t, id.Valid)
		assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", id.String())
	})
}
