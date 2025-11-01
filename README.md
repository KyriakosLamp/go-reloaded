# go-reloaded

A simple text formatter that fixes common writing mistakes and applies transformations using special markers.

## What it does

Takes input text like this:
```
it (cap) was a historic race ,lewis hamilton (up,2) took pole position , scoring " 1E (hex) points " !
```

And turns it into this:
```
It was an historic race, LEWIS HAMILTON took pole position, scoring "30 points"!
```

## How to use

```bash
# Run the program
go run . input.txt output.txt

# Or build and run
go build
./go-reloaded input.txt output.txt
```

## What it fixes

- **Numbers**: `1E (hex)` → `30`, `10 (bin)` → `2`
- **Case**: `hello (up)` → `HELLO` ||  `WORLD (low)` → `world` || `bridge (cap)` → `Bridge`
- **Multiple words**: `this is fun (up, 2)` → `this IS FUN`
- **Punctuation spacing**: `hello ,world !!` → `hello, world!!`
- **Quote spacing**: `' awesome '` → `'awesome'`
- **Grammar**: `a apple` → `an apple`

## Why pipeline architecture?

Each transformation is a separate stage, making it:
- Easy to test individual parts
- Simple to add new rules
- Clean and maintainable code
- Perfect for team development

## Tests

```bash
go test ./tests -v
```

```
=== RUN   TestErrorHandling
=== RUN   TestErrorHandling/Invalid_hex_should_remain_unchanged
WARNING: Invalid hex number: XYZ
=== RUN   TestErrorHandling/Invalid_binary_should_remain_unchanged
WARNING: Invalid binary number: 102
=== RUN   TestErrorHandling/Multiple_invalid_conversions
WARNING: Invalid hex number: GHI
WARNING: Invalid binary number: 123
=== RUN   TestErrorHandling/Mixed_valid_and_invalid_hex
WARNING: Invalid hex number: XYZ
=== RUN   TestErrorHandling/Mixed_valid_and_invalid_binary
WARNING: Invalid binary number: 102
--- PASS: TestErrorHandling (0.00s)
    --- PASS: TestErrorHandling/Invalid_hex_should_remain_unchanged (0.00s)
    --- PASS: TestErrorHandling/Invalid_binary_should_remain_unchanged (0.00s)
    --- PASS: TestErrorHandling/Multiple_invalid_conversions (0.00s)
    --- PASS: TestErrorHandling/Mixed_valid_and_invalid_hex (0.00s)
    --- PASS: TestErrorHandling/Mixed_valid_and_invalid_binary (0.00s)
=== RUN   TestLoggerFunctionality
WARNING: Test warning message
ERROR: Test error message
--- PASS: TestLoggerFunctionality (0.00s)
=== RUN   TestFullPipelineIntegration
--- PASS: TestFullPipelineIntegration (0.00s)
=== RUN   TestComplexCombinations
=== RUN   TestComplexCombinations/Numeric_+_Case
=== RUN   TestComplexCombinations/Case_+_Quotes_+_Punctuation
=== RUN   TestComplexCombinations/Article_+_Case
=== RUN   TestComplexCombinations/Article_+_Case_Multi-word
=== RUN   TestComplexCombinations/Case_+_Punctuation
--- PASS: TestComplexCombinations (0.00s)
    --- PASS: TestComplexCombinations/Numeric_+_Case (0.00s)
    --- PASS: TestComplexCombinations/Case_+_Quotes_+_Punctuation (0.00s)
    --- PASS: TestComplexCombinations/Article_+_Case (0.00s)
    --- PASS: TestComplexCombinations/Article_+_Case_Multi-word (0.00s)
    --- PASS: TestComplexCombinations/Case_+_Punctuation (0.00s)
=== RUN   TestNumericConversion
--- PASS: TestNumericConversion (0.00s)
=== RUN   TestArticleAgreement
--- PASS: TestArticleAgreement (0.00s)
=== RUN   TestCaseTransform
--- PASS: TestCaseTransform (0.00s)
=== RUN   TestQuotation
--- PASS: TestQuotation (0.00s)
=== RUN   TestPunctuation
--- PASS: TestPunctuation (0.00s)
=== RUN   TestPhaseOneIntegration
--- PASS: TestPhaseOneIntegration (0.00s)
=== RUN   TestCLIIntegration
--- PASS: TestCLIIntegration (0.00s)
=== RUN   TestPipelineStructure
--- PASS: TestPipelineStructure (0.00s)
=== RUN   TestPhaseTwoIntegration
--- PASS: TestPhaseTwoIntegration (0.00s)
=== RUN   TestFileIO
--- PASS: TestFileIO (0.00s)
PASS
ok      go-reloaded/tests    (cached)
```

That's it! Simple text formatting made easy. 🚀