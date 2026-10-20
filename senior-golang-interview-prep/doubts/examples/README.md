# Runnable doubt examples

[Study index and commands](../README.md#runnable-examples)

Each subdirectory is a separate `main` package within the parent module. Run from `doubts` using `go run ./examples/NAME`.

| Package | Expected observation | Guide |
|---|---|---|
| [allocations](allocations/main.go) | `7 7`; inspect escape diagnostics separately | [Memory](../01-declarations-assignment-and-memory.md) |
| [values](values/main.go) | Nil slice append succeeds; nil map read misses; typed nil interface is non-nil | [Zero values](../02-zero-values-make-and-new.md) |
| [text](text/main.go) | 6 bytes, 3 runes, offsets 0/1/3; invalid byte decodes with width 1 | [UTF-8](../03-bytes-runes-and-utf8.md) |
| [loops](loops/main.go) | Body 0/2/3; init 1; declared closures 0/1/2; reused closures 3/3/3 | [Loops](../04-increments-and-for-loops.md) |
| [queues](queues/main.go) | First enqueue succeeds; full queue rejects; canceled send returns; closed queue drains | [Buffers](../05-buffers-blocking-and-backpressure.md) |
| [netpoll](netpoll/main.go) | Provide trace path argument; real loopback TCP read expires via deadline | [I/O](../06-network-io-and-blocking-syscalls.md) |
| [slices](slices/main.go) | Caught index panic; observed growth; bounded overflow rejected | [Slice growth](../07-slice-length-capacity-growth-and-limits.md) |

The examples are demonstrations, not a throughput benchmark or an automated correctness test suite. The netpoll example requires loopback socket permission. Its timeout duration is not a performance claim.
