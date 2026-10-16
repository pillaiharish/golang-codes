# 10 — Stacks, recursion and DFS tree traversal

[← Foundations index](../README.md) · [← Problem index](../../README.md)

## Recognition and intuition
Recursion gives tree structure naturally; an explicit stack simulates the function call stack. Both encode the same Left → Node → Right invariant for inorder traversal.

## Operations to know first
- Always check `if node==nil{return}` as recursion base case.
- Recursive inorder: recurse Left, append Node, recurse Right.
- With a slice pointer `*[]int`, write `*out=append(*out,v)`; a closure can instead capture the result slice.
- Stack operations: `append(stack,node)` pushes, last index reads top, `stack=stack[:len(stack)-1]` pops.
- For skewed trees recursion stack and explicit stack may each reach O(n) depth; balanced height is O(log n).

## Runnable Go micro-example
```go
package main
import "fmt"
type Node struct{Val int;Left,Right *Node}
func inorder(root *Node)[]int{
    out:=[]int{}
    stack:=[]*Node{}
    cur:=root
    for cur!=nil||len(stack)>0{
        for cur!=nil{stack=append(stack,cur);cur=cur.Left}
        cur=stack[len(stack)-1];stack=stack[:len(stack)-1]
        out=append(out,cur.Val);cur=cur.Right
    }
    return out
}
func main(){fmt.Println(inorder(&Node{Val:2,Left:&Node{Val:1},Right:&Node{Val:3}}))}
```

## Hand trace
BST 2 with children 1 and 3: stack push 2,1; pop/visit 1, then pop/visit 2, then visit 3 → `[1,2,3]`.

## Recall exercise (no notes)
Code recursive inorder before reading the iterative version. Why can a slice pointer or closure share output across recursive calls?

## Common mistakes
- Traversal orders: inorder LNR, preorder NLR, postorder LRN.
- Append may reallocate backing storage; reassign the slice header when passing pointers.
- The iterative version must progress into the popped node's right subtree after visiting it.

## Stronger problems using this block
- [#94 Binary Tree Inorder Traversal](../../94.%20Binary%20Tree%20Inorder%20Traversal/main.go)
- [#102 Binary Tree Level Order Traversal](../../102.%20Binary%20Tree%20Level%20Order%20Traversal/main.go)
- [#200 Number of Islands](../../200.%20Number%20of%20Islands/main.go)

## Review in 60 seconds
Name the invariant, rewrite the code without looking, then explain one mistake the tests should catch.
