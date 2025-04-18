# Go CPU and Memory Profiling
This is a simple Go program demonstrating how to perform CPU and memory profiling using Go's built-in `pprof` package. It also exposes a pprof HTTP server for live profiling and debugging.

---

## Features

- Starts an HTTP server on `localhost:6060` to serve pprof endpoints for live profiling.
- Creates CPU and heap profile files (`cpu.prof` and `heap.prof`).
- Simulates CPU load by running a loop.
- Forces garbage collection to get accurate heap statistics.
- Writes heap profile to file.
- Keeps the program running indefinitely to allow inspection via HTTP.

---

## Profile Capture
1. **Access live profiling data:**

Open your browser and navigate to [http://localhost:6060/debug/pprof/](http://localhost:6060/debug/pprof/) to see runtime profiling data such as goroutine, heap, threadcreate, and CPU profiles.

2. **Analyze profile files:**

The program generates two files in the working directory:

- `cpu.prof` — CPU profile
- `heap.prof` — Memory (heap) profile

You can analyze these profiles using the Go tool:
```bash
go tool pprof cpu.prof
go tool pprof heap.prof
```
Inside the pprof interactive shell, you can run commands like `top`, `list`, or `web` to visualize profiling data.

---

## Notes

- The CPU profiling starts immediately and stops after the simulated load completes.
- The heap profile is written after forcing garbage collection to ensure accurate memory usage data.
- The HTTP server runs concurrently and remains active, allowing you to connect anytime to inspect live profiling data.
- Use the profiling data to identify performance bottlenecks and optimize your Go applications.

---

## References

- [Go pprof documentation](https://pkg.go.dev/net/http/pprof)
- [Profiling Go programs](https://blog.golang.org/pprof)
- [Go runtime/pprof package](https://pkg.go.dev/runtime/pprof)