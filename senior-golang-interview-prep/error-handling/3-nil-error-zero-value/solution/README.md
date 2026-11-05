# Q3. Understand the zero value of error — worked solution

[Question (try first)](../README.md) · [Error-handling index](../../README.md) · [Go implementation](main.go) · [Tests](main_test.go)

**Level:** Starter | **Concepts:** nil interfaces, successful returns.

## Run

From this `solution/` directory:

```bash
go test ./...
go test -race ./...
go run .
```

For this solution, use Go 1.21 or newer.

## Reasoning and trade-offs

The zero value of an interface is nil (no dynamic type, no dynamic value). The typed-nil trap is explored in Q16.

## Worked implementation

Read [main.go](main.go) and explain every error return and branch. The related [main_test.go](main_test.go) checks both regular and failure behavior.

## Cases your tests must prove

- Zero-value error compares equal to nil.
- A successful function returns nil.

## Common trap to articulate

Explain how this implementation would fail if the cause were discarded, the wrong error were classified, or a caller assumed all errors were retryable. Adapt that explanation to the specific error policy in this exercise.

## Interview recap

In 30–60 seconds, explain **understand the zero value of error**: the caller-visible contract, what remains inspectable after wrapping, the failure-case test, and why this design fits the layer.

## Further practice

Modify one error path (for example, change a `%w` to `%v`, remove cancellation, or change a sentinel to a new value); predict which test fails before running it. Record the result in [NOTES.md](../../NOTES.md).

