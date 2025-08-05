# AI Test Generation Instructions

## 🚨 CRITICAL RULES - READ FIRST

### Test Execution Command

```bash
docker compose run --rm backend make test test_pattern="./path/..."
```

**MANDATORY:** Always use this exact command format. Replace `./path/...` with actual package path.

Examples:

- `./base/...` for base package
- `./signal/...` for signal package
- `./env/...` for env package
- `./model/...` for model package

**NEVER use bare `go test` commands - always use the docker compose wrapper.**

### Build Tag Rules

- ❌ **SKIP** files with `//go:build dev` (development-only)
- ✅ **TEST** files with `//go:build !dev` (production implementations)
- ✅ **TEST** files without build tags (run in all environments)
- When you see dev/prod file pairs, only test the production version

### Testing Scope

- ✅ Test code YOU wrote in the project
- ❌ Don't test third-party library internals (signals.AddListener, GORM methods, etc.)
- ✅ Test integration points where your code interfaces with libraries
- ✅ Test factory functions (Provider()) for initialization behavior
- ✅ Test struct types for zero-value behavior and field initialization

---

## Required Frameworks

```go
import (
"testing"
"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/mock"
"github.com/DATA-DOG/go-sqlmock" // for database tests
)
```

---

## File Structure Rules

- **Test files:** Same package, `{source_file}_test.go` suffix
- **Mock files:** `mocks_test.go` in same package
- **One test file per source file**

---

## Standard Test Structure

### Simple Functions (like TagError, Provider functions)

```go
func TestFunctionName(t *testing.T) {
t.Run("describes what this test does", func (t *testing.T) {
// Arrange
input := "test input"

// Act
result := FunctionName(input)

// Assert
assert.Equal(t, expectedValue, result)
})
}
```

### Database Tests (with setupTest pattern)

```go
func setupTest(t *testing.T) (*Impl, sqlmock.Sqlmock, func ()) {
db, mock, err := sqlmock.New()
require.NoError(t, err)

gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
require.NoError(t, err)

client := &Impl{db: gormDB}
cleanup := func () { db.Close() }

return client, mock, cleanup
}

func TestMethodName(t *testing.T) {
client, mock, cleanup := setupTest(t)
defer cleanup()

t.Run("success case", func (t *testing.T) {
// Setup expectations
mock.ExpectQuery("SELECT").WillReturnRows(...)

// Call method
result, err := client.Method()

// Verify
assert.NoError(t, err)
assert.NoError(t, mock.ExpectationsWereMet())
})
}
```

---

## Required Test Cases

For every function/method, include these test cases when applicable:

1. **✅ Success case** - Happy path
2. **❌ Error cases** - What can go wrong?
3. **🔍 Edge cases** - Empty inputs, nil values, zero values
4. **🗄️ Database errors** (if applicable) - Connection failures, not found, etc.

---

## Database Testing Rules

### SQL Operation Mapping

- `INSERT/UPDATE/DELETE` → Use `mock.ExpectExec()`
- `SELECT` → Use `mock.ExpectQuery()`

### GORM Behaviors

- `Take()`, `Find()` return empty struct + error (NOT nil) when not found
- `Create()`, `Delete()` automatically wrapped in transactions
- GORM adds `created_at`/`updated_at` to INSERT - use `sqlmock.AnyArg()`

### Common Patterns

```go
// CREATE with timestamps
mock.ExpectBegin()
mock.ExpectExec(`INSERT INTO "table_name"`).
WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), userArg, schemaArg).
WillReturnResult(sqlmock.NewResult(1, 1))
mock.ExpectCommit()

// SELECT with WHERE
mock.ExpectQuery(`SELECT \* FROM "table_name" WHERE col = \$1`).
WithArgs(arg).
WillReturnRows(sqlmock.NewRows([]string{"col"}).AddRow(value))
```

---

## HTTP Testing Rules

- ❌ **NEVER send real HTTP requests in tests**
- ✅ Mock all external HTTP calls using `testify/mock` or `httptest`
- ✅ Use `httptest.NewServer()` for integration-style testing
- ✅ Always verify mock expectations with `AssertExpectations(t)`

---

## Mock Implementation Template

```go
type MockClient struct {
mock.Mock
}

func (m *MockClient) Method(ctx context.Context, param Type) (*Model, error) {
args := m.Called(ctx, param)
if args.Get(0) == nil {
return nil, args.Error(1)
}
return args.Get(0).(*Model), args.Error(1)
}
```

---

## Debugging Failed Tests

1. **Run the test** and capture the exact error message
2. **Check operation type** - Are you using `ExpectExec` vs `ExpectQuery` correctly?
3. **Extract SQL pattern** from error log and match exactly
4. **Verify argument count** - Include timestamps for GORM operations
5. **Check transactions** - GORM auto-wraps Create/Delete operations
6. **Update this file** with new patterns you discover

---

## Auto-Update Enabled

When you encounter new patterns, GORM behaviors, or debugging solutions, immediately add them to this file in the
appropriate sections.

**Learning Categories to Update:**

- New GORM operation behaviors
- SQL pattern variations
- Transaction handling edge cases
- Mock setup requirements
- Common pitfalls and resolutions