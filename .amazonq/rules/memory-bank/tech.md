# Technology Stack

## Programming Language
- **Go 1.21**: Modern Go version with latest language features
- **Pure Standard Library**: Zero external dependencies for maximum portability
- **Cross-Platform**: Runs on Linux, macOS, Windows

## Build System
- **Go Modules**: Native dependency management with `go.mod`
- **Standard Go Toolchain**: Uses built-in `go build`, `go run`, `go test`
- **No Build Configuration**: Simple compilation with standard Go commands

## Development Commands

### Running the Application
```bash
# Direct execution
go run . input.txt output.txt

# Build and run
go build
./go-reloaded input.txt output.txt
```

### Testing
```bash
# Run all tests with verbose output
go test ./tests -v

# Run specific test files
go test ./tests/functions_test.go -v

# Run with coverage
go test ./tests -cover
```

### Development Workflow
```bash
# Format code
go fmt ./...

# Lint and vet
go vet ./...

# Build for different platforms
GOOS=linux go build
GOOS=windows go build
GOOS=darwin go build
```

## Project Dependencies

### Standard Library Packages Used
- **regexp**: Pattern matching for text transformations
- **strings**: String manipulation and processing
- **strconv**: String to number conversions (hex, binary)
- **os**: File system operations and command-line arguments
- **fmt**: Formatted I/O operations
- **testing**: Unit and integration testing framework

### No External Dependencies
- **Zero Third-Party Libraries**: Ensures maximum compatibility
- **Minimal Attack Surface**: Reduces security vulnerabilities
- **Easy Deployment**: No dependency management complexity
- **Fast Compilation**: No external package downloads

## Architecture Decisions

### Language Choice: Go
- **Performance**: Compiled language with excellent performance
- **Simplicity**: Clean syntax ideal for text processing
- **Concurrency**: Built-in goroutines (though not used in current implementation)
- **Standard Library**: Rich text processing capabilities

### Design Patterns
- **Interface-Based Design**: `Stage` interface for pipeline components
- **Functional Approach**: Immutable string processing
- **Error Handling**: Go's explicit error handling pattern
- **Testing Strategy**: Table-driven tests with comprehensive coverage

## Development Environment

### Recommended Setup
- **Go 1.21+**: Latest stable Go version
- **IDE Support**: VS Code with Go extension, GoLand, or vim-go
- **Git**: Version control with conventional commit messages
- **Make**: Optional build automation (not currently used)

### Code Quality Tools
- **go fmt**: Automatic code formatting
- **go vet**: Static analysis for common errors
- **golint**: Style guide enforcement (optional)
- **gocyclo**: Cyclomatic complexity analysis (optional)

## Deployment

### Binary Distribution
- **Single Executable**: Self-contained binary with no dependencies
- **Cross-Compilation**: Build for multiple platforms from single source
- **Small Footprint**: Minimal binary size due to no external dependencies

### Integration Options
- **CLI Tool**: Direct command-line usage
- **CI/CD Pipeline**: Text processing in automated workflows
- **Docker Container**: Containerized deployment (Dockerfile not included)
- **Library Usage**: Import as Go module in other projects