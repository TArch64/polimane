package schemas

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"polimane/backend/signal"
)

func TestProvider(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)

	userSchemas := &MockUserSchemas{}
	signals := &signal.Container{}

	client := Provider(gormDB, userSchemas, signals)

	assert.NotNil(t, client)
	assert.IsType(t, &Impl{}, client)

	impl := client.(*Impl)
	assert.Equal(t, gormDB, impl.db)
	assert.Equal(t, userSchemas, impl.userSchemas)
	assert.Equal(t, signals, impl.signals)

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
	assert.NotNil(t, (*Impl).ByUser)
	assert.NotNil(t, (*Impl).Copy)
	assert.NotNil(t, (*Impl).Create)
	assert.NotNil(t, (*Impl).Delete)
	assert.NotNil(t, (*Impl).Update)
}
