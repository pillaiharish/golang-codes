# Q6. Show why ignoring errors hides failure — worked solution

[Question (try first)](../README.md) · [Error-handling index](../../README.md) · [Go implementation](main.go) · [Tests](main_test.go)

**Level:** Starter | **Concepts:** blank identifier, error propagation.

## Run

From this `solution/` directory:

```bash
go test ./...
go test -race ./...
go run .
```

For this solution, use Go 1.21 or newer.

## Reasoning and trade-offs

Discarding an error conflates failure with a legitimate zero result. Deliberate ignores need a documented reason (for example, a best-effort cleanup with no useful recovery).

## Worked implementation

Read [main.go](main.go) and explain every error return and branch. The related [main_test.go](main_test.go) checks both regular and failure behavior.

## Cases your tests must prove

- Valid text is parsed.
- Invalid text is rejected, never treated as zero.

## Common trap to articulate

Explain how this implementation would fail if the cause were discarded, the wrong error were classified, or a caller assumed all errors were retryable. Adapt that explanation to the specific error policy in this exercise.

## Interview recap

In 30–60 seconds, explain **show why ignoring errors hides failure**: the caller-visible contract, what remains inspectable after wrapping, the failure-case test, and why this design fits the layer.

## Further practice

Modify one error path (for example, change a `%w` to `%v`, remove cancellation, or change a sentinel to a new value); predict which test fails before running it. Record the result in [NOTES.md](../../NOTES.md).

