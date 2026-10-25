// Unsolved starter. Judge interface adapted locally; premium Go starter unavailable.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
type Codec struct{}

func Constructor() Codec {
	panic(todoMarker)
}
func (this *Codec) encode(strs []string) string {
	panic(todoMarker)
}
func (this *Codec) decode(s string) []string {
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
	encoded := obj.encode([]string{"a", "", "b#c"})
	fmt.Println("encoded:", encoded)
	fmt.Printf("decoded: %#v\n", obj.decode(encoded))
}

// Local fixture helpers only; do not submit to the judge.
