Yes — **Layer 1 is over only if you can prove it**.

Do not decide by feeling:

> “I think I know basics.”

Decide by a small diagnostic test.

---

# How to know Layer 1 is complete

You are done with Layer 1 if you can do all of these without looking at notes:

## Layer 1 exit test

Can you:

1. Create a Go module?

```bash
go mod init example.com/demo
```

2. Write and run a `main.go`?

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
}
```

3. Use variables?

```go
var x int = 10
y := 20
```

4. Use `if`?

```go
if x > 5 {
    fmt.Println("big")
}
```

5. Use `for`?

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

6. Use `range` over slice?

```go
nums := []int{1, 2, 3}

for i, v := range nums {
    fmt.Println(i, v)
}
```

7. Write a function with return values?

```go
func add(a int, b int) int {
    return a + b
}
```

8. Return an error?

```go
func divide(a float64, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

9. Use map basics?

```go
m := map[string]int{
    "a": 1,
}

m["b"] = 2

v, ok := m["a"]
```

10. Fix common beginner errors?

Examples:

```text
declared and not used
imported and not used
cannot use string as int
undefined variable
```

If yes to almost all, **Layer 1 is over**.

If not, do not fully move on. Use this rule:

```text
80% Layer 2
20% Layer 1 revision
```

Do not restart Layer 1 from zero. Fill gaps while moving forward.

---

# What is Layer 2?

Layer 2 is:

> Building blocks.

Layer 1 was:

> “Can I write basic Go syntax?”

Layer 2 is:

> “Can I represent, store, access, modify, and manipulate data using Go data structures?”

In your original plan, Layer 2 means:

> Make building blocks of concepts with examples.  
> How array is declared and accessed.  
> How linked list is created and accessed.  
> Do this for each data structure.

That is correct.

But we should make Layer 2 more precise.

---

# Layer 2 goal

For every data structure, you should know:

```text
1. What is it?
2. When should I use it?
3. How do I create it in Go?
4. How do I access elements?
5. How do I update elements?
6. How do I delete elements?
7. How do I traverse it?
8. How do I search in it?
9. How do I copy it?
10. How do I compare it?
11. What is its time complexity?
12. What are common traps?
13. How do I test it?
```

This is the Layer 2 template.

Do not just write:

```go
nums := []int{1, 2, 3}
```

You should be able to explain:

```text
What is stored?
How is it accessed?
What happens on append?
What happens on copy?
What is the cost?
What are common mistakes?
```

---

# Layer 2 main topics

Your Layer 2 should cover these data structures and building blocks:

```text
1. Arrays
2. Slices
3. Maps
4. Structs
5. Pointers
6. Methods
7. Strings and runes
8. Singly linked list
9. Doubly linked list
10. Stack
11. Queue
12. Set using map
13. 2D slices / matrix
14. Basic sorting and searching
15. Basic interfaces
```

Let’s go through each.

---

# 1. Arrays

## What to know

Arrays are fixed-size sequences.

## Example

```go
var nums [3]int

nums[0] = 10
nums[1] = 20
nums[2] = 30

fmt.Println(nums)
fmt.Println(nums[0])
fmt.Println(len(nums))
```

Output:

```text
[10 20 30]
10
3
```

Explanation:

- Array length is fixed.
- Index starts at 0.
- `len(nums)` gives size.

---

## Important array behavior

Arrays are value types.

Example:

```go
a := [3]int{1, 2, 3}
b := a

b[0] = 100

fmt.Println(a)
fmt.Println(b)
```

Output:

```text
[1 2 3]
[100 2 3]
```

Explanation:

When you assign an array to another variable, Go copies the whole array.

This is different from slices.

---

## Array comparison

Example:

```go
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}

fmt.Println(a == b)
```

Output:

```text
true
```

Explanation:

Arrays can be compared using `==` if their element type is comparable.

---

## Layer 2 array checklist

You should know:

```text
Declare array
Initialize array
Access element
Update element
Loop over array
Compare arrays
Copy behavior
Fixed size limitation
```

---

# 2. Slices

Slices are the most important collection type in Go.

You must become very comfortable with slices.

---

## Creation

Example:

```go
var s []int              // nil slice
empty := []int{}         // empty slice
nums := []int{1, 2, 3}   // slice literal
made := make([]int, 5)   // length 5, capacity 5
pre := make([]int, 0, 10) // length 0, capacity 10
```

Explanation:

- `nil slice` has no underlying array yet.
- `empty slice` is not nil, but length is 0.
- `make` creates a slice with length and capacity.

---

## Access

Example:

```go
nums := []int{10, 20, 30}

fmt.Println(nums[0])
fmt.Println(nums[1])
fmt.Println(nums[2])
```

Output:

```text
10
20
30
```

---

## Append

Example:

```go
nums := []int{1, 2}
nums = append(nums, 3)

fmt.Println(nums)
```

Output:

```text
[1 2 3]
```

Explanation:

`append` returns the updated slice.

Always assign it back:

```go
nums = append(nums, 3)
```

---

## Length and capacity

Example:

```go
nums := make([]int, 0, 5)

fmt.Println(len(nums))
fmt.Println(cap(nums))
```

Output:

```text
0
5
```

Explanation:

- `len` means current number of elements.
- `cap` means underlying capacity before reallocation is needed.

---

## Slicing

Example:

```go
nums := []int{1, 2, 3, 4, 5}

sub := nums[1:4]

fmt.Println(sub)
```

Output:

```text
[2 3 4]
```

Explanation:

`nums[1:4]` includes indexes:

```text
1, 2, 3
```

It excludes index 4.

---

## Important slice trap

Example:

```go
nums := []int{1, 2, 3, 4}
sub := nums[0:2]

sub = append(sub, 100)

fmt.Println(nums)
fmt.Println(sub)
```

Possible output:

```text
[1 2 100 4]
[1 2 100]
```

Explanation:

`sub` shares the same underlying array as `nums`.

Appending to `sub` can modify `nums` if capacity is shared.

This is a very important Layer 2/3 concept.

---

## Copying slices

Example:

```go
src := []int{1, 2, 3}
dst := make([]int, len(src))

copy(dst, src)

dst[0] = 100

fmt.Println(src)
fmt.Println(dst)
```

Output:

```text
[1 2 3]
[100 2 3]
```

Explanation:

`copy` creates independent values in the destination slice.

---

## Slice equality

Example:

```go
import "slices"

a := []int{1, 2, 3}
b := []int{1, 2, 3}

fmt.Println(slices.Equal(a, b))
```

Output:

```text
true
```

Explanation:

You cannot compare slices using `==`, except comparing to `nil`.

This is invalid:

```go
if a == b {
}
```

Use:

```go
slices.Equal(a, b)
```

---

## Layer 2 slice checklist

You should know:

```text
Create slice
Append
Access
Update
Slice from another slice
Copy
Length
Capacity
Range
Sort
Search
Nil slice vs empty slice
Slice equality
Shared backing array trap
```

---

# 3. Maps

Maps are key-value stores.

---

## Creation

Example:

```go
ages := map[string]int{}

ages["alice"] = 30
ages["bob"] = 25
```

Or:

```go
ages := map[string]int{
    "alice": 30,
    "bob":   25,
}
```

---

## Access

Example:

```go
fmt.Println(ages["alice"])
```

Output:

```text
30
```

---

## Safe access

Example:

```go
age, ok := ages["charlie"]

if ok {
    fmt.Println("found:", age)
} else {
    fmt.Println("not found")
}
```

Explanation:

If key is missing, map returns zero value.

Use comma-ok pattern:

```go
value, ok := m[key]
```

---

## Delete

Example:

```go
delete(ages, "bob")
```

Explanation:

Removes key `"bob"` from map.

---

## Iterate

Example:

```go
for name, age := range ages {
    fmt.Println(name, age)
}
```

Important:

Map iteration order is random.

Do not depend on order.

---

## Map size

Example:

```go
fmt.Println(len(ages))
```

---

## Map keys

Map keys must be comparable.

Valid key types:

```text
string
int
float types mostly
bool
arrays
structs if all fields comparable
```

Invalid key types:

```text
slices
maps
functions
```

Example invalid:

```go
m := map[[]int]string{}
```

This will not compile.

---

## Layer 2 map checklist

You should know:

```text
Create map
Set value
Get value
Check existence
Delete key
Iterate map
Map length
Valid key types
Random iteration order
Zero value for missing key
```

---

# 4. Structs

Structs group related data.

---

## Define struct

Example:

```go
type User struct {
    ID    int
    Name  string
    Email string
}
```

---

## Create struct

Example:

```go
u := User{
    ID:    1,
    Name:  "Alice",
    Email: "alice@example.com",
}
```

---

## Access fields

Example:

```go
fmt.Println(u.Name)
```

---

## Update fields

Example:

```go
u.Email = "alice@example.org"
```

---

## Struct zero value

Example:

```go
var u User
```

Values:

```text
u.ID = 0
u.Name = ""
u.Email = ""
```

Explanation:

Struct fields automatically get zero values.

---

## Structs are value types

Example:

```go
u1 := User{ID: 1, Name: "Alice"}
u2 := u1

u2.Name = "Bob"

fmt.Println(u1)
fmt.Println(u2)
```

Output:

```text
{1 Alice }
{1 Bob }
```

Explanation:

Assigning a struct copies it.

---

## Struct comparison

Example:

```go
type Point struct {
    X int
    Y int
}

a := Point{1, 2}
b := Point{1, 2}

fmt.Println(a == b)
```

Output:

```text
true
```

Explanation:

Structs can be compared if all fields are comparable.

---

## Layer 2 struct checklist

You should know:

```text
Define struct
Create struct
Access fields
Update fields
Copy behavior
Zero values
Compare structs
Embed structs
Use structs as map keys if comparable
```

---

# 5. Pointers

Pointers are essential for understanding mutation.

---

## Basic pointer

Example:

```go
x := 10
p := &x

fmt.Println(p)
fmt.Println(*p)
```

Explanation:

- `&x` gives address of `x`
- `*p` gives value at that address

---

## Modify value through pointer

Example:

```go
x := 10
p := &x

*p = 20

fmt.Println(x)
```

Output:

```text
20
```

---

## Nil pointer

Example:

```go
var p *int

fmt.Println(p == nil)
```

Output:

```text
true
```

Do not dereference nil pointer:

```go
fmt.Println(*p)
```

This panics.

---

## Pointers and functions

Example:

```go
func increment(n *int) {
    *n++
}
```

Call:

```go
x := 1
increment(&x)

fmt.Println(x)
```

Output:

```text
2
```

Explanation:

Because the function receives the address, it can modify the original value.

Without pointer:

```go
func increment(n int) {
    n++
}
```

This would not modify original value.

---

## Layer 2 pointer checklist

You should know:

```text
& operator
* operator
nil pointer
pointer to struct
pointer in function arguments
when to use pointer
when not to use pointer
```

---

# 6. Methods

Methods attach behavior to types.

---

## Value receiver

Example:

```go
type Counter struct {
    count int
}

func (c Counter) Value() int {
    return c.count
}
```

Explanation:

The method receives a copy of the struct.

---

## Pointer receiver

Example:

```go
func (c *Counter) Increment() {
    c.count++
}
```

Usage:

```go
counter := Counter{}

counter.Increment()
counter.Increment()

fmt.Println(counter.Value())
```

Output:

```text
2
```

Explanation:

Pointer receiver allows mutation.

---

## When to use pointer receiver

Use pointer receiver when:

```text
1. You need to modify the receiver.
2. The struct is large and copying is expensive.
3. You want consistent method set with pointers.
```

Use value receiver when:

```text
1. Struct is small.
2. You do not need mutation.
3. You want immutable-style behavior.
```

---

## Layer 2 methods checklist

You should know:

```text
Define method
Value receiver
Pointer receiver
Call method
Difference between function and method
When to use pointer receiver
```

---

# 7. Strings and runes

Strings are common but tricky.

---

## Basic string

Example:

```go
s := "hello"

fmt.Println(len(s))
fmt.Println(s[0])
```

Output:

```text
5
104
```

Explanation:

- `len(s)` counts bytes.
- `s[0]` gives byte, not character.

---

## Runes

Example:

```go
s := "héllo"

for i, r := range s {
    fmt.Println(i, r, string(r))
}
```

Explanation:

`range` over string gives runes.

A rune represents a Unicode code point.

---

## String to rune slice

Example:

```go
s := "hello"
runes := []rune(s)

fmt.Println(runes[0])
```

---

## Rune slice to string

Example:

```go
runes := []rune{'g', 'o'}
s := string(runes)

fmt.Println(s)
```

Output:

```text
go
```

---

## String immutability

Strings cannot be modified by index.

Invalid mindset:

```go
s[0] = 'H'
```

You cannot do this.

Instead:

```go
s := "hello"
runes := []rune(s)
runes[0] = 'H'
s = string(runes)
```

---

## strings.Builder

For building strings efficiently:

```go
import "strings"

var builder strings.Builder

builder.WriteString("Hello")
builder.WriteString(" ")
builder.WriteString("Go")

fmt.Println(builder.String())
```

Output:

```text
Hello Go
```

Explanation:

Use `strings.Builder` when concatenating many strings.

---

## Layer 2 string checklist

You should know:

```text
String length
Byte vs rune
Range over string
String to []byte
String to []rune
String concatenation
strings.Builder
strings.Split
strings.Join
strings.Contains
strings.ToUpper
strings.ToLower
```

---

# 8. Singly linked list

Linked lists are important for pointer practice.

In production Go, you often use slices, but linked lists are useful for interviews and understanding pointers.

---

## Node structure

Example:

```go
type Node struct {
    Value int
    Next  *Node
}
```

Explanation:

Each node contains:

```text
Value
Pointer to next node
```

---

## Create list manually

Example:

```go
head := &Node{Value: 1}
second := &Node{Value: 2}
third := &Node{Value: 3}

head.Next = second
second.Next = third
```

List:

```text
1 -> 2 -> 3 -> nil
```

---

## Traverse list

Example:

```go
func traverse(head *Node) {
    curr := head

    for curr != nil {
        fmt.Println(curr.Value)
        curr = curr.Next
    }
}
```

Explanation:

Start at head.

Move to next until `nil`.

---

## Insert at front

Example:

```go
func insertFront(head *Node, value int) *Node {
    newNode := &Node{Value: value}
    newNode.Next = head
    return newNode
}
```

Usage:

```go
head = insertFront(head, 0)
```

Explanation:

New node points to old head.

New head becomes new node.

---

## Insert at end

Example:

```go
func insertEnd(head *Node, value int) *Node {
    newNode := &Node{Value: value}

    if head == nil {
        return newNode
    }

    curr := head

    for curr.Next != nil {
        curr = curr.Next
    }

    curr.Next = newNode

    return head
}
```

Explanation:

Walk to last node, then attach new node.

---

## Search

Example:

```go
func contains(head *Node, target int) bool {
    curr := head

    for curr != nil {
        if curr.Value == target {
            return true
        }
        curr = curr.Next
    }

    return false
}
```

---

## Delete by value

Example:

```go
func deleteValue(head *Node, target int) *Node {
    dummy := &Node{Next: head}
    curr := dummy

    for curr.Next != nil {
        if curr.Next.Value == target {
            curr.Next = curr.Next.Next
            break
        }
        curr = curr.Next
    }

    return dummy.Next
}
```

Explanation:

Dummy node simplifies deletion, especially deleting head.

---

## Layer 2 linked list checklist

You should know:

```text
Define Node
Create list
Traverse list
Insert front
Insert end
Search
Delete
Use dummy node
Handle nil head
```

---

# 9. Doubly linked list

This is optional but useful.

---

## Node structure

Example:

```go
type DNode struct {
    Value int
    Prev  *DNode
    Next  *DNode
}
```

Explanation:

Each node points to previous and next.

---

## Simple doubly linked list

Example:

```go
type DoublyList struct {
    Head *DNode
    Tail *DNode
}
```

---

## Insert at front

Example:

```go
func (l *DoublyList) InsertFront(value int) {
    newNode := &DNode{Value: value}

    if l.Head == nil {
        l.Head = newNode
        l.Tail = newNode
        return
    }

    newNode.Next = l.Head
    l.Head.Prev = newNode
    l.Head = newNode
}
```

---

## Insert at end

Example:

```go
func (l *DoublyList) InsertEnd(value int) {
    newNode := &DNode{Value: value}

    if l.Tail == nil {
        l.Head = newNode
        l.Tail = newNode
        return
    }

    newNode.Prev = l.Tail
    l.Tail.Next = newNode
    l.Tail = newNode
}
```

---

## When to use doubly linked list

Use when you need:

```text
Move backward
Remove known node quickly
Implement LRU cache
Maintain head and tail efficiently
```

---

# 10. Stack

A stack is LIFO:

```text
Last In, First Out
```

Use slice-based stack for Layer 2.

---

## Stack implementation

Example:

```go
type Stack struct {
    items []int
}
```

---

## Push

Example:

```go
func (s *Stack) Push(v int) {
    s.items = append(s.items, v)
}
```

---

## Pop

Example:

```go
func (s *Stack) Pop() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }

    last := len(s.items) - 1
    value := s.items[last]
    s.items = s.items[:last]

    return value, true
}
```

---

## Peek

Example:

```go
func (s *Stack) Peek() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }

    return s.items[len(s.items)-1], true
}
```

---

## IsEmpty

Example:

```go
func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}
```

---

## Usage

Example:

```go
stack := Stack{}

stack.Push(1)
stack.Push(2)
stack.Push(3)

v, ok := stack.Pop()
fmt.Println(v, ok)
```

Output:

```text
3 true
```

---

# 11. Queue

A queue is FIFO:

```text
First In, First Out
```

---

## Simple slice queue

Example:

```go
type Queue struct {
    items []int
}
```

---

## Enqueue

Example:

```go
func (q *Queue) Enqueue(v int) {
    q.items = append(q.items, v)
}
```

---

## Dequeue

Example:

```go
func (q *Queue) Dequeue() (int, bool) {
    if len(q.items) == 0 {
        return 0, false
    }

    value := q.items[0]
    q.items = q.items[1:]

    return value, true
}
```

Explanation:

Removes first element.

Note:

This simple implementation may not be efficient for large queues because removing from front can cause copying.

For Layer 2, this is fine.

Later you can learn ring buffer.

---

## Usage

Example:

```go
queue := Queue{}

queue.Enqueue(1)
queue.Enqueue(2)
queue.Enqueue(3)

v, ok := queue.Dequeue()
fmt.Println(v, ok)
```

Output:

```text
1 true
```

---

# 12. Set using map

Go does not have a built-in `Set`.

Use:

```go
map[T]struct{}
```

---

## Why `struct{}`?

Because:

```go
struct{}{}
```

takes zero bytes.

---

## Set implementation

Example:

```go
type Set struct {
    items map[string]struct{}
}
```

---

## Create set

Example:

```go
func NewSet() *Set {
    return &Set{
        items: map[string]struct{}{},
    }
}
```

---

## Add

Example:

```go
func (s *Set) Add(item string) {
    s.items[item] = struct{}{}
}
```

---

## Contains

Example:

```go
func (s *Set) Contains(item string) bool {
    _, ok := s.items[item]
    return ok
}
```

---

## Delete

Example:

```go
func (s *Set) Delete(item string) {
    delete(s.items, item)
}
```

---

## Usage

Example:

```go
set := NewSet()

set.Add("go")
set.Add("rust")

fmt.Println(set.Contains("go"))
fmt.Println(set.Contains("java"))
```

Output:

```text
true
false
```

---

# 13. 2D slices / matrix

Many interview problems use matrices.

---

## Create 2D slice

Example:

```go
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
```

---

## Access

Example:

```go
fmt.Println(matrix[0][0]) // 1
fmt.Println(matrix[1][2]) // 6
```

Explanation:

First index is row.

Second index is column.

---

## Update

Example:

```go
matrix[1][1] = 100
```

---

## Traverse

Example:

```go
for i := range matrix {
    for j := range matrix[i] {
        fmt.Println(i, j, matrix[i][j])
    }
}
```

---

## Create empty matrix

Example:

```go
rows := 3
cols := 4

grid := make([][]int, rows)

for i := range grid {
    grid[i] = make([]int, cols)
}
```

Explanation:

Go does not have native 2D arrays like some languages.

Usually we use slice of slices.

---

# 14. Basic sorting and searching

Layer 2 should not implement every sorting algorithm yet.

But you should know how to sort and search common structures.

---

## Sort slice

Example:

```go
import "slices"

nums := []int{3, 1, 2}

slices.Sort(nums)

fmt.Println(nums)
```

Output:

```text
[1 2 3]
```

---

## Sort strings

Example:

```go
words := []string{"banana", "apple", "cherry"}

slices.Sort(words)

fmt.Println(words)
```

Output:

```text
[apple banana cherry]
```

---

## Custom sort

Example:

```go
import (
    "cmp"
    "slices"
)

type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 30},
    {"Bob", 25},
}

slices.SortFunc(people, func(a, b Person) int {
    return cmp.Compare(a.Age, b.Age)
})
```

Explanation:

Sort people by age.

---

## Binary search

Example:

```go
nums := []int{1, 2, 3, 4, 5}

idx, found := slices.BinarySearch(nums, 3)

fmt.Println(idx, found)
```

Output:

```text
2 true
```

Explanation:

Binary search works only on sorted data.

---

# 15. Basic interfaces

Interfaces are not only data structures, but they are important building blocks.

For Layer 2, keep interfaces simple.

---

## Define interface

Example:

```go
type Shape interface {
    Area() float64
}
```

---

## Implement interface

Example:

```go
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

Usage:

```go
func printArea(s Shape) {
    fmt.Println(s.Area())
}
```

Call:

```go
r := Rectangle{Width: 3, Height: 4}
printArea(r)
```

Output:

```text
12
```

Explanation:

Go interfaces are implemented implicitly.

There is no `implements` keyword.

---

# Layer 2 operations matrix

For each data structure, create this table in your notes.

Example:

```text
Data structure: slice

Creation:
    []int{}
    make([]int, 0)

Access:
    s[i]

Update:
    s[i] = x

Insert:
    append(s, x)

Delete:
    append(s[:i], s[i+1:]...)

Traverse:
    for i, v := range s

Search:
    slices.Index
    manual loop

Sort:
    slices.Sort

Copy:
    copy(dst, src)

Equality:
    slices.Equal

Complexity:
    access O(1)
    append amortized O(1)
    delete middle O(n)

Traps:
    shared backing array
    nil vs empty
    cannot compare with ==
```

Do this for every Layer 2 structure.

---

# Layer 2 study material structure

For each data structure, create one file.

Example:

```text
layer2/
  arrays.md
  slices.md
  maps.md
  structs.md
  pointers.md
  methods.md
  strings.md
  linked-list.md
  doubly-linked-list.md
  stack.md
  queue.md
  set.md
  matrix.md
  sorting-searching.md
  interfaces.md
```

Inside each file:

```text
What it is
When to use it
Go syntax
Creation
Access
Update
Delete
Traversal
Search
Copy
Equality
Complexity
Common mistakes
Example code
Small exercise
Flashcards
```

---

# Layer 2 exercises

These exercises will make Layer 2 stick.

---

## Exercise 1: Slice deduplicate

Task:

Remove duplicates from slice.

Example:

```go
func dedupe(nums []int) []int {
    seen := map[int]struct{}{}
    result := make([]int, 0, len(nums))

    for _, n := range nums {
        if _, ok := seen[n]; !ok {
            seen[n] = struct{}{}
            result = append(result, n)
        }
    }

    return result
}
```

---

## Exercise 2: Map word frequency

Task:

Count words.

Example:

```go
func frequency(words []string) map[string]int {
    counts := map[string]int{}

    for _, word := range words {
        counts[word]++
    }

    return counts
}
```

---

## Exercise 3: Reverse slice

Task:

Reverse slice in place.

Example:

```go
func reverse(nums []int) {
    left := 0
    right := len(nums) - 1

    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}
```

---

## Exercise 4: Reverse linked list

Task:

Reverse singly linked list.

Example:

```go
func reverse(head *Node) *Node {
    var prev *Node
    curr := head

    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    return prev
}
```

This is more Layer 4, but it is very useful.

---

## Exercise 5: Stack valid parentheses

Task:

Check if parentheses are balanced.

Example:

```go
func isValid(s string) bool {
    stack := []rune{}

    pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, ch := range s {
        switch ch {
        case '(', '{', '[':
            stack = append(stack, ch)
        case ')', '}', ']':
            if len(stack) == 0 {
                return false
            }

            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            if top != pairs[ch] {
                return false
            }
        }
    }

    return len(stack) == 0
}
```

---

## Exercise 6: Queue using slice

Task:

Implement enqueue/dequeue.

Already shown above.

---

## Exercise 7: Set operations

Task:

Implement union and intersection.

Example intersection:

```go
func intersection(a, b []int) []int {
    seen := map[int]struct{}{}

    for _, v := range a {
        seen[v] = struct{}{}
    }

    result := []int{}

    for _, v := range b {
        if _, ok := seen[v]; ok {
            result = append(result, v)
            delete(seen, v)
        }
    }

    return result
}
```

---

# Layer 2 should not include too much yet

Avoid going deep into:

```text
goroutines
channels
context
mutex
race conditions
GC internals
pprof
system design
advanced LeetCode
dynamic programming
graph algorithms
HLD/LLD
```

Layer 2 is about data structure fluency.

You can do easy LeetCode problems, but do not jump into hard patterns yet.

---

# Suggested Layer 2 schedule

If you have 4 weeks:

## Week 1

```text
Arrays
Slices
Maps
Strings/runes
```

Deliverables:

```text
slice snippet library
map snippet library
string utility examples
```

---

## Week 2

```text
Structs
Pointers
Methods
Basic interfaces
```

Deliverables:

```text
struct examples
pointer mutation examples
value receiver vs pointer receiver examples
small interface example
```

---

## Week 3

```text
Singly linked list
Doubly linked list
Stack
Queue
Set
```

Deliverables:

```text
linked list implementation
stack implementation
queue implementation
set implementation
```

---

## Week 4

```text
2D slices
Sorting
Searching
Revision
Timed recall
Small exercises
```

Deliverables:

```text
matrix traversal examples
sorting examples
binary search examples
Layer 2 checklist completed
```

---

# Layer 2 definition of done

Layer 2 is complete when you can do these without notes:

```text
1. Create and use arrays.
2. Create, append, copy, slice, sort, search slices.
3. Explain slice length and capacity.
4. Explain slice sharing trap.
5. Create, read, update, delete maps.
6. Use comma-ok with maps.
7. Explain map iteration randomness.
8. Define and use structs.
9. Explain struct copying.
10. Use pointers safely.
11. Explain value receiver vs pointer receiver.
12. Implement singly linked list operations.
13. Implement stack.
14. Implement queue.
15. Implement set using map.
16. Traverse 2D slice.
17. Use basic sorting and searching.
18. Write small tests for these structures.
```

---

# Most important Layer 2 advice

Layer 2 should not be passive reading.

For every data structure, you must write:

```text
Creation code
Access code
Update code
Deletion code
Traversal code
Search code
Test code
```

If you only read about slices, you will forget slices.

If you write:

```text
slice examples
slice bugs
slice tests
slice flashcards
```

then slices become permanent.

So yes: if you can pass the Layer 1 exit test, Layer 1 is over.

Now move to Layer 2:

> Data structures and building blocks in Go.