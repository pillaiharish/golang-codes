// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func findItinerary(tickets [][]string) []string {
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
	tickets := [][]string{[]string{"MUC", "LHR"}, []string{"JFK", "MUC"}, []string{"SFO", "SJC"}, []string{"LHR", "SFO"}}
	fmt.Printf("result: %#v\n", findItinerary(tickets))
}

// Local fixture helpers only; do not submit to the judge.
