// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func networkDelayTime(times [][]int, n int, k int) int {
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
	times := [][]int{[]int{2, 1, 1}, []int{2, 3, 1}, []int{3, 4, 1}}
	n := 4
	k := 2
	fmt.Printf("result: %#v\n", networkDelayTime(times, n, k))
}

// Local fixture helpers only; do not submit to the judge.
