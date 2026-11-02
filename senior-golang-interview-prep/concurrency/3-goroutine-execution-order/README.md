# Q3. Prove that goroutine execution order is not guaranteed

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../2-start-five-goroutines-and-wait/README.md) · [Next question](../4-goroutine-panic-and-recover/README.md)

**Level:** 1 — goroutine basics | **Focus:** scheduling versus synchronization.

## Your task

Launch three goroutines: A prints `A`, B prints `B`, and C prints `C`. Wait for all of them. Predict the output, run it repeatedly, and explain why launching A before B before C does **not** guarantee that print order. Then modify the program to guarantee `A -> B -> C`.

## Scenario questions — answer before opening the solution

1. Does `WaitGroup` print goroutines in launch order?
2. Does adding `time.Sleep` in B enforce order?
3. What if one of the ordered goroutines forgets to close its signal?
4. Can results be computed concurrently but printed in order?
5. Does `GOMAXPROCS(1)` force A/B/C?

## Cold-recall prompts

1. Distinguish *concurrency*, *parallelism*, *launch order*, *execution order*, and *completion order*.
2. Name the synchronization edge that guarantees B prints after A.
3. Rewrite the ordered solution using a single coordinator rather than chaining workers.

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q3 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
