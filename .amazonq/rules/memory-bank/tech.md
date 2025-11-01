# Technology Stack - go-reloaded

## Programming Language
- **Go 1.21** - Primary development language
- **Standard Library Only** - No external dependencies

## Module Configuration
```
module go-reloaded
go 1.21
```

## Build System
- **Go Modules** - Dependency management (though no external deps used)
- **Native Go Build** - Standard `go build` and `go run` commands

## Development Commands

### Running the Application
```bash
go run . input.txt output.txt
```

### Testing
```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./functions

# Run with verbose output
go test -v ./...

# Run specific test files
go test main_test.go
go test utils_test.go
```

### Building
```bash
# Build executable
go build -o go-reloaded

# Run built executable
./go-reloaded input.txt output.txt
```

## Project Dependencies
- **Zero External Dependencies** - Uses only Go standard library
- **File I/O**: `os` package for file operations
- **String Processing**: `strings`, `strconv` packages
- **Regular Expressions**: `regexp` package for pattern matching
- **Testing**: `testing` package for unit tests

## Development Environment
- **Go 1.21+** required
- **Linux/Unix** development environment
- **Standard Go toolchain** (go, gofmt, go test)

## Code Organization
- **Package Structure**: Main package with `functions` subpackage
- **Interface-Based Design**: Stage interface for pipeline components
- **Test-Driven Development**: Comprehensive test coverage for all stages
- **Modular Architecture**: Each transformation as separate, testable component