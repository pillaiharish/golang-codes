package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestInsertAt(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		index int
		value int
		want  []int
		ok    bool
	}{
		{"beginning", []int{1, 2}, 0, 9, []int{9, 1, 2}, true},
		{"middle", []int{1, 2, 3}, 1, 9, []int{1, 9, 2, 3}, true},
		{"end", []int{1, 2}, 2, 9, []int{1, 2, 9}, true},
		{"empty", nil, 0, 9, []int{9}, true},
		{"negative", []int{1}, -1, 9, []int{1}, false},
		{"too large", []int{1}, 2, 9, []int{1}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := InsertAt(slices.Clone(tc.input), tc.index, tc.value)
			if ok != tc.ok || !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("InsertAt = (%v,%v); want (%v,%v)", got, ok, tc.want, tc.ok)
			}
		})
	}
}

func TestDeleteAt(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		index int
		want  []int
		ok    bool
	}{
		{"first", []int{1, 2, 3}, 0, []int{2, 3}, true},
		{"middle", []int{1, 2, 3}, 1, []int{1, 3}, true},
		{"last", []int{1, 2, 3}, 2, []int{1, 2}, true},
		{"singleton", []int{1}, 0, []int{}, true},
		{"empty", nil, 0, nil, false},
		{"invalid", []int{1}, 1, []int{1}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := DeleteAt(slices.Clone(tc.input), tc.index)
			if ok != tc.ok || !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("DeleteAt = (%v,%v); want (%v,%v)", got, ok, tc.want, tc.ok)
			}
		})
	}
}

func TestDeleteAtClearsLastPointerSlot(t *testing.T) {
	a, b, c := 1, 2, 3
	backing := []*int{&a, &b, &c}
	got, ok := DeleteAt(backing, 1)
	if !ok || len(got) != 2 || got[1] != &c || backing[2] != nil {
		t.Fatalf("delete should shift and clear removed tail: %v", backing)
	}
}

func TestFilterAndDeduplicate(t *testing.T) {
	a, b, c := 1, 2, 3
	backing := []*int{&a, &b, &c}
	filtered := FilterInPlace(backing, func(p *int) bool { return *p%2 == 1 })
	if len(filtered) != 2 || filtered[0] != &a || filtered[1] != &c || backing[2] != nil {
		t.Fatalf("filter failed or retained tail pointer")
	}
	if got := FilterInPlace([]int(nil), func(n int) bool { return true }); got != nil {
		t.Fatal("nil input should remain nil")
	}
	input := []int{3, 1, 3, 2, 1}
	if got := DeduplicateStable(input); !slices.Equal(got, []int{3, 1, 2}) || !slices.Equal(input, []int{3, 1, 3, 2, 1}) {
		t.Fatalf("deduplicate got %v, input %v", got, input)
	}
}

func TestSortedCopyAndDeepClone(t *testing.T) {
	original := []int{3, 1, 2}
	if got := SortedCopy(original); !slices.Equal(got, []int{1, 2, 3}) || !slices.Equal(original, []int{3, 1, 2}) {
		t.Fatal("sorted copy should not mutate original")
	}
	matrix := [][]int{{1, 2}, nil, {3}}
	clone := DeepCloneMatrix(matrix)
	clone[0][0] = 77
	if matrix[0][0] != 1 || clone[1] != nil {
		t.Fatalf("deep clone wrong: original %v, clone %v", matrix, clone)
	}
	if DeepCloneMatrix(nil) != nil {
		t.Fatal("nil matrix must remain nil")
	}
}

func TestPopReverseAndSearch(t *testing.T) {
	var empty []int
	if got, value, ok := PopLast(empty); ok || got != nil || value != 0 {
		t.Fatalf("empty pop: %v %d %v", got, value, ok)
	}
	values := []int{3, 1, 2}
	ReverseInPlace(values)
	if !slices.Equal(values, []int{2, 1, 3}) {
		t.Fatalf("reverse: %v", values)
	}
	values, last, ok := PopLast(values)
	if !ok || last != 3 || !slices.Equal(values, []int{2, 1}) {
		t.Fatalf("pop: %v %d %v", values, last, ok)
	}
	if idx, found := BinarySearch([]int{1, 3, 5, 7}, 5); !found || idx != 2 {
		t.Fatalf("binary search: %d %v", idx, found)
	}
	if _, found := BinarySearch([]int{1, 3, 5, 7}, 4); found {
		t.Fatal("missing value reported present")
	}
}
