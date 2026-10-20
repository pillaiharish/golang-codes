# Layer 2 — worked Go examples with explanations

[Layer 2 index](README.md) · [Original long course](reference.md) · [Executable examples](examples/README.md)

**How to read:** These are focused fragments; declarations, imports and wrapping functions are intentionally omitted to isolate each idea. Use the compilable [examples module](examples/README.md) for full code and tests. Each example assumes required imports (such as `fmt`, `slices` or `testing`). Predict the output, then read why.

## 01 Arrays

**Purpose:** Fixed-length value collections. **Typical cost:** O(1) index access, O(n) copying.

~~~go
a := [3]int{1, 2, 3}
b := a
b[0] = 9
fmt.Println(a, b) // [1 2 3] [9 2 3]
~~~

**Explanation:** An array copy is independent when its element values do not themselves reference shared mutable data. Array length is part of its type.

**Your exercise:** Try comparison, index traversal, a 2D array, and a copy containing a slice.

---

## 02 Slices

**Purpose:** Resizable views over arrays. **Typical cost:** O(1) indexing; amortized O(1) append; O(n) insert/remove.

~~~go
a := []int{1, 2, 3, 4}
b := a[:2]
b[0] = 9
fmt.Println(a) // [9 2 3 4]
copyOfA := append([]int(nil), a...)
~~~

**Explanation:** Sub-slices share backing storage; cloning makes independent elements for a slice of integers. Changing a referenced element is not a deep copy.

**Your exercise:** Write insert/delete at index using copy; cover nil and capacity, and distinguish delete from clearing removed pointers.

---

## 03 Maps

**Purpose:** Key/value lookups. **Typical cost:** Average O(1) lookup/insert/delete; worst-case not guaranteed.

~~~go
ages := map[string]int{"Ada": 28}
ages["Lin"] = 0
age, ok := ages["Lin"]
fmt.Println(age, ok) // 0 true
delete(ages, "Ada")
~~~

**Explanation:** Comma-ok disambiguates a missing key from a present zero. Map iteration order is unspecified, so sort keys when display order matters.

**Your exercise:** Count words, group values by key, clone, and handle a nil map.

---

## 04 Structs

**Purpose:** Named fields and value composition. **Typical cost:** O(1) field access; copies proportional to struct size.

~~~go
type Person struct { Name string; Age int }
p := Person{Name: "Ada", Age: 28}
copyP := p
copyP.Age++
fmt.Println(p.Age, copyP.Age) // 28 29
~~~

**Explanation:** Assignment copies struct fields. Fields containing pointers, maps or slices may still reference shared data.

**Your exercise:** Construct, edit, embed a field, compare comparable structs and explain zero values.

---

## 05 Pointers and methods

**Purpose:** Addresses and receiver behavior. **Typical cost:** O(1) direct dereference.

~~~go
type Counter struct { n int }
func (c *Counter) Inc() { c.n++ }
func (c Counter) Value() int { return c.n }
c := Counter{}
c.Inc()
fmt.Println(c.Value()) // 1
~~~

**Explanation:** A pointer receiver changes the original Counter; a value receiver receives a copy. Go still passes arguments by value, including pointer values.

**Your exercise:** Write a swap, a constructor, and a method set example.

---

## 06 Strings, bytes and runes

**Purpose:** UTF-8 and immutable strings. **Typical cost:** O(bytes) conversion to runes.

~~~go
s := "hé"
fmt.Println(len(s)) // 3 bytes
for offset, r := range s { fmt.Println(offset, string(r)) }
runes := []rune(s)
fmt.Println(len(runes)) // 2
~~~

**Explanation:** String indexing returns bytes; range decodes UTF-8 code points with byte offsets. Code points are not always user-perceived characters.

**Your exercise:** Reverse by rune; explain combining marks and invalid UTF-8.

---

## 07 Singly linked list

**Purpose:** Pointer-linked nodes. **Typical cost:** O(1) prepend, O(n) traverse/search.

~~~go
type Node struct { Val int; Next *Node }
head := &Node{Val: 2}
head = &Node{Val: 1, Next: head}
for p := head; p != nil; p = p.Next { fmt.Println(p.Val) }
~~~

**Explanation:** Prepending replaces the head pointer. A list has no O(1) random index access.

**Your exercise:** Append, remove head and middle, insert at index, reverse and detect empty list.

---

## 08 Doubly linked list

**Purpose:** Previous/next pointers and head/tail. **Typical cost:** O(1) move/remove if node is known.

~~~go
type DNode struct { Value int; Prev, Next *DNode }
a := &DNode{Value: 1}
b := &DNode{Value: 2, Prev: a}
a.Next = b
fmt.Println(b.Prev.Value) // 1
~~~

**Explanation:** Each mutation must repair both directions plus head/tail. Knowing the node allows O(1) detachment; finding it remains O(n).

**Your exercise:** Insert/delete at head/tail, move an existing node to front, verify all back-links.

---

## 09 Stack (LIFO)

**Purpose:** Last-in, first-out. **Typical cost:** Amortized O(1) push/pop.

~~~go
stack := []int{}
stack = append(stack, 10, 20)
top := stack[len(stack)-1]
stack = stack[:len(stack)-1]
fmt.Println(top) // 20
~~~

**Explanation:** Use the end of a slice as the stack top. Check emptiness before indexing; a reusable stack should return (value, ok).

**Your exercise:** Implement Push/Pop/Peek/Len and balanced parentheses.

---

## 10 Queue and deque (FIFO)

**Purpose:** First-in, first-out and two-ended access. **Typical cost:** Amortized O(1) with an appropriate ring buffer.

~~~go
q := []int{10, 20, 30}
first := q[0]
q[0] = 0
q = q[1:]
fmt.Println(first, q) // 10 [20 30]
~~~

**Explanation:** Reslicing is O(1), not O(n) copying, but continued growth may retain backing memory. Production queues need compaction or circular indexing.

**Your exercise:** Implement queue with reset/compaction; reason about wraparound and empty/singleton.

---

## 11 Set using map

**Purpose:** Unique membership. **Typical cost:** Average O(1) add/contains/remove.

~~~go
seen := map[int]struct{}{}
seen[4] = struct{}{}
_, ok := seen[4]
fmt.Println(ok) // true
delete(seen, 4)
~~~

**Explanation:** An empty struct consumes no per-value payload, but a map still has overhead. Order is unspecified.

**Your exercise:** Implement union, intersection, difference, clone and size.

---

## 12 Multiset / frequency map

**Purpose:** Counting occurrences. **Typical cost:** Average O(1) update/lookup.

~~~go
counts := map[string]int{}
for _, w := range []string{"go", "go", "rust"} { counts[w]++ }
fmt.Println(counts["go"]) // 2
counts["go"]--
~~~

**Explanation:** A multiset permits duplicates by storing counts. On removing the final occurrence, delete the key to keep the map canonical.

**Your exercise:** Find top frequency, add/remove one, account for absent items.

---

## 13 2D slices and matrices

**Purpose:** Rows, columns and jagged input. **Typical cost:** O(rows × columns) to traverse.

~~~go
matrix := [][]int{{1, 2}, {3}}
sum := 0
for _, row := range matrix { for _, n := range row { sum += n } }
fmt.Println(sum) // 6
~~~

**Explanation:** Rows are independent slices and may differ in length. Always use len(row), not len(matrix[0]), for safe traversal.

**Your exercise:** Make rectangular grid; transpose after validating rectangular input.

---

## 14 Sorting and searching

**Purpose:** Ordered comparisons and lookups. **Typical cost:** O(n log n) sort; O(log n) binary search.

~~~go
nums := []int{8, 3, 5}
slices.Sort(nums)
idx, found := slices.BinarySearch(nums, 5)
fmt.Println(nums, idx, found) // [3 5 8] 1 true
~~~

**Explanation:** Binary search requires data sorted using compatible ordering. Comparator mistakes can break custom sorts.

**Your exercise:** Sort structs by two fields and compare linear search with binary search.

---

## 15 Interfaces as building blocks

**Purpose:** Implicit contracts for behavior. **Typical cost:** Dynamic dispatch; complexity of concrete method.

~~~go
type Area interface { Area() int }
type Square struct { Side int }
func (s Square) Area() int { return s.Side * s.Side }
var x Area = Square{Side: 3}
fmt.Println(x.Area()) // 9
~~~

**Explanation:** A type satisfies an interface by its method set; there is no implements keyword. Begin with tiny interfaces.

**Your exercise:** Implement a second shape; discuss pointer vs value method satisfaction later in Layer 3.

---

## 16 Testing structures

**Purpose:** Verify behavior and edges. **Typical cost:** Usually proportional to number of tested operations.

~~~go
func TestPushPop(t *testing.T) {
 s := IntStack{}
 s.Push(7)
 if got, ok := s.Pop(); !ok || got != 7 { t.Fatal(got, ok) }
}
~~~

**Explanation:** Tests exercise contracts, not only happy paths. Include nil, empty, singleton, repeated operation and out-of-range as applicable.

**Your exercise:** Run go test ./... and add table-driven cases once basics are strong.

---

## 17 Complexity and ownership

**Purpose:** Select the right representation. **Typical cost:** Depends on operation and data model.

~~~go
// Slice indexing: O(1)
// Slice middle insertion: O(n)
// Map lookup: O(1) average
// Linked list index search: O(n)
// Known doubly linked node removal: O(1)
~~~

**Explanation:** Always state assumptions: amortized versus worst-case, whether tail pointer exists, and whether a copy shares referenced elements.

**Your exercise:** For each structure justify one example where it beats an alternative.

---

## 18 Snippet library and retention

**Purpose:** Reproduce patterns from memory. **Typical cost:** Measured by recall, not runtime.

~~~go
// From a blank file, implement and test:
// slice insert/delete, map count/group,
// linked list reverse, stack and queue,
// set operations, sorting + binary search.
~~~

**Explanation:** The snippets are *practice prompts*, not a paste-and-memorize library. If you cannot re-create an implementation in a new file, mark it for revision.

**Your exercise:** Review after 1/3/7/14/30 days and log missed cases.


## Source of truth / verification

These small fragments teach *how to think*. The executable implementation in [`examples/collections.go`](examples/collections.go) and tests in [`examples/collections_test.go`](examples/collections_test.go) demonstrate API contracts and edge conditions. Run `go test ./...` from the examples module; treat any gap as a topic to implement yourself, not as a reason to skip revision.
