package main

import "testing"

func TestQueueFIFOAndReuse(t *testing.T) {
	var q Queue[int]
	if _, ok := q.Dequeue(); ok {
		t.Fatal("zero-value queue must report empty")
	}
	for _, v := range []int{5, 3, 11} {
		q.Enqueue(v)
	}
	for _, want := range []int{5, 3, 11} {
		got, ok := q.Dequeue()
		if !ok || got != want {
			t.Fatalf("Dequeue = (%d,%t), want (%d,true)", got, ok, want)
		}
	}
	if q.Len() != 0 || q.items != nil || q.head != 0 {
		t.Fatal("drained queue should reset")
	}
	q.Enqueue(42)
	if got, ok := q.Dequeue(); !ok || got != 42 {
		t.Fatalf("reuse Dequeue = (%d,%t)", got, ok)
	}
}

func TestQueueClearsConsumedPointer(t *testing.T) {
	first, second := 10, 20
	var q Queue[*int]
	q.Enqueue(&first)
	q.Enqueue(&second)
	got, ok := q.Dequeue()
	if !ok || got != &first {
		t.Fatal("wrong first item")
	}
	if q.items[0] != nil {
		t.Fatal("consumed pointer must be cleared")
	}
	got, ok = q.Dequeue()
	if !ok || got != &second || q.items != nil {
		t.Fatal("wrong second item or backing array not released")
	}
}
