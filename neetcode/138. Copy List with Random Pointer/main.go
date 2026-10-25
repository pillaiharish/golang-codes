// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

type Node struct {
	Val          int
	Next, Random *Node
}

// BEGIN JUDGE CODE: implement these TODO bodies.
func copyRandomList(head *Node) *Node {
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
	a := &Node{Val: 4}
	b := &Node{Val: 7}
	a.Next = b
	a.Random = b
	b.Random = a
	fmt.Printf("result: %#v\n", copyRandomList(a))
}

// Local fixture helpers only; do not submit to the judge.
