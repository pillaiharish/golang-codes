# 11 — Structs, pointers, methods and doubly linked lists

[← Foundations index](../README.md) · [← Problem index](../../README.md)

## Recognize the building block
A map gives key lookup, while a doubly linked list gives constant-time removal and recency updates once you already have a node pointer. Structs and pointer receivers let methods mutate shared state.

## Operations you must be able to write
- Define fields with `type Node struct {key,value int;prev,next *Node}`.
- `&Node{...}` constructs a pointer. `*Node` lets multiple structures reference the same node.
- `func (c *Cache) Get(key int) int` has a pointer receiver so the method mutates the original cache.
- A constructor returning a `Codec` value uses `return Codec{}`; returning `&Codec{}` requires a `*Codec` return type. `&` creates a pointer.
- Field initialization needs a **value**: `Codec{encoded: ""}` is valid for a string field; `Codec{encoded: string}` is not, because `string` names a type.
- Choose a field only when later method calls require remembered state. A rate limiter needs a timestamp map; a codec whose `decode` receives its encoded input can use `type Codec struct{}`.
- To remove n: `n.prev.next=n.next`, `n.next.prev=n.prev`; to insert, rewire both neighbors.
- Two sentinels `head` and `tail` eliminate edge cases when the list is empty.

## Complete small Go example
```go
package main
import "fmt"
type Node struct {Value int;Prev,Next *Node}
func main(){
    head,tail:=&Node{},&Node{}
    head.Next=tail;tail.Prev=head
    a:=&Node{Value:7}
    a.Prev=head;a.Next=tail
    head.Next=a;tail.Prev=a
    fmt.Println(head.Next.Value) // 7
    a.Prev.Next=a.Next
    a.Next.Prev=a.Prev
    fmt.Println(head.Next==tail) // true
}
```

## Dry run
Empty list head↔tail; insert node 7 by setting four links; removing it restores head↔tail. In LRU, `tail.Prev` is always least recent.

## Tiny practice task
Draw four pointer assignments to insert a node after head; show which map key to delete after evicting `tail.Prev`.

## Common mistakes
- A field such as `encoded string` that is continually appended to can leak data between repeated `Encode` calls unless deliberately reset. A local variable is simpler when each call is independent.
- A method declared on `*Codec` can be called on an addressable `Codec` variable: Go automatically takes its address for the method call. The constructor does not have to return a pointer just because the method has a pointer receiver.
- A nil `prev` or `next` causes a dereference panic; use sentinels consistently.
- A map of pointers to nodes needs synchronization if mutated concurrently.
- Returning a struct with internal pointers doesn't make independent deep copies.

## Problem links from this repository
- [#146 LRU Cache](../../146.%20LRU%20Cache/main.go)
- [#94 Binary Tree Inorder Traversal](../../94.%20Binary%20Tree%20Inorder%20Traversal/main.go)
- [#102 Binary Tree Level Order Traversal](../../102.%20Binary%20Tree%20Level%20Order%20Traversal/main.go)
- [#353 Design Snake Game](../../353%20-%20Design%20Snake%20Game/main.go)

## Blank-screen check
Close this note. Restate the invariant, retype the operations, then solve the first linked problem without looking.
