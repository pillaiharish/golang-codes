# Pointer ranges and window bookkeeping

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

Two indices can describe a pair, an unprocessed interval, or a moving contiguous window. Their meaning must stay explicit. A half-open window [left,right) has width right-left; an inclusive window [left,right] has width right-left+1. Mixing these conventions causes boundary errors.

A moving window may maintain counts, a sum, or other state as its boundaries change. Each added value must be reflected once and each removed value undone once. Whether a boundary can move monotonically depends on the problem's assumptions; do not assume a window method applies just because the input is an array. Learn range mechanics here, then derive the validity condition yourself.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
x := []int{4, 6, 8, 10}
left, right := 1, 3
fmt.Println(x[left:right], right-left)
left++
fmt.Println(x[left:right], right-left)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **[6 8] 2; then [8] 1**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Maintain a sum while extending and shrinking a chosen window, and check it against direct summation after each move.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Inclusive/exclusive confusion; updating indices but not state; shrinking beyond empty; assuming negatives preserve a sum monotonicity rule.

## Questions that use this building block

- [#125 Valid Palindrome](../../125.%20Valid%20Palindrome/README.md)
- [#167 Two Sum II - Input Array Is Sorted](../../167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/README.md)
- [#15 3Sum](../../15.%203Sum/README.md)
- [#11 Container With Most Water](../../11.%20Container%20With%20Most%20Water/README.md)
- [#42 Trapping Rain Water](../../42.%20Trapping%20Rain%20Water/README.md)
- [#121 Best Time to Buy and Sell Stock](../../121.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock/README.md)
- [#3 Longest Substring Without Repeating Characters](../../3.%20Longest%20Substring%20Without%20Repeating%20Characters/README.md)
- [#424 Longest Repeating Character Replacement](../../424.%20Longest%20Repeating%20Character%20Replacement/README.md)
- [#567 Permutation in String](../../567.%20Permutation%20in%20String/README.md)
- [#76 Minimum Window Substring](../../76.%20Minimum%20Window%20Substring/README.md)
- [#239 Sliding Window Maximum](../../239.%20Sliding%20Window%20Maximum/README.md)
- [#5 Longest Palindromic Substring](../../5.%20Longest%20Palindromic%20Substring/README.md)
- [#647 Palindromic Substrings](../../647.%20Palindromic%20Substrings/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../leetcode/0-problem-solving-building-blocks/06-two-pointers/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
