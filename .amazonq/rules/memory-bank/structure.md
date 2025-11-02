# Project Structure

## Directory Organization

```
go-reloaded/
├── main.go                    # CLI entry point and pipeline orchestration
├── go.mod                     # Go module definition (Go 1.21)
├── README.md                  # User documentation and examples
├── functions/                 # Core transformation stages
│   ├── numeric_conversion.go  # Hex/binary to decimal conversion
│   ├── article_agreement.go   # Article agreement (a -> an)
│   ├── case_transform.go      # Case transformations (up/low/cap)
│   ├── quotation.go          # Quote spacing and formatting
│   ├── punctuation.go        # Punctuation spacing rules
│   ├── logger.go             # Error handling and logging
│   └── utils.go              # File I/O utilities
├── tests/                    # Comprehensive test suite
│   ├── functions_test.go     # Individual stage unit tests
│   ├── integration_test.go   # Multi-stage integration tests
│   ├── main_test.go          # CLI integration tests
│   └── pipeline.go           # Test pipeline definitions
├── docs/                     # GitHub Pages documentation site
│   ├── index.html            # Main documentation page
│   ├── style.css             # Dark theme styling
│   ├── script.js             # Interactive features
│   └── icons/                # VS Code file type icons
└── agent_reports/            # Development task reports
```

## Core Components

### Pipeline Architecture
- **Stage Interface**: Common interface for all transformation stages
- **Pipeline Orchestrator**: Sequential stage execution with error handling
- **Immutable Processing**: Each stage returns new string, no side effects

### Transformation Stages
1. **NumericConversionStage**: Converts `(hex)` and `(bin)` markers to decimal
2. **PunctuationStage**: Normalizes punctuation spacing before transformations
3. **ArticleStage**: Handles `a` to `an` conversion with case-aware patterns
4. **CaseTransformStage**: Applies `(up)`, `(low)`, `(cap)` transformations
5. **QuotationStage**: Final quote formatting and spacing

### Support Systems
- **Logger**: Configurable warning/error reporting system
- **Utils**: File I/O operations with error handling
- **Test Framework**: Comprehensive testing with multiple integration levels

## Architectural Patterns

### Pipeline Pattern
- Sequential processing through independent stages
- Each stage implements common `Process(text string) string` interface
- Stages are stateless and thread-safe
- Easy to add, remove, or reorder stages

### Strategy Pattern
- Different transformation strategies for each text processing rule
- Encapsulated algorithms in separate stage implementations
- Runtime composition of processing pipeline

### Error Handling Strategy
- Graceful degradation: invalid inputs remain unchanged
- Comprehensive logging with configurable levels
- No fatal errors during text processing

## Component Relationships

### Data Flow
```
Input File → NumericConversion → Punctuation → Article → Case → Quotation → Output File
```

### Dependencies
- **main.go** orchestrates all stages and handles CLI
- **functions/** contains independent, reusable stages
- **tests/** validates individual stages and integration scenarios
- **docs/** provides user-facing documentation and examples

### Stage Interdependencies
- **Order Critical**: Stages must run in specific sequence
- **Punctuation First**: Normalizes spacing before case transformations
- **Article Awareness**: Handles both pre and post-case transformation scenarios
- **Quote Last**: Final formatting after all content transformations