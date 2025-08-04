AI_TEST_GENERATION_INSTRUCTIONS:

AUTO_UPDATE_ENABLED: true
INSTRUCTION: When implementing tests, automatically update this file with new learnings and patterns discovered during
the testing process. Add new GORM behaviors, SQL patterns, error cases, and debugging techniques as they are
encountered.

FRAMEWORKS_REQUIRED:

- testify/mock (import alias: tmock "github.com/stretchr/testify/mock")
- sqlmock for database mocking
- standard testing package

FILE_STRUCTURE:

- Tests: same package, {source_file}_test.go suffix
- Mocks: mocks_test.go in same package
- One test file per source file

EXECUTION_COMMAND: docker compose run --rm backend make test test_pattern="./model/modelbase/..."

MANDATORY_TEST_PATTERNS:

1. setupTest() function returning (client, mock, cleanup)
2. t.Run() for each test case
3. defer cleanup() in each test
4. assert.NoError(t, mock.ExpectationsWereMet()) in each test

GORM_SQL_PATTERNS:
Take(): SELECT \* FROM "table" WHERE "table"."id" = \$1 ORDER BY "table"."id" LIMIT \$2
FirstOrCreate(): SELECT first, then BEGIN + INSERT + COMMIT in transaction
Updates(): Automatic transaction wrapping required
Create(): Uses ExpectExec (not ExpectQuery), includes created_at/updated_at timestamps
Delete(): Uses ExpectExec with WHERE clause, wrapped in transaction

CRITICAL_SQL_MOCK_RULES:

- Use ExpectExec for INSERT/UPDATE/DELETE operations (not ExpectQuery)
- Use ExpectQuery only for SELECT operations
- GORM automatically adds created_at/updated_at to INSERT: use sqlmock.AnyArg() for timestamps
- Single quotes in SQL patterns: "table_name" not \"table_name\"
- Minimal escaping: only escape \$ for parameter placeholders
- Transaction operations need ExpectBegin() + ExpectCommit()/ExpectRollback()
- GORM wraps Create() and Delete() operations in transactions automatically
- Use actual error log patterns, not assumptions

CRITICAL_HTTP_MOCK_RULES:

- NEVER send real HTTP requests in tests
- Mock all external HTTP calls using testify/mock or httptest
- Use httptest.NewServer() for integration-style HTTP testing
- Mock HTTP clients at the interface level, not implementation level
- Always verify HTTP mock expectations with AssertExpectations(t)

GORM_ERROR_BEHAVIOR:

- Take(), Find() return empty struct + error (NOT nil)
- Test pattern: assert.Error(t, err) + assert.NotNil(t, result)
- Check actual GORM method behavior in docs

MOCK_IMPLEMENTATION_TEMPLATE:

```
func (m *MockClient) Method(ctx context.Context, param Type) (*Model, error) {
    args := m.Called(ctx, param)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*Model), args.Error(1)
}
```

REQUIRED_TEST_CASES:

- success case
- not_found/empty_result case
- database_error case
- invalid_input case (if applicable)

DEPENDENCIES_INITIALIZATION:

- Use Provider() functions for signal containers
- Mock all external dependencies
- Verify interface compliance with: var _ Interface = (*Implementation)(nil)

COMMON_PATTERNS:

CREATE with timestamps:
mock.ExpectBegin()
mock.ExpectExec(`INSERT INTO "table_name"`).
WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), userArg, schemaArg).
WillReturnResult(sqlmock.NewResult(1, 1))
mock.ExpectCommit()

DELETE with WHERE:
mock.ExpectBegin()
mock.ExpectExec(`DELETE FROM "table_name" WHERE col1 = \$1 AND col2 = \$2`).
WithArgs(arg1, arg2).
WillReturnResult(sqlmock.NewResult(0, 1))
mock.ExpectCommit()

SELECT with context:
mock.ExpectQuery(`SELECT ... FROM "table_name" WHERE col = \$1`).
WithArgs(arg).
WillReturnRows(sqlmock.NewRows([]string{"col"}).AddRow(value))

DEBUG_PROCESS:

1. Run test, capture exact error
2. Check if using correct Expect method (Exec vs Query)
3. Extract actual SQL pattern from error log
4. Update mock expectation to match exact pattern
5. Verify argument count matches GORM behavior (include timestamps)
6. Check transaction expectations if applicable
7. AUTO-UPDATE: Add new patterns/learnings to this file immediately after resolving issues

LEARNING_CATEGORIES_TO_UPDATE:

- New GORM operation behaviors (FirstOrCreate, Updates, etc.)
- SQL pattern variations (different table structures, complex WHERE clauses)
- Transaction handling edge cases
- Argument count mismatches and solutions
- Error handling patterns specific to GORM methods
- Mock setup requirements for different dependency types
- Common pitfalls and their resolutions