// Unsolved starter. Judge interface adapted locally; premium Go starter unavailable.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func canAttendMeetings(intervals [][]int) bool {
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
	intervals := [][]int{[]int{0, 30}, []int{5, 10}, []int{15, 20}}
	fmt.Printf("result: %#v\n", canAttendMeetings(intervals))
}

// Local fixture helpers only; do not submit to the judge.
