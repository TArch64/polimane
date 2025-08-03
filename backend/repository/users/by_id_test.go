package users

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
)

func TestByID(t *testing.T) {
	client, mock, cleanup := setupTest(t)
	defer cleanup()

	ctx := context.Background()
	userID := modelbase.MustStringToID("550e8400-e29b-41d4-a716-446655440000")

	t.Run("success", func(t *testing.T) {
		expectedUser := &model.User{
			Identifiable: &modelbase.Identifiable{ID: userID},
			WorkosID:     "workos_user_123",
		}

		mock.ExpectQuery(`SELECT \* FROM "users"`).
			WithArgs(userID, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "workos_id"}).
				AddRow(userID.String(), "2023-01-01T00:00:00Z", "2023-01-01T00:00:00Z", "workos_user_123"))

		result, err := client.ByID(ctx, userID)

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedUser.ID, result.ID)
		assert.Equal(t, expectedUser.WorkosID, result.WorkosID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("not found", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM "users"`).
			WithArgs(userID, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		result, err := client.ByID(ctx, userID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM "users"`).
			WithArgs(userID, 1).
			WillReturnError(assert.AnError)

		result, err := client.ByID(ctx, userID)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
