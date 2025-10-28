## Task 03: Implement NumericConversionStage (Hex + Bin)
- completed:

**Goal**: Convert preceding hexadecimal or binary values to decimal equivalents.

- **Test Writing**:
  - `"1E (hex)"` → `"30"`
  - `"10 (bin)"` → `"2"`
  - `"1A (hex) and 101 (bin)"` → `"26 and 5"`.
- **Implementation**:
  - Create `NumericConversionStage` to detect `(hex)` or `(bin)` markers.
  - Convert safely and handle invalid numbers gracefully.
- **Validation**: Confirm outputs are correct, and invalid cases log warnings without crashes.
