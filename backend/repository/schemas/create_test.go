package schemas

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
)

func TestCreate(t *testing.T) {
	client, mockUserSchemas, mock, cleanup := setupTest(t)
	defer cleanup()

	ctx := context.Background()
	userID := modelbase.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
	user := &model.User{Identifiable: &modelbase.Identifiable{ID: userID}}

	t.Run("success", func(t *testing.T) {
		mockUserSchemas.On("CreateTx", tmock.Anything, userID, tmock.Anything).Return(nil)

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "schemas"`).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "Test Schema", sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("550e8400-e29b-41d4-a716-446655440001"))
		mock.ExpectCommit()

		result, err := client.Create(&CreateOptions{
			Ctx:     ctx,
			User:    user,
			Name:    "Test Schema",
			Palette: model.SchemaPalette{"#ffffff", "#000000"},
			Content: model.SchemaContent{&model.SchemaPattern{ID: "1", Name: "Pattern 1", Type: "square"}},
		})

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Test Schema", result.Name)
		assert.NotNil(t, result.Palette)
		assert.NotNil(t, result.Content)
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("with nil palette and content", func(t *testing.T) {
		mockUserSchemas.On("CreateTx", tmock.Anything, userID, tmock.Anything).Return(nil)

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "schemas"`).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "Test Schema", sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("550e8400-e29b-41d4-a716-446655440001"))
		mock.ExpectCommit()

		result, err := client.Create(&CreateOptions{
			Ctx:     ctx,
			User:    user,
			Name:    "Test Schema",
			Palette: nil,
			Content: nil,
		})

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Test Schema", result.Name)
		assert.NotNil(t, result.Palette)
		assert.NotNil(t, result.Content)
		assert.Len(t, result.Palette, model.SchemaPaletteSize)
		assert.Len(t, result.Content, 0)
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "schemas"`).
			WillReturnError(assert.AnError)
		mock.ExpectRollback()

		result, err := client.Create(&CreateOptions{
			Ctx:  ctx,
			User: user,
			Name: "Test Schema",
		})

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})
}
