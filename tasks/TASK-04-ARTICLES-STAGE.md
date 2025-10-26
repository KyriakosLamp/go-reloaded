## Task 04: Implement Article Agreement Rule
**Goal**: Replace “a” with “an” when followed by a vowel (`a, e, i, o, u`) or `h`.

- **Test Writing**:
  - `"a orange car"` → `"an orange car"`
  - `"a honest man"` → `"an honest man"`
  - `"a user"` → `"an user"`.
- **Implementation**:
  - Create `ArticleStage` to check the next word’s first character.
- **Validation**: Confirm grammar corrections appear correctly in all test cases.
