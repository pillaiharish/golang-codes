# 2. Zero values, nil, make, and new

[All doubts](README.md) · [Previous](01-declarations-assignment-and-memory.md) · [Next: text](03-bytes-runes-and-utf8.md)

## Direct answer

**Every Go variable has an initialized value**, including composite and reference-containing types. If no initializer is supplied, it receives its type's zero value. The distinction is whether that zero value supports the operation you want. `make` is not required for every “non-primitive” type: it is a built-in specifically for slices, maps, and channels.

## Zero-value operations

| Type | Zero value | Works immediately | Needs attention |
|---|---|---|---|
| Integers, floats, complex numbers | Numeric zero | Arithmetic | Division and overflow rules still apply |
| `bool` | `false` | Conditions | — |
| `string` | `""` | Length, concatenation, range | Indexing empty string panics |
| Array | Elements recursively zeroed | Index within fixed length | No `make` needed |
| Struct | Fields recursively zeroed | Field reads/writes | Type-specific invariants may need a constructor |
| Pointer | `nil` | Compare with nil | Dereference panics |
| Slice | `nil`, len/cap zero | `len`, `range`, `append`, `copy` | Indexing needs sufficient length |
| Map | `nil`, len zero | Lookup, range, `delete` | Assigning an entry panics |
| Channel | `nil` | Compare, len/cap | Send/receive block indefinitely; close panics |
| Function | `nil` | Compare with nil | Calling panics |
| Interface | `nil` | Compare with nil | No dynamic value/method to invoke |

Zero-value usability also depends on an API's contract: a `sync.Mutex` and `bytes.Buffer` are useful without constructors; a user-defined struct with required configuration may not be.

## Slice: nil is already useful

```go
var s []int
fmt.Println(s == nil, len(s), cap(s)) // true 0 0
s = append(s, 7)                    // valid; obtains backing storage as needed
fmt.Println(s)                     // [7]

slots := make([]int, 3)       // length 3, elements [0 0 0]
reserved := make([]int, 0, 3) // length 0, capacity 3
slots[0] = 7                 // valid
// reserved[0] = 7           // PANIC: capacity is not length
reserved = append(reserved, 7)
```

`[]int{}` and `make([]int, 0)` are empty non-nil slices. A nil slice and an empty non-nil slice have the same length, but APIs may distinguish them (for example default JSON encoding produces `null` versus `[]`). Do not rely on a non-nil empty slice to have a unique heap allocation.

## Map: nil reads and initialized writes

```go
var counts map[string]int
value, ok := counts["go"]
fmt.Println(value, ok) // 0 false
// counts["go"]++     // PANIC: assignment to entry in nil map
counts = make(map[string]int, 100)
counts["go"]++        // missing entry reads as 0, then writes 1
```

The `100` is a sizing hint, not a hard capacity or initial number of entries. `len(counts)` starts at zero. `map[string]int{}` also creates a writable map. Neither initialization nor a sizing hint makes concurrent writes safe.

## Channel: nil differs from unbuffered

```go
var disabled chan int        // nil, no usable channel
ready := make(chan int)      // non-nil, unbuffered rendezvous
queue := make(chan int, 3)   // non-nil, capacity 3, length 0
_, _, _ = disabled, ready, queue
```

`make(chan int, 3)` does **not** put three zero values in the queue. Capacity allocates room; only sends enqueue values. A nil channel case is disabled in `select`; with no ready cases and no `default`, the select waits. See [queue semantics](05-buffers-blocking-and-backpressure.md).

## new versus make

| Expression | Result type | Result value / usable state |
|---|---|---|
| `new(int)` | `*int` | Pointer to a zero-valued int |
| `new([]int)` | `*[]int` | Pointer to a nil slice |
| `make([]int, 3)` | `[]int` | Slice of three zero-valued ints |
| `new(map[string]int)` | `*map[string]int` | Pointer to a nil map, not yet writable |
| `make(map[string]int)` | `map[string]int` | Writable empty map |
| `new(chan int)` | `*chan int` | Pointer to a nil channel |
| `make(chan int)` | `chan int` | Usable unbuffered channel |

These examples use the traditional `new(T)` form. Allocation syntax does not promise heap placement. `make` returns a value of the requested slice/map/channel type, not a pointer to that value. It sets up the state needed by that type. `new(T)` gives a pointer to a zero value; using `new(map[K]V)` usually adds indirection without solving initialization.

## Typed nil inside an interface

```go
var p *int
var x any = p
fmt.Println(p == nil) // true
fmt.Println(x == nil) // false: dynamic type is *int, dynamic value is nil
```

This explains the classic typed-nil `error` trap too. A nil pointer placed into an interface is not a nil interface. A method with a nil pointer receiver can choose to handle nil, but that does not make dereferencing nil safe.

## Interview checkpoints

- Why does appending to a nil slice work while writing to a nil map panics?
- Why are channel capacity 3 and channel length 3 different?
- Can `new` replace `make` for a map? What additional step would be necessary?
- Why is a zero-valued mutex useful while a zero-valued channel cannot communicate?

Run [the values example](examples/values/main.go) with `go run ./examples/values` from this directory.

Sources: [zero value](https://go.dev/ref/spec#The_zero_value), [making slices, maps and channels](https://go.dev/ref/spec#Making_slices_maps_and_channels), [allocation](https://go.dev/ref/spec#Allocation), [JSON encoding](https://pkg.go.dev/encoding/json#Marshal).
