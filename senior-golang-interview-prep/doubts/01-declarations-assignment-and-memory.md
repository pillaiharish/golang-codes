# 1. Variable declarations, initialization, assignment, and memory

[All doubts](README.md) · [Next: zero values](02-zero-values-make-and-new.md)

## Direct answer

For variables, Go does not use the C-style distinction “declaration announces a variable; a later definition allocates it.” `var n int` declares a variable and gives it the value `0`. `var n int = 10` declares it with an explicit initializer. `n = 20` assigns a new value to an existing variable. `n := 10` is a **short variable declaration**, allowed inside functions. Calling it “short assignment” hides the important fact that it must declare something new.

A declaration establishes language-level storage and identity. It does not promise a separate physical memory allocation: the compiler may use a register, a stack slot, heap storage, or remove the variable entirely.

## Syntax and scope

| Form | Meaning | Where allowed |
|---|---|---|
| `var n int` | Declare `n`, initialize to zero | Package or function scope |
| `var n int = 10` | Declare with explicit type and value | Package or function scope |
| `var n = 10` | Declare with inferred type `int` | Package or function scope |
| `n := 10` | Short declaration with inferred type | Inside functions |
| `n = 10` | Assign to an existing target | Statement context inside functions |
| `const n = 10` | Declare a constant, not a mutable variable | Package or function scope |

Package variables exist for the program lifetime and are initialized before normal execution of `main`. Local scope and storage lifetime are different: an object may remain reachable after its declaring function returns.

## How short declarations work

```go
n := 10
n, label := 20, "count" // same n, new label
n = 30                // ordinary assignment
_ = label
// n := 40            // INVALID: no new non-blank variables
// n, _ := 40, true   // INVALID: _ does not count as a new variable
```

Redeclaration requires the existing identifier to be in the **same block** (or the function parameter list when declaring in the function body), to retain its type, and at least one other non-blank identifier to be new. `:=` takes identifiers on its left, not `obj.Field` or `slice[0]`.

```go
n := 10
{
    n, label := 20, "inner" // new n shadows outer n
    fmt.Println(n, label)  // 20 inner
}
fmt.Println(n)             // 10
```

A common error is accidentally shadowing `err` inside an `if`, then inspecting the unchanged outer error. Conversely, `value, err := work()` in the same block can reuse `err` while introducing `value`. Prefer names and narrow scopes that make error handling obvious.

The new identifier's scope starts after its declaration, so an inner `n := n + 1` can read the outer `n` on the right. Multiple assignment evaluates the right-hand values before performing the assignments: `a, b = b, a` swaps the values.

## Type inference does not mean dynamic typing

```go
n := 10       // int
f := 1.5      // float64
r := '界'     // rune (int32)
s := "hello"  // string
var small int8 = 10
_ = small
// n = "ten"  // INVALID: n remains int
// n = small  // INVALID: explicit conversion needed
n = int(small)
```

Untyped constants acquire a type from context or their default type. `var small int8 = 128` fails because the constant cannot be represented. Inferred types and constant representability are compile-time rules, not allocation policies.

## Memory at each level

| Level | Example | What you can conclude |
|---|---|---|
| Scalar value | `n := 10` | Value is an `int`; no guaranteed heap allocation |
| Pointer | `p := &n` | Holds an address; taking it alone does not force escape |
| Array / struct | `a := [4]int{}` | Contains its elements / fields as values |
| Slice | `s := make([]int, 4)` | Slice refers to a backing array; header and array have separate placement decisions |
| Map / channel | `make(map[string]int)`, `make(chan int, 4)` | Creates usable runtime-managed state; syntax alone does not specify allocation counts |
| String | `s := "abc"` | String value refers to immutable bytes; literal bytes need not be newly allocated per execution |
| Interface | `var x any = n` | Carries dynamic type/value; boxing does not universally imply heap allocation |
| Package variable | `var total int` | Program-lifetime storage, not automatically a GC heap allocation |

Common implementations represent a slice as pointer/length/capacity and a string as pointer/length. These are explanatory models, not a license to depend on undocumented layout. Copying a slice copies the descriptor, not its elements. Copying an array copies the elements. Copying a struct copies its fields; pointer/slice/map fields can still share referenced data.

```go
func local() int {
    p := new(int)
    *p = 7
    return *p // compiler can avoid a heap allocation
}

func retained() *int {
    n := 7
    return &n // safe: storage survives as needed; often heap without inlining
}
```

Neither `var` versus `:=` nor `new` versus `&T{}` decides stack versus heap. Escape analysis, inlining, size, and compiler limitations matter. Go does not produce a dangling stack pointer when returning `&n`. Heap objects become eligible for collection when unreachable; collection is not an immediate `free` at scope exit.

## Practice: inspect instead of guessing

Run [the allocation example](examples/allocations/main.go):

```sh
go run ./examples/allocations
go build -gcflags='-m=2' -o /tmp/doubts-allocations ./examples/allocations
```

Run from this `doubts` directory. Inspect the local `new(int)` and the pointer stored in a package variable. The example disables inlining on two small helpers to make the comparison easier to inspect. Do not generalize its diagnostic wording or allocation counts to every compiler version. Printing a value can itself affect escape analysis; keep formatting out of the functions being investigated.

## Interview checkpoints

- Explain why `var n int` is already initialized.
- Predict whether `n, err := f()` reuses or shadows `n` in a specified block.
- Explain why two syntax forms can compile to identical machine code.
- Distinguish a slice descriptor from its backing array and scope from lifetime.

Sources: [variable declarations](https://go.dev/ref/spec#Variable_declarations), [short declarations](https://go.dev/ref/spec#Short_variable_declarations), [scope](https://go.dev/ref/spec#Declarations_and_scope), [stack or heap FAQ](https://go.dev/doc/faq#stack_or_heap), [GC guide](https://go.dev/doc/gc-guide).
