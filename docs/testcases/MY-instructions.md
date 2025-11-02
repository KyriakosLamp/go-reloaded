# New Test File - Tricky Cases

## Test 1: Multiple Articles with Different Cases
**Input**: `a (up) apple and a (cap) orange but a (low) umbrella`

**Expected Output**: 
```
AN apple and An orange but an umbrella
```

---

## Test 2: Nested Quotes with Punctuation
**Input**: `he said ' she replied " yes , that is a (cap) excellent idea ! " '`

**Expected Output**: 
```
he said 'she replied "yes, that is An excellent idea!"'
```

---

## Test 3: Mixed Numeric Conversions with Articles
**Input**: `I bought a (cap) 1010 (bin) items and a (up) FF (hex) more`

**Expected Output**: 
```
I bought A 10 items and A 255 more
```

---

## Test 4: Complex Case Transformations
**Input**: `the quick brown fox (up, 2) jumps over a (cap) lazy dog (cap, 3)`

**Expected Output**: 
```
the quick BROWN FOX jumps over A Lazy Dog
```

---

## Test 5: Multiple Punctuation with Quotes
**Input**: `wow !! ' this is a (up) amazing result , right ? ' he asked .`

**Expected Output**: 
```
wow!! 'this is AN amazing result, right?' he asked.
```

---

## Test 6: Edge Case - Article at End
**Input**: `this is a (cap) hour and that was a (up)`

**Expected Output**: 
```
this is An hour and that was A
```

---

## Test 7: Consecutive Transformations
**Input**: `hello (cap) world (up) and a (low) elephant (cap) story`

**Expected Output**: 
```
Hello WORLD and an Elephant story
```

---

## Test 8: Binary/Hex with Case and Articles
**Input**: `a (cap) 101 (bin) errors and a (up) A (hex) warnings occurred`

**Expected Output**: 
```
A 5 errors and A 10 warnings occurred
```

---

## Test 9: Complex Quote Nesting with Cases
**Input**: `' I think " a (cap) apple (up, 2) " is better ' , he said (cap, 2) .`

**Expected Output**: 
```
'I think "AN APPLE" is better', He Said.
```

---

## Test 10: Multiple Articles in Sequence
**Input**: `a (cap) elephant , a (up) ant , and a (low) umbrella walked together`

**Expected Output**: 
```
An elephant, AN ant, and an umbrella walked together
```

---
## Test 11: Tricky case after case
**Input**: `this (cap) is  a awesome day at , the office (cap,3) (up,2) . 101 (bin) more days until we reach " the (up) 64 (hex) days mark "  !!`

**Expected Output**: 
```
This is an awesome day At, THE OFFICE. 5 more days until we reach "THE 100 days mark"!!
```

---
