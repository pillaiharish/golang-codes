# Q58. Benchmark hot-path error creation — worked solution

[Question (try first)](../README.md) · [Error-handling index](../../README.md) · [Go implementation](main.go) · [Tests](main_test.go)

**Level:** Senior | **Concepts:** allocations, sentinel vs formatting, benchmarking.

## Run

From this `solution/` directory:

```bash
go test ./...
go test -race ./...
go run .
```

Go 1.21 or newer is recommended.

## Explanation and trade-offs

Formatting usually does more work than returning a prebuilt error, but measure on the relevant Go/compiler/runtime. Avoid changing public error semantics solely for microbenchmark wins.

## Reference code and tests

Study [main.go](main.go) and [main_test.go](main_test.go) after your own attempt. Explain every failure return, wrapping edge, and caller decision.

### Verify these cases

- Both strategies return valid errors.
- Benchmark functions isolate each path.

## Failure analysis

Modify one error path and predict the failing test. Consider an ignored return, a broken wrapper, a wrong classification, or an unnecessary retry where relevant to this exercise.

## Senior-level explanation

In about one minute: state **benchmark hot-path error creation**, the stable caller contract, the trade-off, how you tested the failure path and what you would do differently at a service boundary.

Record surprises and a confidence score in [NOTES.md](../../NOTES.md).

