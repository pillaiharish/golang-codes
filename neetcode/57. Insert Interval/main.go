// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func insert(intervals [][]int, newInterval []int) [][]int {
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
	intervals := [][]int{[]int{1, 3}, []int{6, 9}}
	newInterval := []int{2, 5}
	fmt.Printf("result: %#v\n", insert(intervals, newInterval))
}

// Local fixture helpers only; do not submit to the judge.
