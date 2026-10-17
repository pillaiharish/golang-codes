# 4. Increment statements and for-loop traps

[All doubts](README.md) · [Previous](03-bytes-runes-and-utf8.md) · [Next: buffers](05-buffers-blocking-and-backpressure.md)

## Direct answer

`i++` and `i--` are statements in Go, not expressions that return a value. They work outside loops too. Go has postfix syntax only, so `++i` is invalid. `i++` acts like `i += 1`, with the operand evaluated once; it does not make a concurrent update atomic.

## Valid and invalid locations

| Code | Valid? | Reason |
|---|---|---|
| `i++` inside a function | Yes | Standalone increment statement |
| `for i := 0; i < 3; i++ {}` | Yes | Post clause accepts this statement |
| `for ; i < 3; { i++ }` | Yes | Body explicitly advances counter |
| `x := i++` | No | No expression value |
| `fmt.Println(i++)` | No | Function arguments need expressions |
| `if i++ < 3 {}` | No | Cannot embed statement in comparison |
| `if i++; i < 3 {}` | Yes | Separate init statement, then condition |
| `++i` | No | No prefix increment |
| `s += "x"` | Yes | String concatenation assignment |
| `s++` | No | Increment requires numeric operand |
| `i += 2` | Yes | Compound assignment statement |
| `i, j = i+1, j-1` | Yes | Multiple assignment, useful as loop post |
| `i++, j--` | No | Not Go's multiple-assignment syntax |

The operand must be addressable or a map index expression, and numeric. `a[k]++`, `m[key]++`, and `(*p)++` can work. `*p++` also parses as `(*p)++` in Go; write parentheses for clarity. Incrementing the pointer itself is invalid: Go has no general pointer arithmetic.

A missing key in a non-nil `map[string]int` reads as zero, so `m[key]++` inserts 1. A nil map still panics when the update attempts to write. Constants and non-addressable numeric return values cannot be incremented.

## for execution order and continue

```go
for i := 0; i < 4; i++ {
    if i == 1 {
        continue
    }
    fmt.Println(i)
}
// 0 2 3
```

The order is init once, condition, body, post, condition again. `continue` runs the post clause of the targeted loop before checking its condition; `break` exits without running the post. A loop with no condition runs until an explicit exit.

```go
i := 0
for i < 4 {
    if i == 1 {
        // continue here would skip i++ and loop forever
    }
    i++
}
```

In a range loop, incrementing the range variable inside the body does not skip an element: the next iteration receives its next value from range. A `for` post statement cannot be a short variable declaration. Braces normally stay on the `for` line because newline semicolon insertion changes parsing.

## Numeric and concurrency traps

```go
var n uint8 = 255
n++                 // wraps to 0
var f float64 = 1.5
f++                 // 2.5
var c complex128 = 2 + 3i
c++                 // 3 + 3i
```

Increment is defined for numeric types, not only integers. Fixed-width integer overflow does not panic. Very large floating-point values may not change after adding 1 because of rounding. A shared `counter++` is a read-modify-write and races without synchronization; use a mutex or an appropriate `sync/atomic` operation. Making a counter atomic does not automatically protect surrounding state.

## Loop variables and Go versions

For code using Go 1.22 or later language semantics, loop variables **declared by the loop** are distinct per iteration. This matters for closures and taking addresses. It does not apply to a variable declared outside and reused with `=`.

```go
var callbacks []func() int
for i := 0; i < 3; i++ {
    callbacks = append(callbacks, func() int { return i })
}
// Calling callbacks in order gives 0, 1, 2 with Go 1.22+ semantics.

var shared int
var reused []func() int
for shared = 0; shared < 3; shared++ {
    reused = append(reused, func() int { return shared })
}
// Calling after the loop gives 3, 3, 3: one shared variable.
```

The module's `go` language version matters, not only the installed toolchain version. Distinct loop variables also do not make every object referenced by those variables independent or thread-safe.

## Interview checkpoints

- Explain why `if i++; i < 3` is valid but `if i++ < 3` is invalid.
- Trace `continue` in a three-clause loop versus a condition-only loop.
- Explain map-entry increment and why it fails on a nil map.
- Explain the remaining closure hazard when the loop uses `=`.

Run `go run ./examples/loops` from this directory: [source](examples/loops/main.go).

Sources: [increment/decrement](https://go.dev/ref/spec#IncDec_statements), [for statements](https://go.dev/ref/spec#For_statements), [loop-variable change](https://go.dev/blog/loopvar-preview), [atomic operations](https://pkg.go.dev/sync/atomic).
