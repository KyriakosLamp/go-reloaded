# Memory Bank - go-reloaded Project

## Project Overview
**go-reloaded** is a Go-based text reformatter that processes input files with transformation markers and produces corrected output files. Uses Pipeline architecture for modular, testable transformations.

## Core Architecture Decision
**Pipeline Pattern** - Each transformation rule is a separate stage, enabling:
- Easy testing and maintenance
- Modular development for team distribution
- Reduced side effects and improved readability
- Future parallel processing optimization

## Transformation Rules Reference

### Number Conversions
- `(hex)` - Convert previous hex word to decimal: `1E (hex)` → `30`
- `(bin)` - Convert previous binary word to decimal: `10 (bin)` → `2`

### Case Transformations
- `(up)` - Uppercase previous word: `go (up)` → `GO`
- `(low)` - Lowercase previous word: `HELLO (low)` → `hello`
- `(cap)` - Capitalize previous word: `bridge (cap)` → `Bridge`
- `(up, n)`, `(low, n)`, `(cap, n)` - Apply to n previous words: `this is fun (up, 2)` → `this IS FUN`

### Punctuation Rules
- Attach to preceding word, space after: `hello ,world !!` → `hello, world!!`
- Group multiple punctuation: `Wait ... what ?` → `Wait... what?`
- Supported: `. , ! ? : ;`

### Quotation Rules
- Remove inner spacing: `' awesome '` → `'awesome'`
- Works with both single ('') and double ("") quotes

### Article Agreement
- `a` → `an` before vowels (a,e,i,o,u) or h: `a orange` → `an orange`

## Critical Implementation Notes

### Processing Order (STRICT SEQUENCE)
1. **Number conversions** (hex, bin) - Must happen first to avoid case changes affecting hex digits
2. **Case transformations** (up, low, cap) - Apply before punctuation to avoid affecting markers
3. **Article agreement** (a → an) - Before punctuation to handle "a" at sentence boundaries
4. **Punctuation spacing** - Fix spacing around punctuation marks
5. **Quotation mark hugging** - Final step to ensure proper quote formatting

### Critical Pattern Recognition
- **Transformation markers** are always in parentheses: `(rule)` or `(rule, n)`
- **Word boundaries** matter: transformations apply to complete words only
- **Case sensitivity** in markers: all markers are lowercase
- **Number parameter** (n) applies to previous n words, not following words

### Critical Edge Cases & Patterns

#### Transformation Conflicts
- **Multiple rules on same word**: Process in pipeline order
- **Overlapping n-word ranges**: Later rules override earlier ones
- **Case + Article**: `a (cap) honest` → `An honest` (article agreement after capitalization)

#### Punctuation Complexities
- **Grouped punctuation**: `!!`, `...`, `!?` stay together
- **Punctuation + quotes**: `' text ' ,` → `'text',`
- **Multiple spaces**: Normalize to single spaces
- **Sentence boundaries**: Handle punctuation at start/end

#### Quote Handling
- **Alternating quotes**: First quote opens, next closes
- **Unmatched quotes**: Leave as-is (don't break text)
- **Quote + punctuation**: `'text',` not `'text' ,`

#### Article Agreement Specifics
- **Vowel sounds**: a,e,i,o,u and h trigger "an"
- **Case insensitive**: `A orange` → `An orange`
- **Word boundaries**: Only standalone "a" gets converted
- **After transformations**: Check final word form, not original

### Test Coverage Requirements
- Single rule applications
- Combined rule scenarios
- Edge cases and malformed input
- Large paragraph processing
- Boundary conditions (first/last words)

## File Structure Expectations
```
go-reloaded/
├── main.go              # Entry point
├── pipeline/            # Pipeline stages
├── transformers/        # Individual transformation rules
├── utils/              # Helper functions
├── test/               # Test files and cases
└── docs/               # Documentation
```

## Implementation Validation Checklist

### Core Functionality
- [ ] **Hex conversion**: Handles uppercase/lowercase hex digits
- [ ] **Binary conversion**: Validates binary format (0s and 1s only)
- [ ] **Case transformations**: Preserves non-alphabetic characters
- [ ] **N-parameter parsing**: Correctly extracts and applies count

### Text Processing
- [ ] **Word tokenization**: Handles punctuation boundaries correctly
- [ ] **Marker removal**: Completely removes transformation markers
- [ ] **Space normalization**: Single spaces between words
- [ ] **Boundary handling**: First/last word edge cases

### Complex Scenarios
- [ ] **Combined rules**: `1A (hex) files (up, 2)` → `26 FILES`
- [ ] **Quote + case**: `' hello ' (up, 2)` → `'HELLO'`
- [ ] **Article + case**: `a hero (cap, 2)` → `An Hero`
- [ ] **Punctuation chains**: Multiple consecutive punctuation marks

### Error Resilience
- [ ] **Invalid hex/bin**: Graceful handling of malformed numbers
- [ ] **Out-of-range n**: Handle n > available words
- [ ] **Unmatched quotes**: Don't break on odd number of quotes
- [ ] **Empty input**: Handle empty files and lines

## Implementation Patterns

### Pipeline Stage Design
```go
type Stage interface {
    Process(text string) string
}

type Pipeline struct {
    stages []Stage
}
```

### Word Processing Pattern
- **Tokenize**: Split on whitespace, preserve positions
- **Transform**: Apply rules to tokens
- **Reconstruct**: Rebuild with proper spacing

### Marker Detection
- **Regex pattern**: `\([a-z]+(?:,\s*\d+)?\)`
- **Backward lookup**: Find previous word(s) to transform
- **Clean removal**: Remove marker completely

### State Management
- **Immutable stages**: Each stage returns new string
- **Context passing**: Minimal state between stages
- **Error propagation**: Fail fast on critical errors

## Performance & Error Strategy

### Optimization
- **Single pass per stage**: Avoid multiple iterations
- **String builder**: Use for efficient concatenation
- **Regex compilation**: Compile patterns once
- **Memory pooling**: Reuse slices for tokenization

### Error Handling
- **Validation first**: Check input before processing
- **Graceful degradation**: Skip malformed markers, continue processing
- **Logging**: Record skipped transformations for debugging
- **Original preservation**: Never lose original text on errors

## Quick Reference - Critical Patterns

### Must-Handle Scenarios (From Test Cases)
```
# Number + Case Combination
"1A (hex) files (up, 2)" → "26 FILES"

# Quote + Punctuation + Case
"He shouted (up, 2): ' run now ' !!" → "HE SHOUTED: 'run now'!!"

# Article + Case Interaction
"a hero (cap, 2)" → "An Hero"
"I met a (cap) honest man." → "I met An honest man."

# Complex Punctuation
"wow (up) ... great ,job !" → "WOW... great, job!"

# Multi-word with Quotes
"' too much 'noise(cap, 3)" → "'Too Much' Noise"

# Chain Processing
"The code (low,2) IS (up) nice." → "the code IS nice."
```

### Transformation Priority Examples
1. **Numbers first**: `1E (hex) (up)` → `30` then `30`
2. **Case before articles**: `a (cap) apple` → `An apple`
3. **Punctuation last**: `hello (up) ,world` → `HELLO, world`

### Common Pitfalls to Avoid
- **Don't process markers twice**: Remove after first use
- **Preserve punctuation grouping**: `!!` stays `!!`, not `! !`
- **Handle quote alternation**: Track open/close state
- **Validate n-parameter bounds**: Don't exceed available words
- **Case-insensitive article detection**: `A apple` → `An apple`

## Testing Strategy
- **Unit tests**: Each transformation rule isolated
- **Integration tests**: Full pipeline with complex scenarios  
- **Property tests**: Verify invariants (no text loss, marker removal)
- **Performance tests**: Large file processing benchmarks
- **Regression tests**: All examples from test cases must pass