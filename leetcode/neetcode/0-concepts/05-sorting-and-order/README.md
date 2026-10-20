# Sorting, comparators, and deterministic ordering

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

Sorting creates an order that later reasoning may depend on. In Go, sort.Ints changes the input; copy first if mutation is forbidden or original positions matter. sort.Slice takes a less-than comparator. Equal elements must not compare less than themselves, and comparisons must be consistent.

For records, define primary and tie-breaking fields explicitly. Stable sorting preserves original order for equal keys only when the selected API promises stability. Lexicographic order compares the first differing component, with a strict prefix preceding the longer sequence. Sorting may destroy original indices, which must be retained separately if the contract asks for positions.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
x := []int{8, 3, 5}
y := append([]int(nil), x...)
sort.Ints(y)
fmt.Println(x, y)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **[8 3 5] [3 5 8]**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Sort records by a score and then name. Verify that equal records return false from the comparator in both directions.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Using <= in Less; forgetting mutation; subtracting large integers in a comparator and overflowing; losing original positions.

## Questions that use this building block

- [#167 Two Sum II - Input Array Is Sorted](../../167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/README.md)
- [#15 3Sum](../../15.%203Sum/README.md)
- [#853 Car Fleet](../../853.%20Car%20Fleet/README.md)
- [#703 Kth Largest Element in a Stream](../../703.%20Kth%20Largest%20Element%20in%20a%20Stream/README.md)
- [#1046 Last Stone Weight](../../1046.%20Last%20Stone%20Weight/README.md)
- [#973 K Closest Points to Origin](../../973.%20K%20Closest%20Points%20to%20Origin/README.md)
- [#215 Kth Largest Element in an Array](../../215.%20Kth%20Largest%20Element%20in%20an%20Array/README.md)
- [#621 Task Scheduler](../../621.%20Task%20Scheduler/README.md)
- [#355 Design Twitter](../../355.%20Design%20Twitter/README.md)
- [#295 Find Median from Data Stream](../../295.%20Find%20Median%20from%20Data%20Stream/README.md)
- [#90 Subsets II](../../90.%20Subsets%20II/README.md)
- [#40 Combination Sum II](../../40.%20Combination%20Sum%20II/README.md)
- [#846 Hand of Straights](../../846.%20Hand%20of%20Straights/README.md)
- [#57 Insert Interval](../../57.%20Insert%20Interval/README.md)
- [#56 Merge Intervals](../../56.%20Merge%20Intervals/README.md)
- [#435 Non-overlapping Intervals](../../435.%20Non-overlapping%20Intervals/README.md)
- [#252 Meeting Rooms](../../252.%20Meeting%20Rooms/README.md)
- [#253 Meeting Rooms II](../../253.%20Meeting%20Rooms%20II/README.md)
- [#1851 Minimum Interval to Include Each Query](../../1851.%20Minimum%20Interval%20to%20Include%20Each%20Query/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../0-problem-solving-building-blocks/05-sorting-and-comparators/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
