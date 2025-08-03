When writing tests for the code, please follow these guidelines:

1. Ensure that each test is independent and can run in isolation.
2. Ensure all tests works correctly and pass without errors.
3. Use testify mock for mocking dependencies.
4. Use sqlmock for mocking database interactions.
5. Focus on writing unit tests that cover the functionality of the code instead of integration tests.
6. Place tests in the same package as the code being tested with a `_test.go` suffix. Each file should contain tests for
   functions in the corresponding file.
7. To run tests use `docker compose run --rm backend make test` which accepts `test_pattern` as an argument to run
   specific tests by file pattern.
8. Place mock implementations in a separate file named `mocks_test.go` in the same package.
9. You can reference `backend/coverage` to check coverage of previously generated tests.

## GORM Testing Best Practices:

10. When testing GORM operations, be aware that GORM may generate different SQL than expected:
    - `Take()` operations use `LIMIT 1` parameter, so expect 2 arguments: `WithArgs(id, 1)`
    - `FirstOrCreate()` operations first do a SELECT, then potentially INSERT in a transaction
    - Some operations like `Updates()` are automatically wrapped in transactions by GORM
    - Use actual SQL patterns from error logs rather than assumptions

11. SQL Mock Expectations:
    - Always escape regex special characters in SQL patterns (e.g., use `\*` for `*`, `\$` for `$`)
    - GORM includes ORDER BY clauses in many queries even when not explicitly specified
    - INSERT queries may have fewer arguments than expected - check actual error logs for exact patterns
    - For transactions, expect `ExpectBegin()`, the query, then `ExpectCommit()` or `ExpectRollback()`

12. Error Handling in Tests:
    - GORM's `Take()` and similar methods return empty structs even on errors, unlike some other methods that return nil
    - When testing error conditions, check the actual behavior of the specific GORM method being tested
    - Some errors may cause GORM to return non-nil results with error values

13. Mock Setup:
    - Use `tmock.Anything` (with alias) for dynamic values like timestamps and auto-generated IDs
    - For testify mocks, use proper type assertions and nil checks in mock implementations
    - Ensure setupTest functions properly initialize all dependencies (e.g., signal containers)