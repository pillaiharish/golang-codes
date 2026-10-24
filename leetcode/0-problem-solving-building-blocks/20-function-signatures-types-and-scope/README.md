# 20 — Function signatures, types, variable scope and shadowing

[← Foundations index](../README.md) · [← LeetCode index](../../README.md) · [Codec attempt concept map](../271-codec-attempt-concept-map.md)

## Recognition trigger and intuition

You know the intended algorithm but the compiler complains about return types, `append`, undefined identifiers, or `:=`. **Read the type contract first**: which values enter a function, which type leaves it, and which variables exist in each scope?

## Go operations and invariants

1. `func Decode(encoded string) []string` receives one `string` and **must return a slice of strings**. `var out string` is the wrong return type; use `out := []string{}` or `var out []string`.
2. `type Codec struct{}` is a valid empty struct. `Codec{}` is a value; `&Codec{}` is a `*Codec` pointer. Match the expression's type to the constructor's declared return type.
3. `string` is a **type**, while `""` is a **value** of that type. A field literal like `encoded: string` cannot initialize a string field.
4. `:=` declares one or more variables (with at least one new name in that scope). `=` assigns to an already declared variable. An inner `{ ... }` block can declare another variable with the same spelling, hiding the outer one (**shadowing**).
5. `append(slice, item)` returns a new slice header; normally write `out = append(out, item)`. It does not accept a string as its first argument.
6. A method has a receiver, such as `func (c *Codec) Collect(...) ...`. The receiver does not automatically mean the object must store state, or that `Constructor` must return `*Codec`.

**Invariant while collecting:** after processing `k` input values, the *outer* output slice contains precisely those `k` values. Accidental shadowing breaks this invariant.

## Complete runnable example

```go
package main

import "fmt"

type Codec struct{}

func Constructor() Codec { return Codec{} }

func (c *Codec) Collect(items []string) []string {
    out := []string{}
    for _, item := range items {
        out = append(out, item) // Assign to the existing outer slice.
    }
    return out
}

func main() {
    codec := Constructor()
    fmt.Printf("%T %T\n", codec, &codec) // main.Codec *main.Codec
    fmt.Println(codec.Collect([]string{"a", "b"})) // [a b]

    total := 0
    for _, n := range []int{1, 2} {
        total := total + n // New variable in the loop body: shadows outer total.
        _ = total
    }
    fmt.Println(total) // 0: the outer variable was never updated.

    var text string // Empty string (zero value), not the type name as a value.
    var names []string // Nil slice, which can still be appended to.
    names = append(names, text)
    fmt.Printf("%q %v\n", text, names)
}
```

## Hand trace

- `Constructor()` returns a `Codec` **value**. The variable `codec` is addressable, so Go can call its `*Codec` receiver method.
- `Collect` appends `"a"` and then `"b"` to the **same** `out`, returning `[]string{"a", "b"}`.
- The inner `total := total + n` shadows the outer `total` separately on each iteration. The printed outer `total` is still `0`.
- `names` starts nil, accepts an append, and ends with a one-element slice containing `""`.

## Mini drill — fix the mistake before reading elsewhere

1. What should `func NewCodec() *Codec` return: `Codec{}` or `&Codec{}`?
2. Write the initialization of a struct field `label string` with an empty string. Explain why `label: string` is invalid.
3. A function returns `[]string`; declare an empty result and append two strings without shadowing it.
4. Explain why `out := append(out, item)` **inside** a `for` body can leave the `out` declared outside the loop unchanged.
5. If the codec stores no fields, why can it still have methods?

## Common traps

- Thinking a pointer receiver forces a pointer-returning constructor. It does not for an addressable variable.
- Confusing a type (`string`, `[]string`, `Codec`) with a value (`""`, `[]string{}`, `Codec{}`).
- Using `:=` in a nested block when the intention is to update an outer variable.
- Assuming `append` modifies a string, or forgetting to assign its returned slice.
- Storing intermediate data in a struct field when a local variable avoids accidental state across calls.

## Apply and recall

- [Structs and state](../11-structs-pointers-methods-and-linked-lists/README.md) deepens constructors, receivers and persistent state.
- [Slices](../01-arrays-and-slices/README.md) explains `append` and collection types.
- [Strings and parsing](../04-strings-bytes-runes-and-parsing/README.md) covers framing and conversion.
- [LeetCode 271 Codec attempt: exact mistake mapping](../271-codec-attempt-concept-map.md) applies all three ideas to one debugging session.
- [Logger Rate Limiter #359](../../359.%20Logger%20Rate%20Limiter/main.go) demonstrates a **stateful** object in contrast to a stateless codec.

**30-second blank-screen recall:** without notes, write `type Codec struct{}`, a value-returning constructor, a pointer receiver method returning `[]string`, and one loop using `out = append(out, s)`. State the type of each expression and why `:=` inside the loop is risky.
