# Grid/matrix mastery — indexes before algorithms

[Mastery system](../MASTERY-SYSTEM.md) · [Queue](QUEUE-MASTERY.md) · [BFS](BFS-MASTERY.md) · [Original #1293](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md)

## Quick recall — 60 seconds

A Go grid is commonly `[][]int`; **row first, column second**. `grid[r][c]` is valid only when both indices are in bounds. A `[][]T` can be **ragged** (rows have different lengths).

```go
dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
for _, d := range dirs {
    nr, nc := r+d[0], c+d[1]
    if nr < 0 || nr >= len(grid) { continue }
    if nc < 0 || nc >= len(grid[nr]) { continue }
    // Safe to use grid[nr][nc]
}
```

**Invariant:** check row bounds **before** indexing `grid[nr]`. Check `len(grid[nr])`, not blindly `len(grid[0])`, unless the problem guarantees a rectangle.

## A–B. Identity, creation, defaults

A grid is a **representation**, not an algorithm. `var grid [][]int` is nil with zero rows; `len(grid)==0`. `make([][]int, rows)` creates nil inner rows; initialize each row separately. `make([]int, cols)` produces zero-valued int elements.

```go
grid := make([][]int, 3)
for r := range grid {
    grid[r] = make([]int, 4)
}
```

`[3][4]int` is a fixed-size array copied by value; `[][]int` is a slice of slice headers, with shared backing arrays possible.

## C. Lifecycle and operations

| Operation | Go form | Cost |
|---|---|---|
| Access/update | `grid[r][c]`, `grid[r][c]=x` | O(1) after bounds |
| Traverse | `for r := range grid { for c := range grid[r] {…} }` | O(total cells) |
| Rows / width | `len(grid)`, `len(grid[r])` | O(1) |
| Append a row | `append(grid, row)` | Amortized O(1), copies row header |
| Find a value | scan every cell | O(total cells) |
| Neighbor generation | up/down/left/right bounds | O(1) candidates |
| Deep clone | allocate each row and `copy` it | O(total cells) |
| Sort | per-row via `slices.Sort`, if appropriate | O(sum rowLen log rowLen) |
| Delete row | `append(grid[:r], grid[r+1:]...)` | O(rows), changes layout |

A grid is not a queue or a graph algorithm. Its cells can *represent* vertices with edges to valid adjacent cells. An adjacency list is optional: neighbors may be **generated from coordinates**.

## D–E. Memory, invariants, traps

- Row slicing and assignment can share arrays; `copy(grid)` only copies **row headers**, not their elements. Deep copy each row to isolate mutation.
- `nil` grid, empty rows, ragged grids, negative/out-of-range indices, and row/column reversal are essential edge cases.
- `clear(grid[r])` zeroes elements without changing row length; dropping references does not synchronously force garbage collection.
- Concurrent mutation of the same cells is not safe without synchronization; ordinary LeetCode grid problems are single-threaded.
- Marking `grid[r][c]` as visited **mutates the input**. Choose whether the input contract allows mutation or allocate a parallel `visited`.

## F–G. Code, trace, LeetCode connections

[**Practice 04: explicit neighbors**](practice-04-grid-neighbors/README.md) generates valid four-direction neighbors for rectangular **and ragged** grids. [Practice 05](practice-05-grid-bfs/README.md) uses it conceptually to walk a grid; [the original #1293](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md) adds obstacle-removal state.

**Trace:** For a 2×3 grid, `(0,1)` has left `(0,0)`, right `(0,2)`, down `(1,1)`; up is out of bounds. Grid BFS is ordinary graph BFS with implicit neighbors rather than `adj[node]`.

Problems: [#36 Sudoku](../../36.%20Valid%20Sudoku/README.md) (3×3 box indexing), [#200 Islands](../../200.%20Number%20of%20Islands/README.md), [#1293](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md).

## H–I. Design and optimization (later)

Possible APIs: `InBounds(Cell)`, `Neighbors(Cell)` and `CloneGrid`, but simple local functions are preferred in LeetCode. Avoid costly copies unless needed; profile memory only once correct.

## J. Cold recall

1. Draw row and column indices and write `grid[r][c]` access.
2. Write four-direction neighbor generation **without notes**.
3. Explain why checking `nc < len(grid[0])` can fail on ragged input.
4. Allocate a 3D `visited[r][c][remaining]` table for #1293.
5. Transfer: solve ordinary grid BFS first, then the existing #1293.

Tests: [neighbors](practice-04-grid-neighbors/main_test.go) · [plain-grid BFS](practice-05-grid-bfs/main_test.go).

Confidence: — · Last independent recall: — · Next review: —
