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
func mergeKLists(lists []*ListNode) *ListNode {
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
	lists := []*ListNode{localList([]int{1, 4, 5}), localList([]int{1, 3, 4}), localList([]int{2, 6})}
	fmt.Printf("result: %#v\n", mergeKLists(lists))
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
