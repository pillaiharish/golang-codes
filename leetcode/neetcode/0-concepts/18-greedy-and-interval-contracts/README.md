# Greedy reasoning and interval boundaries

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A greedy method commits to a local choice without revisiting it. Correctness needs a reason that an optimal answer can retain that choice, often an exchange argument or a dominance invariant. A successful example is not a proof, and sorting alone does not make a greedy choice safe.

Intervals require an endpoint convention. Closed [a,b] includes both endpoints; half-open [a,b) excludes b. Two meetings may share an endpoint without overlapping even though closed-interval union tasks merge touching intervals. Define whether you count duration, integer points, or number of intervals. Keep query output order distinct from any processing order.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
aStart, aEnd := 1, 4
bStart, bEnd := 4, 7
closedOverlap := aStart <= bEnd && bStart <= aEnd
halfOpenOverlap := aStart < bEnd && bStart < aEnd
fmt.Println(closedOverlap, halfOpenOverlap)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **true false**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Give a counterexample to an unjustified local-choice rule, then explicitly label an interval diagram as closed or half-open.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Using the same overlap inequality for every task; confusing integer count with coordinate difference; assuming a greedy strategy from the category name alone.

## Questions that use this building block

- [#53 Maximum Subarray](../../53.%20Maximum%20Subarray/README.md)
- [#55 Jump Game](../../55.%20Jump%20Game/README.md)
- [#45 Jump Game II](../../45.%20Jump%20Game%20II/README.md)
- [#134 Gas Station](../../134.%20Gas%20Station/README.md)
- [#846 Hand of Straights](../../846.%20Hand%20of%20Straights/README.md)
- [#1899 Merge Triplets to Form Target Triplet](../../1899.%20Merge%20Triplets%20to%20Form%20Target%20Triplet/README.md)
- [#763 Partition Labels](../../763.%20Partition%20Labels/README.md)
- [#678 Valid Parenthesis String](../../678.%20Valid%20Parenthesis%20String/README.md)
- [#57 Insert Interval](../../57.%20Insert%20Interval/README.md)
- [#56 Merge Intervals](../../56.%20Merge%20Intervals/README.md)
- [#435 Non-overlapping Intervals](../../435.%20Non-overlapping%20Intervals/README.md)
- [#252 Meeting Rooms](../../252.%20Meeting%20Rooms/README.md)
- [#253 Meeting Rooms II](../../253.%20Meeting%20Rooms%20II/README.md)
- [#1851 Minimum Interval to Include Each Query](../../1851.%20Minimum%20Interval%20to%20Include%20Each%20Query/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../0-problem-solving-building-blocks/14-greedy-intervals-and-proof/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
