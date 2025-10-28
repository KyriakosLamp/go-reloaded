## Test Cases (Natural Language Format)

### 1 Simple Test Cases (One Rule Each)

1. **(hex)** /**Hexadecimal Conversion**  
   Input: “1E (hex) files were added.”  
   Output: “30 files were added.”

2. **(bin)** /**Binary Conversion**  
   Input: “10 (bin) days later.”  
   Output: “2 days later.”

3. **(up)** /**Uppercase Conversion**  
   Input: “go (up) now!”  
   Output: “GO now!”

4. **(low)** /**Lowercase Conversion**  
   Input: “HELLO (low) there.”  
   Output: “hello there.”

5. **(cap)** /**Capitalize Conversion**  
   Input: “bridge (cap) is nice.”  
   Output: “Bridge is nice.”

6. **(cap, n)** /**Multi-word Capitalization**  
   Input: “this is magic (cap, 2)”  
   Output: “this Is Magic”

7. **Punctuation Rules** /**Punctuation Fixing**  
   Input: “Hello ,world !!”  
   Output: “Hello, world!!”

8. **Punctuation Rules** /**Grouped Punctuation**  
   Input: “Wait ... what ?”  
   Output: “Wait... what?”

9. **Quotation Marks ('') ("")** /**Quotation Rule**  
   Input: “He said: ' wow '.”  
   Output: “He said: 'wow'.”

10. **Article Agreement (a → an)**  
   Input: “a amazing car.”  
   Output: “an amazing car.”

---

### 2 Tricky / Combined Test Cases

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

### 3 Medium Paragraph — Formula One Context (All Rules Combined)

**Input:**
```
it (cap) was a historic race ,lewis hamilton (up,2) took pole position  , scoring " 1E (hex) points  " . it (cap) was ' amazing ' to watch a incredible battle on track ! The (low) crowd shouted (up) : go lewis go (cap, 3) !! the fans waved flags (cap, 4) and a HERO (low) returned to the podium . WHAT A INCREDIBLE MOMENT (low, 3) for the sport !

```

**Expected Output:**
```
It was a historic race, LEWIS HAMILTON took pole position, scoring "30 points". It was 'amazing' to watch an incredible battle on track! the crowd SHOUTED: Go Lewis Go!! The Fans Waved Flags and an hero returned to the podium. WHAT an incredible moment for the sport!

```

---

