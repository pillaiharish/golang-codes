# Weights, path costs, and spanning structures

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A weighted edge assigns a cost to a transition. A path's cost may be a sum, a maximum encountered value, or another task-specific aggregate; identify it before choosing an algorithm. Fewest edges is not the same as minimum weighted cost. A spanning tree connects every vertex using V-1 edges and has a different objective from a shortest path between two vertices.

Relaxation means comparing a proposed route against recorded state and replacing it if better. The validity of that comparison depends on the state definition. If routes are limited by stops, a vertex alone may not capture all relevant information. Negative weights invalidate some greedy shortest-path assumptions. Infinity sentinels require care before arithmetic to avoid overflow.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
type Edge struct { To, Cost int }
adj := [][]Edge{{{To: 1, Cost: 6}}, {}}
routeCost := 4
candidate := routeCost + adj[0][0].Cost
fmt.Println(adj[0][0].To, candidate)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **1 10**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Draw two routes: fewer edges but higher cost, and more edges but lower cost. Add a stop bound and explain how feasibility changes.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Confusing stops with edges; using BFS distance for arbitrary weights; adding to an overflow-prone infinity sentinel; confusing shortest paths and minimum spanning trees.

## Questions that use this building block

- [#332 Reconstruct Itinerary](../../332.%20Reconstruct%20Itinerary/README.md)
- [#1584 Min Cost to Connect All Points](../../1584.%20Min%20Cost%20to%20Connect%20All%20Points/README.md)
- [#743 Network Delay Time](../../743.%20Network%20Delay%20Time/README.md)
- [#778 Swim in Rising Water](../../778.%20Swim%20in%20Rising%20Water/README.md)
- [#269 Alien Dictionary](../../269.%20Alien%20Dictionary/README.md)
- [#787 Cheapest Flights Within K Stops](../../787.%20Cheapest%20Flights%20Within%20K%20Stops/README.md)

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
