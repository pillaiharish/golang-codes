// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
type TimeMap struct {
}

func Constructor() TimeMap {
	panic(todoMarker)
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	panic(todoMarker)
}

func (this *TimeMap) Get(key string, timestamp int) string {
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
	obj.Set("foo", "bar", 1)
	fmt.Printf("get: %#v\n", obj.Get("foo", 1))
	fmt.Printf("get: %#v\n", obj.Get("foo", 3))
	obj.Set("foo", "bar2", 4)
	fmt.Printf("get: %#v\n", obj.Get("foo", 4))
	fmt.Printf("get: %#v\n", obj.Get("foo", 5))
}

// Local fixture helpers only; do not submit to the judge.
