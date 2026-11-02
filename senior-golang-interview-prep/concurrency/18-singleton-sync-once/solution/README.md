# Q18. Initialize a singleton using `sync.Once` — worked solution

[Back to Q18 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q17: shared map](../../17-protect-shared-map/README.md) · [Revision notes](../../NOTES.md)

**Level 4 — shared state and synchronization.**

## Task

Let many goroutines call `getInstance()` concurrently, but create the underlying configuration object exactly once.

## Reference answer

```go
var once sync.Once
var instance *Config

func getInstance() *Config {
    once.Do(func() { instance = loadConfig() })
    return instance
}
```

`once.Do(f)` executes `f` at most once for a given `Once`; all callers return after that execution completes. This synchronizes publication of the initialized object. The naive `if instance == nil { instance = loadConfig() }` races and may initialize more than once.

## Scenario Q&A

| Question | Answer |
|---|---|
| What if the initializer panics? | `Once` treats it as having returned; future calls do not retry with that `Once`. |
| Can `f` recursively call `once.Do` on the same `Once`? | It deadlocks waiting for its own initialization to finish. |
| Is the initialized object's mutable state automatically safe? | No. `Once` protects construction/publication, not subsequent writes. |
| What if initialization can return an error? | Capture an `err` in the closure or consider `sync.OnceValues`/explicit retry design; choose whether failures should be cached. |

**Recall:** Distinguish *only once* from *retry until success*. Explain what to do for a configuration that must be refreshed dynamically.

Reference: [`sync.Once`](https://pkg.go.dev/sync#Once).
