package userschemas

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"polimane/backend/model/modelbase"
)

func TestHasAccess(t *testing.T) {
	client, mock, cleanup := setupTest(t)
	defer cleanup()

	ctx := context.Background()
	userID := modelbase.MustStringToID("550e8400-e29b-41d4-a716-446655440001")
	schemaID := modelbase.MustStringToID("550e8400-e29b-41d4-a716-446655440002")

	t.Run("has access", func(t *testing.T) {
		mock.ExpectQuery(`SELECT 1 AS exists FROM "user_schemas" WHERE user_id = \$1 AND schema_id = \$2`).
			WithArgs(userID, schemaID).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

		err := client.HasAccess(ctx, userID, schemaID)

		require.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("no access", func(t *testing.T) {
		mock.ExpectQuery(`SELECT 1 AS exists FROM "user_schemas" WHERE user_id = \$1 AND schema_id = \$2`).
			WithArgs(userID, schemaID).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}))

		err := client.HasAccess(ctx, userID, schemaID)

		assert.Error(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT 1 AS exists FROM "user_schemas" WHERE user_id = \$1 AND schema_id = \$2`).
			WithArgs(userID, schemaID).
			WillReturnError(assert.AnError)

		err := client.HasAccess(ctx, userID, schemaID)

		assert.Error(t, err)
		assert.NotEqual(t, gorm.ErrRecordNotFound, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
