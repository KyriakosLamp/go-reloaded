# Code Analysis and Test Cases  
by klampria

##  Code Analysis

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

##  Architecture Options

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

