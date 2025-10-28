# Project Structure

## Directory Organization

### Root Level
- `main.go` - Entry point and pipeline orchestration
- `utils.go` - Core transformation functions and utilities
- `main_test.go` - Integration tests for main pipeline
- `utils_test.go` - Unit tests for utility functions
- `go.mod` - Go module definition (Go 1.21)
- `README.md` - Project documentation and transformation rules

### Supporting Directories
- `functions/` - Secondary functions used by main.go (modular organization)
- `docs/` - Comprehensive project documentation
  - `AGENTS.md` - AI agent instructions and workflow
  - `Code_Analysis.md` - Technical analysis documentation
  - `go-reloaded_StagesPlan.md` - Step-by-step project plan
  - `Test_Cases.md` - Test case specifications
  - `Agent_Log.md` - Development log and progress tracking
- `tasks/` - Individual task breakdowns (TASK-01 through TASK-13)
- `.amazonq/rules/memory-bank/` - AI context and memory bank files

### Sample Files
- `sample_input.txt` - Example input with transformation markers
- `sample_output.txt` - Expected output after transformations
- `go-reloaded` - Compiled binary executable

## Core Components

### Pipeline Architecture
The project follows a pipeline pattern where text flows through sequential transformation stages:
1. Input processing and tokenization
2. Numeric conversion stage (hex/bin to decimal)
3. Article agreement correction (a → an)
4. Case transformation stage (up/low/cap)
5. Quotation mark normalization
6. Punctuation spacing correction
7. Output generation

### Component Relationships
- `main.go` orchestrates the pipeline and handles file I/O
- `utils.go` contains individual transformation functions
- `functions/` directory houses modular helper functions
- Test files validate each component and integration points
- Documentation provides comprehensive guidance for development and usage