# Project Structure - go-reloaded

## Directory Organization

### Root Level
- `main.go` - Entry point with pipeline orchestration and CLI handling
- `go.mod` - Go module definition (Go 1.21)
- `README.md` - Project documentation and transformation rules
- `*_test.go` - Root-level test files for main functionality

### Core Components

#### `/functions/` - Transformation Stages
Contains all pipeline stage implementations with corresponding tests:
- `numeric_conversion.go/.go_test` - Binary/hex to decimal conversion
- `article_agreement.go/.go_test` - "a" to "an" transformation
- `case_transform.go/.go_test` - Upper/lower/capitalize operations
- `quotation.go/.go_test` - Quotation mark spacing correction
- `punctuation.go/.go_test` - Punctuation spacing normalization
- `utils.go` - Shared utility functions for file I/O

#### `/docs/` - Documentation
- `Agent_Log.md` - Development process documentation
- `AGENTS.md` - AI agent interaction logs
- `Code_Analysis.md` - Technical analysis documentation
- `go-reloaded_StagesPlan.md` - Pipeline stage planning
- `Test_Cases.md` - Test case documentation

#### `/tasks/` - Development Workflow
Sequential task breakdown files (TASK-01 through TASK-13):
- Project setup and core structure
- Individual stage implementations
- Testing phases
- Error handling and logging
- Documentation and polish

#### Test Files
- `sample_input.txt/sample_output.txt` - Main test case
- `test_*.txt` - Individual feature test files with expected outputs

## Architectural Patterns

### Pipeline Architecture
- **Pipeline struct**: Orchestrates sequential stage processing
- **Stage interface**: Defines `Process(text string) string` contract
- **Modular stages**: Each transformation rule as independent stage
- **Sequential processing**: Text flows through stages in defined order

### Stage Processing Order
1. NumericConversionStage - Handle (hex)/(bin) markers
2. ArticleStage - Process "a" to "an" conversions
3. CaseTransformStage - Apply case transformations
4. QuotationStage - Fix quotation mark spacing
5. PunctuationStage - Normalize punctuation spacing

### Component Relationships
- `main.go` creates pipeline and adds stages in specific order
- Each stage in `/functions/` implements Stage interface
- Utility functions in `utils.go` handle file operations
- Test files validate individual stage and integration behavior