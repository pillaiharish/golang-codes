# 04 — Strings, bytes, runes, slicing and parsing

[← All foundations](../README.md) · [← LeetCode index](../../README.md)

## Why this matters
Strings are immutable UTF-8 byte sequences. Sorting anagram characters, extracting fixed-position fields, and indexing a 26-letter alphabet require distinguishing bytes from runes.

## Go operations: learn and explain
- `len(s)` counts bytes; `s[i]` returns a byte. `for _,r := range s` decodes Unicode runes.
- In `for i, r := range s`, `i` is the **byte offset** of the rune's first byte, not an incrementing character position. `r` has type `rune` (`int32`); `s[i]` has type `byte` (`uint8`).
- `unicode.IsDigit` accepts a `rune`, so `unicode.IsDigit(s[i])` does not compile (`byte` is not `rune`). `unicode.IsDigit(rune(s[i]))` compiles but converts just **one byte**; it does not decode a multibyte UTF-8 character.
- For ASCII-only number prefixes use `s[i] >= '0' && s[i] <= '9'` and then `strconv.Atoi` on the prefix. `unicode.IsDigit` also accepts non-ASCII decimal digits that `strconv.Atoi` does not parse as ASCII decimal text.
- `strings.Split(s, "")` splits UTF-8 into code points, not grapheme clusters.
- `strings.Join(parts, "")` builds a sorted signature from sorted pieces.
- `record[11:13]` selects two bytes, provided length is at least 13.
- `age, err := strconv.Atoi(record[11:13])` converts decimal text; handle `err`.
- `strconv.Itoa(n)` converts an `int` to decimal text; `strconv.Atoi(text)` converts decimal text back to an `int`. These operate in opposite directions.
- `s[i]-'a'` is safe only when the input contract limits bytes to lower-case ASCII a–z.

## Runnable example
```go
package main
import ("fmt";"strings";"sort";"strconv";"unicode")
func main() {
    parts := strings.Split("eat", "")
    sort.Strings(parts)
    fmt.Println(strings.Join(parts, "")) // aet
    record := "7868190130M7522"
    age, err := strconv.Atoi(record[11:13])
    fmt.Println(age, err, age > 60) // 75 <nil> true
    fmt.Println(len("é"), len([]rune("é"))) // 2 1

    ascii := "4asdf1a"
    for i := 0; i < len(ascii); i++ {
        if ascii[i] >= '0' && ascii[i] <= '9' {
            fmt.Printf("ASCII digit at byte %d: %c\n", i, ascii[i])
        }
    }
    for i, r := range ascii {
        if unicode.IsDigit(r) {
            fmt.Printf("Unicode digit at byte %d: %c\n", i, r)
        }
    }
    mixed := "4٤a" // The middle character is Arabic-Indic digit four.
    fmt.Println("bytes / runes:", len(mixed), len([]rune(mixed))) // 4 3
    for i, r := range mixed {
        if unicode.IsDigit(r) {
            fmt.Printf("Unicode digit at byte %d: %c\n", i, r)
        }
    }
}
```

## Dry run
`eat` → `e,a,t` → `a,e,t` → `aet`; the fixed-format record positions 11 and 12 contain `75`.

For `"4asdf1a"`, both ASCII byte testing and rune testing find `4` and `1`; they have the same numerical codes, but different Go types. For `"4٤a"`, UTF-8 byte indexing visits 4 bytes, while `range` visits 3 runes at byte offsets **0, 1, 3**. `unicode.IsDigit` on decoded runes finds both `4` and `٤`; ASCII byte testing finds only `4`.

## Byte versus rune: the decision rule

| Task | Use | Reason |
|---|---|---|
| Read a protocol delimiter or ASCII decimal length prefix | `s[i]` and `'0' <= s[i] && s[i] <= '9'` | Protocol syntax uses ASCII bytes; slicing offsets are bytes. |
| Inspect arbitrary Unicode decimal digits | `for _, r := range s` with `unicode.IsDigit(r)` | A Unicode digit may occupy multiple UTF-8 bytes. |
| Get the number of bytes in an encoded payload | `len(s)` | Matches Go string-slice boundaries `s[start:end]`. |
| Count code points (not visual grapheme clusters) | `utf8.RuneCountInString(s)` or `len([]rune(s))` | Counts decoded runes, not bytes. |

Do not confuse a digit's **character code** with its numeric value: ASCII `'4'` is byte value `52`, not integer `4`. After checking it is between `'0'` and `'9'`, subtract `'0'` to get the numeric digit (`'4' - '0' == 4`). Without the digit check, `'a' - '0'` also produces a number (`49`), but it is not a decimal digit.

## Tiny exercise (do before opening a hard problem)
Extract the age from `5303914400F9211`; answer 92. Explain why `[26]int` is not valid for an arbitrary Unicode anagram.

Then, without running code, predict both loops for `"A😊B"`: its byte length, its rune count, and the byte offsets from `for i, r := range s`. Explain why casting `rune(s[1])` does not produce `😊`.

## Typical mistakes
- `strings.Split(encoded, "#")` is unsafe for arbitrary payload strings: a payload can itself contain `#`. A reversible framing scheme is `length#payload`: read the decimal length up to the **first** `#`, then consume exactly that many bytes, even if those bytes include `#`.
- `strconv.Itoa(len(s)) + s + "#"` has no explicit boundary between a potentially multi-digit length and payload. Use `strconv.Itoa(len(s)) + "#" + s` instead.
- `len("é") == 2` because Go counts UTF-8 bytes. A byte-length prefix is correct if decoding uses byte offsets and slicing, rather than rune counts.
- `unicode.IsDigit(s[i])` is a type mismatch: `s[i]` is a byte, while `IsDigit` requires a rune. Even after `rune(s[i])`, non-ASCII digits are not reconstructed from UTF-8 bytes.
- Repeated `answer += piece` is fine for a small learning example, but can repeatedly copy accumulated data; prefer `strings.Builder` for substantial output.
- Unvalidated byte slicing may panic on too-short strings.
- Unicode code points differ from user-perceived letters/combining sequences.
- Do not use `s[i]` to represent a Unicode character unless the alphabet is ASCII.

## Apply to the existing repository
- [49. Group Anagrams](../../49.%20Group%20Anagrams/main.go)
- [242. Valid Anagram](../../242.%20Valid%20Anagram/main.go)
- [208. Implement Trie (Prefix Tree)](../../208.%20Implement%20Trie%20%28Prefix%20Tree%29/main.go)
- [2273. Find Resultant Array After Removing Anagrams](../../2273.%20Find%20Resultant%20Array%20After%20Removing%20Anagrams/main.go)
- [2678. Number of Senior Citizens](../../2678.%20Number%20of%20Senior%20Citizens/main.go)

## 30-second closed-book recall
Say what state this building block stores, how to traverse/update it, and which boundary case could invalidate your solution. Then rewrite the snippet without copying.
