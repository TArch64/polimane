package userschemas

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"polimane/backend/model"
)

func TestDeleteTx(t *testing.T) {
	client, mock, cleanup := setupTest(t)
	defer cleanup()

	userID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440001")
	schemaID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440002")

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "user_schemas" WHERE user_id = \$1 AND schema_id = \$2`).
			WithArgs(userID, schemaID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		err := client.DeleteTx(client.db, userID, schemaID)

		require.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "user_schemas" WHERE user_id = \$1 AND schema_id = \$2`).
			WithArgs(userID, schemaID).
			WillReturnError(assert.AnError)
		mock.ExpectRollback()

		err := client.DeleteTx(client.db, userID, schemaID)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("no rows affected", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "user_schemas" WHERE user_id = \$1 AND schema_id = \$2`).
			WithArgs(userID, schemaID).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		err := client.DeleteTx(client.db, userID, schemaID)

		require.NoError(t, err) // GORM doesn't return error for 0 affected rows
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
