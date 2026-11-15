# 09 — Queues, breadth-first search and 2D grids

[← Foundations index](../README.md) · [← Problem index](../../README.md)

## Recognition and intuition
BFS processes nodes in order of distance/layers using a FIFO queue. Matrix traversal adds bounds, visited-state and neighbor direction operations.

## Hands-on prerequisites from the learner: queue → BFS → #1293

The following are **standalone, runnable examples written by the learner**, retained without changing their core algorithms. Study them in sequence; this bridges the gap between a FIFO data structure and the stateful grid search in [#1293](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md).

1. [**Practice 01 — dynamic queue with a Go struct**](practice-01-dynamic-queue/README.md) ([source](practice-01-dynamic-queue/main.go), [tests](practice-01-dynamic-queue/main_test.go)): `Enqueue` calls `append`; `Dequeue` reads `items[0]` and removes it with `items[1:]`. Pointer receivers preserve changes to the original struct.
2. [**Practice 02 — graph BFS with an adjacency list**](practice-02-graph-bfs/README.md) ([source](practice-02-graph-bfs/main.go), [tests](practice-02-graph-bfs/main_test.go)): combine the same FIFO operations with a `visited` map; mark nodes visited **when enqueuing** to avoid cycles and duplicate visits.

### Queue trace: FIFO, not LIFO

| Action | Queue contents | Returned |
|---|---|---|
| Start | `[]` | — |
| `Enqueue(5)`, `Enqueue(3)`, `Enqueue(11)` | `[5 3 11]` | — |
| `Dequeue()` | `[3 11]` | `5, true` |
| Continue dequeuing | `[]` | `3, true`; `11, true` |
| Dequeue when empty | `[]` | `0, false` |

The `bool` is a **success flag**, not a Go `error`. The learner's demo calls the variable `err`, which works syntactically but is less descriptive than `ok`. `q.items = q.items[1:]` is an O(1) slice operation; the old backing array may remain allocated.

### BFS trace: what the queue and visited set do together

For the learner's graph `0:[1,3], 1:[0,2,4], 2:[1], 3:[0,4], 4:[1,3]`:

| Dequeued | Newly discovered | Queue after enqueueing neighbors | Processed so far |
|---|---|---|---|
| Start | — | `[0]` | — |
| 0 | 1, 3 | `[1 3]` | `0` |
| 1 | 2, 4 | `[3 2 4]` | `0 1` |
| 3 | — | `[2 4]` | `0 1 3` |
| 2 | — | `[4]` | `0 1 3 2` |
| 4 | — | `[]` | `0 1 3 2 4` |

The resulting traversal is `0 1 3 2 4`, in BFS layers `{0}`, `{1,3}`, `{2,4}`. A **queue** ensures nodes at a smaller number of edges from the start are processed first. The **visited map** prevents rediscovering those nodes through cycles.

### Transfer: ordinary graph BFS versus obstacle-elimination grid BFS

| Ordinary graph example | LeetCode #1293 |
|---|---|
| Node is an integer | Search state is `(row,col,remainingEliminations)` |
| Neighbors from `graph[curr]` | Generate four valid grid neighbors |
| `visited[node]` | `visited[row][col][remaining]` |
| Print every reachable node | Return shortest steps to bottom-right |
| FIFO guarantees nondecreasing distance | Count one `steps` increment per **complete BFS layer** |

A cell reached with **one** elimination remaining cannot be treated as the same state as that cell with **zero** remaining. This distinction is the key additional concept in #1293.

Run the two packages separately; each contains its own `package main`:

```sh
cd leetcode/0-problem-solving-building-blocks/09-queues-bfs-and-grids/practice-01-dynamic-queue
GO111MODULE=off go test -race -count=1 -v .
cd ../practice-02-graph-bfs
GO111MODULE=off go test -race -count=1 -v .
```

## Operations to know first
- Start `queue := []*TreeNode{root}`; dequeue front with `node:=queue[0]; queue=queue[1:]`.
- For level order, **snapshot `levelSize:=len(queue)`** before processing that level; newly queued children belong to the next level.
- For grids, a cell is `(row,col)` and neighbors are `(+1,0),(-1,0),(0,+1),(0,-1)`.
- Mark visited **when enqueuing**, so the same cell is not enqueued repeatedly.
- When future moves depend on a limited resource, the BFS state must include it. For [#1293](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md), use `(row,col,remaining)` and a 3D visited set; the same cell with different budgets represents different states.
- If you overwrite grid `'1'` with `'0'`, note that the algorithm mutates its input.

## Runnable Go micro-example
```go
package main
import "fmt"
func countIslands(grid [][]byte) int {
    rows:=len(grid); if rows==0{return 0}
    dirs:=[4][2]int{{1,0},{-1,0},{0,1},{0,-1}}
    count:=0
    for r:=range grid {for c:=range grid[r]{
        if grid[r][c]!='1' {continue}
        count++;queue:=[][2]int{{r,c}};grid[r][c]='0'
        for len(queue)>0 {
            cell:=queue[0];queue=queue[1:]
            for _,d:=range dirs {
                nr,nc:=cell[0]+d[0],cell[1]+d[1]
                if nr<0||nr>=len(grid)||nc<0||nc>=len(grid[nr]){continue}
                if grid[nr][nc]=='1' {grid[nr][nc]='0';queue=append(queue,[2]int{nr,nc})}
            }
        }
    }}
    return count
}
func main(){g:=[][]byte{{'1','1','0'},{'0','1','0'},{'1','0','1'}};fmt.Println(countIslands(g))}
```

## Hand trace
For grid `110 / 010 / 101`, BFS consumes the connected three land cells as one island; the isolated bottom-left and bottom-right cells add two more. Answer 3.

## Recall exercise (no notes)
For tree `[3,9,20,null,null,15,7]`, write each level size: 1,2,2. Explain why `len(queue)` cannot be re-read as the inner-loop limit while adding children.

## Common mistakes
- A queue implemented by reslicing is logically FIFO but can retain backing arrays.
- Index `grid[nr]` only after testing `0<=nr<len(grid)`; ragged grids require row-specific width checks.
- Grid BFS and tree BFS use the same queue but different neighbor-generation rules.

## Stronger problems using this block
- [#102 Binary Tree Level Order Traversal](../../102.%20Binary%20Tree%20Level%20Order%20Traversal/main.go)
- [#200 Number of Islands](../../200.%20Number%20of%20Islands/main.go)
- [#353 Design Snake Game](../../353%20-%20Design%20Snake%20Game/main.go)
- [#1293 Shortest Path in a Grid with Obstacles Elimination](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md)

## Review in 60 seconds
Name the invariant, rewrite the code without looking, then explain one mistake the tests should catch.
