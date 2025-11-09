## Task 05: Implement CaseTransformStage (Unified)
- completed: ✅

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
