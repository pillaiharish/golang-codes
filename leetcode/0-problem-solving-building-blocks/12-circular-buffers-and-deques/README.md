# 12 — Circular buffers, modulo, head and tail

[← Foundations index](../README.md) · [← Problem index](../../README.md)

## Recognize the building block
A ring buffer reuses a fixed array and can guarantee bounded memory for deque operations; modulo wraps logical positions onto physical slots.

## Operations you must be able to write
- Allocate `capacity+1` slots to disambiguate full versus empty without a separate length counter.
- When front==rear, deque is empty; when `(rear+1)%size==front`, deque is full.
- Push front: `(front-1+size)%size`; push back: assign at rear, then `(rear+1)%size`.
- Pop front: `(front+1)%size`; pop back: `(rear-1+size)%size`.
- Peek rear at `(rear-1+size)%size`; modulo with negative values needs normalized arithmetic.

## Complete small Go example
```go
package main
import "fmt"
func main(){
    size:=4 // capacity 3, one slot unused
    a:=make([]int,size)
    front,rear:=0,0
    a[rear]=10;rear=(rear+1)%size
    a[rear]=20;rear=(rear+1)%size
    front=(front-1+size)%size;a[front]=5
    fmt.Println(a[front],a[(rear-1+size)%size]) // 5 20
    fmt.Println((rear+1)%size==front) // true (full)
}
```

## Dry run
With size=4: put 10 at slot 0 (rear 1), 20 at slot 1 (rear 2), then put 5 at front slot 3. Occupied order is [5,10,20].

## Tiny practice task
Capacity=3, insert last 1, insert last 2, insert front 3; predict full; delete rear, insert front 4, predict order [4,3,1].

## Common mistakes
- Off-by-one: `size` here includes a deliberately unused slot.
- `(front-1)%size` can be negative in Go; add size first.
- A circular buffer has fixed allocated capacity unlike a slice queue that may retain backing storage.

## Problem links from this repository
- [#641 Design Circular Deque](../../641.%20Design%20Circular%20Deque/main.go)
- [#239 Sliding Window Maximum](../../239.%20Sliding%20Window%20Maximum/main.go)
- [#102 Binary Tree Level Order Traversal](../../102.%20Binary%20Tree%20Level%20Order%20Traversal/main.go)

## Blank-screen check
Close this note. Restate the invariant, retype the operations, then solve the first linked problem without looking.
