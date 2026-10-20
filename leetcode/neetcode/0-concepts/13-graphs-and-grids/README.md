# Graph representations, coordinates, and visited state

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A graph has vertices and edges. Directed edges have a source and destination; an undirected edge permits both directions. An adjacency list stores neighbors per vertex. A grid is an implicit graph whose allowed moves define edges. Four-neighbor and eight-neighbor connectivity differ, as do blocked cells and traversable cells.

A visited set records a clearly stated event such as discovered or fully processed. Marking only after removal from a queue can enqueue the same vertex repeatedly. A graph may have isolated vertices even if no edges mention them. Distinguish vertex count from edge count in O(V+E) traversal bounds. Copying adjacency slices is not automatically a deep copy of a graph's nodes.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
adj := make([][]int, 3)
adj[0] = append(adj[0], 2)
adj[2] = append(adj[2], 0)
fmt.Println(adj[0], adj[1], adj[2])
cell := [2]int{1, 2}
seen := map[[2]int]bool{cell: true}
fmt.Println(seen[cell])
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **[2] [] [0]; true**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Write valid neighbors for a corner, edge, and middle cell. Include an isolated vertex in a separate adjacency list.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Accidentally reversing directed edges; diagonal moves not allowed by the contract; missing isolated vertices; reusing a row slice for every grid row.

## Questions that use this building block

- [#51 N-Queens](../../51.%20N-Queens/README.md)
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
- [#1584 Min Cost to Connect All Points](../../1584.%20Min%20Cost%20to%20Connect%20All%20Points/README.md)
- [#743 Network Delay Time](../../743.%20Network%20Delay%20Time/README.md)
- [#778 Swim in Rising Water](../../778.%20Swim%20in%20Rising%20Water/README.md)
- [#269 Alien Dictionary](../../269.%20Alien%20Dictionary/README.md)
- [#787 Cheapest Flights Within K Stops](../../787.%20Cheapest%20Flights%20Within%20K%20Stops/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../0-problem-solving-building-blocks/09-queues-bfs-and-grids/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
