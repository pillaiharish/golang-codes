// Unsolved starter. Judge interface adapted locally; premium Go starter unavailable.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func alienOrder(words []string) string {
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
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}
	fmt.Printf("result: %#v\n", alienOrder(words))
}

// Local fixture helpers only; do not submit to the judge.
