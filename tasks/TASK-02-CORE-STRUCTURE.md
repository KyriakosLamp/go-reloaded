## Task 02: Core Structure and CLI Integration
- completed: ✅

**Goal**: Create the program’s entry point, pipeline skeleton, and argument parsing.

- **Test Writing**:
  - Test CLI behavior using mocked arguments: `go run . input.txt output.txt`
  - Verify both files exist and output file gets created.
- **Implementation**:
  - Implement `main.go` with argument parsing (`os.Args`).
  - Validate file paths and call the pipeline processor.
  - Create an initial `Pipeline` struct for sequential stage execution.
- **Validation**: Run integration test to confirm CLI + core structure work end-to-end.
