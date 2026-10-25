// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func multiply(num1 string, num2 string) string {
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
	num1 := "2"
	num2 := "3"
	fmt.Printf("result: %#v\n", multiply(num1, num2))
}

// Local fixture helpers only; do not submit to the judge.
