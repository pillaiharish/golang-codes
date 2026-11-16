package main

import (
	"slices"
	"testing"
)

func TestAppendAliasingVsIsolation(t *testing.T) {
	base, result := SharedAppend()
	if !slices.Equal(base, []int{1, 2, 9, 4}) || !slices.Equal(result, []int{1, 2, 9}) {
		t.Fatalf("shared append: base %v, result %v", base, result)
	}
	for _, fn := range []func() ([]int, []int){IsolatedAppend, CapacityLimitedAppend} {
		base, result = fn()
		if !slices.Equal(base, []int{1, 2, 3, 4}) || !slices.Equal(result, []int{1, 2, 9}) {
			t.Fatalf("isolated append: base %v, result %v", base, result)
		}
	}
}

func TestZeroLengthCopyAndJSON(t *testing.T) {
	if got := ZeroLengthCopy(); got != 0 {
		t.Fatalf("copy should copy zero elements; got %d", got)
	}
	if a, b := NilEmptyJSON(); a != "null" || b != "[]" {
		t.Fatalf("nil vs empty JSON: %q %q", a, b)
	}
}

func TestRangeValueCopyAndShallowNestedClone(t *testing.T) {
	type User struct{ Score int }
	users := []User{{Score: 1}}
	for _, u := range users {
		u.Score = 99 // changes a copy, not users[0]
	}
	if users[0].Score != 1 {
		t.Fatal("range value unexpectedly modified original")
	}
	for i := range users {
		users[i].Score = 99
	}
	if users[0].Score != 99 {
		t.Fatal("indexed mutation did not change original")
	}
	nested := [][]int{{1}, {2}}
	outerClone := slices.Clone(nested)
	outerClone[0][0] = 88
	if nested[0][0] != 88 {
		t.Fatal("outer clone should still share inner slice backing arrays")
	}
}

func TestClearTailThenShrinkDoesNotKeepPointersInOldArray(t *testing.T) {
	a, b := 1, 2
	backing := []*int{&a, &b}
	front := backing[:1] // retains the same backing array
	backing[1] = nil     // explicitly release object reference in unused tail
	if front[0] != &a || backing[1] != nil {
		t.Fatal("pointer tail must be cleared")
	}
}
