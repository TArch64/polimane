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

func TestMakeCopyName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "first copy",
			input:    "Original Schema",
			expected: "Original Schema (1)",
		},
		{
			name:     "increment existing counter",
			input:    "Schema (1)",
			expected: "Schema  (2)",
		},
		{
			name:     "increment higher counter",
			input:    "Schema (10)",
			expected: "Schema  (11)",
		},
		{
			name:     "parentheses without number",
			input:    "Schema (name)",
			expected: "Schema (name) (1)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := makeCopyName(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCopy(t *testing.T) {
	client, mockUserSchemas, mock, cleanup := setupTest(t)
	defer cleanup()

	ctx := context.Background()
	userID := modelbase.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
	user := &model.User{Identifiable: &modelbase.Identifiable{ID: userID}}
	schemaID := modelbase.MustStringToID("550e8400-e29b-41d4-a716-446655440001")

	t.Run("success", func(t *testing.T) {
		// Mock HasAccess for ByID call
		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)

		// Mock ByID query
		mock.ExpectQuery(`SELECT \* FROM "schemas"`).
			WithArgs(schemaID, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "palette", "content"}).
				AddRow("550e8400-e29b-41d4-a716-446655440001", "Original Schema", `["#ffffff", "#000000"]`, `[{"id": "1", "name": "Pattern 1", "type": "square"}]`))

		// Mock Create transaction
		mockUserSchemas.On("CreateTx", tmock.Anything, userID, tmock.Anything).Return(nil)
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "schemas"`).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), "Original Schema (1)", sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("550e8400-e29b-41d4-a716-446655440002"))
		mock.ExpectCommit()

		result, err := client.Copy(&CopyOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
		})

		require.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Original Schema (1)", result.Name)
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})

	t.Run("original not found", func(t *testing.T) {
		// Mock HasAccess for ByID call
		mockUserSchemas.On("HasAccess", ctx, userID, schemaID).Return(nil)

		// Mock ByID query returning error
		mock.ExpectQuery(`SELECT \* FROM "schemas"`).
			WithArgs(schemaID, 1).
			WillReturnError(assert.AnError)

		result, err := client.Copy(&CopyOptions{
			Ctx:      ctx,
			User:     user,
			SchemaID: schemaID,
		})

		assert.Error(t, err)
		assert.Nil(t, result) // Should return nil on error in Copy function
		assert.NoError(t, mock.ExpectationsWereMet())
		mockUserSchemas.AssertExpectations(t)
	})
}
