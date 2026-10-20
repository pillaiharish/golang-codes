// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
type DetectSquares struct {
}

func Constructor() DetectSquares {
	panic(todoMarker)
}

func (this *DetectSquares) Add(point []int) {
	panic(todoMarker)
}

func (this *DetectSquares) Count(point []int) int {
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
	obj.Add([]int{3, 10})
	obj.Add([]int{11, 2})
	obj.Add([]int{3, 2})
	fmt.Printf("count: %#v\n", obj.Count([]int{11, 10}))
	fmt.Printf("count: %#v\n", obj.Count([]int{14, 8}))
	obj.Add([]int{11, 2})
	fmt.Printf("count: %#v\n", obj.Count([]int{11, 10}))
}

// Local fixture helpers only; do not submit to the judge.
