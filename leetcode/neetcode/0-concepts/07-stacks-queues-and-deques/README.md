# Stacks, queues, deques, and element lifetime

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A stack removes the latest item; a queue removes the oldest. A deque supports both ends. A slice can model a stack with append and truncation. A queue often uses a head index to avoid repeatedly shifting elements. Both approaches still need empty checks and a policy for retaining consumed backing storage.

For monotonic structures, separately track what ordering the values maintain and what age/index information determines expiration. Those are generic invariants; the direction and removal rule must come from the particular task. A slice itself provides no synchronization or fixed-size bound. For pointer entries, stale backing-array slots may keep objects reachable after logical removal.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
stack := []int{2, 7}
stack = append(stack, 9)
last := stack[len(stack)-1]
stack = stack[:len(stack)-1]
queue := []int{3, 5, 8}
head := 0
first := queue[head]; head++
fmt.Println(last, stack, first, queue[head:])
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **9 [2 7] 3 [5 8]**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Implement a tiny deque interface on paper and specify what each operation does when empty.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Popping an empty slice; confusing FIFO and LIFO; keeping pointers in consumed slots; treating every nested pop loop as quadratic without counting total removals.

## Questions that use this building block

- [#20 Valid Parentheses](../../20.%20Valid%20Parentheses/README.md)
- [#155 Min Stack](../../155.%20Min%20Stack/README.md)
- [#150 Evaluate Reverse Polish Notation](../../150.%20Evaluate%20Reverse%20Polish%20Notation/README.md)
- [#22 Generate Parentheses](../../22.%20Generate%20Parentheses/README.md)
- [#739 Daily Temperatures](../../739.%20Daily%20Temperatures/README.md)
- [#853 Car Fleet](../../853.%20Car%20Fleet/README.md)
- [#84 Largest Rectangle in Histogram](../../84.%20Largest%20Rectangle%20in%20Histogram/README.md)
- [#102 Binary Tree Level Order Traversal](../../102.%20Binary%20Tree%20Level%20Order%20Traversal/README.md)
- [#199 Binary Tree Right Side View](../../199.%20Binary%20Tree%20Right%20Side%20View/README.md)
- [#200 Number of Islands](../../200.%20Number%20of%20Islands/README.md)
- [#133 Clone Graph](../../133.%20Clone%20Graph/README.md)
- [#695 Max Area of Island](../../695.%20Max%20Area%20of%20Island/README.md)
- [#417 Pacific Atlantic Water Flow](../../417.%20Pacific%20Atlantic%20Water%20Flow/README.md)
- [#130 Surrounded Regions](../../130.%20Surrounded%20Regions/README.md)
- [#994 Rotting Oranges](../../994.%20Rotting%20Oranges/README.md)
- [#286 Walls and Gates](../../286.%20Walls%20and%20Gates/README.md)
- [#207 Course Schedule](../../207.%20Course%20Schedule/README.md)
- [#210 Course Schedule II](../../210.%20Course%20Schedule%20II/README.md)
- [#684 Redundant Connection](../../684.%20Redundant%20Connection/README.md)
- [#323 Number of Connected Components in an Undirected Graph](../../323.%20Number%20of%20Connected%20Components%20in%20an%20Undirected%20Graph/README.md)
- [#261 Graph Valid Tree](../../261.%20Graph%20Valid%20Tree/README.md)
- [#127 Word Ladder](../../127.%20Word%20Ladder/README.md)
- [#332 Reconstruct Itinerary](../../332.%20Reconstruct%20Itinerary/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../0-problem-solving-building-blocks/12-circular-buffers-and-deques/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
