# Layer 3 notebook — arrays, slices, aliasing and memory correctness

[Array](ARRAY-MASTERY.md) · [Slice](SLICE-MASTERY.md) · [Runnable trap lab](practice-03-aliasing-and-memory/README.md) · [Recall](RECALL-DRILLS.md)

**Routine:** before running anything, predict output or compile error; explain *why*; run the tested implementation; repair; then recall without notes. Broken examples below are labeled intentionally broken.

## 1. Slice equality is not array equality

**Wrong assumption:** `[]int{1,2} == []int{1,2}` checks element equality.

```go
// INTENTIONALLY DOES NOT COMPILE; do not put in a Go .go file:
a, b := []int{1, 2}, []int{1, 2}
_ = a == b
```

Slices are not comparable except against `nil`. **Fix:** `slices.Equal(a,b)`. By contrast, `[2]int` can be compared with `==` and used as a map key because its elements are comparable.

## 2. Capacity is not an index bound

```go
s := make([]int, 0, 4)
fmt.Println(len(s), cap(s)) // 0 4
// s[0] = 5                  // INTENTIONALLY PANICS: index out of range
s = append(s, 5)            // now len=1
```

**Fix:** use `make([]int,n)` for n addressable elements, or append to grow the length.

## 3. Append can mutate the original subslice parent

```go
base := []int{1, 2, 3, 4}
sub := base[:2]             // len=2, cap=4
sub = append(sub, 9)        // reuse available capacity
fmt.Println(base)           // [1 2 9 4]
```

**Why:** the slice header points to the same array; write at index 2 is in capacity. **Fix to isolate:** `sub=slices.Clone(base[:2])` before append. **Partial fix:** `sub=base[:2:2]` forces allocation *on this append*, but elements before it are still shared.

## 4. An outer clone is not deep nested cloning

```go
matrix := [][]int{{1,2}}
clone := slices.Clone(matrix)
clone[0][0] = 99
fmt.Println(matrix[0][0]) // 99
```

**Fix for [][]int:** allocate a new outer slice and clone each row, as shown in [practice 02](practice-02-slice-operations/main.go). For slices of structs containing pointers, maps or slices, define the required clone depth explicitly.

## 5. Copy into zero-length destination copies zero elements

```go
src := []int{1,2,3}
dst := make([]int,0,3)
fmt.Println(copy(dst,src)) // 0
dst = make([]int,len(src))
fmt.Println(copy(dst,src)) // 3
```

`copy` uses `min(len(dst), len(src))` and permits overlapping source and destination. Capacity is not the copy limit.

## 6. Range loop element is a copy

```go
users := []struct{ Score int }{{1}}
for _, user := range users { user.Score = 99 }
fmt.Println(users[0].Score) // 1
for i := range users { users[i].Score = 99 }
fmt.Println(users[0].Score) // 99
```

Go 1.22 changed **loop variable lifetime/closure capture** for newly declared loop variables; it did **not** turn ranged struct values into references to slice elements. When iterating pointer elements, copied pointers still point to the same objects.

## 7. Nil vs empty slice are different API values

```go
var a []int
b := []int{}
fmt.Println(a == nil, b == nil) // true false
// json.Marshal(a) -> "null"; json.Marshal(b) -> "[]"
```

Both have len 0, can append, and have equal sequences under `slices.Equal`. Preserve distinction when JSON/API semantics require it.

## 8. Slice deletion and pointer memory retention

```go
items := []*int{&a, &b, &c}
copy(items[1:],items[2:])
items[len(items)-1] = nil
items = items[:len(items)-1]
```

**Why:** shifting may leave a stale pointer in backing storage, which can retain an object while the array stays reachable. Clear vacated slots **before shrinking**. Built-in `clear(items)` zeros only accessible elements; it does not change len/cap. `slices.Delete` from Go 1.22 onward zeros obsolete tail slots. Nil-ing a slice descriptor does not immediately free storage held by other aliases.

**Large backing array retention:** `small:=huge[:10]` can retain the whole original array. Clone only the required part if that matters, and do not claim precise GC release times.

## 9. Appending inside a function changes a local descriptor

```go
func mutate(s []int) {
    s[0] = 5         // shared element changes
    s = append(s, 9) // function's local slice header changes
}
```

The caller sees element mutation, but not a new length unless the resulting slice is returned and assigned. Append may overwrite shared spare storage in either scope.

## 10. Array of slices copies headers, not deep data

```go
a := [2][]int{{1},{2}}
b := a
b[0][0] = 7
fmt.Println(a[0][0]) // 7
```

Array assignment copied the array value; its slice-valued elements still reference common storage.

## 11. Common false claims to reject

- "Slice dequeue by `s=s[1:]` is O(n)." **False:** reslicing is O(1); retention is the issue.
- "Full-slice expression performs a deep copy." **False:** existing elements still alias.
- "`clear` releases slice allocation and sets len to zero." **False:** it overwrites accessible elements.
- "Map and slice fields inside copied structs become independent." **False:** descriptors/references are copied.
- "`append` always doubles capacity." **False:** no specified growth factor.
- "Go immediately frees memory when setting `s=nil`." **False:** GC and other references determine reachability.
- "Binary search also works on unsorted slices." **False:** sorted order is a precondition.

## Memory diagram

```text
Array copy:    a=[1|2|3] b=[1|2|3]    independent values
Slice copy:    a=(ptr X,len 3,cap 4)
               b=(ptr X,len 3,cap 4) shared elements in X
Clone []int:   b=(ptr Y,...)         independent integer storage
Clone [][]int: outer array cloned, but inner row backing storage still shared
```

## Blank-editor interview questions

Why do `[]int` and `[3]int` differ in comparability? When is `cap` larger than `len`? What does a function copy when passed a slice? Why is deleted pointer cleanup needed? What does `copy` use as its bound? Explain using code, not only terminology.

Try [trap lab](practice-03-aliasing-and-memory/README.md) and [recall plan](RECALL-DRILLS.md). Confidence and review dates stay blank until a real cold attempt.
