# Practice 02 — Breadth-first search on a graph

[← Lesson 09](../README.md) · [Your complete BFS Go code](main.go) · [Tests](main_test.go) · [Previous: queue](../practice-01-dynamic-queue/README.md) · [Apply to #1293](../../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md)

## Recognition and intuition

BFS explores an **unweighted graph level by level** from a starting node. Its central ingredients are a FIFO queue (from practice 01) and a visited set.

Your `BFS(graph map[int][]int, start int)` implementation is preserved as a runnable program. This version **prints traversal order**; it does not yet return distances. Nodes not reachable from the chosen start are not processed.

## Data structures and invariant

- `map[int][]int` is an **adjacency list**. Each key is a node and its slice lists neighbors in a chosen order.
- `queue := []int{start}`, then `queue[0]` plus `queue = queue[1:]`, implements FIFO removal; `append` adds new neighbors.
- `visited := make(map[int]bool)` acts as a set. Reading an absent key returns `false`.
- Mark **when enqueuing**, not when dequeuing, to prevent cycles and multiple incoming edges from enqueuing a node repeatedly.

**Invariant:** Each reachable node is enqueued at most once, and nodes at distance d are processed before nodes at distance d+1. The exact order **within a layer** depends on the adjacency-list order.

## Trace of your graph

```text
0 --- 1 --- 2
|     |
3 --- 4
```

Starting with `queue=[0]` and `visited={0}`:

| Dequeued | Newly enqueued | Queue afterwards | Output so far |
|---|---|---|---|
| 0 | 1, 3 | `[1 3]` | `0` |
| 1 | 2, 4 | `[3 2 4]` | `0 1` |
| 3 | — (already visited) | `[2 4]` | `0 1 3` |
| 2 | — | `[4]` | `0 1 3 2` |
| 4 | — | `[]` | `0 1 3 2 4` |

**Levels:** `{0}` (distance 0), `{1,3}` (distance 1), `{2,4}` (distance 2).

Printed output: `BFS Traversal: 0 1 3 2 4`.

## Why BFS works, and where it doesn't

A FIFO queue processes all nodes one edge away before nodes two edges away. Therefore BFS can find **shortest paths measured in number of edges**, when each edge has equal cost. A stack would create a depth-first traversal that doesn't guarantee shortest paths.

For V reachable nodes and E outgoing edges, time is O(V+E) and visited plus live queue uses O(V) space. The sliced queue may retain backing memory, as discussed in the queue practice.

## Practice

1. Trace queue and visited after each pop, without running Go.
2. Swap the order of `graph[0]`'s neighbors and predict the new traversal.
3. Add a cycle or duplicate neighbor; explain why marking visited on **enqueue** matters.
4. Add a disconnected component; why isn't it printed?
5. If a function needs the shortest **distance**, where should a `steps` or level boundary be added?

```sh
cd leetcode/0-problem-solving-building-blocks/09-queues-bfs-and-grids/practice-02-graph-bfs
GO111MODULE=off go run main.go
GO111MODULE=off go test -race -count=1 -v .
```

## Direct link to LeetCode #1293

Your graph BFS uses `visited[node]` and `queue []int`. In [#1293](../../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md), each state is **`(row,col,remainingEliminations)`**, neighbors are four possible grid moves, and visited must track **all three state components**. Visiting cell `(r,c)` with zero budget does **not** make visiting it with one budget redundant. Process one full queue layer before increasing `steps` to compute the shortest route.
