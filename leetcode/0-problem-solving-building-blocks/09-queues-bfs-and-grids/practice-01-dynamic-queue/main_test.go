package main

import (
	"reflect"
	"testing"
)

func TestDynamicQueueFIFO(t *testing.T) {
	q := dynamicQueue{}
	for _, x := range []int{5, 3, 11} {
		q.Enqueue(x)
	}
	if !reflect.DeepEqual(q.items, []int{5, 3, 11}) {
		t.Fatalf("enqueue: got %v", q.items)
	}
	for _, want := range []int{5, 3, 11} {
		got, ok := q.Dequeue()
		if !ok || got != want {
			t.Fatalf("Dequeue()=(%d,%v), want (%d,true)", got, ok, want)
		}
	}
	if got, ok := q.Dequeue(); ok || got != 0 {
		t.Fatalf("empty dequeue = (%d,%v), want (0,false)", got, ok)
	}
}

func TestQueueZeroAndReuse(t *testing.T) {
	var q dynamicQueue
	q.Enqueue(0)
	if got, ok := q.Dequeue(); !ok || got != 0 {
		t.Fatalf("zero is a valid value: (%d,%v)", got, ok)
	}
	q.Enqueue(42)
	if got, ok := q.Dequeue(); !ok || got != 42 {
		t.Fatalf("reused queue: (%d,%v)", got, ok)
	}
}

func TestQueueInstancesIndependent(t *testing.T) {
	var a, b dynamicQueue
	a.Enqueue(10)
	b.Enqueue(20)
	if got, _ := a.Dequeue(); got != 10 {
		t.Fatalf("a=%d", got)
	}
	if got, _ := b.Dequeue(); got != 20 {
		t.Fatalf("b=%d", got)
	}
}
