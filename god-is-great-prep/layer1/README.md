# Layer 1 — Go syntax and everyday fundamentals

[Prep home](../README.md) · [Learning mandate](../STUDY-MANDATE.md) · [Complete preserved course](reference.md) · [Layer 2](../layer2/README.md)

**Goal:** Write, run, debug and explain basic Go without looking up syntax. Do not study runtime internals or advanced DSA here.

## Suggested order / coverage tracker

| Module | What to demonstrate | Recalled without notes? |
|---|---|---|
| [1. Environment and tooling](reference.md#1-environment-and-tooling) | Basics + one example + error/edge-case note | ☐ |
| [2. Anatomy of a Go program](reference.md#2-anatomy-of-a-go-program) | Basics + one example + error/edge-case note | ☐ |
| [3. Printing and formatting output](reference.md#3-printing-and-formatting-output) | Basics + one example + error/edge-case note | ☐ |
| [4. Variables, constants, and zero values](reference.md#4-variables-constants-and-zero-values) | Basics + one example + error/edge-case note | ☐ |
| [5. Basic types and type conversion](reference.md#5-basic-types-and-type-conversion) | Basics + one example + error/edge-case note | ☐ |
| [6. Operators and expressions](reference.md#6-operators-and-expressions) | Basics + one example + error/edge-case note | ☐ |
| [7. Conditional statements: `if` and `switch`](reference.md#7-conditional-statements-if-and-switch) | Basics + one example + error/edge-case note | ☐ |
| [8. Loops: `for` and `range`](reference.md#8-loops-for-and-range) | Basics + one example + error/edge-case note | ☐ |
| [9. Functions](reference.md#9-functions) | Basics + one example + error/edge-case note | ☐ |
| [10. Basic error handling](reference.md#10-basic-error-handling) | Basics + one example + error/edge-case note | ☐ |
| [11. Basic collections: arrays, slices, maps](reference.md#11-basic-collections-arrays-slices-maps) | Basics + one example + error/edge-case note | ☐ |
| [12. Strings and runes basics](reference.md#12-strings-and-runes-basics) | Basics + one example + error/edge-case note | ☐ |
| [13. Structs and pointers preview](reference.md#13-structs-and-pointers-preview) | Basics + one example + error/edge-case note | ☐ |
| [14. Packages, modules, and imports](reference.md#14-packages-modules-and-imports) | Basics + one example + error/edge-case note | ☐ |
| [15. Basic input and command-line arguments](reference.md#15-basic-input-and-command-line-arguments) | Basics + one example + error/edge-case note | ☐ |
| [16. Basic testing](reference.md#16-basic-testing) | Basics + one example + error/edge-case note | ☐ |
| [17. Debugging and common beginner compile errors](reference.md#17-debugging-and-common-beginner-compile-errors) | Basics + one example + error/edge-case note | ☐ |
| [18. Layer 1 practice exercises](reference.md#18-layer-1-practice-exercises) | Basics + one example + error/edge-case note | ☐ |

Each module is explained, with example code, in the [complete Layer 1 reference](reference.md). Start there rather than copying another beginner tutorial.

## Go-specific mandatory essentials (do not skip)

- Declarations: `var`, `const`, `:=`, scope, shadowing and zero values.
- Types: `int`, `float64`, `bool`, `string`, `byte`, `rune`, explicit conversions.
- Branches: `if`, `else`, `switch`, short statement, switch without expression, no ternary.
- Loops: counted/condition-only/infinite `for`, `range`, `break`, `continue`, `_` discard.
- Functions: declarations, named parameters, returns, multiple returns, variadic functions, elementary `defer`.
- Error values: `(value, error)`, `if err != nil`, `errors.New`, `fmt.Errorf`.
- Surface collection syntax: array, slice, map, struct and pointer; the *operations and ownership* belong to Layer 2.
- Strings: bytes versus runes, `len`, `range`, formatting and `strconv`.
- Go packages/modules: `go.mod`, imports, exported names, `go fmt`, `go vet`, `go test`.
- Simple CLI input, `os.Args`, file read awareness and a minimal table-free test.

## Layer 1 practical checkpoint

From a blank folder, write an executable that accepts an integer argument; rejects missing or nonnumeric input, classifies it using `if`/`switch`, loops from 1 to N, stores running totals in a slice/map, prints with `fmt.Printf`, and has a test for a pure helper.

**Pass:** executable works, helper test passes, no compiler error, and you can explain each line and its zero value. Do not mark complete because the course was read.

## Boundaries

The original notes include some deeper previews. Treat those as **orientation**, not required mastery. Array/slice aliasing, linked lists, method sets, interfaces, sorting and complexity move to Layer 2 or Layer 3. See the [study mandate](../STUDY-MANDATE.md) for recall rules.
