# Go Error Handling — Learning Index and Revision Guide

An index of the **existing code examples** under `error-handling/`, ordered for learning and interview revision. Each link goes to the actual source file; no example has been rewritten for this index.

**Covered:** 13 runnable `main.go` files and one existing written guide. **Progress:** not assessed until you have tried the recall questions yourself.

## Learning map

| # | Existing example | What it demonstrates | What to learn / explain |
| --- | --- | --- | --- |
| 1 | [Error basics](1-errors-basics/main.go) | `divide(a, b) (int, error)`, `errors.New`, success returning `nil` | The built-in `error` interface, multiple return values, checking a returned error |
| 2 | [Static vs dynamic errors](2-errors-static-vs-dynamic-errors/main.go) | `errors.New` and `fmt.Errorf` | Fixed messages vs formatted errors containing runtime values |
| 3 | [Error data type](3-error-data-type/main.go) | `%T` and `reflect.TypeOf` on two errors | Interface type vs concrete dynamic type (`*errors.errorString`, file-related `*fs.PathError`) |
| 4 | [Sentinel errors](4-sentinel-errors/main.go) | `ErrNotAuthorized` and `ErrLoginRequired` | Named, reusable error values representing recognizable conditions |
| 5 | [Wrapping errors](5-wrapping-errors/main.go) | `fmt.Errorf` with two `%w` verbs and a `%v` value | Add context, preserve causes, distinguish wrapped errors from formatted text |
| 6 | [Error comparison](6-error-type-compare/main.go) | `errors.Is(finalError, ErrNotAuthorized)` | Match a sentinel inside wrapped errors rather than only comparing the outer value |
| 7 | [`errors.As` example](7-error-as-example/main.go) | Wrapped `insufficientFundsError`, extraction with `errors.As` | Extract a concrete error type, access its fields, calculate shortfall |
| 8 | [Custom errors](8-custom-errors/main.go) | `APIError` with `Code` and `Message`, `Error() string` | Implement `error` with a struct; inspect structured data with `errors.As` |
| 9 | [File error exercise](9-example_1/main.go) | `os.ReadFile`, error wrapping, `errors.Is(err, os.ErrNotExist)` | Distinguish file-not-found from other I/O failures after adding context |
| 10 | [`defer` basics](defer-keyword_1/main.go) | Deferred print at function exit | Cleanup timing and execution order |
| 11 | [`defer` rules](defer-keyword_2/main.go) | Three defers, mutable `i` | Last-in-first-out execution and immediate argument evaluation |
| 12 | [`panic` example](panic-keyword_1/main.go) | Explicit panic after failed `os.Open` | Panic unwinding, stack traces, and stopping normal execution |
| 13 | [`defer` + `recover` example](defer-panic-recover_example1/main.go) | Deferred closure recovers from divide-by-zero panic | `recover` in a deferred function and modifying a named return value |
| Guide | [`defer-panic-recover.md`](defer-panic-recover.md) | Written explanation and walkthrough | When to use `defer`, what `panic` does, and how `recover` works |

## 1. Foundations: error values and creation

**[1 — Error basics](1-errors-basics/main.go)**

- Understand that the built-in `error` type is an interface exposing `Error() string`.
- Read the `divide` function and trace the failure and success paths: `(0, err)` for a zero divisor, and `(quotient, nil)` for success.
- **Recall:** Why return an error value instead of an exception? What does `err == nil` communicate?

**[2 — Static vs dynamic errors](2-errors-static-vs-dynamic-errors/main.go)**

- Compare `errors.New("something went wrong")` with `fmt.Errorf("failed to connect to port %d", 8080)`.
- Distinguish constructing a formatted error message from *wrapping* an existing error (covered in example 5).
- **Recall:** How would you include a failing port in your error without losing useful context?

**[3 — Error data type](3-error-data-type/main.go)**

- Examine `fmt.Printf("%T", err)` and `reflect.TypeOf(err)`.
- Compare the concrete error from `errors.New` with the error returned by opening a missing file.
- **Recall:** How can two values have the static type `error` but different concrete types?

**[4 — Sentinel errors](4-sentinel-errors/main.go)**

- See package-level `ErrNotAuthorized` and `ErrLoginRequired`.
- Explain why sharing one named value is useful when the caller must recognize a particular failure.
- **Recall:** Would recreating an error with the same message necessarily create the same error value?

## 2. Inspecting wrapped errors

**[5 — Wrapping errors](5-wrapping-errors/main.go)**

- See how `fmt.Errorf` adds login/authorization context.
- Notice the example uses **two `%w` verbs** for errors and `%v` for the username. The error values remain inspectable, while the username is merely formatted text.
- **Recall:** Why can `%w` be preferable to `%v` for an error produced by a lower layer?

**[6 — `errors.Is`](6-error-type-compare/main.go)**

- Read the call `errors.Is(finalError, ErrNotAuthorized)`.
- Learn that testing the outer error with `==` is not equivalent to searching through wrapping.
- **Source accuracy note:** This file executes `errors.Is`, but only *comments on* `errors.As`. Its commented output does not match the actual boolean printed by the `errors.Is` call. Use examples 7 and 8 for runnable `errors.As` demonstrations.
- **Recall:** What exactly does `errors.Is` look for, and why does the wrapped sentinel still match?

**[7 — `errors.As` with a wrapped custom error](7-error-as-example/main.go)**

- Inspect the `insufficientFundsError` struct, its pointer-receiver `Error() string`, and `withdraw`.
- Follow `fmt.Errorf("Transaction declined %w", customErr)` and the subsequent `errors.As(err, &fundsErr)` call.
- See how retrieving the concrete error lets the caller calculate the shortfall from structured fields.
- **Recall:** How do `errors.Is` and `errors.As` differ? Why pass `&fundsErr`?

**[8 — Custom API error](8-custom-errors/main.go)**

- Inspect `APIError{Code, Message}`, `(*APIError).Error()`, and `callAPIErr() error`.
- Observe extracting `*APIError` with `errors.As` even without an extra wrapping layer.
- **Recall:** Why expose fields in a custom error instead of parsing a message string?

**[9 — Missing configuration file](9-example_1/main.go)**

- Follow `os.ReadFile` → contextual `fmt.Errorf(... %w ...)` → `errors.Is(err, os.ErrNotExist)`.
- Trace the two branches: known missing-file condition vs unexpected error.
- **Recall:** At which layer should you add context, and which layer decides how to handle the failure?
- **Run note:** This example exercises the missing-file path when `./config.json` is absent. Exact filesystem messages vary by system.

## 3. `defer`, `panic`, and `recover`

The folder already has a standalone [written guide](defer-panic-recover.md) explaining the relationship among these three features.

**[`defer` example 1](defer-keyword_1/main.go)**

- Follow the order of immediate and deferred `fmt.Println` statements.
- Understand that a deferred call runs when the surrounding function exits.
- **Recall:** Does `defer` execute the function immediately when the statement is evaluated?

**[`defer` example 2](defer-keyword_2/main.go)**

- Predict the print order before running.
- Explain **LIFO** ordering among deferred calls.
- Explain why the arguments to `fmt.Println` reflect `i` at the corresponding `defer` statement, not its final value.
- **Recall:** What is captured immediately, and in which order will the calls run?

**[`panic` example](panic-keyword_1/main.go)**

- See the program call `panic` when a deliberately nonexistent file cannot be opened.
- Understand that normal execution stops and an unrecovered panic terminates the program with a stack trace.
- **Recall:** What happens to deferred cleanup during panic unwinding? Should an ordinary recoverable file-open failure usually be returned as an `error` instead?

**[`defer` + `recover` example](defer-panic-recover_example1/main.go)**

- Read the named return `(result int)`, the deferred anonymous function, the explicit `panic`, and `recover()`.
- See how recovery changes `result` to zero and allows the caller to continue.
- **Recall:** Why is `recover()` invoked from the deferred function? Why is the named result relevant here?

**Important boundary:** The panic files are demonstrations, not a recommendation to use `panic`/`recover` for everyday failures. The source's [guide](defer-panic-recover.md) recommends ordinary error returns for normal application errors.

## 4. Run and revise

Run from the **repository root**, choosing the source file you want:

```sh
go run error-handling/1-errors-basics/main.go
go run error-handling/5-wrapping-errors/main.go
go run error-handling/6-error-type-compare/main.go
go run error-handling/7-error-as-example/main.go
go run error-handling/9-example_1/main.go
go run error-handling/defer-keyword_2/main.go
go run error-handling/defer-panic-recover_example1/main.go
```

**Special case:** `go run error-handling/panic-keyword_1/main.go` is expected to exit with an unrecovered panic when its sample file is missing.

### Interview recall checklist

Try answering these *before* reviewing the source:

- [ ] Explain the `error` interface, `(value, error)` returns, and `nil` success. ([1](1-errors-basics/main.go))
- [ ] Distinguish `errors.New` and `fmt.Errorf` and identify a concrete error type. ([2](2-errors-static-vs-dynamic-errors/main.go), [3](3-error-data-type/main.go))
- [ ] Explain named sentinels and the difference between `%w` and `%v`. ([4](4-sentinel-errors/main.go), [5](5-wrapping-errors/main.go))
- [ ] Demonstrate `errors.Is` on a wrapped sentinel. ([6](6-error-type-compare/main.go))
- [ ] Demonstrate `errors.As` with a wrapped typed error. ([7](7-error-as-example/main.go), [8](8-custom-errors/main.go))
- [ ] Explain how to distinguish `os.ErrNotExist` from another file error. ([9](9-example_1/main.go))
- [ ] Predict both `defer` ordering and argument values. ([defer 1](defer-keyword_1/main.go), [defer 2](defer-keyword_2/main.go))
- [ ] Explain panic unwinding, recovery, and named return values. ([panic](panic-keyword_1/main.go), [recover](defer-panic-recover_example1/main.go))

### Confidence tracker

Use **1** = unfamiliar, **2** = recognize, **3** = need notes, **4** = explain without notes, **5** = demonstrate and defend trade-offs.

| Area | Source examples | Confidence 1–5 | Last revision | What I need to revisit |
| --- | --- | --- | --- | --- |
| Basic error values and construction | 1–3 | — | — | |
| Sentinel, wrapping and `errors.Is` | 4–6 | — | — | |
| `errors.As`, custom errors, file failures | 7–9 | — | — | |
| `defer` timing and LIFO | defer 1–2 | — | — | |
| `panic` / `recover` | panic 1, recover 1 | — | — | |

**Revision workflow:** recall → check the linked code → run it → record the mistake → revisit the example before the next interview.

---

*This index maps the material currently present in this folder. It does not claim the examples have been independently tested or that every existing comment/output annotation is accurate.*

