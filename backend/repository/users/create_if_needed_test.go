package users

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"polimane/backend/model"
)

func TestCreateIfNeeded(t *testing.T) {
	client, mock, cleanup := setupTest(t)
	defer cleanup()

	ctx := context.Background()
	workosID := "workos_user_123"

	t.Run("creates new user when not exists", func(t *testing.T) {
		userID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")

		// Expect SELECT query to check if user exists (returns no rows)
		mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"\."workos_id" = \$1 ORDER BY "users"\."id" LIMIT \$2`).
			WithArgs(workosID, 1).
			WillReturnRows(sqlmock.NewRows([]string{}))

		// Expect transaction for INSERT
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "users"`).
			WithArgs(workosID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(userID.String()))
		mock.ExpectCommit()

		result, err := client.CreateIfNeeded(ctx, workosID)

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, workosID, result.WorkosID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("returns existing user when found", func(t *testing.T) {
		userID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")

		// Expect SELECT query to check if user exists (returns existing user)
		mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"\."workos_id" = \$1 ORDER BY "users"\."id" LIMIT \$2`).
			WithArgs(workosID, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "workos_id"}).
				AddRow(userID.String(), "2023-01-01T00:00:00Z", "2023-01-01T00:00:00Z", workosID))

		result, err := client.CreateIfNeeded(ctx, workosID)

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, workosID, result.WorkosID)
		assert.Equal(t, userID, result.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("database error on select", func(t *testing.T) {
		// Expect SELECT query to fail
		mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"\."workos_id" = \$1 ORDER BY "users"\."id" LIMIT \$2`).
			WithArgs(workosID, 1).
			WillReturnError(assert.AnError)

		result, err := client.CreateIfNeeded(ctx, workosID)

		assert.Error(t, err)
		assert.NotNil(t, result) // GORM returns empty struct even on error
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("database error on insert", func(t *testing.T) {
		// Expect SELECT query to check if user exists (returns no rows)
		mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"\."workos_id" = \$1 ORDER BY "users"\."id" LIMIT \$2`).
			WithArgs(workosID, 1).
			WillReturnRows(sqlmock.NewRows([]string{}))

		// Expect transaction for INSERT that fails
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "users"`).
			WithArgs(workosID).
			WillReturnError(assert.AnError)
		mock.ExpectRollback()

		result, err := client.CreateIfNeeded(ctx, workosID)

		assert.Error(t, err)
		assert.NotNil(t, result) // GORM returns empty struct even on error
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
