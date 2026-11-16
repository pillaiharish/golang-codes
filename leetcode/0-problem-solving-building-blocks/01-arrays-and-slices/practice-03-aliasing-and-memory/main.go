package main

import (
	"encoding/json"
	"fmt"
	"slices"
)

// SharedAppend demonstrates a subslice appending into available shared capacity.
func SharedAppend() (base, result []int) {
	base = []int{1, 2, 3, 4}
	part := base[:2] // len 2, cap 4
	part = append(part, 9)
	return base, part
}

// IsolatedAppend clones before append, so it cannot overwrite base's elements.
func IsolatedAppend() (base, result []int) {
	base = []int{1, 2, 3, 4}
	part := slices.Clone(base[:2])
	part = append(part, 9)
	return base, part
}

// CapacityLimitedAppend requires reallocation because the slice has no spare cap.
func CapacityLimitedAppend() (base, result []int) {
	base = []int{1, 2, 3, 4}
	part := base[:2:2]
	part = append(part, 9)
	return base, part
}

// ZeroLengthCopy illustrates that copy counts len(dst), not cap(dst).
func ZeroLengthCopy() int {
	src := []int{1, 2, 3}
	dst := make([]int, 0, 3)
	return copy(dst, src)
}

// NilEmptyJSON distinguishes a nil slice from allocated empty slice during JSON encoding.
func NilEmptyJSON() (string, string) {
	var nilSlice []int
	emptySlice := []int{}
	a, _ := json.Marshal(nilSlice) // always succeeds for []int
	b, _ := json.Marshal(emptySlice)
	return string(a), string(b)
}

func main() {
	base, part := SharedAppend()
	fmt.Println(base, part)       // [1 2 9 4] [1 2 9]
	fmt.Println(ZeroLengthCopy()) // 0
	a, b := NilEmptyJSON()
	fmt.Println(a, b) // null []
}
