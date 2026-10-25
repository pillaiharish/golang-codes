# Binary search bounds and monotonic predicates

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

Binary search requires an ordered domain or a predicate that changes truth value at most once. State what every remaining candidate could mean before selecting a midpoint. A half-open [lo,hi) interval can become empty when lo equals hi; a closed interval requires different updates and termination.

Use lo+(hi-lo)/2 when endpoints are nonnegative and the difference fits. Every update must strictly shrink the unresolved range. Equality may mean return immediately or continue toward a first/last boundary depending on the contract. Binary search over a possible answer range also needs a justified feasibility predicate, which you must derive for the problem rather than assume.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
lo, hi := 2, 10
mid := lo + (hi-lo)/2
fmt.Println(lo, mid, hi)
hi = mid
fmt.Println("remaining width", hi-lo)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **2 6 10; remaining width 4**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Trace both [lo,hi) and [lo,hi] conventions on ranges of length zero, one, and two without writing a search solution.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

No-progress updates; mixing boundary conventions; assuming rotated input is globally sorted; choosing a non-monotonic predicate.

## Questions that use this building block

- [#704 Binary Search](../../704.%20Binary%20Search/README.md)
- [#74 Search a 2D Matrix](../../74.%20Search%20a%202D%20Matrix/README.md)
- [#875 Koko Eating Bananas](../../875.%20Koko%20Eating%20Bananas/README.md)
- [#153 Find Minimum in Rotated Sorted Array](../../153.%20Find%20Minimum%20in%20Rotated%20Sorted%20Array/README.md)
- [#33 Search in Rotated Sorted Array](../../33.%20Search%20in%20Rotated%20Sorted%20Array/README.md)
- [#981 Time Based Key-Value Store](../../981.%20Time%20Based%20Key-Value%20Store/README.md)
- [#4 Median of Two Sorted Arrays](../../4.%20Median%20of%20Two%20Sorted%20Arrays/README.md)
- [#300 Longest Increasing Subsequence](../../300.%20Longest%20Increasing%20Subsequence/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../leetcode/0-problem-solving-building-blocks/07-binary-search-and-invariants/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
