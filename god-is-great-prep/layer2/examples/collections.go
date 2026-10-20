// Package examples contains small, non-concurrent Go collection exercises.
// It is deliberately isolated from the unrelated learning programs in this repository.
package examples

import "slices"

// ArrayCopy proves the array itself is copied on assignment.
func ArrayCopy() ([3]int, [3]int) {
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 9
	return a, b
}

// InsertAt returns an independent slice and accepts positions 0 through len(in).
// An invalid position leaves the input unchanged and returns false.
func InsertAt(in []int, pos, value int) ([]int, bool) {
	if pos < 0 || pos > len(in) {
		return nil, false
	}
	out := make([]int, len(in)+1)
	copy(out, in[:pos])
	out[pos] = value
	copy(out[pos+1:], in[pos:])
	return out, true
}

// DeleteAt returns an independent slice and removes one indexed element.
func DeleteAt(in []int, pos int) ([]int, bool) {
	if pos < 0 || pos >= len(in) {
		return nil, false
	}
	out := make([]int, len(in)-1)
	copy(out, in[:pos])
	copy(out[pos:], in[pos+1:])
	return out, true
}

// WordCounts is a frequency map; missing keys have count zero.
func WordCounts(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

// CloneCounts copies the map. Value type int is not a nested reference.
func CloneCounts(source map[string]int) map[string]int {
	if source == nil {
		return nil
	}
	out := make(map[string]int, len(source))
	for key, value := range source {
		out[key] = value
	}
	return out
}

// User illustrates struct values; a pointer receiver can mutate the original.
type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User { return &User{Name: name, Age: age} }
func (u *User) Birthday()               { u.Age++ }
func (u User) Label() string             { return u.Name }

// ReverseRunes reverses Unicode code points, not grapheme clusters.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Node is a singly linked node.
type Node struct {
	Value int
	Next  *Node
}

func Prepend(head *Node, value int) *Node { return &Node{Value: value, Next: head} }

func AppendNode(head *Node, value int) *Node {
	n := &Node{Value: value}
	if head == nil {
		return n
	}
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = n
	return head
}

func DeleteFirst(head *Node, value int) *Node {
	dummy := &Node{Next: head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Value == value {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	return dummy.Next
}

func ReverseList(head *Node) *Node {
	var prev *Node
	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}

func ListValues(head *Node) []int {
	out := []int{}
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Value)
	}
	return out
}

// DNode and DList illustrate doubly-linked list invariants.
// Remove and MoveToFront require a node currently belonging to this list.
type DNode struct {
	Value int
	Prev  *DNode
	Next  *DNode
}

type DList struct {
	Head *DNode
	Tail *DNode
	size int
}

func (l *DList) Len() int { return l.size }

func (l *DList) PushFront(v int) *DNode {
	n := &DNode{Value: v, Next: l.Head}
	if l.Head != nil {
		l.Head.Prev = n
	} else {
		l.Tail = n
	}
	l.Head = n
	l.size++
	return n
}

func (l *DList) PushBack(v int) *DNode {
	n := &DNode{Value: v, Prev: l.Tail}
	if l.Tail != nil {
		l.Tail.Next = n
	} else {
		l.Head = n
	}
	l.Tail = n
	l.size++
	return n
}

func (l *DList) Remove(n *DNode) {
	if n.Prev != nil {
		n.Prev.Next = n.Next
	} else {
		l.Head = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	} else {
		l.Tail = n.Prev
	}
	n.Prev, n.Next = nil, nil
	l.size--
}

func (l *DList) MoveToFront(n *DNode) {
	if n == l.Head {
		return
	}
	l.Remove(n)
	n.Next = l.Head
	if l.Head != nil {
		l.Head.Prev = n
	} else {
		l.Tail = n
	}
	l.Head = n
	l.size++
}

func (l *DList) Values() []int {
	out := []int{}
	for n := l.Head; n != nil; n = n.Next {
		out = append(out, n.Value)
	}
	return out
}

// IntStack uses the slice tail as LIFO top.
type IntStack struct{ values []int }

func (s *IntStack) Push(v int) { s.values = append(s.values, v) }
func (s *IntStack) Len() int   { return len(s.values) }

func (s *IntStack) Peek() (int, bool) {
	if len(s.values) == 0 {
		return 0, false
	}
	return s.values[len(s.values)-1], true
}

func (s *IntStack) Pop() (int, bool) {
	n := len(s.values)
	if n == 0 {
		return 0, false
	}
	v := s.values[n-1]
	s.values = s.values[:n-1]
	return v, true
}

// IntQueue uses a head index with compaction so it does not grow indefinitely
// as elements are dequeued. Its operations are amortized O(1).
type IntQueue struct {
	values []int
	head   int
}

func (q *IntQueue) Enqueue(v int) { q.values = append(q.values, v) }
func (q *IntQueue) Len() int      { return len(q.values) - q.head }

func (q *IntQueue) Peek() (int, bool) {
	if q.Len() == 0 {
		return 0, false
	}
	return q.values[q.head], true
}

func (q *IntQueue) Dequeue() (int, bool) {
	if q.Len() == 0 {
		return 0, false
	}
	v := q.values[q.head]
	q.values[q.head] = 0
	q.head++
	switch {
	case q.head == len(q.values):
		q.values, q.head = nil, 0
	case q.head >= 64 && q.head*2 >= len(q.values):
		copy(q.values, q.values[q.head:])
		q.values = q.values[:len(q.values)-q.head]
		q.head = 0
	}
	return v, true
}

// IntSet stores unique ints with map[int]struct{}.
type IntSet map[int]struct{}

func NewIntSet(values ...int) IntSet {
	s := make(IntSet)
	for _, value := range values {
		s.Add(value)
	}
	return s
}
func (s IntSet) Add(v int)                  { s[v] = struct{}{} }
func (s IntSet) Remove(v int)               { delete(s, v) }
func (s IntSet) Has(v int) bool             { _, ok := s[v]; return ok }
func (s IntSet) Len() int                   { return len(s) }
func (s IntSet) Union(other IntSet) IntSet {
	out := NewIntSet()
	for v := range s { out.Add(v) }
	for v := range other { out.Add(v) }
	return out
}
func (s IntSet) Intersection(other IntSet) IntSet {
	out := NewIntSet()
	for v := range s { if other.Has(v) { out.Add(v) } }
	return out
}
func (s IntSet) Difference(other IntSet) IntSet {
	out := NewIntSet()
	for v := range s { if !other.Has(v) { out.Add(v) } }
	return out
}

// Multiset is a frequency map. Its zero value must be initialized before Add.
type Multiset struct{ counts map[string]int }

func NewMultiset() *Multiset { return &Multiset{counts: make(map[string]int)} }
func (m *Multiset) Add(key string) { m.counts[key]++ }
func (m *Multiset) Count(key string) int { return m.counts[key] }
func (m *Multiset) RemoveOne(key string) bool {
	if m.counts[key] == 0 { return false }
	if m.counts[key] == 1 { delete(m.counts, key) } else { m.counts[key]-- }
	return true
}

// MatrixSum permits ragged rows.
func MatrixSum(grid [][]int) int {
	total := 0
	for _, row := range grid { for _, v := range row { total += v } }
	return total
}

// Transpose requires a rectangular grid; empty input yields empty output.
func Transpose(grid [][]int) ([][]int, bool) {
	if len(grid) == 0 { return [][]int{}, true }
	cols := len(grid[0])
	for _, row := range grid { if len(row) != cols { return nil, false } }
	out := make([][]int, cols)
	for j := 0; j < cols; j++ {
		out[j] = make([]int, len(grid))
		for i := range grid { out[j][i] = grid[i][j] }
	}
	return out, true
}

// SortedSearch sorts an independent clone, then binary-searches for target.
func SortedSearch(values []int, target int) ([]int, int, bool) {
	clone := slices.Clone(values)
	slices.Sort(clone)
	i, ok := slices.BinarySearch(clone, target)
	return clone, i, ok
}

// Shape is an implicitly satisfied interface.
type Shape interface { Area() int }
type Square struct { Side int }
type Rectangle struct { Width, Height int }
func (s Square) Area() int { return s.Side * s.Side }
func (r Rectangle) Area() int { return r.Width * r.Height }
func TotalArea(shapes []Shape) int {
	total := 0
	for _, s := range shapes { total += s.Area() }
	return total
}
