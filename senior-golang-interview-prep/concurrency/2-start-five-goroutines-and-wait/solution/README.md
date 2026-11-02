# Q2. Start five goroutines and wait for all of them — worked solution

[Back to Q2 question](../README.md) · [Concurrency index](../../README.md)

## Run the reference program

From this `solution/` directory:

```sh
go run .
go run -race .
```


The code is a learning reference; compare it against your own attempt before reading the explanations below.

[Concurrency index](../../README.md) · [Q1: one goroutine](../../1-start-goroutine-and-wait/README.md) · [Runnable solution](main.go) · [Revision notes](../../NOTES.md)

**Level:** 1 — fundamentals | **Concepts:** multiple goroutines, `sync.WaitGroup`, loop parameters, nondeterministic scheduling, completion versus ordering.

## Interview question

Write a standalone Go program that:

1. Launches **exactly five goroutines** in a loop.
2. Makes each goroutine print its **index**, from the set `{0, 1, 2, 3, 4}`.
3. Uses **`sync.WaitGroup`** to wait until **all five** have finished.
4. Prints `All goroutines completed` only after the five workers finish.

Do not solve this with `time.Sleep`, and do not assume the five worker lines will be ordered.

### One valid output

```text
Goroutine: 4
Goroutine: 1
Goroutine: 0
Goroutine: 3
Goroutine: 2
All goroutines completed
```

Many permutations of the first five lines are valid. **The contract is one appearance of every index, with the completion line last.**

## Try it first

Before opening [the solution](main.go), answer:

- How does `main` register five tasks before waiting?
- How does each goroutine get the right index?
- Why can completion order differ from launch order?
- What is the invariant immediately before `Wait` returns?

## Reference solution

The [standalone program](main.go) uses this essential sequence:

```go
var wg sync.WaitGroup
wg.Add(5)

for i := 0; i < 5; i++ {
    go func(index int) {
        defer wg.Done()
        fmt.Println("Goroutine:", index)
    }(i)
}

wg.Wait()
fmt.Println("All goroutines completed")
```

The `index` parameter receives a copy of `i` for that invocation. This works consistently across Go versions. Go 1.22 also changed the semantics of loop variables declared by the loop (`for i := ...`) to be per-iteration in modules using Go 1.22 or later; explicitly passing `i` is still clear and portable.

Run from this folder:

```sh
go run main.go
go run -race main.go
```

The counter starts at `0`, becomes `5` before workers launch, decreases by one for every completed worker, and eventually reaches `0`. `Wait` guarantees all five completion signals have occurred; it does **not** make the print order predictable.

## Scenario variations interviewers may ask

| Scenario | Question to reason through | Correct conclusion |
|---|---|---|
| Run the same program 10 times | Must indexes print `0,1,2,3,4`? | No; scheduler interleavings vary. Validate the set, not the order. |
| `wg.Add(4)`, but five workers call `Done` | Is it safe? | No. `Wait` may return prematurely, and a fifth `Done` can drive the counter negative and panic. |
| `wg.Add(6)`, but only five workers call `Done` | Does `Wait` finish? | No, unless some sixth completion is signaled; this is a typical deadlock. |
| `wg.Add(1)` inside each launched goroutine | Can `Wait` return early? | Yes; initial count may still be zero when `main` reaches `Wait`. |
| Remove the `(i)` argument | Is that always a loop-capture bug? | Not for a loop declaring `i` in a Go 1.22+ module; older loop semantics could capture one changing variable. Check module language version. |
| Change five workers to zero | What does `Wait` do? | With a zero counter, it returns immediately. |
| Require increasing output order | Does a WaitGroup do that? | No. Collect results and print them in order after joining, or design an explicit ordering protocol. |
| Each worker appends to a shared slice | Is `WaitGroup` enough for thread safety? | No. Concurrent appends to shared slice state require synchronization or disjoint indexed writes. |
| Launch 100,000 workers | Is it a reasonable default for a server? | Not necessarily. Use bounded concurrency/admission to control memory and downstream load. |
| A worker panics | Will `WaitGroup` recover it? | No. Deferred `Done` can execute, but an unrecovered panic still brings down the process. |

## Understanding the lifecycle

| Event | Outstanding counter |
|---|---:|
| Create zero-value `WaitGroup` | 0 |
| `wg.Add(5)` before launching | 5 |
| First worker completes | 4 |
| Three more workers complete | 1 |
| Last worker completes | 0 |
| `wg.Wait()` returns | 0 |

The intermediate *identity* of the completing worker is deliberately unspecified. This is a conceptual trace, not a promise about which worker runs first.

## Acceptance criteria and a quick manual check

- [ ] Exactly five worker lines are printed.
- [ ] The indexes, after sorting, are `[0,1,2,3,4]` with no duplicates.
- [ ] `All goroutines completed` is always last.
- [ ] There is no sleep-based synchronization.
- [ ] You can explain why changing the loop to 100 goroutines scales the number of tasks, but does not impose an execution order.

The program performs `O(n)` task launches and `O(n)` total printing work for `n` goroutines; it creates `O(n)` goroutines, while the WaitGroup's counter itself uses constant space. This does not claim constant *total* memory.

## Cold-recall checks

1. Why is `wg.Add(5)` before the loop valid? What alternative placement is safe?
2. Why pass `i` as a parameter, and what changed in Go 1.22?
3. Can `wg.Wait()` return while a worker is still inside its function if `Done()` is deferred?
4. Does protecting `append` with a WaitGroup eliminate a data race? Why not?
5. How would you preserve result order without serializing the worker computations?

**Mastery evidence:** Rebuild the program without copying; validate multiple runs; explain one wrong-counter failure schedule and record the results in [NOTES.md](../../NOTES.md). Having this reference solution in the repository is not a claim of completion.

## References

- [`sync.WaitGroup` reference](https://pkg.go.dev/sync#WaitGroup)
- [Go 1.22 loop-variable change](https://go.dev/blog/loopvar-preview)
- [Go memory model](https://go.dev/ref/mem)

