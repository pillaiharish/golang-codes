package main

import "fmt"

type TreeNode struct {
    Val int
    Left, Right *TreeNode
}

// Primary approach from notes: pass pointer to a shared output slice.
// Left -> Node -> Right. Time O(n); recursion O(h); result O(n).
func inorderTraversal(root *TreeNode) []int {
    result:=[]int{}
    inorder(root,&result)
    return result
}

func inorder(root *TreeNode, result *[]int) {
    if root == nil { return }
    inorder(root.Left,result)
    *result = append(*result,root.Val)
    inorder(root.Right,result)
}

// Variant: closure captures the output slice in outer scope.
func inorderClosure(root *TreeNode) []int {
    result:=[]int{}
    var walk func(*TreeNode)
    walk=func(n *TreeNode) {
        if n==nil {return}
        walk(n.Left)
        result=append(result,n.Val)
        walk(n.Right)
    }
    walk(root)
    return result
}

// Variant: iterative stack. Explicitly simulates recursion.
func inorderIterative(root *TreeNode) []int {
    result:=[]int{}
    stack:=[]*TreeNode{}
    current:=root
    for current!=nil || len(stack)>0 {
        for current!=nil {
            stack=append(stack,current)
            current=current.Left
        }
        current=stack[len(stack)-1]
        stack=stack[:len(stack)-1]
        result=append(result,current.Val)
        current=current.Right
    }
    return result
}

func main() {
    root:=&TreeNode{Val:1,Right:&TreeNode{Val:2,Left:&TreeNode{Val:3}}}
    fmt.Println(inorderTraversal(root)) // [1 3 2]
}
