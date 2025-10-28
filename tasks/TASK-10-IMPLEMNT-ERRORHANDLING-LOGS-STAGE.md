## Task 10: Implement Error Handling & Logging
- completed:

**Goal**: Ensure robust error handling throughout all stages.

- **Test Writing**:
  - Invalid `(hex)` input logs warning, e.g., `"XYZ (hex)"`.
  - Missing numbers or markers skipped safely.
- **Implementation**:
  - Add `Logger` utility for controlled warnings.
  - Wrap parsing in try-safe blocks.
- **Validation**: Confirm system logs issues without stopping execution.
