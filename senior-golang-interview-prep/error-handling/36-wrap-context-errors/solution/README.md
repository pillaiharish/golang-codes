# Q36. Wrap context errors without changing their policy — worked solution

[Question (try first)](../README.md) · [Error-handling index](../../README.md) · [Go implementation](main.go) · [Tests](main_test.go)

**Level:** Upper intermediate | **Concepts:** cancellation, retry policy, context error identity.

## Run

From this `solution/` directory:

```bash
go test ./...
go test -race ./...
go run .
```

Go 1.21 or newer is recommended.

## Explanation and trade-offs

Wrapping adds operational context without losing the cancellation signal. Do not treat caller cancellation as a transient dependency failure or log it as a 500 by default.

## Reference code and tests

Study [main.go](main.go) and [main_test.go](main_test.go) after your own attempt. Explain every failure return, wrapping edge, and caller decision.

### Verify these cases

- Wrapped canceled error matches context.Canceled.
- Successful context returns nil.

## Failure analysis

Modify one error path and predict the failing test. Consider an ignored return, a broken wrapper, a wrong classification, or an unnecessary retry where relevant to this exercise.

## Senior-level explanation

In about one minute: state **wrap context errors without changing their policy**, the stable caller contract, the trade-off, how you tested the failure path and what you would do differently at a service boundary.

Record surprises and a confidence score in [NOTES.md](../../NOTES.md).

