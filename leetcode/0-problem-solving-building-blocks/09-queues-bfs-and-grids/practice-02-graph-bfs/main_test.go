package main

import (
	"io"
	"os"
	"testing"
)

// Capture the printed BFS without modifying the learner's implementation.
// Do not parallelize tests while redirecting process-wide stdout.
func captureBFS(t *testing.T, graph map[int][]int, start int) string {
	t.Helper()
	original := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	defer func() { os.Stdout = original; r.Close() }()
	BFS(graph, start)
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

func TestGraphBFS(t *testing.T) {
	tests := []struct {
		name  string
		graph map[int][]int
		start int
		want  string
	}{
		{"original graph", map[int][]int{0: {1, 3}, 1: {0, 2, 4}, 2: {1}, 3: {0, 4}, 4: {1, 3}}, 0, "BFS Traversal: 0 1 3 2 4 \n"},
		{"cycle", map[int][]int{1: {2}, 2: {1, 3}, 3: {1}}, 1, "BFS Traversal: 1 2 3 \n"},
		{"duplicate neighbor", map[int][]int{0: {1, 1}, 1: {0}}, 0, "BFS Traversal: 0 1 \n"},
		{"disconnected", map[int][]int{0: {1}, 1: {0}, 2: {3}, 3: {2}}, 0, "BFS Traversal: 0 1 \n"},
		{"isolated", map[int][]int{}, 9, "BFS Traversal: 9 \n"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := captureBFS(t, tc.graph, tc.start); got != tc.want {
				t.Fatalf("BFS printed %q; want %q", got, tc.want)
			}
		})
	}
}
