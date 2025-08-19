# Help

```sh
API sync tool that imports OpenAPI documentation to Postman collections.

Options:
  -doc-api-key string
        The OpenAPI doc API key
  -pm-api-key string
        The Postman API key
  -pm-workspace-id string
        The Postman workspace ID

```

# How to run

```sh
go run . --doc-api-key=xxx --pm-api-key=xxx --pm-workspace-id=xxx
```

## Testing

### Running Tests

To run all tests:
```sh
./test.sh
```

Or run tests manually:
```sh
# Unit tests
go test -v ./...

# With coverage
go test -cover ./...

# Benchmarks
go test -bench=. ./...
```

### Integration Tests

Integration tests require real API credentials and are skipped by default. To run them:

```sh
export RUN_INTEGRATION_TESTS=1
export TEST_DOC_API_KEY=your_doc_api_key
export TEST_PM_API_KEY=your_postman_api_key
export TEST_PM_WORKSPACE_ID=your_workspace_id
go test -v -run TestMainIntegration
```

### Test Coverage

The test suite includes:
- **Unit tests** for command-line argument parsing
- **Integration tests** for end-to-end functionality (requires API keys)
- **Component tests** for dependency injection and initialization
- **Error handling tests** for various failure scenarios
- **Benchmarks** for performance monitoring
