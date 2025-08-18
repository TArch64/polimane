package schemas

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/datatypes"

	"polimane/backend/model"
)

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	userID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
	user := &model.User{Identifiable: &model.Identifiable{ID: userID}}
	schemaID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440001")

	t.Run("success", func(t *testing.T) {
		client, mockUserSchemas, mock, cleanup := setupTest(t)
		defer cleanup()

		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)

		palette := model.SchemaPalette{"#000000", "#ffffff"}
		updates := &model.Schema{
			Name:    "Updated Schema",
			Palette: datatypes.NewJSONType(palette),
		}

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "schemas" SET`).
			WithArgs(sqlmock.AnyArg(), "Updated Schema", sqlmock.AnyArg(), schemaID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		err := client.Update(&UpdateOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
			Updates:  updates,
		})

		require.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {
		client, mockUserSchemas, mock, cleanup := setupTest(t)
		defer cleanup()

		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)

		updates := &model.Schema{
			Name: "Updated Schema",
		}

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "schemas" SET`).
			WithArgs(sqlmock.AnyArg(), "Updated Schema", schemaID).
			WillReturnError(assert.AnError)
		mock.ExpectRollback()

		err := client.Update(&UpdateOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
			Updates:  updates,
		})

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("no rows affected", func(t *testing.T) {
		client, mockUserSchemas, mock, cleanup := setupTest(t)
		defer cleanup()

		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)

		updates := &model.Schema{
			Name: "Updated Schema",
		}

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "schemas" SET`).
			WithArgs(sqlmock.AnyArg(), "Updated Schema", schemaID).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		err := client.Update(&UpdateOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
			Updates:  updates,
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

		updates := &model.Schema{
			Name: "Updated Schema",
		}

		err := client.Update(&UpdateOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
			Updates:  updates,
		})

		assert.Error(t, err)
		assert.Equal(t, accessError, err)
		mockUserSchemas.AssertExpectations(t)
	})
}
