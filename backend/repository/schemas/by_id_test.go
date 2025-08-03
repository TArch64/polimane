package schemas

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"polimane/backend/model"
	"polimane/backend/model/modelbase"
	"polimane/backend/signal"
)

func setupTest(t *testing.T) (*Impl, *MockUserSchemas, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)

	mockUserSchemas := &MockUserSchemas{}
	client := &Impl{
		db:          gormDB,
		userSchemas: mockUserSchemas,
		signals:     signal.Provider(),
	}

	cleanup := func() {
		db.Close()
	}

	return client, mockUserSchemas, mock, cleanup
}

func TestByID(t *testing.T) {
	client, mockUserSchemas, mock, cleanup := setupTest(t)
	defer cleanup()

	ctx := context.Background()
	userID := modelbase.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
	user := &model.User{Identifiable: &modelbase.Identifiable{ID: userID}}
	schemaID := modelbase.MustStringToID("550e8400-e29b-41d4-a716-446655440001")

	t.Run("success", func(t *testing.T) {
		expectedSchema := &model.Schema{
			Identifiable: &modelbase.Identifiable{ID: schemaID},
			Name:         "Test Schema",
		}

		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)

		mock.ExpectQuery(`SELECT \* FROM "schemas"`).
			WithArgs(schemaID, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
				AddRow("550e8400-e29b-41d4-a716-446655440001", "Test Schema"))

		result, err := client.ByID(&ByIDOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
		})

		require.NoError(t, err)
		assert.Equal(t, expectedSchema.ID, result.ID)
		assert.Equal(t, expectedSchema.Name, result.Name)
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("with select fields", func(t *testing.T) {
		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)

		mock.ExpectQuery(`SELECT "name" FROM "schemas"`).
			WithArgs(schemaID, 1).
			WillReturnRows(sqlmock.NewRows([]string{"name"}).
				AddRow("Test Schema"))

		result, err := client.ByID(&ByIDOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
			Select:   []string{"name"},
		})

		require.NoError(t, err)
		assert.Equal(t, "Test Schema", result.Name)
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)

		mock.ExpectQuery(`SELECT \* FROM "schemas"`).
			WithArgs(schemaID, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		result, err := client.ByID(&ByIDOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
		})

		assert.Error(t, err)
		assert.NotNil(t, result) // GORM returns empty struct even on error
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})
}
