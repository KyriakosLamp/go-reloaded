# Code Analysis and Test Cases  
by klampria

## 1. Code Analysis

The **go-reloaded** project is a Go-based text reformatter and editor designed to process an input text file and produce a corrected or enhanced output file. The program applies a series of transformation rules to the text to correct case formatting, punctuation spacing, and linguistic details. It also handles number conversions (binary, hexadecimal) and article agreement (“a” → “an”).

### General Functionality
- **Input:** a text file containing text with transformation markers (e.g., `(hex)`, `(up, 2)`, `(low)`).
- **Output:** a new text file with all corrections applied.

### Transformation Rules

1. **(hex)** — Convert the previous hexadecimal word into decimal.  
   - Example: `1E (hex)` → `30`
2. **(bin)** — Convert the previous binary word into decimal.  
   - Example: `10 (bin)` → `2`
3. **(up)** — Convert the previous word to uppercase.  
   - Example: `go (up)` → `GO`
4. **(low)** — Convert the previous word to lowercase.  
   - Example: `HELLO (low)` → `hello`
5. **(cap)** — Capitalize the previous word (first letter uppercase).  
   - Example: `bridge (cap)` → `Bridge`
6. **(up, n)**, **(low, n)**, **(cap, n)** — Apply transformation to *n* previous words.  
   - Example: `this is fun (up, 2)` → `this IS FUN`
7. **Punctuation Rules:**  
   - Punctuation (`. , ! ? : ;`) should be attached to the preceding word and separated from the following one by a space.  
   - Multiple punctuations (e.g., `!?`, `...`) stay grouped together.  
   - Example: `hello ,world !!` → `hello, world!!`
8. **Quotation Marks ('') ("")** — Should hug the quoted words with no inner spacing.  
   - Example: `' awesome '` → `'awesome'`
   - Example: `" awesome"` → `"awesome"`
9. **Article Agreement (a → an)** — If the next word starts with a vowel (`a, e, i, o, u`) or `h`, replace `a` with `an`.  
   - Example: `a orange car` → `an orange car`
   - Example: `i need a hero` → `i need an hero`

This program emphasizes text normalization and formatting consistency, suitable for auto-editing or preprocessing text before further analysis.

---

## 2. Architecture Options

### File Management System (FMS)
An FMS architecture organizes the program around file operations. It reads the entire input file, processes its content as a single data entity, and writes the result to the output file. Each transformation rule can be applied sequentially in a monolithic way, where the main focus is file handling and manipulation.

- **Pros:** Simple flow, easy to implement, and minimal concurrency.
- **Cons:** Harder to scale, debug, or parallelize. Every transformation is dependent on the previous one’s completion.

### Pipeline Architecture
The pipeline approach splits the text processing into independent **stages**, where each stage applies a specific transformation (e.g., case changes, punctuation cleanup, conversions). The output of one stage becomes the input for the next. This makes the code modular, readable, and more efficient for concurrent processing in Go (via channels and goroutines).

- **Pros:** Modular, scalable, easier debugging.
- **Cons:** Slightly more complex setup due to inter-stage data flow.

---

### Why I Chose Pipeline

I chose the **Pipeline architecture** because it fits my flow of thinking perfectly. Each rule or transformation can act as a separate stage, making the code easier to test, extend, and maintain. Pipelines also help avoid large loops and  makes it easy to break the project into smaller ones and distribute to each team member,  in case of a team project. This approach improves readability, reduces the risk of side effects, and allows future optimizations such as parallel rule application.


---

## 3. Test Cases (Natural Language Format)

### 3.1 Simple Test Cases (One Rule Each)

1. **Hexadecimal Conversion**  
   Input: “1E (hex) files were added.”  
   Output: “30 files were added.”

2. **Binary Conversion**  
   Input: “10 (bin) days later.”  
   Output: “2 days later.”

3. **Uppercase Conversion**  
   Input: “go (up) now!”  
   Output: “GO now!”

4. **Lowercase Conversion**  
   Input: “HELLO (low) there.”  
   Output: “hello there.”

5. **Capitalize Conversion**  
   Input: “bridge (cap) is nice.”  
   Output: “Bridge is nice.”

6. **Multi-word Capitalization**  
   Input: “this is magic (cap, 2)”  
   Output: “this Is Magic”

7. **Punctuation Fixing**  
   Input: “Hello ,world !!”  
   Output: “Hello, world!!”

8. **Grouped Punctuation**  
   Input: “Wait ... what ?”  
   Output: “Wait... what?”

9. **Quotation Rule**  
   Input: “He said: ' wow '.”  
   Output: “He said: 'wow'.”

10. **Article Agreement**  
   Input: “a amazing car.”  
   Output: “an amazing car.”

---

### 3.2 Tricky / Combined Test Cases

1. “1A (hex) files and 10 (bin) more (up)” → “26 files and 2 MORE”  
2. “He shouted (up, 2): ' run now ' !!” → “HE SHOUTED: 'run now'!!”  
3. “I met a (cap) honest man.” → “I met An honest man.”  
4. “That was amazing (cap, 3) ,right ?” → “That Was Amazing, right?”  
5. “This is (low, 2) 101 (bin) test.” → “this is 5 test.”  
6. “a old hero (cap, 3)” → “An Old Hero”  
7. “wow (up) ... great ,job !” → “WOW... great, job!”  
8. “The code (low,2) IS (up) nice.” → “the code IS nice.”  
9. “‘ too much 'noise(cap, 3)” → “‘Too Much' Noise”  
10. “He said ' hello world ' ,and (up,5) SMILED (low)  .” → “HE SAID 'HELLO WORLD', AND smiled.”

---

### 3.3 Medium Paragraph — Formula One Context (All Rules Combined)

**Input:**
```
it (cap) was a historic race ,lewis hamilton (up,2) took pole position  , scoring " 1E (hex) points  " . it (cap) was ' amazing ' to watch a incredible battle on track ! The (low) crowd shouted (up) : go lewis go (cap, 3) !! the fans waved flags (cap, 4) and a HERO (low) returned to the podium . WHAT A INCREDIBLE MOMENT (low, 3) for the sport !

```

**Expected Output:**
```
It was a historic race, LEWIS HAMILTON took pole position, scoring "30 points". It was 'amazing' to watch an incredible battle on track! the crowd SHOUTED: Go Lewis Go!! The Fans Waved Flags and an hero returned to the podium. WHAT an incredible moment for the sport!

```

---


