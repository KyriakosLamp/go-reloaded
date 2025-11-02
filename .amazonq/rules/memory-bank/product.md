# go-reloaded - Text Formatter

## Project Purpose
A command-line text formatter that automatically fixes common writing mistakes and applies text transformations using special marker syntax. The tool processes input files and outputs clean, properly formatted text.

## Value Proposition
- **Automated Text Processing**: Eliminates manual text formatting tasks
- **Pipeline Architecture**: Modular design allows easy extension and maintenance
- **Zero Dependencies**: Pure Go standard library implementation
- **Comprehensive Testing**: 16+ test cases ensure reliability
- **Production Ready**: Handles edge cases and error scenarios gracefully

## Key Features

### Text Transformations
- **Case Transformations**: `(up)`, `(low)`, `(cap)` with multi-word support
- **Numeric Conversions**: Hexadecimal `(hex)` and binary `(bin)` to decimal
- **Article Agreement**: Automatic `a` to `an` conversion before vowels
- **Punctuation Spacing**: Proper spacing around punctuation marks
- **Quote Formatting**: Smart quote spacing with opening/closing detection

### Advanced Capabilities
- **Multi-word Processing**: `word1 word2 (up, 2)` transforms last N words
- **Nested Quotes**: Handles complex quote nesting scenarios
- **Error Recovery**: Invalid inputs remain unchanged with warnings
- **Smart Spacing**: Context-aware spacing around punctuation and quotes

## Target Users

### Primary Users
- **Content Writers**: Automated text formatting and grammar fixes
- **Developers**: Text processing in build pipelines or documentation
- **Students**: Academic writing assistance and formatting

### Use Cases
- **Document Processing**: Batch formatting of text files
- **Content Management**: Automated text cleanup workflows
- **Educational Tools**: Teaching proper grammar and formatting
- **Development Workflows**: Text preprocessing in CI/CD pipelines

## Command Line Interface
```bash
# Basic usage
go run . input.txt output.txt

# Build and run
go build
./go-reloaded input.txt output.txt
```

## Example Transformation
**Input:**
```
it (cap) was a historic race ,lewis hamilton (up,2) took pole position , scoring " 1E (hex) points " !
```

**Output:**
```
It was an historic race, LEWIS HAMILTON took pole position, scoring "30 points"!
```