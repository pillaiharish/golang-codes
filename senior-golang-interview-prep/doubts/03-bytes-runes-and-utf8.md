# 3. Bytes, runes, UTF-8, and visible characters

[All doubts](README.md) · [Previous](02-zero-values-make-and-new.md) · [Next: increments](04-increments-and-for-loops.md)

## Direct answer

A `rune` is an alias for `int32`; it conventionally represents a Unicode code point. It does **not** store the UTF-8 encoding of that code point. A `byte` is an alias for `uint8` and represents one byte. UTF-8 encodes a valid Unicode scalar value using one to four bytes. A Go string is an immutable sequence of bytes and can contain invalid UTF-8.

## Comparison

| Property | `byte` | `rune` |
|---|---|---|
| Alias | `uint8` | `int32` |
| Storage per value | 1 byte | 4 bytes |
| Integer range | 0 through 255 | -2^31 through 2^31-1 |
| Text role | One encoded byte | One decoded code point |
| String indexing / range value | `s[i]` | `for _, r := range s` |
| Typical use | Binary protocols, ASCII parsing | Unicode decoding and classification |

Aliases are the same types, not wrapper types. `var r rune = 300` is ordinary integer storage. A rune can hold negative integers or surrogate values; its type alone does not guarantee a valid Unicode scalar value. Use `utf8.ValidRune` when validating such input.

## Walk through A, é, and 界

```go
s := "Aé界"
fmt.Println(len(s))                    // 6 bytes
fmt.Println(utf8.RuneCountInString(s))  // 3 code points
fmt.Printf("% x\n", []byte(s))         // 41 c3 a9 e7 95 8c
for offset, r := range s {
    fmt.Printf("%d %U %c\n", offset, r, r)
}
// 0 U+0041 A
// 1 U+00E9 é
// 3 U+754C 界
```

Range indexes are **byte offsets**, not rune indexes. `s[1]` is the first byte `0xC3` of `é`, not the whole code point. Slicing `s[1:2]` is legal but produces an incomplete UTF-8 encoding. Use a rune conversion or decode boundaries when splitting text by code point.

```go
r := '界'              // numeric code point U+754C, decimal 30028
fmt.Println(int32(r))  // 30028
fmt.Println(string(r)) // 界, encoded as UTF-8 in the resulting string
fmt.Println(string(65)) // A, NOT the decimal text "65"
fmt.Println(strconv.Itoa(65)) // 65
```

`byte(r)` truncates the integer to eight bits; it does not UTF-8 encode it. Conversely, converting one UTF-8 byte to `rune` does not decode a multi-byte character.

## Code points are not grapheme clusters

`"é"` has one code point, while `"e\u0301"` has two code points that may render as one visible character. Emoji sequences can contain several code points joined into one grapheme cluster. Consequently:

- Byte count measures encoded size.
- Rune count measures code points, not necessarily visible characters.
- User-facing cursor movement and truncation may require grapheme segmentation.
- Equal-looking text may need an explicit Unicode normalization policy before equality checks.

Converting to `[]rune` does not perform normalization or grapheme segmentation.

## Invalid UTF-8

```go
s := string([]byte{0xff, 'A'})
fmt.Println(utf8.ValidString(s)) // false
for offset, r := range s {
    fmt.Printf("%d %U\n", offset, r)
}
// 0 U+FFFD
// 1 U+0041
```

Invalid encoding encountered by range yields `utf8.RuneError` (U+FFFD) and advances one byte. A correctly encoded U+FFFD is also RuneError as a value; `utf8.DecodeRuneInString` returns a width of 1 for the invalid-byte case, versus 3 for a valid encoded replacement character. Converting invalid input to `[]rune` and back does not preserve the original bytes.

## Memory and choosing a representation

A `[]rune` uses four bytes per element, plus descriptor/backing-storage overhead. Converting ASCII text to runes can expand element storage about fourfold. Iterating with `range` decodes incrementally without requiring a full rune slice. `[]byte(s)` and `[]rune(s)` provide mutable sequence values, with copy/allocation behavior subject to compiler optimization while preserving language semantics.

Use bytes for Kafka payloads, network frames, ASCII-delimited formats, and raw hashing. Use runes for Unicode-aware classification (`unicode.IsLetter`) or code-point algorithms. For DNS/security features, explicitly define whether “length” means bytes, code points, labels, or a normalized representation; changing that choice changes the feature.

## Interview checkpoints

- Predict the byte offsets from ranging over `"Aé界"`.
- Explain why `rune(s[1])` does not decode `é`.
- Explain why reversing runes can still damage a visible emoji sequence.
- Choose bytes or runes for a binary protocol and justify the memory cost.

Run `go run ./examples/text` from this directory: [source](examples/text/main.go).

Sources: [Go text model](https://go.dev/blog/strings), [unicode/utf8](https://pkg.go.dev/unicode/utf8), [unicode](https://pkg.go.dev/unicode), [strconv](https://pkg.go.dev/strconv).
