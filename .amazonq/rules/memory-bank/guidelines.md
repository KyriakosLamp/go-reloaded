# Development Guidelines

## Code Quality Standards

### Formatting and Style
- **Package Declaration**: All files use `package main` (4/4 files)
- **Import Grouping**: Standard library imports grouped together with proper formatting
- **Function Naming**: CamelCase with descriptive names (ReadFile, WriteFile, NewPipeline)
- **Variable Naming**: Clear, descriptive names (inputFile, outputFile, testContent)
- **File Permissions**: Consistent use of 0644 for file operations

### Error Handling Patterns
- **Explicit Error Checking**: Every function that can fail returns an error (4/4 files)
- **Early Return Pattern**: Check errors immediately and return/exit on failure
- **Error Messages**: Descriptive error messages with context (e.g., "Error reading input file: %v")
- **Exit Codes**: Use os.Exit(1) for fatal errors in main function

### Testing Standards
- **Test Function Naming**: TestFunctionName pattern (TestCLIIntegration, TestFileIO)
- **Cleanup Pattern**: Use defer statements for test cleanup (defer os.Remove())
- **Test Isolation**: Each test creates and cleans up its own test files
- **Error Assertions**: Use t.Fatalf() for setup failures, t.Errorf() for assertion failures
- **Test Coverage**: Both unit tests (utils_test.go) and integration tests (main_test.go)

## Architectural Patterns

### Pipeline Architecture
- **Interface-Based Design**: Stage interface defines Process(text string) string
- **Composition Pattern**: Pipeline struct contains slice of Stage interfaces
- **Builder Pattern**: AddStage method for pipeline construction
- **Immutable Processing**: Each stage returns new string, doesn't modify input

### Struct Design
```go
type Pipeline struct {
    stages []Stage
}

type Stage interface {
    Process(text string) string
}
```

### Constructor Pattern
- **Factory Functions**: NewPipeline() returns pointer to initialized struct
- **Zero Value Safety**: Initialize slices to avoid nil pointer issues

## Common Implementation Patterns

### File I/O Operations
```go
// Standard file reading pattern
func ReadFile(filename string) (string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

// Standard file writing pattern
func WriteFile(filename, content string) error {
    return os.WriteFile(filename, []byte(content), 0644)
}
```

### Command Line Argument Handling
```go
if len(os.Args) != 3 {
    fmt.Println("Usage: go run . input.txt output.txt")
    os.Exit(1)
}
```

### Test Setup and Cleanup
```go
// Create test files
testFile := "test_file.txt"
defer os.Remove(testFile)  // Always cleanup

// Test the functionality
err := WriteFile(testFile, testContent)
if err != nil {
    t.Fatalf("Setup failed: %v", err)
}
```

## Development Standards

### Function Signatures
- **Consistent Error Returns**: Functions return (result, error) or just error
- **String Processing**: Text processing functions take and return strings
- **Pointer Receivers**: Use pointer receivers for struct methods that modify state

### Code Organization
- **Single Responsibility**: Each function has one clear purpose
- **Minimal Dependencies**: Only use Go standard library
- **Modular Design**: Separate concerns (main.go for orchestration, utils.go for utilities)
- **Interface Segregation**: Small, focused interfaces (Stage interface has single method)

### Documentation
- **Function Comments**: Brief, clear descriptions of function purpose
- **Package Comments**: Explain overall package functionality
- **Example Usage**: Provide usage examples in comments or README