package main

import "fmt"

type dynamicQueue struct {
	items []int
}

func (q *dynamicQueue) Enqueue(data int) {
	q.items = append(q.items, data)
}

// pop is from front.
// return popped int number and if it was removed then true
func (q *dynamicQueue) Dequeue() (int, bool) {
	// queue is empty and pop is false
	if len(q.items) == 0 {
		return 0, false
	}
	// pop from front
	front := q.items[0]
	q.items = q.items[1:]
	return front, true
}

func main() {
	queue := dynamicQueue{}
	queue.Enqueue(5)
	queue.Enqueue(3)
	queue.Enqueue(11)
	fmt.Println("Out", queue.items)
	out, err := queue.Dequeue()
	if !err {
		fmt.Println("Error in Dequeue")
	} else {
		fmt.Println("First Dequeue", out, queue.items)
	}
}
