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
type Codec struct {
}

func Constructor() Codec {
	panic(todoMarker)
}

func (this *Codec) serialize(root *TreeNode) string {
	panic(todoMarker)
}

func (this *Codec) deserialize(data string) *TreeNode {
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
	obj := Constructor()
	root := &TreeNode{Val: 4, Right: &TreeNode{Val: 7}}
	encoded := obj.serialize(root)
	fmt.Println("encoded:", encoded)
	fmt.Printf("decoded: %#v\n", obj.deserialize(encoded))
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
