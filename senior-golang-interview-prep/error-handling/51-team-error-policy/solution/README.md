# Q51. Define and enforce a team error policy — worked solution

[Question (try first)](../README.md) · [Error-handling index](../../README.md) · [Go implementation](main.go) · [Tests](main_test.go)

**Level:** Senior | **Concepts:** return, wrap, log, panic, public vs internal.

## Run

From this `solution/` directory:

```bash
go test ./...
go test -race ./...
go run .
```

Go 1.21 or newer is recommended.

## Explanation and trade-offs

Policy should require actionable error contracts, context only at meaningful boundaries, one logging owner, and explicit panic exceptions. Linters enforce unchecked results; reviewers judge taxonomy and leaked details.

## Reference code and tests

Study [main.go](main.go) and [main_test.go](main_test.go) after your own attempt. Explain every failure return, wrapping edge, and caller decision.

### Verify these cases

- An error is preserved at a boundary.
- A nil error stays nil.

## Failure analysis

Modify one error path and predict the failing test. Consider an ignored return, a broken wrapper, a wrong classification, or an unnecessary retry where relevant to this exercise.

## Senior-level explanation

In about one minute: state **define and enforce a team error policy**, the stable caller contract, the trade-off, how you tested the failure path and what you would do differently at a service boundary.

Record surprises and a confidence score in [NOTES.md](../../NOTES.md).

