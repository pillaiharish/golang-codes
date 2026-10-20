// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// BEGIN JUDGE CODE: implement these TODO bodies.
func hasCycle(head *ListNode) bool {
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
	head := localList([]int{3, 2, 0, -4})
	// Fixture pos wires the cycle; it is not a judge function argument.
	pos := 1
	if pos >= 0 {
		nodes := []*ListNode{}
		for node := head; node != nil; node = node.Next {
			nodes = append(nodes, node)
		}
		nodes[len(nodes)-1].Next = nodes[pos]
	}
	fmt.Printf("result: %#v\n", hasCycle(head))
}

// Local fixture helpers only; do not submit to the judge.
func localList(values []int) *ListNode {
	var head, tail *ListNode
	for _, value := range values {
		node := &ListNode{Val: value}
		if head == nil {
			head = node
		} else {
			tail.Next = node
		}
		tail = node
	}
	return head
}
