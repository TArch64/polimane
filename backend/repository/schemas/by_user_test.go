package schemas

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"polimane/backend/model"
)

func TestByUser(t *testing.T) {
	client, _, mock, cleanup := setupTest(t)
	defer cleanup()

	ctx := context.Background()
	userID := model.MustStringToID("550e8400-e29b-41d4-a716-446655440000")
	user := &model.User{Identifiable: &model.Identifiable{ID: userID}}

	t.Run("success", func(t *testing.T) {
		schemaID1 := model.MustStringToID("550e8400-e29b-41d4-a716-446655440001")
		schemaID2 := model.MustStringToID("550e8400-e29b-41d4-a716-446655440002")
		fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

		mock.ExpectQuery(`SELECT "schemas"\."id","schemas"\."created_at","schemas"\."updated_at","schemas"\."name","schemas"\."palette","schemas"\."content" FROM "schemas" JOIN user_schemas ON user_schemas\.schema_id = schemas\.id AND user_schemas\.user_id = \$1`).
			WithArgs(user.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "palette", "content"}).
				AddRow("550e8400-e29b-41d4-a716-446655440001", fixedTime, fixedTime, "Schema 1", `[]`, `[]`).
				AddRow("550e8400-e29b-41d4-a716-446655440002", fixedTime, fixedTime, "Schema 2", `[]`, `[]`))

		result, err := client.ByUser(&ByUserOptions{
			Ctx:  ctx,
			User: user,
		})

		require.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, schemaID1, result[0].ID)
		assert.Equal(t, "Schema 1", result[0].Name)
		assert.Equal(t, schemaID2, result[1].ID)
		assert.Equal(t, "Schema 2", result[1].Name)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("with select fields", func(t *testing.T) {
		mock.ExpectQuery(`SELECT "name" FROM "schemas" JOIN user_schemas ON user_schemas\.schema_id = schemas\.id AND user_schemas\.user_id = \$1`).
			WithArgs(user.ID).
			WillReturnRows(sqlmock.NewRows([]string{"name"}).
				AddRow("Schema 1"))

		result, err := client.ByUser(&ByUserOptions{
			Ctx:    ctx,
			User:   user,
			Select: []string{"name"},
		})

		require.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "Schema 1", result[0].Name)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("no schemas found", func(t *testing.T) {
		mock.ExpectQuery(`SELECT "schemas"\."id","schemas"\."created_at","schemas"\."updated_at","schemas"\."name","schemas"\."palette","schemas"\."content" FROM "schemas" JOIN user_schemas ON user_schemas\.schema_id = schemas\.id AND user_schemas\.user_id = \$1`).
			WithArgs(user.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "palette", "content"}))

		result, err := client.ByUser(&ByUserOptions{
			Ctx:  ctx,
			User: user,
		})

		require.NoError(t, err)
		assert.Len(t, result, 0)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
