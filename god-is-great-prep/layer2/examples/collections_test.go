package examples

import (
	"reflect"
	"testing"
)

func TestArrayCopy(t *testing.T) {
	a, b := ArrayCopy()
	if a != [3]int{1, 2, 3} || b != [3]int{9, 2, 3} { t.Fatalf("unexpected arrays: %v %v", a, b) }
}

func TestSliceOperations(t *testing.T) {
	in := []int{1, 3}
	inserted, ok := InsertAt(in, 1, 2)
	if !ok || !reflect.DeepEqual(inserted, []int{1, 2, 3}) { t.Fatal(inserted, ok) }
	inserted[0] = 8
	if in[0] != 1 { t.Fatal("insert accidentally aliased input") }
	deleted, ok := DeleteAt(inserted, 1)
	if !ok || !reflect.DeepEqual(deleted, []int{8, 3}) { t.Fatal(deleted, ok) }
	if _, ok := DeleteAt(in, -1); ok { t.Fatal("negative index accepted") }
	if _, ok := InsertAt(in, 3, 4); ok { t.Fatal("out of bounds accepted") }
	empty, ok := InsertAt(nil, 0, 7)
	if !ok || !reflect.DeepEqual(empty, []int{7}) { t.Fatal(empty, ok) }
}

func TestMapsStructsAndRunes(t *testing.T) {
	counts := WordCounts([]string{"go", "rust", "go"})
	if counts["go"] != 2 || counts["missing"] != 0 { t.Fatal(counts) }
	clone := CloneCounts(counts)
	clone["go"] = 10
	if counts["go"] != 2 { t.Fatal("map clone shares state") }
	if CloneCounts(nil) != nil { t.Fatal("nil input should yield nil map") }
	u := NewUser("Ada", 28)
	v := *u
	u.Birthday()
	if u.Age != 29 || v.Age != 28 || u.Label() != "Ada" { t.Fatal(u, v) }
	if ReverseRunes("hé🙂") != "🙂éh" { t.Fatal("rune reversal failed") }
}

func TestSinglyLinkedList(t *testing.T) {
	var head *Node
	head = AppendNode(head, 2)
	head = Prepend(head, 1)
	head = AppendNode(head, 3)
	if got := ListValues(head); !reflect.DeepEqual(got, []int{1, 2, 3}) { t.Fatal(got) }
	head = DeleteFirst(head, 1)
	head = DeleteFirst(head, 999)
	head = ReverseList(head)
	if got := ListValues(head); !reflect.DeepEqual(got, []int{3, 2}) { t.Fatal(got) }
	if ReverseList(nil) != nil { t.Fatal("reverse(nil) != nil") }
}

func TestDoublyLinkedList(t *testing.T) {
	var list DList
	first := list.PushFront(1)
	last := list.PushBack(3)
	middle := list.PushBack(2)
	list.MoveToFront(middle)
	if got := list.Values(); !reflect.DeepEqual(got, []int{2, 1, 3}) { t.Fatal(got) }
	list.Remove(first)
	list.Remove(last)
	if list.Len() != 1 || list.Head != middle || list.Tail != middle || middle.Prev != nil || middle.Next != nil {
		t.Fatal("head/tail/back-link invariant broken")
	}
	list.Remove(middle)
	if list.Head != nil || list.Tail != nil || list.Len() != 0 { t.Fatal("nonempty after removal") }
}

func TestStackAndQueue(t *testing.T) {
	var stack IntStack
	if _, ok := stack.Pop(); ok { t.Fatal("empty stack pop succeeded") }
	stack.Push(1); stack.Push(2)
	if got, ok := stack.Peek(); !ok || got != 2 { t.Fatal(got, ok) }
	if got, ok := stack.Pop(); !ok || got != 2 { t.Fatal(got, ok) }
	if stack.Len() != 1 { t.Fatal(stack.Len()) }

	var q IntQueue
	if _, ok := q.Dequeue(); ok { t.Fatal("empty queue dequeue succeeded") }
	for i := 0; i < 200; i++ { q.Enqueue(i) }
	for i := 0; i < 200; i++ {
		got, ok := q.Dequeue()
		if !ok || got != i { t.Fatalf("position %d got %d %v", i, got, ok) }
	}
	if q.Len() != 0 || len(q.values) != 0 { t.Fatal("queue did not release backing values") }
	q.Enqueue(300)
	if got, ok := q.Dequeue(); !ok || got != 300 { t.Fatal(got, ok) }
}

func TestSetAndMultiset(t *testing.T) {
	a, b := NewIntSet(1, 2), NewIntSet(2, 3)
	if !a.Union(b).Has(3) || !a.Intersection(b).Has(2) || a.Difference(b).Has(2) { t.Fatal("set operation incorrect") }
	a.Remove(1)
	if a.Has(1) { t.Fatal("remove failed") }
	m := NewMultiset()
	m.Add("go"); m.Add("go")
	if m.Count("go") != 2 || !m.RemoveOne("go") || m.Count("go") != 1 { t.Fatal("multiset decrement") }
	m.RemoveOne("go")
	if m.Count("go") != 0 || m.RemoveOne("go") { t.Fatal("multiset empty") }
}

func TestMatrixSortAndInterface(t *testing.T) {
	if MatrixSum([][]int{{1, 2}, {3}}) != 6 { t.Fatal("ragged matrix sum") }
	tr, ok := Transpose([][]int{{1, 2}, {3, 4}})
	if !ok || !reflect.DeepEqual(tr, [][]int{{1, 3}, {2, 4}}) { t.Fatal(tr, ok) }
	if _, ok := Transpose([][]int{{1}, {2, 3}}); ok { t.Fatal("ragged transpose accepted") }
	in := []int{3, 1, 2}
	sorted, idx, found := SortedSearch(in, 2)
	if !found || idx != 1 || !reflect.DeepEqual(sorted, []int{1, 2, 3}) || in[0] != 3 { t.Fatal(sorted, idx, found) }
	if _, _, found := SortedSearch(nil, 2); found { t.Fatal("nil search found value") }
	area := TotalArea([]Shape{Square{Side: 3}, Rectangle{Width: 2, Height: 4}})
	if area != 17 { t.Fatal(area) }
}
