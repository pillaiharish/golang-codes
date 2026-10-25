// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func minInterval(intervals [][]int, queries []int) []int {
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
	intervals := [][]int{[]int{1, 4}, []int{2, 4}, []int{3, 6}, []int{4, 4}}
	queries := []int{2, 3, 4, 5}
	fmt.Printf("result: %#v\n", minInterval(intervals, queries))
}

// Local fixture helpers only; do not submit to the judge.
