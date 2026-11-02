# Q17. Protect a shared map — worked solution

[Back to Q17 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q16: data race](../../16-data-race-counter/README.md) · [Revision notes](../../NOTES.md)

**Level 4 — shared state and synchronization.**

## Task

Read and write one ordinary Go map from several goroutines. Observe the unsafe behavior with the race detector, then build versions using `sync.Mutex`, `sync.RWMutex`, and `sync.Map`.

## Reference answer

```go
type Store struct {
    mu sync.RWMutex
    values map[string]int
}
func (s *Store) Set(k string, v int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.values[k] = v
}
func (s *Store) Get(k string) (int, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    v, ok := s.values[k]
    return v, ok
}
```

Initialize `values` with `make(map[string]int)` before writing. Plain maps are not safe for unsynchronized concurrent reads and writes. `Mutex` is the simpler default; `RWMutex` allows concurrent readers but can add overhead. `sync.Map` has specialized optimizations (such as write-once/read-many or mostly disjoint key sets), and its API/semantics differ from a normal map.

## Scenario Q&A

| Question | Answer |
|---|---|
| Can many goroutines read a map with no writer? | Concurrent reads of a map that is not being mutated are safe. |
| Does `RWMutex` always outperform `Mutex`? | No. Benchmark the workload; lock hold time, contention and write ratio matter. |
| Is `sync.Map` always ideal for read-heavy traffic? | No. Its documented specialization and workload shape matter; use a typed map plus lock for ordinary cases. |
| What about `if _, ok := m[k]; !ok { m[k] = v }`? | Lock the whole check-then-act operation, not individual reads/writes separately. |

**Recall:** Explain the zero-value map versus an initialized map, and why protecting one field but not a related invariant can remain incorrect.

References: [`sync.Map`](https://pkg.go.dev/sync#Map), [`sync.RWMutex`](https://pkg.go.dev/sync#RWMutex).
