# Executable Layer 2 Go examples

[Layer 2 index](../README.md) · [18 explained examples](../WORKED-EXAMPLES.md) · [Study mandate](../../STUDY-MANDATE.md)

This folder is an **isolated Go module**: it deliberately does not add a `go.mod` to the repository root.

~~~sh
cd god-is-great-prep/layer2/examples
go test ./...
go test -v ./...
go vet ./...
gofmt -w collections.go collections_test.go
~~~

## APIs to reproduce without notes

- Arrays/slices/maps: `ArrayCopy`, `InsertAt`, `DeleteAt`, `WordCounts`, `CloneCounts`
- Structs, pointers, methods, strings/runes: `User`, `Birthday`, `ReverseRunes`
- Singly linked list: `Node`, `Prepend`, `AppendNode`, `DeleteFirst`, `ReverseList`, `ListValues`
- Doubly linked list: `DList` with head/tail, add, remove and move-to-front
- LIFO/FIFO: `IntStack`, `IntQueue` including reset/compaction
- Maps as sets/multisets: `IntSet`, `Multiset`
- Matrices, sort and search: `MatrixSum`, `Transpose`, `SortedSearch`
- Interfaces: `Shape` implemented by `Square` and `Rectangle`

## Deliberate limitations

These are learning examples, **not production-safe concurrent collections**. Linked-list `Remove` and `MoveToFront` require a live node owned by that list; a real library would enforce/document ownership. Slice operations return `false` for invalid indices and never mutate the input. `ReverseRunes` reverses code points, not grapheme clusters. The queue compacts consumed slots to avoid unbounded backing-array retention.

To study: run the tests, close the file, reimplement a function from its name, and add your own edge-case test. Later layers address concurrency, data races and benchmark tradeoffs.
