# Q42. Cancel sibling tasks using errgroup — worked solution

[Question (try first)](../README.md) · [Error-handling index](../../README.md) · [Go implementation](main.go) · [Tests](main_test.go)

**Level:** Advanced | **Concepts:** golang.org/x/sync/errgroup, derived context.

## Run

From this `solution/` directory:

```bash
go test ./...
go test -race ./...
go run .
```

Go 1.21 or newer is recommended. This example declares an external module dependency.

## Explanation and trade-offs

errgroup.WithContext cancels derived context when a task returns an error (or Wait returns). It does not forcibly stop goroutines; they must cooperatively check context. errgroup returns one error, not all errors.

## Reference code and tests

Study [main.go](main.go) and [main_test.go](main_test.go) after your own attempt. Explain every failure return, wrapping edge, and caller decision.

### Verify these cases

- First error is returned.
- Sibling observes cancellation and exits.

## Failure analysis

Modify one error path and predict the failing test. Consider an ignored return, a broken wrapper, a wrong classification, or an unnecessary retry where relevant to this exercise.

## Senior-level explanation

In about one minute: state **cancel sibling tasks using errgroup**, the stable caller contract, the trade-off, how you tested the failure path and what you would do differently at a service boundary.

Record surprises and a confidence score in [NOTES.md](../../NOTES.md).

