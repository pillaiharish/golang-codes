# BFS mastery — adjacency list → grid → stateful shortest path

[Mastery system](../MASTERY-SYSTEM.md) · [Queue](QUEUE-MASTERY.md) · [Grid](GRID-MASTERY.md) · [Original #1293](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md)

## Quick recall — 60 seconds

**BFS is an algorithm; queue is its FIFO data structure.** In an unweighted graph, enqueue start at distance zero, mark it visited, and process FIFO. Each neighbor has distance `currentDistance+1`; the first time the goal is reached is shortest.

```go
queue := []int{start}
head := 0
visited := map[int]bool{start: true}
for head < len(queue) {
    node := queue[head]
    head++
    for _, neighbor := range graph[node] {
        if visited[neighbor] { continue }
        visited[neighbor] = true // at enqueue, not dequeue
        queue = append(queue, neighbor)
    }
}
```

**Invariant:** nodes are processed in nondecreasing distance from the start; each visited *state* is discovered at its shortest distance.

## A–C. Identity, representation and operations

BFS isn't a collection; zero value, map-key comparability, insertion/deletion, sorting and copying BFS itself are **N/A**. Its working state is `queue` + `visited` + optional `distance/steps` and a way to enumerate neighbors.

| Representation | Neighbor source | Visited key |
|---|---|---|
| Graph adjacency list | `graph[node]` | `node` |
| 2D unweighted grid | four valid `(r,c)` cells | `(r,c)` or 2D bool |
| Grid with obstacle budget | four cells + budget change | `(r,c,remaining)` |

For unweighted shortest distance, track per-state distance or process one complete BFS level at a time:

```go
steps := 0
for head < len(queue) {
    levelEnd := len(queue) // freeze this distance layer
    for head < levelEnd {
        current := queue[head]
        head++
        // Visit neighbors and append next-level states.
        _ = current
    }
    steps++
}
```

`levelEnd` matters because `append` grows the queue while processing a layer. The exact return point determines whether a destination found as a neighbor is at `steps+1` or a destination dequeued at this level is at `steps`.

## D–E. Memory, invariants and Layer 3 traps

1. Mark an unseen state **before enqueueing**, otherwise cycles and multiple parents cause duplicates.
2. Your visited key must include **everything affecting future legal moves**. For #1293 it is `row,col,remaining`; `visited[row][col]` is wrong.
3. A graph with cycles needs visited tracking; an adjacency-list map returns a nil slice for absent keys, so ranging over it is safe.
4. Graph/BFS order across neighbors can vary; only distances are guaranteed, not one unique traversal sequence (unless neighbor order is fixed).
5. Plain BFS shortest-path guarantee requires **equal nonnegative edge costs**. Use weighted shortest-path algorithms for varying edge weights.
6. Stop when queue exhausted → unreachable; single-cell start==goal → distance 0.
7. Runtime for graph adjacency list: O(V+E); memory O(V). For rectangular obstacle-budget grid: up to O(rows×cols×(k+1)) states; each has ≤4 transitions.
8. Head-index queue keeps previous consumed slots until released; details in [Queue mastery](QUEUE-MASTERY.md).

## F. Implementation progression

1. [Original learner-written adjacency-list BFS](practice-02-graph-bfs/README.md) — FIFO, visited map, graph cycles.
2. [Explicit grid neighbors](practice-04-grid-neighbors/README.md) — adjacency list is generated on demand, not stored.
3. [Plain grid shortest path](practice-05-grid-bfs/README.md) — BFS layers, start/goal validation, mark-on-enqueue.
4. [Existing #1293 solution and variants](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md) — **refer to this; do not duplicate it here**.

## G. Recognition, problems and transfer

- "Fewest steps" + equal-cost moves → BFS.
- "Connected components" → repeated DFS/BFS with visited.
- "All sources start simultaneously" → multi-source BFS.
- "Budget, keys, mask, remaining fuel" → augment state; may need higher-dimensional visited.

Practice: [#102 tree levels](../../102.%20Binary%20Tree%20Level%20Order%20Traversal/README.md), [#200 Islands](../../200.%20Number%20of%20Islands/README.md), [#1293](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md).

## H–I. Design and performance (later)

Design question: choose a queue API or a plain slice; decide whether to return length, path, or traversal. Avoid memory optimization before correctness. State-space pruning (such as best remaining budget per cell) is a later improvement; it must be proved to preserve shortest paths.

## J. Recall sequence

1. Empty editor: code BFS on `map[int][]int` with a cycle.
2. Swap neighbor generation to four grid directions, without constructing an adjacency list.
3. Add level counting for shortest distance and prove why BFS layers are shortest.
4. Change visited from 2D to 3D for a resource-constrained grid.
5. Re-solve #1293 **in its existing folder** without reading the solution, and record the gap in [revision log](../../REVISION-LOG.md).

Confidence: — · Last independent recall: — · Next review: —
