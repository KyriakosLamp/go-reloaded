## Task 06: Implement Quotation Formatting
- completed:

**Goal**: Clean spacing and ensure quotes hug inner text properly.

- **Test Writing**:
  - `"He said: ' wow '."` → `"He said: 'wow'."`
  - `"\" hello \""` → `"\"hello\""`.
- **Implementation**:
  - Create `QuotationStage` to trim spaces inside `' '` and `" "`.
  - Handle nested quotes gracefully.
- **Validation**: Confirm formatting correctness across all quotation variations.
