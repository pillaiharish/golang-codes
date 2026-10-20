# Layer 2 — Go data structures and reusable building blocks

[Prep home](../README.md) · [Mandate](../STUDY-MANDATE.md) · [Original detailed course](reference.md) · [Worked examples](WORKED-EXAMPLES.md)

**Goal:** From a blank file, represent data, perform meaningful operations, test corner cases, and justify complexity + copy/alias semantics. Layer 1 teaches syntax; Layer 2 teaches what to *build* with it.

## Study modules (ordered)

| Topic | Required outcome | Recalled at 7 days? |
|---|---|---|
| [01. Arrays](WORKED-EXAMPLES.md#01-arrays) | Fixed-length value collections | ☐ |
| [02. Slices](WORKED-EXAMPLES.md#02-slices) | Resizable views over arrays | ☐ |
| [03. Maps](WORKED-EXAMPLES.md#03-maps) | Key/value lookups | ☐ |
| [04. Structs](WORKED-EXAMPLES.md#04-structs) | Named fields and value composition | ☐ |
| [05. Pointers and methods](WORKED-EXAMPLES.md#05-pointers-and-methods) | Addresses and receiver behavior | ☐ |
| [06. Strings, bytes and runes](WORKED-EXAMPLES.md#06-strings-bytes-and-runes) | UTF-8 and immutable strings | ☐ |
| [07. Singly linked list](WORKED-EXAMPLES.md#07-singly-linked-list) | Pointer-linked nodes | ☐ |
| [08. Doubly linked list](WORKED-EXAMPLES.md#08-doubly-linked-list) | Previous/next pointers and head/tail | ☐ |
| [09. Stack (LIFO)](WORKED-EXAMPLES.md#09-stack-lifo) | Last-in, first-out | ☐ |
| [10. Queue and deque (FIFO)](WORKED-EXAMPLES.md#10-queue-and-deque-fifo) | First-in, first-out and two-ended access | ☐ |
| [11. Set using map](WORKED-EXAMPLES.md#11-set-using-map) | Unique membership | ☐ |
| [12. Multiset / frequency map](WORKED-EXAMPLES.md#12-multiset-frequency-map) | Counting occurrences | ☐ |
| [13. 2D slices and matrices](WORKED-EXAMPLES.md#13-2d-slices-and-matrices) | Rows, columns and jagged input | ☐ |
| [14. Sorting and searching](WORKED-EXAMPLES.md#14-sorting-and-searching) | Ordered comparisons and lookups | ☐ |
| [15. Interfaces as building blocks](WORKED-EXAMPLES.md#15-interfaces-as-building-blocks) | Implicit contracts for behavior | ☐ |
| [16. Testing structures](WORKED-EXAMPLES.md#16-testing-structures) | Verify behavior and edges | ☐ |
| [17. Complexity and ownership](WORKED-EXAMPLES.md#17-complexity-and-ownership) | Select the right representation | ☐ |
| [18. Snippet library and retention](WORKED-EXAMPLES.md#18-snippet-library-and-retention) | Reproduce patterns from memory | ☐ |

## Practical operations mandate

Apply operations **where they make sense**, rather than forcing sorting or “move middle” onto a map.

| Structure | Core operations to implement |
|---|---|
| Array | create/index/update/range/copy/compare |
| Slice | create/append/insert/remove/search/copy/compare/sort/reverse |
| Map | create/lookup comma-ok/update/delete/iterate/count/group/clone/sorted keys |
| Struct/pointer/method | construct/read/update/copy/embed/value-vs-pointer receiver |
| String | index byte/range rune/convert/compare/split/join/reverse by rune |
| Singly linked list | prepend/append/traverse/find/insert/remove/reverse |
| Doubly linked list | push front/back/remove known node/move front/back/traverse both ways |
| Stack/queue/deque | push/pop/peek and enqueue/dequeue; empty cases and retention |
| Set/multiset | membership/union/intersection/difference and frequency add/remove |
| Matrix | construct/access/update/jagged traverse/search/transpose |
| Sort/search | comparison sort, custom comparator, linear vs binary search |
| Interfaces | one small behavior contract and two concrete implementations |

## How to work through this layer

1. Read [one worked example](WORKED-EXAMPLES.md) and predict the result before looking at its explanation.
2. Type it into a new file with correct package/imports, and run it.
3. Build the operation list above with **no notes**; test an empty and a non-empty case.
4. Compare to the reference, name one edge case and analyze big-O.
5. Mark the topic “retrievable” only after re-coding it at day 7 and day 30.

The original [Layer 2 course](reference.md) is preserved and available for additional narrative. For an executable reference, use the scoped [examples module](examples/README.md), not standalone fragments that omit package/import statements.

**Run verification:**

~~~sh
cd god-is-great-prep/layer2/examples
go test ./...
go vet ./...
~~~

## Layer 2 vs later layers

- **Layer 3:** aliasing gotchas, nil-interface, method sets, tricky range/defer behavior, races.
- **Layer 4:** fluency drill for every insert/delete/move/reverse operation and production-ready variants.
- **Layers 5+:** goroutines, GC, memory profiling, rate limiters, HTTP, Kafka, distributed systems, DSA and design.

**Exit check:** implement a singly linked list, a LIFO stack, a FIFO queue, a set, a frequency map, and sorted search without reference; test boundary cases; explain costs and ownership.
