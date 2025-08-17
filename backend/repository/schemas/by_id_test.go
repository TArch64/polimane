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
	mockS3 := &MockS3Client{}
	client := &Impl{
		db:          gormDB,
		userSchemas: mockUserSchemas,
		signals:     signal.Provider(),
		s3:          mockS3,
	}

	cleanup := func() {
		db.Close()
	}

	return client, mockUserSchemas, mock, cleanup
}

func TestByID(t *testing.T) {
	ctx := context.Background()
	userID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
	user := &model.User{Identifiable: &model.Identifiable{ID: userID}}
	schemaID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440001")

	t.Run("success", func(t *testing.T) {
		client, mockUserSchemas, mock, cleanup := setupTest(t)
		defer cleanup()

		expectedSchema := &model.Schema{
			Identifiable: &model.Identifiable{ID: schemaID},
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
		client, mockUserSchemas, mock, cleanup := setupTest(t)
		defer cleanup()

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
		client, mockUserSchemas, mock, cleanup := setupTest(t)
		defer cleanup()

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

	t.Run("access denied", func(t *testing.T) {
		client, mockUserSchemas, _, cleanup := setupTest(t)
		defer cleanup()

		accessError := assert.AnError

		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(accessError)

		result, err := client.ByID(&ByIDOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
		})

		assert.Error(t, err)
		assert.Equal(t, accessError, err)
		assert.Nil(t, result)
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("database connection error", func(t *testing.T) {
		client, mockUserSchemas, mock, cleanup := setupTest(t)
		defer cleanup()

		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)

		mock.ExpectQuery(`SELECT \* FROM "schemas"`).
			WithArgs(schemaID, 1).
			WillReturnError(assert.AnError)

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

	t.Run("with select fields - access denied", func(t *testing.T) {
		client, mockUserSchemas, _, cleanup := setupTest(t)
		defer cleanup()

		accessError := assert.AnError

		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(accessError)

		result, err := client.ByID(&ByIDOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
			Select:   []string{"name"},
		})

		assert.Error(t, err)
		assert.Equal(t, accessError, err)
		assert.Nil(t, result)
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("with select fields - database error", func(t *testing.T) {
		client, mockUserSchemas, mock, cleanup := setupTest(t)
		defer cleanup()

		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)

		mock.ExpectQuery(`SELECT "name" FROM "schemas"`).
			WithArgs(schemaID, 1).
			WillReturnError(assert.AnError)

		result, err := client.ByID(&ByIDOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
			Select:   []string{"name"},
		})

		assert.Error(t, err)
		assert.NotNil(t, result) // GORM returns empty struct even on error
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})
}
