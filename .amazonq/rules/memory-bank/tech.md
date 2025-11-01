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
# Run all tests (16 tests, all passing)
go test ./tests -v

# Run specific test categories
go test ./tests -run TestErrorHandling
go test ./tests -run TestFullPipelineIntegration
go test ./tests -run TestComplexCombinations

# Test output shows comprehensive coverage:
# - Error handling with warnings
# - Individual stage functionality
# - Multi-stage integration
# - CLI end-to-end testing
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
- **String Processing**: `strings`, `strconv` packages for transformations
- **Regular Expressions**: `regexp` package for pattern matching
- **Error Handling**: `fmt`, `log` packages for logging system
- **Testing**: `testing` package with table-driven tests

## GitHub Pages Deployment
- **Location**: `/docs/` folder ready for GitHub Pages
- **Features**: Dark theme, responsive design, VS Code icons
- **Technology**: Pure HTML/CSS/JavaScript, no build process required
- **Sections**: Interactive carousel, smooth scrolling, file tree display

## Development Environment
- **Go 1.21+** required
- **Linux/Unix** development environment
- **Standard Go toolchain** (go, gofmt, go test)

## Code Organization
- **Package Structure**: Main package with `functions` subpackage, `tests` package
- **Interface-Based Design**: Stage interface for pipeline components
- **Test-Driven Development**: 16 comprehensive tests covering all functionality
- **Modular Architecture**: Each transformation as separate, testable component
- **Error Recovery**: Graceful degradation with configurable logging
- **Clean Repository**: No temporary files, proper .gitignore practices