# #1293 — Shortest Path in a Grid with Obstacles Elimination: cold recall

[Full guide](README.md) · [Primary solution](main.go) · [Variants](variants.go) · [Tests](main_test.go) · [LeetCode index](../README.md)

**Trigger:** Shortest unweighted grid path, but you have a limited obstacle-removal resource.

**State:** `(r, c, remaining)`—same cell with different budgets is **not** the same state.

**Invariant:** Every state in the current BFS queue is exactly `steps` moves from the start. Exact states are marked when enqueued, so the first destination has minimum distance.

| Step | Recall | Why |
|---|---|---|
| 1 | `1×1` → `0`; if `k >= m+n-3` → `m+n-2` | Manhattan lower bound attainable with enough removals |
| 2 | 3D `visited[r][c][rem]` | Preserve resource state |
| 3 | Enqueue `(0,0,k)`, mark it | BFS begins at distance zero |
| 4 | Process one complete level | Count one move per level |
| 5 | For 4 neighbors: bounds, obstacle cost, visited, enqueue | Only valid unseen states |
| 6 | Return destination distance or `-1` if exhausted | BFS shortest-path guarantee |

**Trace:** In `[[0,1,0],[0,0,1],[1,1,0]]`, `k=1`, `(1,1)` is reached with both `0` and `1` elimination left. Only the second budget can finish in four moves.

**Alternatives:** DFS is exponentially slow (for learning). The optional 2D `maxRemaining` dominance table saves tracking memory, but **queue memory can still be O(mn(k+1))**.

**Drill:** From a blank editor, code exact-state BFS. Then explain why `visited[r][c]` alone can fail, and why the first destination is shortest.

Confidence `—` · Last cold recall `—` · Next recall `—`
