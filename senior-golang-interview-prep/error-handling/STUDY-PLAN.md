# Six-week error-handling study plan

[Back to questions](README.md) · [Capstone projects](projects/README.md) · [Revision notes](NOTES.md)

| Week | Questions | Foundation | Build |
|---|---|---|---|
| 1 | Q1–Q10 | interface, nil, errors.New/Errorf, context, defer, CLI | CLI calculator |
| 2 | Q11–Q20 | sentinel, typed errors, Is/As, typed nil, panic, table tests | User store |
| 3 | Q21–Q30 | %w, Unwrap, chains, Join, regression tests | Config loader |
| 4 | Q31–Q40 | structural tests, slog, context, adapters, linters | File processor |
| 5 | Q41–Q50 | goroutines, errgroup, retry, API/DB mapping, observability | URL health checker |
| 6 | Q51–Q60 | taxonomy, compatibility, SLO, error storms, migration, design review | Order/payment service |

## Study loop per question

1. Cold recall: 5 minutes reading only the question, then articulate the error contract.
2. Implement: write code in your own scratch module.
3. Attack: test at least one failure path and one success path. Include cancellation, unwrap behavior or a typed assertion where relevant.
4. Reveal: compare the worked solution; identify one actual misconception.
5. Retest: rerun without hints after 1, 3, 7 and 14 days. Record confidence 1–5 in NOTES.md.

## Run one solution

```bash
cd 16-typed-nil-error/solution
go test ./...
go vet ./...
go test -race ./...
go run .
```

## Run all 60 independent modules

From the `error-handling/` track directory:

```bash
bash verify.sh
```

Each solution has its own `go.mod`, so running `go test ./...` from the track root alone does **not** visit the nested modules. The script changes into each solution directory. Use `bash verify.sh --lint` to invoke `staticcheck` and `errcheck` when they are installed.

The two dependency-backed examples (Q42 `golang.org/x/sync/errgroup` and Q48 `google.golang.org/grpc/status`) may require downloading modules on the first run; if Go requests missing checksums, run `go mod tidy` in those individual solution directories.

