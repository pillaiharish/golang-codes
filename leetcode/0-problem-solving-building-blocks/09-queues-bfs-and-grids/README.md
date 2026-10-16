# 09 — Queues, breadth-first search and 2D grids

[← Foundations index](../README.md) · [← Problem index](../../README.md)

## Recognition and intuition
BFS processes nodes in order of distance/layers using a FIFO queue. Matrix traversal adds bounds, visited-state and neighbor direction operations.

## Operations to know first
- Start `queue := []*TreeNode{root}`; dequeue front with `node:=queue[0]; queue=queue[1:]`.
- For level order, **snapshot `levelSize:=len(queue)`** before processing that level; newly queued children belong to the next level.
- For grids, a cell is `(row,col)` and neighbors are `(+1,0),(-1,0),(0,+1),(0,-1)`.
- Mark visited **when enqueuing**, so the same cell is not enqueued repeatedly.
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

## Review in 60 seconds
Name the invariant, rewrite the code without looking, then explain one mistake the tests should catch.
