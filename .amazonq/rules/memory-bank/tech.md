# Technology Stack

## Programming Language
- **Go 1.21** - Primary development language
- Modern Go features and standard library usage
- No external dependencies beyond Go standard library

## Build System
- **Go Modules** - Dependency management with go.mod
- Native Go build tools and commands
- Standard Go project structure

## Development Commands

### Building
```bash
go build -o go-reloaded main.go utils.go
```

### Testing
```bash
go test -v                    # Run all tests with verbose output
go test ./...                 # Run tests in all subdirectories
go test -cover               # Run tests with coverage report
```

### Running
```bash
./go-reloaded input.txt output.txt    # Run compiled binary
go run main.go utils.go input.txt output.txt    # Direct execution
```

### Development Tools
```bash
go fmt ./...                 # Format all Go files
go vet ./...                 # Static analysis
go mod tidy                  # Clean up dependencies
```

## Project Configuration
- **Module Name**: go-reloaded
- **Go Version**: 1.21 (minimum required)
- **Architecture**: Single module with multiple source files
- **Testing Framework**: Go's built-in testing package
- **No External Dependencies**: Pure Go standard library implementation

## File Processing
- Text file input/output operations
- String manipulation and transformation
- Regular expression support (if needed)
- Pipeline-based processing architecture