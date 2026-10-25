// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// BEGIN JUDGE CODE: implement these TODO bodies.
func kthSmallest(root *TreeNode, k int) int {
	panic(todoMarker)
}

// END JUDGE CODE

// Local harness: catches only the intentional unimplemented marker.
// Unexpected panics are rethrown. A TODO run is not a correctness test.
type localTODO struct{}

var todoMarker = &localTODO{}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == todoMarker {
				fmt.Println("TODO: implement the judge function/methods, then run again.")
				return
			}
			panic(r)
		}
	}()
	localDemo()
}
func localDemo() {
	root := localTree("[3, 1, 4, null, 2]")
	k := 1
	fmt.Printf("result: %#v\n", kthSmallest(root, k))
}

// Local fixture helpers only; do not submit to the judge.
func localTree(data string) *TreeNode {
	var values []*int
	if err := json.Unmarshal([]byte(data), &values); err != nil {
		panic(err)
	}
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	i := 1
	for head := 0; head < len(queue) && i < len(values); head++ {
		node := queue[head]
		if values[i] != nil {
			node.Left = &TreeNode{Val: *values[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(values) {
			if values[i] != nil {
				node.Right = &TreeNode{Val: *values[i]}
				queue = append(queue, node.Right)
			}
			i++
		}
	}
	return root
}
