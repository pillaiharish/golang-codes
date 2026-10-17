# 6. Goroutine blocking on network I/O versus blocking system calls

[All doubts](README.md) · [Previous](05-buffers-blocking-and-backpressure.md)

## Direct answer

Network I/O also uses system calls. The useful distinction is **poller-integrated I/O that parks the goroutine** versus **an operation that blocks its OS thread in the kernel or foreign code**. Both can look synchronous to the Go caller. Do not equate a blocking Go API with a thread blocked for the entire wait.

## G, M, and P

- **G:** a goroutine, including its stack and scheduling state.
- **M:** an OS thread capable of executing Go or being inside a syscall.
- **P:** runtime resources/permission required to execute ordinary Go code. `GOMAXPROCS` limits how many Ps execute Go code simultaneously, not total goroutines, sockets, or OS threads.

A parked goroutine still retains its stack and reachable data. “Does not occupy an OS thread during the wait” does not mean “has no resource cost.”

## Typical network read on Unix

For an ordinary poller-managed TCP connection, conceptually:

1. G runs on M with a P and calls `conn.Read(buf)`.
2. The runtime/library attempts a read on a nonblocking descriptor.
3. If bytes are ready, the read may finish immediately.
4. If it would block (`EAGAIN`), the poller integration arranges readiness notification and parks G.
5. M can run another runnable G. There is no dedicated blocked OS thread for each idle socket.
6. Readiness wakes G, which becomes runnable and retries the operation when scheduled.

Readiness is not the same as completed application processing. The implementation coordinates registration, retries and wakeups to avoid lost notifications; these steps are a conceptual model, not a literal syscall-by-syscall trace. Linux commonly uses epoll, BSD/macOS kqueue; Windows uses a different completion-based integration. Details depend on OS and Go version.

The poller itself may wait in an OS call, but it multiplexes many descriptors. That is different from one thread waiting for each connection. Writes can also wait when socket send buffers fill.

## A truly blocking syscall

A regular-file operation may block an M in the kernel. For a runtime-aware syscall path, the runtime records the transition. The P can be handed off or retaken so another M can execute Go work. A known-blocking path can hand off promptly; other syscall transitions permit later P retaking. Do not assume every syscall immediately creates a thread or immediately releases its P in exactly the same way.

When the syscall returns, the G/M needs a P to resume ordinary Go execution. If none is available, G may need scheduling before it can continue. Meanwhile the original OS thread really was occupied by the syscall. Large numbers of concurrent blocking calls can therefore increase thread counts and resource pressure.

## Comparison and exceptions

| Operation | Typical wait | What remains occupied | Caveat |
|---|---|---|---|
| Poller-managed TCP read/write | Park G until readiness/completion | G state, buffers, socket | Fast path may not wait at all |
| Channel send/receive | Park G until counterpart/space/value | G state and referenced values | No kernel I/O needed for the channel wait |
| Regular-file read | Can block M in syscall | G plus OS thread for duration | Page cache can make it quick; disk/filesystem matter |
| Pollable pipe | Can use poller | Depends on descriptor/platform | Not every file descriptor behaves like a regular disk file |
| Blocking cgo call | M occupied in foreign code | OS thread and call state | Runtime can schedule Go elsewhere; foreign code has its own behavior |
| CPU-bound Go loop | G runs on M with P | CPU execution slot | Preemption helps scheduling, not service capacity |

A name like `Read` does not identify the mechanism. Inspect the concrete implementation and platform. `net.Pipe` is an in-memory connection and cannot prove OS network-poller behavior. DNS resolution can use a Go resolver or platform/cgo mechanisms. Raw low-level syscalls can bypass normal runtime integration; do not infer behavior from the word “syscall” alone.

## Cancellation and deadlines

A context does not magically cancel every blocking call. Use APIs that honor it, such as a context-aware dial/request. For a `net.Conn`, set read/write deadlines or close the connection as appropriate. Closing a connection unblocks pending reads/writes with errors according to the API contract. Regular file deadlines are not universally supported and a blocked file/cgo operation may not be cancelable in the same way.

A timeout around “wait for a result channel” can stop the caller waiting while leaving the underlying worker blocked. Always ask who stops the operation and who owns cleanup.

## Reproducible observation lab

[The TCP example](examples/netpoll/main.go) opens a real loopback TCP connection. Its peer holds the connection open without sending bytes; the reader uses a deadline to end the wait. It sets `GOMAXPROCS(1)` to make the concurrency question concrete and records a runtime trace.

From this directory:

```sh
go run ./examples/netpoll /tmp/doubts-netpoll.trace
go tool trace /tmp/doubts-netpoll.trace
```

Expected behavior: the read ends with a timeout, the program exits, and the trace shows a network wait rather than CPU execution for that goroutine. The exact wait duration and trace labels are not assertions. Loopback avoids an external service dependency. The program cleans up its sockets and waits for its helper goroutine.

For the syscall comparison on Linux, use an isolated test environment and a deliberately slow regular-file/filesystem workload. Capture a runtime trace plus an OS thread/syscall trace (for example `strace -f -T` if installed). Compare where the goroutine waits, the syscall duration, and thread counts. Do not use `time.Sleep` as proof of blocked disk I/O, and do not assume cached reads will block long enough to observe. This repository does not include a fake deterministic disk-blocking demo because storage timing is environment-dependent.

Useful tools answer different questions: goroutine dumps show wait sites; execution traces show scheduling/I/O waits; CPU profiles mostly show executing work, not time parked; OS tools show threads and kernel calls. Record Go version, OS, concrete API, and workload when reporting a finding.

## Why this does not fix a serial Kafka consumer

A consumer waiting for Kafka may release its OS thread through the network poller while the **consumer goroutine itself remains unable to read the next item**. Efficient scheduling saves threads; it does not make that serial worker process several messages simultaneously. Batching or bounded concurrency must be explicitly supported by the application/client, with ordering and acknowledgement trade-offs checked.

## Interview checkpoints

- Why can 10,000 idle TCP connections use far fewer than 10,000 threads?
- Why does blocking disk I/O not necessarily stop all Go work with `GOMAXPROCS=1`?
- Why does “network calls are syscalls” not contradict poller integration?
- Why can low CPU utilization coexist with queue overflow in a serial publisher?
- Which mechanism actually stops your canceled I/O operation?

Related: [scheduler module](../1-goroutines-and-go-scheduler/README.md), [shutdown module](../4-context-cancellation-and-shutdown/README.md).

Sources: [net.Conn contract](https://pkg.go.dev/net#Conn), [runtime poller source](https://go.dev/src/runtime/netpoll.go), [runtime scheduling/syscall source](https://go.dev/src/runtime/proc.go), [internal poll I/O](https://go.dev/src/internal/poll/fd_unix.go), [os.File deadlines](https://pkg.go.dev/os#File.SetDeadline), [execution tracing](https://pkg.go.dev/runtime/trace). Runtime implementation links describe implementation details, not permanent language guarantees.
