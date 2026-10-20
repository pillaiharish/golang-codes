# Prefix, suffix, and aggregate bookkeeping

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A prefix summarizes values before or through a boundary; a suffix summarizes values after or from a boundary. State precisely whether that boundary element is included. A prefix-sum array of length n+1 can represent the sum before each of n+1 boundaries, making the empty prefix explicit. Aggregate operations have different properties: sum can be undone by subtraction, but maximum cannot generally be undone by removing the latest value.

Products introduce zeros and sign changes, so division is not a universally valid inverse. Overflow must be checked against the problem's numeric promises. A running aggregate stores only the current summary, while an array of summaries supports queries at earlier boundaries. This lesson constructs a small sum table; derive your own required prefix/suffix state for a problem rather than treating every aggregate as interchangeable.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
values := []int{4, 1, 7}
prefix := make([]int, len(values)+1)
for i, v := range values { prefix[i+1] = prefix[i] + v }
fmt.Println(prefix)
fmt.Println(prefix[3]-prefix[1])
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **[0 4 5 12]; 8**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Mark the empty prefix and the boundary before each element. Explain why subtracting two boundary sums gives a half-open range total.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Including a boundary element twice; using an inverse that does not exist; product division through zero; overflowing before storage.

## Questions that use this building block

- [#238 Product of Array Except Self](../../238.%20Product%20of%20Array%20Except%20Self/README.md)
- [#42 Trapping Rain Water](../../42.%20Trapping%20Rain%20Water/README.md)
- [#152 Maximum Product Subarray](../../152.%20Maximum%20Product%20Subarray/README.md)
- [#53 Maximum Subarray](../../53.%20Maximum%20Subarray/README.md)
- [#134 Gas Station](../../134.%20Gas%20Station/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../0-problem-solving-building-blocks/15-subarrays-running-sums-and-modular-arithmetic/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
