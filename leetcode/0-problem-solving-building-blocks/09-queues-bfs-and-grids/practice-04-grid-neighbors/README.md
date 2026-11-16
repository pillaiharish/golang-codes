# Practice 04 — grid neighbor indexing (Go)

[← Grid mastery](../GRID-MASTERY.md) · [Next: plain-grid BFS](../practice-05-grid-bfs/README.md) · [Existing #1293](../../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md)

## Problem

Given a `[][]int` and coordinate `(row,col)`, produce valid four-direction neighboring cells. Do not assume every row has the same width.

**Trap:** check the next row first; `grid[nr]` would panic if `nr` is invalid. Then check `nc < len(grid[nr])`.

[main.go](main.go) demonstrates `type Cell struct{ Row, Col int }`, a 4-element directions array, explicit bounds, and a `[]Cell` result.

### Trace

For `[[0,0,0],[0,0]]` and `(0,1)`, up is invalid; down `(1,1)`, left `(0,0)`, right `(0,2)` are valid.

Complexity: exactly 4 candidates → O(1) time and O(1) extra result size per query. No grid mutation; order follows the direction array.

## Tests and run

```sh
GO111MODULE=off go run .
GO111MODULE=off go test -race -count=1 -v .
GO111MODULE=off go vet .
```

[Tests](main_test.go) cover center, boundary, ragged grid, invalid coordinates and empty grid.

## Cold recall

Without notes, write `Neighbors` with row bounds **before** column bounds. Then explain why grid neighbors do not require a stored graph adjacency list.
