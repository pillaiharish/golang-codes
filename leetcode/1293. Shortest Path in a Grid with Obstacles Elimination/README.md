# 1293 — Shortest Path in a Grid with Obstacles Elimination (Go)

[← LeetCode index](../README.md) · [Primary BFS solution](main.go) · [DFS and memory-optimized BFS variants](variants.go) · [Table-driven tests](main_test.go) · [Cold-recall card](REVISION-CARD.md) · [Official problem](https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/) · [Revision log](../REVISION-LOG.md)

**Pattern:** Shortest path in an unweighted graph → BFS with extra resource state · **Difficulty:** Hard · **Primary solution:** O(m·n·(k+1)) time and space.

## Problem statement (paraphrased)

Given an `m × n` binary grid, move from `(0,0)` to `(m-1,n-1)` in four directions. A `0` is free to enter, and a `1` costs one of your **at most `k` obstacle eliminations**. Return the fewest moves needed; return `-1` if no valid route exists. You count **moves**, not visited cells.

```go
func shortestPath(grid [][]int, k int) int
```

**Input contract:** The grid is nonempty, rectangular, and contains only `0` and `1`; the start and destination are `0`. The official limits are `1 <= m,n <= 40` and `1 <= k <= m*n` (see the [problem page](https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/)). These implementations also handle `k=0`, which is useful for practice tests but lies outside the official limits. None modifies the grid.

### Examples and expected results

| Grid (`;` separates rows) | `k` | Output | Reason |
|---|---:|---:|---|
| `[[0,0,0];[1,1,0];[0,0,0];[0,1,1];[0,0,0]]` | 1 | `6` | One obstacle can be removed on a six-move route |
| `[[0,1,1];[1,1,1];[1,0,0]]` | 1 | `-1` | Every route requires more removals |
| `[[0]]` | 0 | `0` | Already at the destination |
| `[[0,1,0];[0,0,1];[1,1,0]]` | 1 | `4` | The budget at a cell matters, not just its coordinates |

## Prerequisites

1. [Arrays and slices](../0-problem-solving-building-blocks/01-arrays-and-slices/README.md): index `grid[r][c]`, build 2D and 3D slices.
2. [Loops and bounds](../0-problem-solving-building-blocks/02-loops-and-bounds/README.md): four directions, boundary checks, level loops.
3. [Queues and grid BFS](../0-problem-solving-building-blocks/09-queues-bfs-and-grids/README.md): FIFO exploration, mark visited on enqueue, distance layers.
4. [Complexity and constraints](../0-problem-solving-building-blocks/17-complexity-amortization-and-input-contracts/README.md): count states and account for the queue as well as visited memory.

## 1. Recognition trigger

**Fewest steps + four-direction moves + every move costs one step** suggests BFS. Ordinary grid BFS usually marks `visited[row][col]`, but here **the elimination budget is part of your position in the search space**.

At `(r,c)` with `0` eliminations left and at `(r,c)` with `1` left, the future possibilities differ. We therefore search over `(row, col, remaining)` states instead of just cells.

## 2. Start with brute-force DFS, then change the search order

An intuitive first solution enumerates all simple paths: recursively try the four directions, pay for an obstacle, mark the current path, and backtrack. Any shortest valid path can be made simple by removing cycles, which cannot increase its length or obstacle cost. Thus DFS does find the answer if it examines every possibility.

But the number of simple paths may grow exponentially with grid size. The [small-input DFS reference](variants.go) uses a best-so-far answer and a Manhattan-distance pruning bound; it is for understanding and test cross-checks, **not** for LeetCode submission.

**Key improvement:** BFS explores by distance. Level 0 is the start, level 1 is one move away, and so on. The first destination state dequeued has the minimum moves—provided we do not merge states with different budgets.

### Two states at the same cell

Consider:

```text
0 1 0
0 0 1
1 1 0       k = 1
```

The cell `(1,1)` can be reached in two moves:

- Right → down: enter an obstacle first; arrive with `0` removals left.
- Down → right: avoid that obstacle; arrive with `1` removal left.

From `(1,1)`, either direct continuation to the goal crosses another obstacle. The second state can finish; the first cannot. **A single `visited[1][1]` flag would lose useful information.**

## 3. Exact-state invariant and correctness

For the primary algorithm, `visited[r][c][rem]` means that the state `(r,c,rem)` has already been enqueued. The queue at the start of each outer iteration contains states reached in exactly `steps` moves; `nextQueue` contains states reached in `steps+1` moves.

On moving to `(nr,nc)`, subtract `grid[nr][nc]` from the remaining budget. Reject a negative budget. An already visited **identical triple** need not be revisited: BFS first enqueued it using no more moves than any later route to that triple, and its future options are identical. Because every valid move costs one step, the first time BFS removes the destination is the shortest valid route; exhausting the queue proves none exists.

### Manhattan shortcut — when it is safe

Ignoring obstacles, the shortest possible route requires `m+n-2` moves. It visits `m+n-1` cells. The start and end are guaranteed empty, so any such route has at most `m+n-3` obstacle cells between them.

Therefore, after handling a `1×1` grid, if `k >= m+n-3`, **return `m+n-2` immediately**. This is a sufficient condition—not a requirement for every short route. It depends on the input promise that both endpoints are empty.

## 4. Six-step reconstruction recipe

1. Let `m,n` be grid dimensions. If there is one cell, return `0`. Apply the Manhattan shortcut if the budget is large enough.
2. Define a state `(row,col,remaining)` and allocate `visited[m][n][k+1]`.
3. Enqueue `(0,0,k)` and mark it visited. Start `steps = 0`.
4. For every state in the current BFS layer, return `steps` if it is the destination.
5. Try all four neighbors: reject out-of-bounds coordinates, subtract the next cell's obstacle cost, reject negative budget or previously seen exact state, and mark and enqueue new states.
6. Replace the current queue with the next layer and increment `steps`. If no states remain, return `-1`.

**Important:** Mark visited *when enqueuing*, rather than when popping, to avoid duplicate work. Submit [the primary `shortestPath` implementation](main.go), not the DFS.

## 5. Short hand trace

For `[[0,1,0],[0,0,1],[1,1,0]]` with `k=1`:

| Level | Notable states `(row,col,remaining)` | What it teaches |
|---:|---|---|
| 0 | `(0,0,1)` | Start has the full budget |
| 1 | `(0,1,0)`, `(1,0,1)` | Going right spends an elimination; going down does not |
| 2 | `(1,1,0)` and `(1,1,1)` | **Same cell, different valid states** |
| 3 | `(1,2,0)` from `(1,1,1)` | Use the remaining elimination |
| 4 | `(2,2,0)` | Destination reached after 4 moves |

This table shows selected states, not every item in each queue. Since Manhattan distance is 4, no route could use fewer than 4 moves.

## 6. Complexity, limits, and alternatives

| Approach | Time | Tracking memory | Total auxiliary memory | When to use |
|---|---|---|---|---|
| Backtracking DFS | Exponential (loose upper bound O(4^(mn))) | O(mn) path-visited | O(mn) recursion + visited | Understand the search space; small test oracle only |
| **BFS + exact 3D state** | O(mn(k+1)) | O(mn(k+1)) | O(mn(k+1)) | Primary, easy-to-prove submission |
| BFS + best remaining budget per cell | O(mn(k+1)) worst case | **O(mn)** table | **O(mn(k+1))** worst case including queued states | Optional practical memory reduction |

The third method uses `maxRemaining[r][c]`: if a cell has been enqueued previously at an earlier or equal BFS level with **at least** this much budget, a new state is dominated and can be discarded. That earlier state has no greater distance and can follow every route the new state could.

**Memory nuance:** Replacing a 3D `visited` table with a 2D `maxRemaining` table reduces *tracking-table* memory to O(mn). It does **not** prove total auxiliary memory is O(mn): the queue may still hold multiple states per cell, and up to O(mn(k+1)) states can be generated. See [the optional variant](variants.go).

Other traps:

- Plain `visited[r][c]` is insufficient because budgets differ.
- Use the **remaining** budget when marking states, not just the initial `k`.
- Enqueueing a neighbor costs exactly one step; don't increment `steps` once per neighbor.
- Count moves, so one cell has distance `0`.
- The shortcut is valid because the endpoints are empty; don't silently apply it to different contracts.
- `O(mn(k+1))` is a safe bound; after the shortcut, the nontrivial cases satisfy `k < m+n-3`.
- No mutation of the grid is necessary.

## 7. 45-second interview explanation and cold recall

> This is shortest path on an unweighted grid, so I use BFS. But reaching one cell with different obstacle eliminations left is not equivalent. My BFS state is row, column, and remaining budget. I mark each exact state on enqueue and process one distance layer at a time. Moving into a wall uses one elimination. The first destination dequeued gives the fewest moves. There are at most `m*n*(k+1)` states, each with four edges, so time and space are O(mnk). If `k` can remove every obstacle on a Manhattan path, I return its length immediately.

Before opening the Go code, answer:

1. Why is BFS better suited than enumerating all simple paths?
2. Give an example where the same cell should be visited with two remaining budgets.
3. Why do we mark a **triple** as visited at enqueue time?
4. Why does `k >= m+n-3` guarantee the Manhattan distance?
5. Why does the 2D dominance table **not automatically** make total space O(mn)?

**Personal review:** Confidence `— / 5` · Last cold recall `—` · Next recall `—`. Log an actual cold attempt in [REVISION-LOG.md](../REVISION-LOG.md).

## Run locally and submit

```sh
cd 'leetcode/1293. Shortest Path in a Grid with Obstacles Elimination'
GO111MODULE=off go test -count=1 -race -v
GO111MODULE=off go run main.go
```

The demo prints `6`. For LeetCode, submit only `shortestPath` from [main.go](main.go), omitting the demo and `fmt` import. `variants.go` is optional study material and is tested against the primary BFS on small grids.
