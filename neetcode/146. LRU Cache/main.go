// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
type LRUCache struct {
}

func Constructor(capacity int) LRUCache {
	panic(todoMarker)
}

func (this *LRUCache) Get(key int) int {
	panic(todoMarker)
}

func (this *LRUCache) Put(key int, value int) {
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
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Printf("get: %#v\n", obj.Get(1))
	obj.Put(3, 3)
	fmt.Printf("get: %#v\n", obj.Get(2))
	obj.Put(4, 4)
	fmt.Printf("get: %#v\n", obj.Get(1))
	fmt.Printf("get: %#v\n", obj.Get(3))
	fmt.Printf("get: %#v\n", obj.Get(4))
}

// Local fixture helpers only; do not submit to the judge.
