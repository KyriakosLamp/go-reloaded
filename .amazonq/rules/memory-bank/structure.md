# Project Structure - go-reloaded

## Directory Organization

### Root Level
- `main.go` - Entry point with pipeline orchestration and CLI handling
- `go.mod` - Go module definition (Go 1.21)
- `README.md` - Project documentation with examples and test output
- `AGENTS.md` - AI agent instructions and development guidelines

### Core Components

#### `/functions/` - Transformation Stages
Contains all pipeline stage implementations (tests moved to /tests/):
- `numeric_conversion.go` - Binary/hex to decimal conversion with error logging
- `article_agreement.go` - "a" to "an" transformation
- `case_transform.go` - Upper/lower/capitalize operations with count support
- `quotation.go` - Quotation mark spacing correction
- `punctuation.go` - Punctuation spacing normalization
- `logger.go` - Configurable warning and error logging system
- `utils.go` - File I/O utilities (ReadFile, WriteFile)

#### `/tests/` - All Test Files
Consolidated test suite with comprehensive coverage:
- `functions_test.go` - Individual stage unit tests
- `integration_test.go` - Phase one integration tests
- `phase_two_test.go` - Phase two integration tests
- `full_integration_test.go` - Complete pipeline with Formula One test
- `error_handling_test.go` - Error scenarios and logging tests
- `main_test.go` - CLI integration tests
- `utils_test.go` - File I/O tests
- `pipeline.go` - Test pipeline definitions

#### `/docs/` - GitHub Pages Site
- `index.html` - Main site with dark theme and interactive features
- `style.css` - Responsive CSS with section navigation
- `script.js` - Carousel, smooth scrolling, section detection
- `icons/` - VS Code file type icons for project tree

#### `/Agent_reports/` - Implementation Reports
- `Task-01-report.md` through `Task-12-report.md` - Detailed implementation summaries

#### `/docs/` (Additional Documentation)
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

#### Temporary Files (Cleaned Up)
- All temporary test files removed for clean repository
- Tests create/cleanup their own temporary files during execution

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