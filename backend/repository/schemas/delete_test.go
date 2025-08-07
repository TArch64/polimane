package schemas

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"polimane/backend/model"
)

func TestDelete(t *testing.T) {
	ctx := context.Background()
	userID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
	user := &model.User{Identifiable: &model.Identifiable{ID: userID}}
	schemaID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440001")

	t.Run("success", func(t *testing.T) {
		client, mockUserSchemas, mock, cleanup := setupTest(t)
		defer cleanup()

		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)
		mockUserSchemas.On("DeleteTx", tmock.Anything, userID, schemaID).Return(nil)

		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "schemas"`).
			WithArgs(schemaID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		err := client.Delete(&DeleteOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
		})

		require.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {
		client, mockUserSchemas, mock, cleanup := setupTest(t)
		defer cleanup()

		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)

		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "schemas"`).
			WithArgs(schemaID).
			WillReturnError(assert.AnError)
		mock.ExpectRollback()

		err := client.Delete(&DeleteOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
		})

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("no rows affected", func(t *testing.T) {
		client, mockUserSchemas, mock, cleanup := setupTest(t)
		defer cleanup()

		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)
		mockUserSchemas.On("DeleteTx", tmock.Anything, userID, schemaID).Return(nil)

		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "schemas"`).
			WithArgs(schemaID).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		err := client.Delete(&DeleteOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
		})

		require.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("access denied", func(t *testing.T) {
		client, mockUserSchemas, _, cleanup := setupTest(t)
		defer cleanup()

		accessError := assert.AnError
		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(accessError)

		err := client.Delete(&DeleteOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
		})

		assert.Error(t, err)
		assert.Equal(t, accessError, err)
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("delete tx error", func(t *testing.T) {
		client, mockUserSchemas, mock, cleanup := setupTest(t)
		defer cleanup()

		deleteTxError := assert.AnError
		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)
		mockUserSchemas.On("DeleteTx", tmock.Anything, userID, schemaID).Return(deleteTxError)

		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "schemas"`).
			WithArgs(schemaID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectRollback()

		err := client.Delete(&DeleteOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
		})

		assert.Error(t, err)
		assert.Equal(t, deleteTxError, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("transaction error handling", func(t *testing.T) {
		client, mockUserSchemas, mock, cleanup := setupTest(t)
		defer cleanup()

		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)
		mockUserSchemas.On("DeleteTx", tmock.Anything, userID, schemaID).Return(nil)

		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "schemas"`).
			WithArgs(schemaID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit().WillReturnError(assert.AnError)

		err := client.Delete(&DeleteOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
		})

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})
}
