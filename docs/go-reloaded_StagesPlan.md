# GoReloaded – 13-Stage TDD Plan (Sequential Pipeline)

## Task 01: Project Setup & Test Framework
**Goal**: Initialize the Go module, directories, and test framework.

- **Test Writing**: `TestFileIO` — ensure input/output file reading and writing works.
- **Implementation**: Add `utils.go` for file handling (`ReadFile`, `WriteFile`).
- **Validation**: Run `go test` and verify file read/write functionality.

## Task 02: Core Structure and CLI Integration
**Goal**: Create the program’s entry point, pipeline skeleton, and argument parsing.

- **Test Writing**:
  - Test CLI behavior using mocked arguments: `go run . input.txt output.txt`
  - Verify both files exist and output file gets created.
- **Implementation**:
  - Implement `main.go` with argument parsing (`os.Args`).
  - Validate file paths and call the pipeline processor.
  - Create an initial `Pipeline` struct for sequential stage execution.
- **Validation**: Run integration test to confirm CLI + core structure work end-to-end.

## Task 03: Implement NumericConversionStage (Hex + Bin)
**Goal**: Convert preceding hexadecimal or binary values to decimal equivalents.

- **Test Writing**:
  - `"1E (hex)"` → `"30"`
  - `"10 (bin)"` → `"2"`
  - `"1A (hex) and 101 (bin)"` → `"26 and 5"`.
- **Implementation**:
  - Create `NumericConversionStage` to detect `(hex)` or `(bin)` markers.
  - Convert safely and handle invalid numbers gracefully.
- **Validation**: Confirm outputs are correct, and invalid cases log warnings without crashes.

## Task 04: Implement Article Agreement Rule
**Goal**: Replace “a” with “an” when followed by a vowel (`a, e, i, o, u`) or `h`.

- **Test Writing**:
  - `"a orange car"` → `"an orange car"`
  - `"a honest man"` → `"an honest man"`
  - `"a user"` → `"an user"`.
- **Implementation**:
  - Create `ArticleStage` to check the next word’s first character.
- **Validation**: Confirm grammar corrections appear correctly in all test cases.

## Task 05: Implement CaseTransformStage (Unified)
**Goal**: Handle `(up)`, `(low)`, `(cap)` and their numeric forms `(up, n)`, `(low, n)`, `(cap, n)`.

- **Test Writing**:
  - `"go (up)"` → `"GO"`
  - `"HELLO (low)"` → `"hello"`
  - `"bridge (cap)"` → `"Bridge"`
  - `"this is fun (up, 2)"` → `"this IS FUN"`.
- **Implementation**:
  - Parse `(rule, n)` pattern; default `n=1`.
  - Apply transformation to last `n` words accordingly.
- **Validation**: Run all case transformation tests and verify behavior across single and multiple words.

## Task 06: Implement Quotation Formatting
**Goal**: Clean spacing and ensure quotes hug inner text properly.

- **Test Writing**:
  - `"He said: ' wow '."` → `"He said: 'wow'."`
  - `"\" hello \""` → `"\"hello\""`.
- **Implementation**:
  - Create `QuotationStage` to trim spaces inside `' '` and `" "`.
  - Handle nested quotes gracefully.
- **Validation**: Confirm formatting correctness across all quotation variations.

## Task 07: Implement Punctuation Handling
**Goal**: Fix punctuation spacing and grouping.

- **Test Writing**:
  - `"Hello ,world !!"` → `"Hello, world!!"`
  - `"Wait ... what ?"` → `"Wait... what?"`.
- **Implementation**:
  - Implement `PunctuationStage` to:
    - Attach punctuation to previous word.
    - Ensure one space follows it.
    - Preserve grouped marks like `!?` and `...`.
- **Validation**: Verify results in single and multiple words/sentences.

## Task 08: Integration Test – Numeric + Article + Case - Phase One
**Goal**: Test early rule interactions.

- **Test Writing**:
  - `"1A (hex) files and 10 (bin) more (up, 2)"` → `"26 FILES AND 2 MORE"`.
- **Implementation**:
  - Run pipeline: `NumericConversion` → `Article` → `Case`.
- **Validation**: Confirm sequential transformation correctness.

## Task 09: Integration Test – Quotes + Punctuation - Phase Two
**Goal**: Validate structural formatting rules.

- **Test Writing**:
  - `"He said ' hello world ' ,and smiled ."` → `"He said 'hello world', and smiled."`.
- **Implementation**:
  - Chain `Quotation` → `Punctuation` stages.
- **Validation**: Confirm quote trimming and punctuation spacing correctness.

## Task 10: Implement Error Handling & Logging
**Goal**: Ensure robust error handling throughout all stages.

- **Test Writing**:
  - Invalid `(hex)` input logs warning, e.g., `"XYZ (hex)"`.
  - Missing numbers or markers skipped safely.
- **Implementation**:
  - Add `Logger` utility for controlled warnings.
  - Wrap parsing in try-safe blocks.
- **Validation**: Confirm system logs issues without stopping execution.

## Task 11: Implement Formula One Full Integration Test - Phase Final
**Goal**: Validate all rules together on a complex paragraph.

- **Test Writing**: Use the Formula One text case from the analysis doc.
- **Implementation**: Run full pipeline with optimized order:
  - `NumericConversion` → `Punctuation` → `Article` → `Case` → `Quotation`.
- **Validation**: Confirm the output matches expected final formatted result exactly.

## Task 12: Refactor and Polish
**Goal**: Ensure readability, consistent naming, and reusable code.

- **Test Writing**: Re-run all existing tests (unit + integration).
- **Implementation**: Refactor duplicated logic, clean imports, add comments.
- **Validation**: Confirm zero behavioral changes (all tests still green).

## Task 13: Documentation and AI Test Integration
**Goal**: Finalize project documentation and AI-readable tests.

- **Test Writing**: Export test cases in JSON/YAML for future automated testing.
- **Implementation**:
  - Write docstrings for each stage.
  - Update `README.md` to describe CLI usage, pipeline order, and examples.
- **Validation**: Ensure docs are complete and test exports validate successfully.

## Final Sequential Pipeline Flow (OPTIMIZED)
`NumericConversion` → `Punctuation` → `Article` → `Case` → `Quotation`

##### Why this optimized order works:
1. **Numeric conversions first**: Convert markers before any text transformations.
2. **Punctuation normalization**: Fix spacing issues (like comma spacing) before case transformations.
3. **Article agreement**: Handle a→an conversion with robust keyword skipping.
4. **Case transformations**: Apply after punctuation is normalized to avoid interference.
5. **Quote formatting last**: Final cleanup of internal quote spacing after all content changes.
6. Simple, sequential logic — no concurrency, no synchronization, just clean transformations.