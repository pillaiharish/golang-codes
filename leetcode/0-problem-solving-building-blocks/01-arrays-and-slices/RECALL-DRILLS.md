# Arrays + slices — repetition until independent recall

[Array mastery](ARRAY-MASTERY.md) · [Slice mastery](SLICE-MASTERY.md) · [Traps](TRAPS-AND-MEMORY.md) · [Existing revision log](../../REVISION-LOG.md)

**Rule:** try a blank editor **before** reopening the notes or code. If you forget a step, read only that subsection, close it, implement it again, then test. Repeating the explanation while looking at it is not cold recall.

## Five mastery levels

| Stage | Evidence |
|---|---|
| Recognize | Choose array vs slice in a new problem |
| Reproduce | Complete operations with the reference and pass tests |
| Recall | Write operations and boundary checks with no reference |
| Transfer | Use the building block in an unfamiliar problem |
| Defend | Explain memory sharing, invariants, complexity and tradeoffs |

## Spaced review roadmap

| Checkpoint | Exercise | Pass gate |
|---|---|---|
| D0 | `var`, literal, `make(len)`, `make(0,cap)`, array vs slice | correct zero values, len/cap, valid indices |
| D1 | Implement middle insert, middle delete, and stack pop | first, middle, last, invalid and empty cases |
| D3 | Copy an array, clone a slice, clone nested [][]int | explain which backing storage is shared |
| D7 | Predict append alias, full-slice cap, zero-length copy, nil JSON | explain each output from code semantics |
| D14 | Solve existing [#238](../../238.%20Product%20of%20Array%20Except%20Self/README.md) or [#11](../../11.%20Container%20With%20Most%20Water/README.md) | reproduce indexed traversal and invariant |
| D30 | Solve [#41](../../41.%20First%20Missing%20Positive/README.md) or a new in-place array problem | boundary correctness + complexity explanation |

Failing a review means retrying sooner, not simply marking the box. Try pointers, nested slices, and zero-length slices rather than memorizing only one demonstration.

## Eight predict-before-running questions

1. `var s []int`: nil? len? cap? append possible?
2. `s:=make([]int,0,4)`: can you index `s[0]`?
3. `a:=[3]int{1,2,3}; b:=a; b[0]=9`: what is `a[0]`?
4. `a:=[]int{1,2,3,4}; b:=a[:2]; b=append(b,9)`: what is `a`?
5. `copy(make([]int,0,3),[]int{1,2})`: how many elements?
6. Why clear the tail after deleting from a slice of pointers?
7. Is `slices.BinarySearch` meaningful before sorting?
8. Does `slices.Clone([][]int)` recursively deep-copy the integers?

### Answers — reveal only after attempting

1. Yes nil, 0, 0, append works.
2. No: indexability requires `i < len(s)`.
3. 1; copied array value.
4. `[1 2 9 4]`; capacity shared.
5. Zero, because destination length is zero.
6. Stale backing references may retain otherwise-unreachable objects.
7. No; sorted ascending is a precondition.
8. No; it clones outer row headers, so each row needs its own clone.

## Test commands

```sh
for dir in leetcode/0-problem-solving-building-blocks/01-arrays-and-slices/practice-*; do
  (cd "$dir" && GO111MODULE=off go test -race -count=1 . && GO111MODULE=off go vet .)
done
```

Use the already solved problem folders. [#36](../../36.%20Valid%20Sudoku/README.md) checks arrays, [#238](../../238.%20Product%20of%20Array%20Except%20Self/README.md) and [#41](../../41.%20First%20Missing%20Positive/README.md) test slice indexing, and [#1293](../../1293.%20Shortest%20Path%20in%20a%20Grid%20with%20Obstacles%20Elimination/README.md) connects to nested slices/queue.

Log **actual** attempts in [REVISION-LOG.md](../../REVISION-LOG.md), not invented progress. Example: `YYYY-MM-DD | Concept: slice-delete | Cold? | Tests? | Exact missed step | Next review`.

Confidence — · Last cold recall — · Next review —
