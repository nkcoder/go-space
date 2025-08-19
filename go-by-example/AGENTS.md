# Agent Guidelines for go-by-example

## Build/Test Commands
- `task run` or `go run .` - Run the main application
- `task format` or `gofmt -s -w .` - Format all Go code
- `go test ./...` - Run all tests (if any exist)
- `go test -run TestName` - Run specific test
- `go build` - Build the application

## Code Style Guidelines
- **Package**: All files in `topics/` use `package topics`
- **Imports**: Standard library imports only, group imports logically
- **Naming**: Use camelCase for functions, lowercase for unexported, PascalCase for exported
- **Functions**: Exported functions start with capital letter (e.g., `ErrorMain()`, `Function()`)
- **Error Handling**: Return errors as last value, check with `if err != nil`
- **Comments**: Use `//` for single line, document exported functions
- **Structs**: Use struct literals with field names for clarity
- **Methods**: Use pointer receivers for mutating methods, value receivers for read-only
- **Variables**: Use short variable names in small scopes, descriptive names for larger scopes
- **Error Types**: Create custom error types implementing `Error() string` method
- **Sentinel Errors**: Use `var ErrName = fmt.Errorf("message")` pattern

## Project Structure
- `main.go` - Entry point, calls topic functions
- `topics/` - Individual Go concept examples as separate files
- Each topic file contains educational examples with detailed comments