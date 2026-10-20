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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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
	l1 := localList([]int{2, 4, 3})
	l2 := localList([]int{5, 6, 4})
	fmt.Printf("result: %#v\n", addTwoNumbers(l1, l2))
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
