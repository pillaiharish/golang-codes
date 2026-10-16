# 18 — Interfaces, type assertions and standard-library containers

[← Foundations index](../README.md) · [← LeetCode index](../../README.md)

## Mental model
LeetCode design tasks may store values in `any` and recover the concrete type. `container/list` is a doubly linked list with `Element.Value` of type `any`.

## Operations and reusable rules
- `list.New()` creates a doubly linked list; `PushBack(value)` inserts; `Front()` / `Back()` return *list.Element.
- `lst.Remove(element)` removes that element and returns its stored value as `any`.
- `coord := element.Value.([2]int)` asserts a concrete array type; wrong types panic.
- Safe assertion: `coord,ok := value.([2]int)` to handle mismatches.
- A fixed coordinate array `[2]int` is comparable and safe as a `map[[2]int]bool` key.
- For a hot cache, manually linked `*Node` avoids type assertions and lets the map point directly to nodes.

## Runnable example
```go
package main
import ("container/list";"fmt")
func main(){
    body:=list.New()
    body.PushBack([2]int{0,0})
    body.PushBack([2]int{0,1})
    tailAny:=body.Remove(body.Front())
    tail,ok:=tailAny.([2]int)
    fmt.Println(tail,ok) // [0 0] true
    occupied:=map[[2]int]bool{{0,1}:true}
    fmt.Println(occupied[[2]int{0,1}]) // true
}
```

## Trace it on paper
Body starts `[0,0]→[0,1]`; removing front returns an interface value holding `[0,0]`, which the type assertion safely unwraps.

## Pre-problem drill
Why can `[2]int` be a map key but `[]int` cannot? What happens if `Remove` returns `string` and you attempt `.([2]int)`?

## Pitfalls to recall
- A failed one-value type assertion panics; check comma-ok when type isn't guaranteed.
- Do not confuse slice `[]int` with array `[2]int`; both syntax and comparability differ.
- `container/list` stores values as `any`; a custom typed linked list avoids assertion overhead and mistakes.

## Problems that depend on it
- [#353 Design Snake Game](../../353%20-%20Design%20Snake%20Game/main.go)
- [#146 LRU Cache](../../146.%20LRU%20Cache/main.go)
- [#641 Design Circular Deque](../../641.%20Design%20Circular%20Deque/main.go)

## Cold recall
Write the minimum state and invariant, then rewrite the example in Go and test one edge case.
