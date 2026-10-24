# 04 — Strings, bytes, runes, slicing and parsing

[← All foundations](../README.md) · [← LeetCode index](../../README.md)

## Why this matters
Strings are immutable UTF-8 byte sequences. Sorting anagram characters, extracting fixed-position fields, and indexing a 26-letter alphabet require distinguishing bytes from runes.

## Go operations: learn and explain
- `len(s)` counts bytes; `s[i]` returns a byte. `for _,r := range s` decodes Unicode runes.
- `strings.Split(s, "")` splits UTF-8 into code points, not grapheme clusters.
- `strings.Join(parts, "")` builds a sorted signature from sorted pieces.
- `record[11:13]` selects two bytes, provided length is at least 13.
- `age, err := strconv.Atoi(record[11:13])` converts decimal text; handle `err`.
- `strconv.Itoa(n)` converts an `int` to decimal text; `strconv.Atoi(text)` converts decimal text back to an `int`. These operate in opposite directions.
- `s[i]-'a'` is safe only when the input contract limits bytes to lower-case ASCII a–z.

## Runnable example
```go
package main
import ("fmt";"strings";"sort";"strconv")
func main() {
    parts := strings.Split("eat", "")
    sort.Strings(parts)
    fmt.Println(strings.Join(parts, "")) // aet
    record := "7868190130M7522"
    age, err := strconv.Atoi(record[11:13])
    fmt.Println(age, err, age > 60) // 75 <nil> true
    fmt.Println(len("é"), len([]rune("é"))) // 2 1
}
```

## Dry run
`eat` → `e,a,t` → `a,e,t` → `aet`; the fixed-format record positions 11 and 12 contain `75`.

## Tiny exercise (do before opening a hard problem)
Extract the age from `5303914400F9211`; answer 92. Explain why `[26]int` is not valid for an arbitrary Unicode anagram.

## Typical mistakes
- `strings.Split(encoded, "#")` is unsafe for arbitrary payload strings: a payload can itself contain `#`. A reversible framing scheme is `length#payload`: read the decimal length up to the **first** `#`, then consume exactly that many bytes, even if those bytes include `#`.
- `strconv.Itoa(len(s)) + s + "#"` has no explicit boundary between a potentially multi-digit length and payload. Use `strconv.Itoa(len(s)) + "#" + s` instead.
- `len("é") == 2` because Go counts UTF-8 bytes. A byte-length prefix is correct if decoding uses byte offsets and slicing, rather than rune counts.
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
