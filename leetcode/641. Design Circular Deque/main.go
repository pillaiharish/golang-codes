package main

import "fmt"

type MyCircularDeque struct {
	elements []int
	size     int
	front    int
	rear     int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{elements: make([]int, k+1), size: k + 1, front: 0, rear: 0}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.front = (this.front - 1 + this.size) % this.size
	this.elements[this.front] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.elements[this.rear] = value
	this.rear = (this.rear + 1) % this.size

	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % this.size
	// this.elements = this.elements[1:len(this.elements)-1]
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.rear = (this.rear - 1 + this.size) % this.size
	// this.elements = this.elements[0:len(this.elements)-2]
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	return this.elements[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.elements[(this.rear-1+this.size)%this.size]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.front == this.rear
}

func (this *MyCircularDeque) IsFull() bool {
	return (this.rear+1)%this.size == this.front
}

func main() {
	obj := Constructor(3)
	fmt.Println(obj.InsertLast(1))  // return true
	fmt.Println(obj.InsertLast(2))  // return true
	fmt.Println(obj.InsertFront(3)) // return true
	fmt.Println(obj.InsertFront(4)) // return false, the queue is full
	fmt.Println(obj.GetRear())      // return 2
	fmt.Println(obj.IsFull())       // return true
	fmt.Println(obj.DeleteLast())   // return true
	fmt.Println(obj.InsertFront(4)) // return true
	fmt.Println(obj.GetFront())     // return 4
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

// Output:
// harish $ go run main.go 
// true
// true
// true
// false
// 2
// true
// true
// true
// 4