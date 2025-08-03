package users

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTest(t *testing.T) (*Impl, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)

	client := &Impl{
		db: gormDB,
	}

	cleanup := func() {
		db.Close()
	}

	return client, mock, cleanup
}

func TestProvider(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)

	client := Provider(gormDB)

	assert.NotNil(t, client)
	assert.IsType(t, &Impl{}, client)

	impl := client.(*Impl)
	assert.Equal(t, gormDB, impl.db)

	// Verify that client implements the Client interface
	var _ Client = client

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestImplImplementsClientInterface(t *testing.T) {
	// This test ensures that Impl struct implements all methods of the Client interface
	var impl *Impl
	var _ Client = impl

	// Test that all interface methods exist
	assert.NotNil(t, (*Impl).ByID)
	assert.NotNil(t, (*Impl).CreateIfNeeded)
}
