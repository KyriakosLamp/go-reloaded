# Development Guidelines - go-reloaded

## Code Quality Standards

### Naming Conventions
- **Struct Types**: PascalCase with descriptive suffixes (e.g., `CaseTransformStage`, `NumericConversionStage`)
- **Interface Names**: Simple, action-oriented names (e.g., `Stage`)
- **Function Names**: PascalCase for exported, camelCase for internal (e.g., `NewPipeline`, `Process`)
- **Variable Names**: camelCase with clear, descriptive names (e.g., `inputFile`, `outputFile`, `wordsPart`)
- **Package Names**: Lowercase, single word when possible (e.g., `functions`)

### Documentation Standards
- **Struct Comments**: Brief description of purpose (e.g., `// CaseTransformStage handles case transformations`)
- **Function Comments**: Action-oriented descriptions (e.g., `// Process applies case transformations`)
- **Interface Comments**: Clear contract definition (e.g., `// Stage interface for pipeline stages`)
- **No inline comments**: Code should be self-documenting through clear naming

### Error Handling Patterns
- **Explicit Error Checking**: Always check and handle errors immediately after function calls
- **Error Propagation**: Return errors to caller with context using `fmt.Printf` for user-facing messages
- **Graceful Degradation**: Return original input when transformation fails (e.g., numeric conversion failures)
- **Exit on Critical Errors**: Use `os.Exit(1)` for unrecoverable CLI errors

## Structural Conventions

### Package Organization
- **Main Package**: Contains only pipeline orchestration and CLI logic
- **Functions Package**: All transformation logic isolated in separate package
- **Single Responsibility**: Each file handles one specific transformation type
- **Test Co-location**: Test files alongside implementation files with `_test.go` suffix

### Interface Design
- **Minimal Interfaces**: Single method interfaces (e.g., `Stage` with only `Process` method)
- **String Processing Contract**: All stages follow `Process(text string) string` signature
- **Stateless Design**: Stages contain no mutable state, only behavior

### Pipeline Architecture
- **Sequential Processing**: Stages process text in defined order without parallelization
- **Immutable Text Flow**: Each stage returns new string, original unchanged
- **Stage Independence**: No inter-stage communication or shared state
- **Flexible Composition**: Pipeline can be built with any combination of stages

## Implementation Patterns

### Regular Expression Usage
- **Compile Once**: Use `regexp.MustCompile()` for pattern definitions
- **Descriptive Patterns**: Complex regex with clear variable names (e.g., `hexPattern`, `binPattern`)
- **Capture Groups**: Use parentheses for extracting specific parts of matches
- **Iterative Processing**: Use loops with `FindStringSubmatch` for multiple replacements

### String Processing Idioms
- **Field Splitting**: Use `strings.Fields()` for whitespace-based word separation
- **Case Transformations**: Standard library functions (`strings.ToUpper`, `strings.ToLower`)
- **String Building**: Use `strings.Join()` for reconstructing processed text
- **Replacement Strategy**: Use `strings.Replace()` with count parameter for controlled substitution

### Testing Patterns
- **Table-Driven Tests**: Use struct slices with input/expected pairs
- **Descriptive Test Names**: Function name + behavior being tested (e.g., `TestArticleAgreement`)
- **Stage Isolation**: Test individual stages independently
- **Integration Testing**: Separate tests for full pipeline behavior
- **Cleanup Pattern**: Use `defer` for test file cleanup

### Error Recovery Strategies
- **Conversion Failures**: Return original text when numeric conversion fails
- **Pattern Matching**: Graceful handling when regex doesn't match expected format
- **File Operations**: Explicit error messages with context for user debugging
- **Validation**: Input validation at entry points (CLI argument checking)

## API Usage Patterns

### Standard Library Utilization
- **strconv Package**: Use `ParseInt` with explicit base parameters for number conversions
- **regexp Package**: Compile patterns once, reuse for multiple operations
- **strings Package**: Prefer standard library functions over manual string manipulation
- **os Package**: Direct file argument access via `os.Args` for CLI applications

### Memory Management
- **String Immutability**: Leverage Go's string immutability for safe concurrent access
- **Slice Operations**: Use `append()` for dynamic stage collection building
- **No Premature Optimization**: Focus on clarity over performance in text processing

### Concurrency Considerations
- **Stateless Design**: All stages are inherently thread-safe due to no shared state
- **Sequential Processing**: Current implementation prioritizes simplicity over parallel execution
- **Future Extensibility**: Architecture supports future parallel stage execution