# Bit operations, signedness, and masks

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A bit mask selects positions in a fixed-width representation. AND keeps common set bits, OR combines them, XOR marks differences, and AND-NOT clears selected bits. A shift moves bit positions; signed and unsigned interpretation matters, especially for right shifts and conversions. Specify the intended width instead of relying on the host int width.

Bitwise reasoning is about the representation of values, not their printed decimal digits. A loop over exactly 32 bit positions differs from looping until a signed number becomes zero. Current LeetCode signatures may differ from older unsigned versions, so use each folder's verified interface. This lesson demonstrates isolated operations without implementing any bit-manipulation question.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
var x uint32 = 10
mask := uint32(1) << 1
fmt.Printf("%04b %04b %04b\n", x&mask, x|1, x^mask)
fmt.Println(x >> 1)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **0010 1011 1000; 5**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Write the 8-bit patterns for two numbers, apply each operator by hand, then repeat after converting a negative int32 to uint32.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Using architecture-sized width by accident; confusing XOR with exponentiation; signed shifts with unintended stopping conditions; shifting by invalid negative amounts.

## Questions that use this building block

- [#50 Pow(x, n)](../../50.%20Pow%28x%2C%20n%29/README.md)
- [#136 Single Number](../../136.%20Single%20Number/README.md)
- [#191 Number of 1 Bits](../../191.%20Number%20of%201%20Bits/README.md)
- [#338 Counting Bits](../../338.%20Counting%20Bits/README.md)
- [#190 Reverse Bits](../../190.%20Reverse%20Bits/README.md)
- [#268 Missing Number](../../268.%20Missing%20Number/README.md)
- [#371 Sum of Two Integers](../../371.%20Sum%20of%20Two%20Integers/README.md)
- [#7 Reverse Integer](../../7.%20Reverse%20Integer/README.md)

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
