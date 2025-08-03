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