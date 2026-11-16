package main

import "fmt"

// Queue uses a slice and a moving head, so dequeue never shifts live elements.
// Its zero value is usable. For a long-lived queue, add compaction or a ring.
type Queue[T any] struct {
	items []T
	head  int
}

func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

func (q *Queue[T]) Len() int {
	return len(q.items) - q.head
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T
	if q.Len() == 0 {
		return zero, false
	}
	value := q.items[q.head]
	q.items[q.head] = zero // avoid retaining pointers in consumed slots
	q.head++
	if q.head == len(q.items) {
		q.items = nil // release storage when fully drained
		q.head = 0
	}
	return value, true
}

func main() {
	var queue Queue[int]
	for _, v := range []int{5, 3, 11} {
		queue.Enqueue(v)
	}
	first, ok := queue.Dequeue()
	fmt.Println(first, ok, queue.Len()) // 5 true 2
}
