# Dependency order and disjoint-set concepts

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A dependency graph asks which events must happen before others. Indegree counts incoming edges. A topological order exists only for a directed acyclic graph; multiple orders may be valid. An undirected connectivity problem is different from directed reachability, so do not interchange their cycle rules.

A disjoint-set representation partitions vertices into groups. A representative identifies a group, not necessarily a meaningful original vertex. Find asks for the representative; union joins groups. Path compression and union-by-size/rank are optimizations whose invariants should be learned independently. This lesson initializes state only: derive the actual union/find or scheduling steps yourself.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
parent := make([]int, 4)
size := make([]int, 4)
for i := range parent { parent[i] = i; size[i] = 1 }
fmt.Println(parent, size)
indegree := make([]int, 3)
indegree[2]++
fmt.Println(indegree)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **[0 1 2 3] [1 1 1 1]; [0 0 1]**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Draw three separate components, join two on paper, and update representatives and sizes without writing a complete algorithm.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Applying undirected cycle tests to directed graphs; double-counting edges; assuming representative values encode meaningful ordering; mutating size on non-roots.

## Questions that use this building block

- [#207 Course Schedule](../../207.%20Course%20Schedule/README.md)
- [#210 Course Schedule II](../../210.%20Course%20Schedule%20II/README.md)
- [#684 Redundant Connection](../../684.%20Redundant%20Connection/README.md)
- [#323 Number of Connected Components in an Undirected Graph](../../323.%20Number%20of%20Connected%20Components%20in%20an%20Undirected%20Graph/README.md)
- [#261 Graph Valid Tree](../../261.%20Graph%20Valid%20Tree/README.md)
- [#1584 Min Cost to Connect All Points](../../1584.%20Min%20Cost%20to%20Connect%20All%20Points/README.md)
- [#269 Alien Dictionary](../../269.%20Alien%20Dictionary/README.md)

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
