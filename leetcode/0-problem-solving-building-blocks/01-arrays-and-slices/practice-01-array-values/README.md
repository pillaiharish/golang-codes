# Practice 01 — arrays as values

[Array mastery](../ARRAY-MASTERY.md) · [Slice mastery](../SLICE-MASTERY.md) · [Traps](../TRAPS-AND-MEMORY.md)

## Task

From scratch: initialize `[3]int`; copy it and change one element; implement a function taking an array by value, a function taking a pointer to an array, and a map keyed by `[3]int`. Finally copy an array containing slices, and predict whether inner elements can still be shared.

## Progressive explanation and trace

- `ChangeCopy(original [3]int)` receives a **copy**. Changing it does not change caller array.
- `ChangeThroughPointer(original *[3]int)` mutates the caller's array through a pointer.
- `map[[3]int]int` is valid because arrays of comparable ints can be map keys.
- `[2][]int` still aliases nested slice backing arrays after an array-value copy.

```text
original [1 2 3] -> ChangeCopy -> returned [99 2 3]
caller still [1 2 3]
```

**Costs:** array copy O(N), indexing O(1), comparing arrays up to O(N); hashing an array key costs work proportional to the number of elements, even though map lookup is average constant time for a fixed-size key.

## Run

```sh
GO111MODULE=off go run .
GO111MODULE=off go test -race -count=1 -v .
GO111MODULE=off go vet .
```

[Source](main.go) · [Table-driven and edge tests](main_test.go).

## Closed-book challenge

Write the entire program without notes. Explain why array values are distinct but the slice *elements* in a copied array may still reference shared storage. Transfer to [#36 Sudoku](../../../36.%20Valid%20Sudoku/README.md).
