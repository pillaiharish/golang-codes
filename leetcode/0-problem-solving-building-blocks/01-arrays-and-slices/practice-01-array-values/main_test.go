package main

import "testing"

func TestArrayValueAndPointer(t *testing.T) {
	var zero [3]int
	if zero != [3]int{0, 0, 0} {
		t.Fatal("array elements must have their zero values")
	}
	original := [3]int{1, 2, 3}
	got := ChangeCopy(original)
	if original != [3]int{1, 2, 3} || got != [3]int{99, 2, 3} {
		t.Fatalf("copy changed caller: original %v, got %v", original, got)
	}
	ChangeThroughPointer(&original)
	if original[0] != 99 {
		t.Fatal("pointer operation did not modify caller")
	}
}

func TestArrayKeyAndSliceSharing(t *testing.T) {
	if got := ArrayKey([3]int{5, 6, 7}); got != 2 {
		t.Fatalf("key count = %d, want 2", got)
	}
	original := [2][]int{{1}, {2}}
	copied := original // array values copied; each element is still a slice header
	copied[0][0] = 44
	if original[0][0] != 44 {
		t.Fatal("nested slice elements should still share their backing storage")
	}
}
