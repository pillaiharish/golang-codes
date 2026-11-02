# Q2. Start five goroutines and wait for all of them

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../1-start-goroutine-and-wait/README.md) · [Next question](../3-goroutine-execution-order/README.md)

**Level:** 1 — fundamentals | **Concepts:** multiple goroutines, `sync.WaitGroup`, loop parameters, nondeterministic scheduling, completion versus ordering.

## Your task

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

## Scenario questions — answer before opening the solution

1. Run the same program 10 times — Must indexes print `0,1,2,3,4`?
2. `wg.Add(4)`, but five workers call `Done` — Is it safe?
3. `wg.Add(6)`, but only five workers call `Done` — Does `Wait` finish?
4. `wg.Add(1)` inside each launched goroutine — Can `Wait` return early?
5. Remove the `(i)` argument — Is that always a loop-capture bug?
6. Change five workers to zero — What does `Wait` do?
7. Require increasing output order — Does a WaitGroup do that?
8. Each worker appends to a shared slice — Is `WaitGroup` enough for thread safety?
9. Launch 100,000 workers — Is it a reasonable default for a server?
10. A worker panics — Will `WaitGroup` recover it?

## Cold-recall prompts

1. Why is `wg.Add(5)` before the loop valid? What alternative placement is safe?
2. Why pass `i` as a parameter, and what changed in Go 1.22?
3. Can `wg.Wait()` return while a worker is still inside its function if `Done()` is deferred?
4. Does protecting `append` with a WaitGroup eliminate a data race? Why not?
5. How would you preserve result order without serializing the worker computations?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q2 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
