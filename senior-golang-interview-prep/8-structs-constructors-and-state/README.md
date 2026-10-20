[← Senior Go preparation index](../README.md) · [Practice code](examples/) · [Personal notes](NOTES.md) · [Revision log](../REVISION-LOG.md)

# 8 — Structs, constructors, object creation, and persistent state

**Goal:** Given a blank `main.go`, create an instance with its own state; initialize maps and other fields; call methods correctly; predict copying/receiver behavior; and explain how this design changes under concurrent use. This is foundational Go object construction, not a memorization exercise.

**Prerequisites:** `var`, `:=`, `map[K]V`, `make`, pointers, functions. For shorter refreshers, see [declarations](../doubts/01-declarations-assignment-and-memory.md), [zero values, `make`, and `new`](../doubts/02-zero-values-make-and-new.md), and [Layer 2 structs/pointers](../../god-is-great-prep/layer2/README.md).

## 8.1 A complete Counter, from the first line

Corrected from a handwritten, from-scratch standalone example (Go syntax needs `struct {`, the package/import lines separated, and `fmt.Println`, not `fmt.PrintLn`):

```go
package main

import "fmt"

type Counter struct {
    counts map[string]int
}

func NewCounter() Counter {
    return Counter{counts: make(map[string]int)}
}

func (c *Counter) Increment(key string) {
    c.counts[key]++
}

func (c *Counter) GetKey(key string) int {
    return c.counts[key]
}

func main() {
    ct := NewCounter()
    ct.Increment("a")
    ct.Increment("a")
    ct.Increment("a")
    ct.Increment("a")
    fmt.Println(ct.GetKey("a")) // 4
}
```

The identical working version is split into [counter.go](examples/counter.go) and [main.go](examples/main.go), with [tests](examples/counter_test.go). Run it from `examples/` on macOS/Linux using `go run .`.

**The four pieces have distinct roles:**

1. `type Counter struct { counts map[string]int }` defines a **type and its fields**. It does not instantiate a map.
2. `NewCounter()` is an **ordinary Go function** returning a Counter whose `counts` field refers to a writable empty map.
3. `func (c *Counter) Increment(...)` is a **method**: its receiver chooses the Counter instance to operate on.
4. `ct := NewCounter()` binds an instance to `ct`; later calls access the same state through `ct.counts`.

The example is `O(1)` average time per increment/read; space grows as `O(K)` for `K` distinct keys.

## 8.2 Lifetime and ownership: local variables vs struct fields

A fresh `counts := make(map[string]int)` **inside** `Increment` would create a different map on each invocation. The next call would not have access to the old one. The problem is **state lifetime**, not a rule that “maps must live in structs.”

```go
// LOGIC BUG: each call discards its newly created map.
func brokenIncrement(key string) {
    counts := make(map[string]int)
    counts[key]++
}
```

A field, a map passed between functions, or a map captured by a closure can all preserve state. A struct is useful because it bundles the state and its behavior into one instance, makes dependencies explicit, and permits two instances to be independent:

```go
a := NewCounter()
b := NewCounter()
a.Increment("x")
// a.GetKey("x") == 1; b.GetKey("x") == 0
```

**Model:** `a` is a struct value containing a map field. That map field refers to underlying map storage. Repeated calls on `a` keep accessing that same map. Calling `NewCounter` again allocates a separate empty map.

## 8.3 Creation syntax and zero values

| Form | What it creates | `counts` writable? | Explanation |
|---|---|---|---|
| `var c Counter` | Zero-valued Counter variable | **No** | Map field is `nil` |
| `c := Counter{}` | Zero-valued Counter literal | **No** | Constructors are not invoked implicitly |
| `c := new(Counter)` | `*Counter` pointing at a zero-valued Counter | **No** | `new` allocates zeroed storage; it does not initialize the map |
| `c := &Counter{}` | Pointer to zero-valued Counter literal | **No** | Still has a nil map |
| `c := Counter{counts: make(map[string]int)}` | Fully initialized Counter value | **Yes** | Keyed composite literal |
| `c := &Counter{counts: make(map[string]int)}` | Pointer to initialized Counter | **Yes** | Often used in pointer-returning constructors |
| `c := NewCounter()` | Initialized Counter value | **Yes** | Calls the explicit initialization function |
| `c := Counter{counts: map[string]int{}}` | Initialized Counter value | **Yes** | Empty map literal also works |

The **zero value** of a map is `nil`:

```go
var c Counter
fmt.Println(c.GetKey("absent")) // 0: reading a nil map is valid
c.Increment("absent")            // PANIC: assignment to entry in nil map
```

The panic is at the map **write**, not when declaring the struct or reading the missing key. The `int` value type has zero value `0`, so `counts[key]++` works on a valid map even when the key was not previously inserted.

**`make` vs `new`:** `make` initializes maps, slices, and channels and returns a value of that type; `new(T)` returns `*T` whose `T` is zero-valued. `make(Counter)` does not compile. `new(Counter)` compiles but leaves `counts` nil.

## 8.4 Struct declarations, literals, visibility, and embedding

- Define a named type: `type Counter struct { counts map[string]int }`. The field name `counts` is unexported, accessible to code in the **same package**.
- Keyed literal: `Counter{counts: make(map[string]int)}`. Prefer keyed fields for readability and resilience to field reordering.
- Positional literals like `Counter{make(map[string]int)}` compile in this package, but are fragile and are disallowed for some cross-package struct literals. Prefer keyed.
- Anonymous struct: `cfg := struct{ Limit int }{Limit: 10}` is a distinct unnamed struct type, useful for small local groupings.
- Exported identifiers begin with uppercase letters (`Counter`, `NewCounter`, `Increment`); `counts` is package-private, not class-private.
- Struct embedding gives promoted fields/methods, **not class inheritance**. Composition (holding dependencies as fields) is idiomatic Go. An embedded value still has its own initialization requirements.

**Common declaration trap:** `type Counter struct.` is invalid syntax. Braces and separated declarations are required; use `gofmt -w` on saved Go files.

## 8.5 Constructor functions: design, return types, and invariants

Go has no special constructor keyword and **does not automatically call `NewCounter` or `Constructor`** when you write `Counter{}`. Conventional constructors are ordinary functions such as `NewCounter`, `NewClient`, `NewServer`.

```go
// Value-returning style — mirrors the exercise.
func NewCounter() Counter {
    return Counter{counts: make(map[string]int)}
}

// Alternative, pointer-returning style (choose one API, not both in one package):
// func NewCounter() *Counter {
//     return &Counter{counts: make(map[string]int)}
// }
```

**Why choose a constructor?** It communicates and enforces a usable starting configuration for cooperative callers: initialized maps, dependencies, limits, options, resource ownership and validation. It does **not** make raw struct literals illegal. Whether `Counter{}` should work is an intentional API design choice.

Some constructors validate inputs:

```go
type Config struct { Capacity int }
type Buffer struct { capacity int }

func NewBuffer(cfg Config) (*Buffer, error) {
    if cfg.Capacity <= 0 {
        return nil, fmt.Errorf("capacity must be positive: %d", cfg.Capacity)
    }
    return &Buffer{capacity: cfg.Capacity}, nil
}
```

Here `fmt` must be imported, and the caller checks `err` before using the pointer. Constructor names are conventions, not overloading: Go doesn't have overloaded functions/methods based on parameter lists.

**Value vs pointer return:** Returning `Counter` is valid and efficient enough for this small example. Returning `*Counter` is common when callers need shared mutable struct state, stable identity, or the type contains locks/noncopyable fields. Do not claim that a pointer return automatically means “heap” or is always faster: Go's escape analysis determines allocation placement.

## 8.6 Receiver methods, implicit address-taking, and interfaces

```go
func (c *Counter) Increment(key string) { c.counts[key]++ }
func (c *Counter) GetKey(key string) int { return c.counts[key] }
```

The receiver is an argument with special syntax. `c` points to the selected `Counter`. In `ct := NewCounter(); ct.Increment("a")`, `ct` is an **addressable variable**, so Go can implicitly use `(&ct).Increment("a")`.

But `NewCounter().Increment("a")` **does not compile** for this pointer-receiver method: the returned temporary value is not addressable. Assign it to a variable first, or use a pointer-returning constructor.

A **value receiver** receives a copy of the struct. A **pointer receiver** can directly update fields of the original struct:

```go
type Gauge struct{ value int }

func (g Gauge) TrySetByValue(v int) { g.value = v } // changes only copy
func (g *Gauge) SetByPointer(v int) { g.value = v }  // changes original
```

But note the crucial map nuance: copying a struct with a map field copies the map *value*, which still refers to the same underlying map data. Thus, a value-receiver method can mutate entries of an initialized map. Pointer receivers are still the better consistent choice for a mutable stateful API, especially if a method might replace the map field.

**Method-set trap:** if `Increment` exists only on `*Counter`, then `*Counter` satisfies an interface requiring `Increment(string)`, but `Counter` does not. Calling a pointer method on an addressable value is permitted; interface method-set satisfaction is a separate rule.

## 8.7 Copying a struct is not cloning its state

```go
a := NewCounter()
b := a // copies the struct value, not the underlying map
a.Increment("x")
b.Increment("x")
// a.GetKey("x") == 2 and b.GetKey("x") == 2
```

Contrast **two constructor calls** (two maps) with **one constructor plus assignment** (shared map). To make a true independent copy, explicitly clone map entries, for example using `maps.Clone` (Go 1.21+) or a manual loop, then assign the cloned map to the new struct. This is a shallow clone of map entries; nested reference-like values may still share storage.

If a struct contains a `sync.Mutex`, copying it **after first use** is unsafe. This is one reason stateful synchronized objects normally use pointer-returning constructors and pointer receivers. `go vet -copylocks` can detect some accidental copies.

## 8.8 Designing a usable zero value

The exercise's `Counter` intentionally requires `NewCounter` before the first increment. Go also encourages zero-value-usable APIs when practical.

See [LazyCounter](examples/variants.go), whose `Increment` does:

```go
if c.counts == nil {
    c.counts = make(map[string]int)
}
c.counts[key]++
```

Now `var c LazyCounter; c.Increment("x")` works. That only makes **sequential** initialization convenient: concurrent first calls can race on both initialization and writes. The choice is a contract:

- **Constructor-required**: force known configuration; document that zero-value writes are invalid.
- **Zero-value-ready**: lazily initialize or use fields whose zero values are already useful; test the zero value.

Avoid pretending that an unexported field prevents all invalid construction when the exported type itself is still instantiable.

## 8.9 Concurrency: object construction does not imply thread safety

The baseline `Counter` is **not safe** for concurrent increments. Map writes from multiple goroutines and `counts[key]++` (a read-modify-write operation) need synchronization.

[SafeCounter](examples/variants.go) owns a `sync.Mutex` and a map, uses a pointer-returning `NewSafeCounter`, locks both updates and reads, and is tested with multiple goroutines. A constructor initializes its map; the mutex itself has a useful zero value, but it does not initialize the map.

Run `go test -race ./...`. Race detection can expose unsynchronized access when exercised; passing once is not a formal proof for every schedule. For more, use [Module 3 — synchronization](../3-synchronization-and-race-conditions/README.md).

## 8.10 The complete instance lifecycle

```text
type declaration      -> names fields and their types (no instance)
var / literal / new   -> creates a zero-valued instance unless fields supplied
NewCounter() call     -> constructs usable state with a writable map
method call           -> accesses or changes that instance's fields
struct assignment     -> copies fields; may still share referenced data
out of scope          -> object becomes eligible for GC when unreachable
```

There is no automatic C++/Java-style constructor/destructor dispatch for Go structs. Garbage collection reclaims unreachable memory; external resources such as files/connections still need explicit `Close` / cleanup at an appropriate ownership boundary.

A constructor that creates a `*T` does not by itself prove the value lives on the heap. Check escape diagnostics and profiles when performance actually matters; use [Module 5 — memory](../5-memory-escape-analysis-and-gc/README.md).

## 8.11 Common syntax mistakes and runtime traps

| Snippet | Verdict | Why / correction |
|---|---|---|
| `type Counter struct.` | Compile error | Use `type Counter struct { ... }` |
| `fmt.PrintLn(4)` | Compile error | Go identifier is `fmt.Println` (lowercase `ln`) |
| `Counter{}` then `Increment("x")` | Runtime panic | Nil map write; construct/init first |
| `new(Counter).Increment("x")` | Runtime panic | `new` did not initialize map |
| `Counter{counts: make(map[string]int)}` | Valid | Explicit initialized composite literal |
| `NewCounter().Increment("x")` | Compile error | Pointer method on unaddressable returned value |
| `a := NewCounter(); b := a` | Valid, shared map | Struct copy does not deep-clone map |
| `a := NewCounter(); b := NewCounter()` | Valid, independent maps | Separate constructor calls |
| `make(Counter)` | Compile error | `make` accepts only slices, maps, channels |
| `var m map[string]int; _ = m["x"]` | Valid | Reading nil map returns zero value |
| `var m map[string]int; m["x"]++` | Runtime panic | Writes to nil map are invalid |

These are **distinct failure phases**: compiler errors, runtime panics, and silently incorrect logic. State which category a bug belongs to before suggesting a fix.

## 8.12 Interview questions (answer without opening code)

1. Explain `var c Counter`, `Counter{}`, `new(Counter)`, `&Counter{}`, and `NewCounter()` including their types and map state.
2. Why does a local `make(map[string]int)` inside a method fail to preserve counts?
3. Does `Counter{}` automatically call `NewCounter()`? Why not?
4. What happens when reading and writing an uninitialized map? Why does `counts["x"]++` work on an initialized empty map?
5. What exactly is copied by `copy := original` when the struct contains a map, pointer, slice, or mutex?
6. Why can an addressable `Counter` call pointer-receiver methods, while it cannot satisfy an interface whose method set requires that pointer method?
7. Can a value-receiver method mutate entries of a map field? Can it replace the original map field? Explain each.
8. When return `Counter` versus `*Counter`? Does either guarantee heap or stack allocation?
9. When should a constructor return `(*T, error)`? Which invariants should it establish?
10. How would you make a zero-value Counter usable? How would you make it concurrency-safe?
11. If ten goroutines call `Increment("x")`, is `counts["x"]++` atomic? How would you test your answer?
12. What owns the lifetime of an external resource acquired during construction? When does `Close` run?
13. Describe the difference between construction, initialization, method invocation, field mutation, struct copying, and garbage collection in Go.
14. In Go, what does embedding provide, and what does it **not** provide compared to class inheritance?

## 8.13 Hands-on exercises and recall checkpoints

1. **Blank file:** reconstruct `Counter`, `NewCounter`, `Increment`, `GetKey`, and `main`. Expect output `4`.
2. **Two instances:** increment only `a`, prove `b` is unchanged, and explain why.
3. **Zero-value trap:** predict the exact failing operation in `var c Counter; c.Increment("x")`. Test with a recover only in an isolated test.
4. **Copy trap:** assign `b := a`, increment `b`, then prove `a` sees the same map update. Clone to make truly independent state.
5. **Receiver trap:** implement a struct with an `int` field and both value and pointer receiver methods; predict their effect.
6. **API design:** rewrite the constructor to return `*Counter`, add a rejected configuration with an error, and test error handling.
7. **Concurrency:** run the synchronized variant with `go test -race`, then reason about what would fail without locks. Do not place intentional race examples in the passing test suite.

**Run locally (no external dependencies or CI workflow):**

```sh
cd senior-golang-interview-prep/8-structs-constructors-and-state/examples
go run .
go test -v ./...
go test -race ./...
go vet ./...
gofmt -d *.go
```

Expected `go run .` output: `4`. See [runnable files](examples/) for the two alternative designs and tests.

**Revision gates:**

- [ ] Explain where state lives and why it survives separate method calls.
- [ ] Predict nil map reads, nil map writes, and constructor bypasses.
- [ ] Recreate the corrected Counter with no reference.
- [ ] Explain pointer/value receivers and map aliasing without handwaving.
- [ ] Distinguish value vs pointer construction, method set, and interface satisfaction.
- [ ] Run and interpret the passing tests/race detector; explain what they do not prove.
- [ ] Give a 60-second senior-level explanation and defend a zero-value or constructor-required API.
- [ ] Repeat from memory after 1, 3, 7, 14 and 30 days as needed, recording misses in [NOTES.md](NOTES.md) and [REVISION-LOG.md](../REVISION-LOG.md).

**Primary references:** [Go spec — types and struct types](https://go.dev/ref/spec#Struct_types), [composite literals](https://go.dev/ref/spec#Composite_literals), [method sets](https://go.dev/ref/spec#Method_sets), [Go blog — effective Go allocation](https://go.dev/doc/effective_go#allocation_new), [Go memory model](https://go.dev/ref/mem), [sync package](https://pkg.go.dev/sync).
