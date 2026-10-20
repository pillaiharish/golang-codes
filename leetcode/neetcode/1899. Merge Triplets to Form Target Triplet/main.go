// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func mergeTriplets(triplets [][]int, target []int) bool {
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
	triplets := [][]int{[]int{2, 5, 3}, []int{1, 8, 4}, []int{1, 7, 5}}
	target := []int{2, 7, 5}
	fmt.Printf("result: %#v\n", mergeTriplets(triplets, target))
}

// Local fixture helpers only; do not submit to the judge.
