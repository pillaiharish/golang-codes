# #271 Encode and Decode Strings — 90-second cold recall card

[Full reasoning](README.md) · [Original attempt and concept gaps](../0-problem-solving-building-blocks/271-codec-attempt-concept-map.md) · [Commented code](main.go) · [Tests](main_test.go)

**Trigger:** Multiple arbitrary strings need unambiguous serialization, even if content contains a delimiter.

**Invariants:** Encoder emits `<byte count>#<payload>` records. Decoder index `i` always begins the next record, never the middle of a payload.

| Step | Reconstruct | Why |
|---|---|---|
| 1 | `type Codec struct{}`; `Constructor() Codec { return Codec{} }` | Stateless API; value type matches signature |
| 2 | Write `strconv.Itoa(len(str))`, `'#'`, then `str` | Byte length frames unrestricted payload data |
| 3 | Decoder scans `j` from `i` to first `#` | Isolates variable-length decimal prefix |
| 4 | `L := Atoi(s[i:j]); start:=j+1; end:=start+L` | Computes byte boundaries (end exclusive) |
| 5 | `append(result,s[start:end]); i=end` | Extracts exact payload then resumes at next record |

**Trace:** `["a#b",""]` → `3#a#b0#` → record 1 spans bytes `[2:5]`, record 2 `[7:7]`.

**Your personal syntax traps:** `Codec{}` vs `&Codec{}` types; local builder vs persistent struct field; string vs `[]string`; `result = append` vs shadowing with `:=`; `strconv.Itoa` vs `Atoi`; `len` bytes vs `range` runes; `unicode.IsDigit(byte)` type mismatch.

**Boundary drill:** Distinguish `[]string{}` → `""` from `[]string{""}` → `0#`; explain why embedded `#` is safe; trace `"é"` as two UTF-8 bytes; test repeated Encode calls on one Codec.

**Complexity:** linear in total framed bytes for both methods; only valid Encode-generated data is decoded (malformed-input validation would be a separate contract).

Confidence: `—` · Last cold recall: `—` · Next recall: `—`
