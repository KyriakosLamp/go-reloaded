# Development Guidelines

## Code Quality Standards

### Naming Conventions
- **Struct Types**: PascalCase with descriptive suffixes (e.g., `PunctuationStage`, `NumericConversionStage`)
- **Interface Names**: Simple, action-oriented names (e.g., `Stage` interface)
- **Function Names**: PascalCase for exported, camelCase for private (e.g., `Process()`, `handleQuoteSpacing()`)
- **Variable Names**: Descriptive camelCase (e.g., `singleQuoteOpen`, `multiplePunctPattern`)
- **Constants**: PascalCase for exported, camelCase for package-level

### Documentation Standards
- **Package Comments**: Brief description of package purpose
- **Type Comments**: Explain the role and responsibility of each struct
- **Function Comments**: Start with function name, describe behavior and parameters
- **Inline Comments**: Explain complex regex patterns and business logic

### Code Formatting
- **Standard Go Format**: Use `go fmt` for consistent formatting
- **Import Organization**: Standard library first, then project imports
- **Line Length**: Keep lines readable, break long function calls appropriately
- **Whitespace**: Use blank lines to separate logical sections

## Structural Conventions

### Package Organization
```go
// Standard package structure
package functions

import (
    "regexp"      // Standard library imports first
    "strings"
)

// Type definitions
type StageName struct{}

// Public methods
func (s *StageName) Process(text string) string {
    // Implementation
}

// Private helper methods
func (s *StageName) helperMethod(input string) string {
    // Implementation
}
```

### Interface Implementation
- **Single Method Interfaces**: Prefer simple, focused interfaces
- **Consistent Signatures**: All stages implement `Process(text string) string`
- **Stateless Design**: Stages should not maintain internal state
- **Immutable Processing**: Always return new strings, never modify input

### Error Handling Patterns
```go
// File operations with explicit error handling
content, err := functions.ReadFile(inputFile)
if err != nil {
    fmt.Printf("Error reading input file: %v\n", err)
    os.Exit(1)
}

// Graceful degradation for invalid input
if invalidInput {
    GlobalLogger.LogWarning("Invalid input detected")
    return originalText // Return unchanged
}
```

## Implementation Patterns

### Pipeline Pattern Usage
```go
// Pipeline construction with ordered stages
pipeline := NewPipeline()
pipeline.AddStage(&functions.NumericConversionStage{})
pipeline.AddStage(&functions.PunctuationStage{})
pipeline.AddStage(&functions.ArticleStage{})
pipeline.AddStage(&functions.CaseTransformStage{})
pipeline.AddStage(&functions.QuotationStage{})

// Sequential processing
result := pipeline.Process(content)
```

### Regular Expression Patterns
- **Compile Once**: Use `regexp.MustCompile()` for patterns used multiple times
- **Descriptive Names**: Name patterns clearly (e.g., `colonQuotePattern`, `multiplePunctPattern`)
- **Comment Complex Patterns**: Explain regex logic with inline comments
- **Escape Properly**: Use raw strings for regex patterns when possible

### String Processing Idioms
```go
// Pattern: Remove spaces before punctuation
pattern := regexp.MustCompile(`\s+([.,:;!?]+)`)
text = pattern.ReplaceAllString(text, "$1")

// Pattern: Clean up multiple spaces
multiSpacePattern := regexp.MustCompile(`\s+`)
text = multiSpacePattern.ReplaceAllString(text, " ")

// Pattern: Trim final result
return strings.TrimSpace(text)
```

### Testing Conventions
```go
// Table-driven test structure
func TestStageName(t *testing.T) {
    stage := &functions.StageType{}
    
    tests := []struct {
        input    string
        expected string
    }{
        {"input1", "expected1"},
        {"input2", "expected2"},
    }
    
    for _, test := range tests {
        result := stage.Process(test.input)
        if result != test.expected {
            t.Errorf("Input: %q, Expected: %q, Got: %q", 
                test.input, test.expected, result)
        }
    }
}
```

## Architectural Guidelines

### Stage Design Principles
- **Single Responsibility**: Each stage handles one type of transformation
- **Order Independence**: Stages should not depend on internal state of other stages
- **Idempotent Operations**: Running a stage multiple times should be safe
- **Error Resilience**: Invalid input should not crash the stage

### Pipeline Ordering Rules
1. **Numeric Conversions First**: Convert markers before text transformations
2. **Punctuation Early**: Normalize spacing before case transformations
3. **Articles Before Case**: Handle agreement before capitalization
4. **Case Transformations**: Apply after content is normalized
5. **Quotes Last**: Final formatting after all content changes

### Memory Management
- **Avoid String Concatenation**: Use `strings.Builder` for multiple operations
- **Rune Slices for Complex Operations**: Use `[]rune` for character-level manipulation
- **Return New Strings**: Never modify input parameters directly

### Logging Strategy
```go
// Global logger for consistent behavior
var GlobalLogger = NewLogger(false) // Disabled by default for clean CLI

// Warning for recoverable issues
GlobalLogger.LogWarning("Invalid hex number: " + input)

// Error for critical failures
GlobalLogger.LogError("File not found", true) // Exits program
```

## Performance Considerations

### Regex Optimization
- **Compile Once**: Store compiled patterns as package variables when possible
- **Specific Patterns**: Use precise patterns rather than overly broad ones
- **Avoid Backtracking**: Design patterns to minimize regex engine backtracking

### String Processing Efficiency
- **Minimize Allocations**: Reuse string builders and buffers when possible
- **Process in Order**: Design pipeline to minimize string copying
- **Batch Operations**: Group related transformations in single stage

### Testing Performance
- **Benchmark Critical Paths**: Use Go's benchmarking for performance-critical stages
- **Profile Memory Usage**: Monitor allocations in text processing functions
- **Test with Large Inputs**: Validate performance with realistic file sizes