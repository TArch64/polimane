package userschemas

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"polimane/backend/model/modelbase"
)

func TestCreateTx(t *testing.T) {
	client, mock, cleanup := setupTest(t)
	defer cleanup()

	userID := modelbase.MustStringToID("550e8400-e29b-41d4-a716-446655440001")
	schemaID := modelbase.MustStringToID("550e8400-e29b-41d4-a716-446655440002")

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "user_schemas"`).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), userID, schemaID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := client.CreateTx(client.db, userID, schemaID)

		require.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "user_schemas"`).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), userID, schemaID).
			WillReturnError(assert.AnError)
		mock.ExpectRollback()

		err := client.CreateTx(client.db, userID, schemaID)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
