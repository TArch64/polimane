AI_TEST_GENERATION_INSTRUCTIONS:

FRAMEWORKS_REQUIRED:

- testify/mock (import alias: tmock "github.com/stretchr/testify/mock")
- sqlmock for database mocking
- standard testing package

FILE_STRUCTURE:

- Tests: same package, {source_file}_test.go suffix
- Mocks: mocks_test.go in same package
- One test file per source file

EXECUTION_COMMAND: docker compose run --rm backend make test test_pattern="./path/..."

MANDATORY_TEST_PATTERNS:

1. setupTest() function returning (client, mock, cleanup)
2. t.Run() for each test case
3. defer cleanup() in each test
4. assert.NoError(t, mock.ExpectationsWereMet()) in each test

GORM_SQL_PATTERNS:
Take(): SELECT \* FROM "table" WHERE "table"."id" = \$1 ORDER BY "table"."id" LIMIT \$2
FirstOrCreate(): SELECT first, then BEGIN + INSERT + COMMIT in transaction
Updates(): Automatic transaction wrapping required
Create(): INSERT with transaction

CRITICAL_SQL_MOCK_RULES:

- Escape ALL regex chars: \* for *, \$ for $, \" for "
- GORM adds ORDER BY clauses automatically
- GORM adds LIMIT 1 for Take() operations
- Transaction operations need ExpectBegin() + ExpectCommit()/ExpectRollback()
- Use actual error log patterns, not assumptions

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

DEBUG_PROCESS:

1. Run test, capture exact error
2. Extract actual SQL pattern from error
3. Update mock expectation to match exact pattern
4. Verify argument count matches GORM behavior
5. Check transaction expectations if applicable