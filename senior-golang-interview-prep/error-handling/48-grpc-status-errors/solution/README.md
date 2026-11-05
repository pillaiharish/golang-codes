# Q48. Return idiomatic gRPC status codes — worked solution

[Question (try first)](../README.md) · [Error-handling index](../../README.md) · [Go implementation](main.go) · [Tests](main_test.go)

**Level:** Advanced | **Concepts:** grpc/status, codes, retry semantics.

## Run

From this `solution/` directory:

```bash
go test ./...
go test -race ./...
go run .
```

Go 1.21 or newer is recommended. This example declares an external module dependency.

## Explanation and trade-offs

gRPC status codes are transport-level classifications; attach safe details via status.WithDetails when needed. Retrying Unavailable can be appropriate for idempotent calls, but never blindly retry canceled or invalid requests.

## Reference code and tests

Study [main.go](main.go) and [main_test.go](main_test.go) after your own attempt. Explain every failure return, wrapping edge, and caller decision.

### Verify these cases

- NotFound is classifiable.
- InvalidArgument remains distinct from Internal.

## Failure analysis

Modify one error path and predict the failing test. Consider an ignored return, a broken wrapper, a wrong classification, or an unnecessary retry where relevant to this exercise.

## Senior-level explanation

In about one minute: state **return idiomatic grpc status codes**, the stable caller contract, the trade-off, how you tested the failure path and what you would do differently at a service boundary.

Record surprises and a confidence score in [NOTES.md](../../NOTES.md).

