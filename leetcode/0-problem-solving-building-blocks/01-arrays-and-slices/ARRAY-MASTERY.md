# Arrays — full mastery card (Layer 2 → Layer 3)

[← Existing lesson](README.md) · [Slices](SLICE-MASTERY.md) · [Traps](TRAPS-AND-MEMORY.md) · [Recall](RECALL-DRILLS.md) · [Layer 2](../../../god-is-great-prep/layer2/README.md)

## Quick recall (60 seconds)

```go
a := [3]int{1, 2, 3}
b := a              // whole array is a copied value
b[0] = 9             // a[0] remains 1
s := a[:]            // slice view, shares a's storage
s[0] = 7             // a[0] now 7
```

**Model:** `[N]T` is a fixed-length composite value. Array length is part of the type. An array's zero value is usable without constructors. Arrays are comparable if and only if their element type is comparable.

## A. Identity, creation and default values

| Expression | Meaning |
|---|---|
| `var a [3]int` | `[0 0 0]`; all elements initialized |
| `var b [2]bool` | `[false false]` |
| `a := [3]int{1,2,3}` | fixed array |
| `a := [...]int{1,2,3}` | compiler infers 3 |
| `a := [4]int{2:9}` | `[0 0 9 0]` |
| `a := [2][3]int{}` | fixed 2×3 grid |
| `p := new([3]int)` | pointer to zero-valued array |
| `s := a[:]` | slice view; not a copy |
| `make([3]int)` | **invalid**; make works with slices, maps, channels |

`[3]int` and `[4]int` are distinct types. Array length is permanently fixed; use `[]int` for variable-length data.

## B. Full operation lifecycle

| Operation | Syntax | Time |
|---|---|---|
| Create / zero | `var a [N]T` | O(N) elements |
| Index / update | `a[i]` / `a[i]=v` | O(1) |
| Traverse | `for i,v:=range a` | O(N) |
| Copy value | `b:=a` | O(N) in array size |
| Compare | `a==b` if elements comparable | up to O(N) |
| Search | indexed loop | O(N) |
| Sort | `slices.Sort(a[:])` | O(N log N) |
| Size / capacity | `len(a)` / `cap(a)` both N | O(1) |
| Clear | `clear(a[:])` or assign zero array | O(N) |
| Convert to slice | `a[:]` | O(1), shares storage |
| Insert / remove / append | **N/A**; fixed length | use slice |

Range returns a **copy of the element**; `for _,v:=range a { v++ }` doesn't mutate `a`. Update `a[i]`. When ranging over an array value, the range expression may copy the array itself (unlike a slice view). Avoid claiming a particular stack/heap placement without measurements.

## C. Value semantics and memory: a precise diagram

```text
a: [1 | 2 | 3]     b := a      b: [1 | 2 | 3]
  independent array values

s := a[:]
s: (ptr -> a[0], len=3, cap=3) // element storage shared
s[0]=8   => modifies a, NOT b
```

**Comparability:** `map[[26]int][]string` is valid, useful for lower-case alphabet frequency signatures. `[2][]int` is **not** comparable because slices are noncomparable; it cannot be a map key. Arrays of pointers compare addresses, not pointee contents.

**Shallow nested copy:** `a := [2][]int{{1},{2}}; b:=a; b[0][0]=9` changes `a[0][0]` too. The array was copied, but the **slice elements are descriptors** sharing backing arrays. The same issue applies to structs holding maps or pointers.

**Pass to a function:** `func F(a [3]int)` takes a copied value. `func F(a *[3]int)` can mutate caller storage through a pointer.

## D. Invariants and traps

- Valid indices: `0 <= i < N`. `a[N]` is out of bounds.
- `[3]int` cannot be assigned to `[4]int`.
- Arrays compare only when their element types compare.
- Slicing an array does not copy it.
- A copied array of slices may still alias nested storage.
- Large arrays can be expensive to copy; avoid premature performance conclusions.

## E. Implementation and transfer

[Practice 01: runnable array values](practice-01-array-values/README.md) tests default values, full copies, pointer mutation, composite keys and the nested-slice trap. Trace: `[1 2 3]` passed by value returns `[99 2 3]` while caller remains `[1 2 3]`.

Related existing problems: [#36 Sudoku](../../36.%20Valid%20Sudoku/README.md), [#49 Group Anagrams](../../49.%20Group%20Anagrams/README.md), [#1293 Grid BFS](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md) uses a fixed four-direction array.

## F. Design/optimization (later)

Use fixed arrays when a fixed length belongs to the problem model, dynamic slices otherwise. Profiling, escape analysis and allocations belong to a later optimization pass.

## G. Closed-book recall

Rebuild `[3]int` zero value, full-copy example, `*[3]int` mutation, map key and slice view. Explain why array of slices is not recursively deep copied. Then solve #36's fixed-array indexing from memory. Revisit D1/D3/D7/D14/D30; log only real attempts in [revision log](../../REVISION-LOG.md).

Confidence — · Last cold recall — · Next review —
