# Q20. Write table-driven tests for ParseAge
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../19-must-pattern/README.md) · [Next](../21-percent-v-vs-percent-w/README.md)

**Level:** Beginner | **Concepts:** subtests, success/failure paths, errors.As.

## Your task — code first

1. Implement ParseAge with a wrapped strconv error.
2. Write table-driven subtests for valid and invalid input.
3. Assert a typed *strconv.NumError with errors.As.

## Failure cases and acceptance tests

- [ ] Valid input succeeds.
- [ ] Invalid and overflow inputs produce typed parse errors.

## Explain before revealing

- What is the behavior on both success and failure?
- Which error identity or typed details can the caller inspect?
- Where does wrapping, logging, cancellation or retry policy belong?
- Which apparently reasonable implementation could conceal a failure?

## Practice workflow

1. Implement independently without opening the answer.
2. Write tests for expected failures and successes.
3. Compare with the [worked solution](solution/README.md), then refactor.
4. Run tests and static checks; record evidence in [revision notes](../NOTES.md).

**Mastery is earned by reproducing the solution and explaining its trade-offs**, not by reading the reference.

