# Slice mastery — complete lifecycle, ownership and LeetCode usage

[← Existing lesson](README.md) · [Array mastery](ARRAY-MASTERY.md) · [Traps and memory](TRAPS-AND-MEMORY.md) · [Recall drills](RECALL-DRILLS.md) · [Layer 2](../../../god-is-great-prep/layer2/README.md)

## Quick recall (60 seconds)

```go
var a []int                 // nil, len=0, cap=0; usable
b := []int{}                // non-nil empty
c := make([]int, 3)        // len 3, [0 0 0]
d := make([]int, 0, 3)     // len 0, cap 3: cannot index d[0]
d = append(d, 10)          // capture append's result
v := c[1:3]                // shares backing storage
clone := slices.Clone(c)   // independent []int backing
```

**Mental model:** A slice is a descriptor (conceptually a pointer to an array, length, capacity). Assignment and parameter passing copy the **descriptor**, not the elements. **Invariant:** `0 <= len(s) <= cap(s)`; indexability ends at `len(s)`, not `cap(s)`.

## A. Zero values and initialization

| Form | Nil? | Len | Cap | Meaning |
|---|---|---:|---:|---|
| `var s []int` | yes | 0 | 0 | append and range work |
| `s:=[]int{}` | no | 0 | 0 | non-nil empty |
| `s:=make([]int,3)` | no | 3 | 3 | zero-valued elements |
| `s:=make([]int,0,3)` | no | 0 | 3 | reserved but not indexable |
| `s:=[]int{3,1,2}` | no | 3 | 3 | literal |
| `s:=arr[1:3]` | no | 2 | depends on underlying suffix | shared storage |
| `s:=arr[1:3:3]` | no | 2 | 2 | capacity-limited; still shares existing elements |

`make([]int,3)` returns a usable slice value. `new([]int)` returns a pointer to a nil slice descriptor; a pointer isn't needed for most Go collection code. Nil slice marshals to JSON `null`, non-nil empty slice to `[]`; both compare equal using `slices.Equal` (though `nil==empty` is not a permitted slice comparison).

## B. Full operations and time costs (Layer 2)

| Operation | Go technique | Complexity |
|---|---|---|
| Index / update | `s[i]` / `s[i]=x` | O(1) |
| Traverse | `for i,v:=range s` | O(n) |
| Append | `s=append(s,x)` | amortized O(1); may allocate |
| Insert middle | extend then shift right via overlapping `copy` | O(n) |
| Delete middle | `copy(s[i:],s[i+1:])`, clear tail, shorten | O(n) |
| Stack push / pop last | append; read last, zero slot, shorten | amortized O(1) / O(1) |
| Linear lookup | `slices.Index(s,x)` | O(n) |
| Sorted lookup | `slices.BinarySearch(s,x)` | O(log n), sorted prerequisite |
| Sort | `slices.Sort(s)` | O(n log n), mutates |
| Reverse | `slices.Reverse(s)` | O(n), mutates |
| Compare | `slices.Equal(a,b)` | O(n), `a==b` invalid |
| Clone | `slices.Clone(s)` or make + copy | O(n), shallow for nested objects |
| Filter | read/write indices and tail clear | O(n), often in place |
| Stable dedupe | map of seen values | avg O(n), extra O(n) |
| Subslice | `s[i:j]` | O(1), shares backing |
| Clear elements | `clear(s)` | O(len(s)), keeps len/cap |
| Drop this reference | `s=nil` | O(1), doesn't force GC |
| Len and capacity | `len(s)` / `cap(s)` | O(1) |

For circular FIFO behavior, use a queue with head index or a ring buffer, not repeated shifting. See existing [Queue](../09-queues-bfs-and-grids/QUEUE-MASTERY.md) and [Circular deque](../12-circular-buffers-and-deques/README.md) lessons.

## C. Memory-level semantics (Layer 2/3: correctness before optimization)

```text
s := []int{1,2,3,4}
s=(ptr X,len 4,cap 4) -> X:[1|2|3|4]

t:=s[:2]            -> (ptr X,len 2,cap 4)
t[0]=8              -> s[0] becomes 8
t=append(t,9)       -> s[2] becomes 9 (shared spare capacity)

u:=s[:2:2]          -> (ptr X,len 2,cap 2)
u=append(u,9)       -> append allocates another backing array
v:=slices.Clone(s[:2]) -> independent backing for []int
```

- **Copy vs alias:** `t:=s` or passing `[]int` to a function copies the header, not the data. Element updates can affect aliases. Appending inside a function updates its *local slice header*; return and assign it if caller needs new length. If append reuses backing, element effects may still be visible.
- **Three-index full slice expression:** restricting capacity can force reallocation on append, but it does **not** isolate edits to existing elements. Use a clone to isolate.
- **Nested clone:** `slices.Clone([][]int)` duplicates outer slice of row headers, not inner row arrays; clone each row when you need an independent integer grid.
- **Pointer retention:** deleting or filtering pointer elements may leave references in vacated backing slots. Clear old tail before shrinking; `clear` works only on current slice elements.
- **Large subslice retention:** a small view can keep a very large backing array reachable. Clone the necessary portion when storage independence matters.
- **Safety:** concurrent reads without writes are fine. Unsynchronized mutation of shared elements/append storage can race. Ordinary slices do not provide concurrent queue semantics.
- **Cleanup:** `s=nil` only drops this descriptor and does not force immediate collection. Go's GC manages allocation lifetime. Don't confuse memory correctness with later allocation optimization.

## D. Traps / edge cases (Layer 3)

1. Slices are not comparable to each other with `==`; compare to nil or use `slices.Equal`.
2. `make([]int,0,10)` has no indexable elements, only capacity.
3. `copy(dst,src)` moves up to `min(len(dst),len(src))`, not destination cap.
4. `for _,v:=range s` gets a copy of each value; mutating `v` doesn't change structs in `s`.
5. Subslice + append can overwrite original data.
6. Deleting by reslicing alone may retain references in backing storage.
7. `clear(s)` zeros elements but does **not** change len/cap.
8. Sorting a slice mutates it and potentially other aliases.
9. `slices.Clone` is not recursively deep for nested reference types.
10. Go does not promise a fixed slice growth factor; don't assume capacity doubles.
11. Binary search requires sorted input.
12. `slices.Delete` in Go 1.22+ clears obsolete tail elements; a manual delete must clear its tail if referenced elements can be retained.

All proof snippets and fixes are in [Traps and memory](TRAPS-AND-MEMORY.md) and the [runnable trap lab](practice-03-aliasing-and-memory/README.md).

## E. Executable progressive practice

| Stage | Files and tests | Target |
|---|---|---|
| 1 | [Array values](practice-01-array-values/README.md) | copying and comparable keys |
| 2 | [Slice operations](practice-02-slice-operations/README.md) | insert/delete/pop/filter/dedupe/sort/reverse/search/2D clone |
| 3 | [Aliasing/memory](practice-03-aliasing-and-memory/README.md) | append aliasing, full-slice capacity, nil JSON, range copying, tail clearing |

**Important:** The existing numbered lesson `README.md` still has *exactly one* complete Go example, as required by `verify-examples.py`. These are independent `package main` subfolders.

```sh
for dir in leetcode/0-problem-solving-building-blocks/01-arrays-and-slices/practice-*; do
  (cd "$dir" && GO111MODULE=off go test -race -count=1 . && GO111MODULE=off go vet .)
done
```

## F. LeetCode problem map (link, never duplicate solutions)

| Recognition clue | Slice building block | Existing folder |
|---|---|---|
| all except element at i | prefix/suffix passes with allocated output | [#238](../../238.%20Product%20of%20Array%20Except%20Self/README.md) |
| place each x at index x-1 | indexed swaps and mutation | [#41](../../41.%20First%20Missing%20Positive/README.md) |
| two positions define container area | left/right indices | [#11](../../11.%20Container%20With%20Most%20Water/README.md) |
| grid shortest path with state | 2D/3D slice initialization + queue | [#1293](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md) |
| fixed 9×9 membership | array rather than dynamic slice | [#36](../../36.%20Valid%20Sudoku/README.md) |
| ordered or rotated search | binary-search invariants and bounds | [#153](../../153.%20Find%20Minimum%20in%20Rotated%20Sorted%20Array/README.md) |

## G. Go design and optimization (intentionally later)

Wrap a slice in a struct only when an API/invariant benefits from it. Go methods/interfaces and stack/queue design are later depth passes. Profiling allocations, `pprof`, `sync.Pool`, cache behavior and benchmarks are **not** Layer 2 exit requirements.

## H. Cold recall

After [recall drills](RECALL-DRILLS.md), you should be able to build inserts/deletes with tail cleanup, explain len vs cap, predict both append behaviors, clone a nested matrix and solve #238 or #41 from scratch.

Confidence — · Last cold recall — · Next review —
