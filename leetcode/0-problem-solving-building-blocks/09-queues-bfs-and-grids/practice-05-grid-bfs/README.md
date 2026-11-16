# Practice 05 — shortest path on a plain grid

[← BFS mastery](../BFS-MASTERY.md) · [Queue V2](../practice-03-head-index-queue/README.md) · [Grid neighbors](../practice-04-grid-neighbors/README.md) · [Existing #1293](../../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md)

## Problem

Given a grid of `0` (passable) and `1` (blocked), return the fewest **unit-cost** four-direction moves from `start` to `goal`. Return `-1` if invalid or unreachable. This is **not** a second implementation of #1293: **no obstacle elimination** here.

## Progressive solution

1. Convert a cell to `Cell{Row,Col}`.
2. Validate start and goal before allocating `visited`.
3. Build `visited[r]` with `len(grid[r])` to handle ragged rows.
4. Enqueue the start, mark visited when enqueued.
5. Freeze `levelEnd := len(queue)`; process exactly that distance layer.
6. Generate up to four valid, passable neighbors; return `steps+1` when discovering goal.
7. Exhaustion means unreachable.

[main.go](main.go) uses `head` instead of repeated `queue=queue[1:]`. Complexity: O(V+E) with V = traversable cells and E ≤ 4V; O(total cells) visited/queue storage.

**Invariant:** `visited[r][c]` indicates cell already enqueued; during each outer loop, every queued state in that layer has the same distance. Marking on enqueue prevents cycles.

### Trace

```text
0 1 0
0 0 0
1 1 0
```

`(0,0) -> (1,0) -> (1,1) -> (1,2) -> (2,2)` takes 4 moves. With obstacles blocking every path, return -1.

## Tests and run

```sh
GO111MODULE=off go run .
GO111MODULE=off go test -race -count=1 -v .
GO111MODULE=off go vet .
```

[Tests](main_test.go) cover reachable, blocked, start=goal, blocked start, invalid goal, empty grid, and ragged grid.

## Transfer to existing #1293

This visited representation is correct **only when the cell alone describes future moves**. In #1293 the next legal moves depend on remaining eliminations, so a cell must be distinguished by `(row,col,remaining)`. Revisit the [existing #1293 README](../../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md), not a cloned solution.

## Cold recall

Write the entire solution from a blank editor; explain row/column checks, `levelEnd`, mark-on-enqueue and `steps+1`. Then attempt #1293 independently.
