# 7. Slice length, capacity, growth, and explicit limits

[All doubts](README.md) · [Zero values and make](02-zero-values-make-and-new.md) · [Bounded buffers](05-buffers-blocking-and-backpressure.md)

## Answers to the three scenarios

Starting with:

```go
var slice1 []int                 // nil; length 0, capacity 0
slice1 = make([]int, 0, 3)       // non-nil; length 0, capacity 3
```

| Scenario | Answer | Reason |
|---|---|---|
| `slice1[0] = 30` | Panics at runtime | Indexing requires `0 <= index < len(slice1)`; capacity does not grant index access |
| Append beyond capacity 3 | Append obtains a larger backing array and returns an updated slice | Growth is not specified as always double or always plus three |
| Keep appending after full | Can continue while resources and representable sizes permit | Capacity is not an application maximum; reject or otherwise handle excess explicitly |

Appending does not grow an array in place as a language-level operation: a slice refers to an array, and an append exceeding capacity obtains a sufficiently large new backing array. Always keep the returned slice, normally with `slice1 = append(slice1, value)`.

## Length and capacity describe different things

- **Length:** number of elements currently in this slice; controls indexing, range, and how many values `copy` can copy into it.
- **Capacity:** number of elements from this slice's start through the available backing-array region; controls how far it can be resliced and whether append can reuse that array.
- **Maximum allowed length:** your application rule; there is no separate built-in slice field for it.

`make([]int, 0, 3)` creates room for three zero-valued integers but exposes none as current slice elements. The slice prints as `[]`, not `[0 0 0]`. Zero-initialized storage and current logical contents are different concepts.

| Construction | Length | Capacity | Valid indexed write now? |
|---|---:|---:|---|
| `var s []int` | 0 | 0 | None |
| `make([]int, 0, 3)` | 0 | 3 | None |
| `make([]int, 3)` | 3 | 3 | `s[0]`, `s[1]`, `s[2]` |
| `make([]int, 1, 3)` | 1 | 3 | `s[0]` only |
| `var a [3]int` | 3 | 3 | All three array elements |

## Scenario 1: three ways to write the first element

Choose according to whether the element already logically exists:

```go
// Grow the logical contents by one element.
s := make([]int, 0, 3)
s = append(s, 30)             // [30], len=1, cap=3

// Start with three actual elements, then replace one.
t := make([]int, 3)
t[0] = 30                    // [30 0 0], len=3, cap=3

// Explicitly expose part of already available backing storage.
u := make([]int, 0, 3)
u = u[:1]                    // valid because 1 <= cap(u)
u[0] = 30                    // [30], len=1, cap=3
```

Reslicing past capacity panics; reslicing does not allocate a larger array. Increasing a slice's length by reslicing reveals the existing values in that region. They are initially zero after `make`, but may be old data if the backing array was used previously. Shrinking and extending a slice does not reset its elements.

A related trap: `copy(make([]int, 0, 3), []int{30})` copies zero elements because the destination length is zero. Capacity is not destination length.

## Scenario 2: what happens at the growth boundary

For successive single-element appends:

| Operation | Length | Capacity | Backing-array behavior |
|---|---:|---|---|
| `s := make([]int, 0, 3)` | 0 | 3 | Initial storage |
| Append 10 | 1 | 3 | Reused |
| Append 20 | 2 | 3 | Reused |
| Append 30 | 3 | 3 | Reused; now full |
| Append 40 | 4 | At least 4 | New array; existing elements copied |

Length becomes exactly old length plus the number appended. Capacity may exceed that length to leave room for later appends. The original `3` is neither a permanent limit nor a remembered growth increment.

### Language guarantee versus runtime policy

The contract is enough capacity for the resulting elements, reusing the existing array when sufficient. **Do not write correctness checks that require a particular new capacity.**

As an implementation example, the Go 1.25.1 runtime's `nextslicecap` uses doubling for small capacities when the requested length fits that growth; from a threshold of 256 it transitions toward roughly 1.25x growth. A large multi-element append can request more than doubling would provide. Allocation size-class rounding and element size can further change the capacity observed. The current runtime source uses this same general growth strategy, but it is not a language promise.

For `[]int` on typical 64-bit Go 1.25.1, filling capacity 3 and appending one element commonly yields capacity 6. Learn **why it grows**, not a universal sequence of capacities. Re-run the demonstration on your toolchain and record its version and architecture.

### Why extra capacity helps

If every append allocated exactly one extra slot and copied all old elements, building n elements could require about `1 + 2 + ... + (n-1)` copies: quadratic work. Geometric growth amortizes these copies across many appends; under the usual growth strategy, appending one fixed-size element has amortized O(1) cost, though a growth operation can cost O(n). This is a performance model, not a worst-case latency guarantee in the Go specification.

Preallocation reduces reallocations when an expected count is known. It trades more reserved storage for fewer copies and allocations; overestimating wastes memory. `make` does not itself promise heap placement: escape analysis still matters.

## Aliasing and why append returns a slice

A slice value describes backing storage plus length and capacity. Assigning it to another variable copies that description, not its elements.

```go
s := make([]int, 1, 3)
s[0] = 10
alias := s
s = append(s, 20)  // same array; alias still has length 1
s[0] = 99
fmt.Println(alias[0]) // 99
fmt.Println(len(alias), len(s)) // 1 2

s = append(s, 30)  // len=3, cap=3
s = append(s, 40)  // exceeds capacity: new backing array
s[0] = 7
fmt.Println(alias[0], s[0]) // 99 7
```

Old aliases remain valid and can keep the old array alive. They do not follow the new slice to its new array. Appending in a function also does not automatically update a caller's slice header; return and assign the result. Mutating existing shared elements is a different operation and can be visible to the caller.

Concurrent append to a shared slice requires synchronization. Spare capacity does not make concurrent writes safe.

## Scenario 3: growth is possible, not infinite

Repeated append can exhaust memory, hit address-space/runtime size limits, overflow representable lengths, or exceed container limits. Some failures terminate the process rather than returning a normal error; trying to recover from out-of-memory is not an admission-control policy. Growing can temporarily retain old and new arrays while copying, and aliases may retain old arrays longer.

A logical item limit also differs from a byte limit: three `[]byte` values can reference three enormous payloads. Account for element size, referenced data, in-flight work, and retained aliases when setting a memory budget.

### Enforce a logical maximum explicitly

```go
// Single-owner helper; shared callers need synchronization around check+append.
func AppendBounded(s []int, max int, values ...int) ([]int, error) {
    if max < 0 || len(s) > max {
        return s, fmt.Errorf("invalid limit %d for length %d", max, len(s))
    }
    if len(values) > max-len(s) {
        return s, fmt.Errorf("maximum length %d exceeded", max)
    }
    return append(s, values...), nil
}
```

The subtraction check avoids overflowing `len(s)+len(values)`. The operation rejects the whole batch before modifying the slice, preserving it on error. At exactly the limit, appending no elements is fine; appending one fails. This caps **logical length**, not necessarily backing capacity or payload bytes. Runtime growth can reserve extra space when starting with a smaller capacity.

For a known small maximum, initialize with `make([]int, 0, max)` and require all growth to pass through the helper; accepted appends then fit the reserved storage. Avoid preallocating huge untrusted limits. For a reusable abstraction, keep the slice private and expose only bounded operations. Returning a raw slice lets callers bypass the policy. This helper is deliberately not a concurrent queue.

On reaching the limit, decide whether to reject, flush a batch, drop according to a documented policy, overwrite an old ring-buffer entry, or wait for a consumer. A slice itself cannot choose that policy.

### Restricting capacity does not prohibit append

```go
s := []int{10, 20, 30}
restricted := s[:len(s):len(s)] // full slice expression: cap equals len
larger := append(restricted, 40) // legal: allocates a new backing array
fmt.Println(len(restricted), len(larger)) // 3 4
```

Full slice expressions limit access to spare capacity, which is useful when passing a view to code that may append. They do not impose an absolute maximum. Existing elements are still shared until growth copies them; this is not an immutable view. `slices.Clip` similarly limits capacity and does not forbid append or necessarily release the underlying allocation.

## Should you switch to an array?

| Need | Suitable choice | Why |
|---|---|---|
| Exactly three components as part of the data model | `[3]int` | Fixed length, length is part of the type |
| Up to three values, with an actual count | Bounded slice API or array plus count | Separates occupied entries from the maximum |
| Unknown result count | `[]T` with append | Convenient dynamic sequence |
| Collect at most N records then send | Slice batch plus flush/limit policy | Fits batching with explicit memory/latency bounds |
| Fixed-memory queue with repeated insert/remove | Ring buffer | Reuses slots and defines full/empty behavior |
| Concurrent bounded handoff | Buffered channel or synchronized queue | Adds coordination beyond storage |

An array has fixed length and cannot be passed directly as the first argument to `append`. But slicing the array does not create a maximum-enforcing container:

```go
var a [3]int
s := a[:]
s = append(s, 40) // legal; new slice array, original a remains length 3
```

Use arrays for truly fixed-size values, not as an automatic replacement for every bounded collection. A fixed-size ring buffer may use an array, but it still needs head/tail/count and overflow rules.

## Where the design is useful

A filter may inspect 100 input items and keep only 12. `make([]Result, 0, len(input))` reserves an upper-bound capacity but starts with no results; appending matches the count of matches. Starting at `make([]Result, len(input))` and appending would leave 100 unwanted initial entries before the matches.

Conversely, if you will produce exactly 100 indexed results, `make([]Result, 100)` is appropriate. A telemetry batch can use length as its current count, capacity to reuse storage, and a separate batch limit to trigger a flush. A generic growing slice is not equivalent to the bounded channel from the producer-consumer guide.

After processing, `s = s[:0]` resets the logical length and retains storage for reuse. It does not erase values or free memory. For pointer-containing elements, clear the used entries before shortening if you want to release their references: `clear(s); s = s[:0]`. Unseen stale references beyond the current length need separate ownership-aware handling. `s = nil` removes this slice's reference, but other aliases may still retain the array and GC timing is not immediate.

## Practice and interview checkpoints

Run from `senior-golang-interview-prep/doubts`:

```sh
go run ./examples/slices
go vet ./...
go build ./...
```

[The runnable example](examples/slices/main.go) isolates and recovers only the deliberate indexing panic, observes growth without asserting exact capacities, shows aliasing and capacity restriction, and checks bounded-append edge cases. Recovery here is for demonstration, not ordinary bounds validation.

Before running, predict:

1. Why does `make([]int, 0, 3)[0]` panic while reslicing to length one enables index zero?
2. What differs between `len`, `cap`, and a business maximum?
3. Which properties of the fourth append are guaranteed, and which are observations?
4. Why does `s[:len(s):len(s)]` cause growth rather than reject it?
5. When does an alias see updates, and when can it retain an older array?
6. When would you choose `make([]T, n)`, `make([]T, 0, n)`, an array, or a bounded queue?

Sources: [index expressions](https://go.dev/ref/spec#Index_expressions), [slice expressions](https://go.dev/ref/spec#Slice_expressions), [append contract](https://go.dev/ref/spec#Appending_and_copying_slices), [Go 1.25.1 slice runtime](https://github.com/golang/go/blob/go1.25.1/src/runtime/slice.go), [current runtime source](https://go.dev/src/runtime/slice.go), [slices.Clip](https://pkg.go.dev/slices#Clip). Growth-complexity discussion and bounded helper are explanatory design examples, not additional language guarantees.
