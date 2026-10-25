// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
type KthLargest struct {
}

func Constructor(k int, nums []int) KthLargest {
	panic(todoMarker)
}

func (this *KthLargest) Add(val int) int {
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
	obj := Constructor(3, []int{4, 5, 8, 2})
	fmt.Printf("add: %#v\n", obj.Add(3))
	fmt.Printf("add: %#v\n", obj.Add(5))
	fmt.Printf("add: %#v\n", obj.Add(10))
	fmt.Printf("add: %#v\n", obj.Add(9))
	fmt.Printf("add: %#v\n", obj.Add(4))
}

// Local fixture helpers only; do not submit to the judge.
