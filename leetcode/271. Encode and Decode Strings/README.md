# 271 — Encode and Decode Strings (Go)

[← LeetCode index](../README.md) · [Commented solution](main.go) · [Table-driven tests](main_test.go) · [Quick revision card](REVISION-CARD.md) · **[My original mistakes mapped to Go concepts](../0-problem-solving-building-blocks/271-codec-attempt-concept-map.md)** · [Problem link](https://leetcode.com/problems/encode-and-decode-strings/) · [Revision log](../REVISION-LOG.md)

**Pattern:** Self-delimiting length-prefix framing · **Difficulty:** Medium (premium question).

## Problem contract and test cases

Design `Encode([]string) string` and `Decode(string) []string` so decoding an encoding recovers **exactly the same list in the same order**. Input strings may contain separator characters, digits, empty strings, and Unicode. The encoding itself does **not** have to match one special output representation; the round trip is what matters.

```go
type Codec struct{}
func Constructor() Codec
func (c *Codec) Encode(strs []string) string
func (c *Codec) Decode(s string) []string
```

Our chosen format is **`<decimal byte length>#<raw string bytes>`**, repeated with no separator between records.

| Input | This Codec's encoding | Decoded result |
|---|---|---|
| `[]string{"neet","code"}` | `4#neet4#code` | `[]string{"neet","code"}` |
| `[]string{"a#b",""}` | `3#a#b0#` | `[]string{"a#b",""}` |
| `[]string{}` | `""` | `[]string{}` |
| `[]string{""}` | `0#` | `[]string{""}` |
| `[]string{"é"}` | `2#é` | `[]string{"é"}` |

These are examples of **our own encoding format**, not a prescribed LeetCode serialization. A valid input to `Decode` is produced by `Encode`; malformed arbitrary encoded strings do not have a specified recovery behavior in this implementation.

## Before solving: foundational Go concepts

1. [01 — Slices and append](../0-problem-solving-building-blocks/01-arrays-and-slices/README.md): `[]string` return values and `append`.
2. [02 — Indexing and exclusive bounds](../0-problem-solving-building-blocks/02-loops-and-bounds/README.md): cursor `i`, separator index `j`, `s[start:end]`.
3. [04 — Strings, bytes, runes, and `strconv`](../0-problem-solving-building-blocks/04-strings-bytes-runes-and-parsing/README.md): `len`, `Atoi`, `Itoa`, bytes versus runes.
4. [11 — Structs and methods](../0-problem-solving-building-blocks/11-structs-pointers-methods-and-linked-lists/README.md): `Codec` receiver and constructor.
5. [20 — Function contracts and scope](../0-problem-solving-building-blocks/20-function-signatures-types-and-scope/README.md): type `Codec` versus `*Codec`; `:=` shadowing; function parameters.

## 1. Recognition trigger

**Question:** Multiple strings, including strings containing delimiters, must be serialized and reconstructed with no ambiguity. **Trigger:** Give each payload its own length, then read exactly that many bytes instead of splitting on content.

## 2. Intuition: why a delimiter alone is insufficient

`strings.Join(strs, "#")` cannot distinguish `[]string{"a#b"}` from `[]string{"a","b"}`; both can yield `a#b`. A delimiter *after* an unrestricted payload is ambiguous. In `3#a#b`, the first `#` closes the length **3**, after which the decoder consumes precisely three bytes `a#b`; the embedded `#` is merely data.

Your two submitted versions use the same correct approach. One constructs `&Codec{}` directly; the other exposes `Constructor() Codec`. The main solution keeps the **constructor form** for a recognizable design-question API and is still stateless: no encoded string needs to be stored on the struct. `codec := Constructor()` can call pointer-receiver methods because `codec` is an addressable local value. `(&Codec{}).Encode(...)` works too.

## 3. Invariant / proof of round-trip correctness

**Encode invariant:** after encoding the first `k` strings, the builder contains `k` complete records, each with its exact decimal byte count and payload.

**Decode invariant:** at the beginning of each iteration, `i` points to the start of the **next length prefix**; all bytes before it have been decoded into complete strings. Find `j` at the prefix separator, parse the ASCII decimal length `L`, then slice `s[j+1:j+1+L]`. Set `i` to that exclusive end. Because the decoder trusts the byte count, `#` and Unicode *inside the payload* cannot move the boundary.

## 4. Numbered implementation recipe (and why)

**Encode**

1. Start a `strings.Builder` — append text without repeatedly constructing whole new strings.
2. For each `str`, append `strconv.Itoa(len(str))` — record the **byte** size before the payload.
3. Append `'#'`, then `str` — the delimiter ends a variable-width decimal length, while payload bytes stay untouched.

**Decode**

4. Set `i := 0`. While `i < len(s)`, move `j` to the **next** `#` — this separates digits from data.
5. Parse `L := strconv.Atoi(s[i:j])`; set `start := j+1`, `end := start+L` — compute exact exclusive payload bounds.
6. `append(result, s[start:end])` and `i = end` — return one original string and start the next record at the correct offset.

The code discards the `Atoi` error **only because the problem passes strings produced by Encode**. For untrusted network traffic, validate separator existence, decimal syntax, lengths, and integer bounds and return errors explicitly.

## 5. Dry run: embedded separator and empty payload

Input `[]string{"a#b", ""}` encodes as `3#a#b0#`.

| Record | `i` | `j` | Parsed bytes `L` | `start` | `end` | Slice |
|---|---:|---:|---:|---:|---:|---|
| 1 | 0 | 1 | 3 | 2 | 5 | `"a#b"` |
| 2 | 5 | 6 | 0 | 7 | 7 | `""` |

After record 2, `i == len(s)`, so stop. Notice the decoder **didn't** treat the `#` at byte index 3 as a new record boundary.

## 6. Complexity and personal pitfalls

- Encoding and decoding each take O(B + D) time, where `B` is the total payload bytes and `D` is the total ASCII length-prefix bytes. Encoder output uses O(B + D) memory; decoder's result slice and strings require space proportional to decoded output. `strings.Builder` prevents repeated full-string concatenation.
- `len("é") == 2` in UTF-8, while `for _, r := range "é"` yields **one rune**. `s[i]` returns one byte and `range s` yields decoded runes (`int32`) and their starting byte offsets. A byte length must be paired with **byte slicing**, not a rune-count loop.
- `unicode.IsDigit(s[i])` **does not compile** because `s[i]` is `byte` but the function accepts `rune`; merely converting a UTF-8 byte to `rune` does not decode a Unicode character. The encoded prefix is ASCII from `Itoa`, so scan for `#` and pass its decimal substring to `Atoi`.
- `str` (`string`) is not a `[]string`. `append` adds to a **slice** and must be assigned back: `result = append(result, ...)`, not a nested shadow `result := append(...)`.
- `func Constructor() Codec` returns `Codec{}`, **not** `&Codec{}` (which has type `*Codec`). The encoder should use local builder state; a persistent `strEncoded` field could accidentally leak between calls.
- **Never use `strings.Split(s,"#")`**: an arbitrary payload may contain `#`. A single empty string encodes as `0#`, while an empty list encodes as `""`.

For the exact original mistakes and their corrected Go building blocks, revisit the [attempt concept map](../0-problem-solving-building-blocks/271-codec-attempt-concept-map.md); it preserves the failed stage separately from the solved implementation.

## 7. 45-second answer and cold recall

> I use length-prefix framing so payload delimiters don't matter. Each record is its decimal UTF-8 byte length, a `#` marker, and the raw string. The decoder scans the length prefix, computes the exact payload slice and jumps to its end. An empty string has a length-zero record, unlike an empty list. Both directions are linear in encoded bytes; my Codec is stateless.

1. Encode/decode `[]string{"a#b", "", "é"}` by hand with byte offsets.
2. Why is `len("é")` not the number of runes? Which loop is correct for byte offsets?
3. What breaks if you use `strings.Split` or forget `i = end`?
4. Why do `Codec{}` and `&Codec{}` have different types? Why can the value call pointer methods?
5. Distinguish `[]string{}` from `[]string{""}` and explain `result = append(...)`.

**Personal review:** Confidence `— / 5` · Last cold recall `—` · Next recall `—`. Record attempts, not assumed mastery, in [REVISION-LOG.md](../REVISION-LOG.md).

Run locally: `cd 'leetcode/271. Encode and Decode Strings' && GO111MODULE=off go test -race -v`.
