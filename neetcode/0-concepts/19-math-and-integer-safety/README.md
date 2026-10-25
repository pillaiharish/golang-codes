# Arithmetic, geometry, and integer boundaries

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

Go integer division truncates toward zero. Remainder follows the dividend's sign; it is not automatically a nonnegative modulo index. An int is architecture-sized, while int32 fixes a signed width. Overflow does not automatically become an error. Multiplication may overflow before a later conversion, so convert operands first when wider arithmetic is permitted by the task.

For geometry, identify whether distance is Euclidean, squared Euclidean, or Manhattan. Squaring or taking an absolute value can overflow at extreme values. Floating-point equality needs a tolerance when the contract expects approximate output. Decimal digit operations work on representation and should not silently convert a whole huge numeric string to a machine integer.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
fmt.Println(-7/3, -7%3)
x, y := int64(3), int64(4)
fmt.Println(x*x+y*y)
fmt.Println(((-1 % 5) + 5) % 5)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **-2 -1; 25; 4**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Before computing a product, derive the largest possible magnitude from the input bounds. State whether int64 is permitted.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Converting after overflow; assuming abs(MinInt) fits; confusing squared distance and distance; using float equality for approximate results.

## Questions that use this building block

- [#150 Evaluate Reverse Polish Notation](../../150.%20Evaluate%20Reverse%20Polish%20Notation/README.md)
- [#853 Car Fleet](../../853.%20Car%20Fleet/README.md)
- [#4 Median of Two Sorted Arrays](../../4.%20Median%20of%20Two%20Sorted%20Arrays/README.md)
- [#39 Combination Sum](../../39.%20Combination%20Sum/README.md)
- [#1584 Min Cost to Connect All Points](../../1584.%20Min%20Cost%20to%20Connect%20All%20Points/README.md)
- [#1899 Merge Triplets to Form Target Triplet](../../1899.%20Merge%20Triplets%20to%20Form%20Target%20Triplet/README.md)
- [#48 Rotate Image](../../48.%20Rotate%20Image/README.md)
- [#54 Spiral Matrix](../../54.%20Spiral%20Matrix/README.md)
- [#73 Set Matrix Zeroes](../../73.%20Set%20Matrix%20Zeroes/README.md)
- [#202 Happy Number](../../202.%20Happy%20Number/README.md)
- [#66 Plus One](../../66.%20Plus%20One/README.md)
- [#50 Pow(x, n)](../../50.%20Pow%28x%2C%20n%29/README.md)
- [#43 Multiply Strings](../../43.%20Multiply%20Strings/README.md)
- [#2013 Detect Squares](../../2013.%20Detect%20Squares/README.md)
- [#136 Single Number](../../136.%20Single%20Number/README.md)
- [#191 Number of 1 Bits](../../191.%20Number%20of%201%20Bits/README.md)
- [#338 Counting Bits](../../338.%20Counting%20Bits/README.md)
- [#190 Reverse Bits](../../190.%20Reverse%20Bits/README.md)
- [#268 Missing Number](../../268.%20Missing%20Number/README.md)
- [#371 Sum of Two Integers](../../371.%20Sum%20of%20Two%20Integers/README.md)
- [#7 Reverse Integer](../../7.%20Reverse%20Integer/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../leetcode/0-problem-solving-building-blocks/15-subarrays-running-sums-and-modular-arithmetic/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
