# LeetCode 271 — My Codec attempt: mistake → underlying Go concept

[← Foundations index](README.md) · [LeetCode 271 problem](https://leetcode.com/problems/encode-and-decode-strings/)

This is a **debugging and revision record**, not a claim that the problem has been mastered. The first attempt had sensible building blocks—`Codec`, `Constructor`, `len`, `strconv.Itoa`, `append` and `Split`—but several type/scope mistakes and one critical serialization ambiguity.

## Mistakes grouped by concept (study the lesson, not the error message)

| Attempted expression | What goes wrong | Underlying concept / existing lesson |
|---|---|---|
| `func Constructor() Codec { return &Codec{...} }` | `&Codec{}` has type `*Codec`, not the declared `Codec` return type. | [11. Structs, pointers, methods](11-structs-pointers-methods-and-linked-lists/README.md), [20. Function signatures and types](20-function-signatures-types-and-scope/README.md) |
| `Codec{strEncoded: string}` | `string` is a type, not an expression/value. An empty string value is `""`. | [20. Types versus values](20-function-signatures-types-and-scope/README.md) |
| Storing every encoded output in `this.strEncoded` | Repeated calls could accumulate previous results. Encoding does not require persistent state. | [11. Stateful versus stateless structs](11-structs-pointers-methods-and-linked-lists/README.md) |
| `append(this.strEncoded, ...)` | `append` takes a slice, not a `string`. | [01. Slice operations](01-arrays-and-slices/README.md) |
| `strconv.Itoa(len(s)) + s + "#"` | Decoder cannot reliably locate the end of a multi-digit length; the separator is in the wrong place. | [04. Length-prefix framing](04-strings-bytes-runes-and-parsing/README.md) |
| `strings.Split(strEncoded, "#")` | `strEncoded` is not a local name; use the `s` argument (or a receiver field if appropriate). Even `strings.Split(s, "#")` destroys payload boundaries if the payload contains `#`. | [20. Scope and receivers](20-function-signatures-types-and-scope/README.md), [04. Delimiter collision](04-strings-bytes-runes-and-parsing/README.md) |
| `var out string` in a function returning `[]string` | A single string is not a slice of strings. | [01. Slice declarations](01-arrays-and-slices/README.md), [20. Return contracts](20-function-signatures-types-and-scope/README.md) |
| `out := append(out, s)` inside the loop | Once `out` is correctly a slice, `:=` inside the loop creates a shadow variable instead of updating the outer result. | [20. `:=` versus `=` and shadowing](20-function-signatures-types-and-scope/README.md) |
| No byte-offset scan in `decode` | The decoder needs indices to read the length, then exactly that many payload bytes; slicing ends are exclusive. | [02. Loops and bounds](02-loops-and-bounds/README.md), [04. Byte slicing](04-strings-bytes-runes-and-parsing/README.md) |

## Algorithm invariant — the information you must preserve

Use `decimal byte length + # + exact payload bytes`, for example:

```text
["a#bcd", ""]  →  "5#a#bcd0#"
```

During decoding, `#` only separates a length from its payload. Once a length is known, take **exactly** that number of bytes, whether those bytes include `#` or not. Never split the entire encoded string on `#`.

| Record | `i` at record start | `j` at length separator | Parsed length | `start = j + 1` | `end = start + length` | Extract `s[start:end]` |
|---|---:|---:|---:|---:|---:|---|
| First | 0 | 1 | 5 | 2 | 7 | `"a#bcd"` |
| Second | 7 | 8 | 0 | 9 | 9 | `""` |

After each record, set `i = end`. Loop invariant: all bytes before `i` have been decoded into complete original strings.

## What was already covered vs what was added

- **Already covered:** 01 had basic slice append; 02 had exclusive slice bounds; 04 had `Atoi` and bytes versus runes; 11 had struct fields and pointer receivers. Those explanations remain.
- **Clarified in those lessons:** strings versus slices; `Itoa` versus `Atoi`; delimiter collisions; constructor return types; whether object state is necessary.
- **Previously missing:** an explicit foundation for function signatures, type names versus values, reassignment, nested scope, and shadowing. See the new [20. Function signatures, types and scope](20-function-signatures-types-and-scope/README.md).

## Practice and recall (attempt cold)

1. Correct a constructor with signature `func Constructor() Codec` and explain the type of `Codec{}` versus `&Codec{}`.
2. Explain why `var out string` and `out := []string{}` represent different types, even when both are empty.
3. Predict the result when you use `out := append(out, item)` inside a loop block instead of `out = append(out, item)`.
4. Encode and dry-run decode for `[]string{"ab#cd", "", "é"}`. Remember: `len("é")` is **2 bytes**.
5. Test encoding and decoding an empty slice and a slice containing one empty string. They are different inputs.

After practicing, record actual progress—not assumed mastery—in [REVISION-LOG.md](../REVISION-LOG.md).
