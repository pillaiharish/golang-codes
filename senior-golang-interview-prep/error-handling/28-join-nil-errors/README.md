# Q28. Verify errors.Join with nil inputs
 
[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../27-join-validation-errors/README.md) · [Next](../29-multiple-percent-w/README.md)

**Level:** Intermediate | **Concepts:** nil filtering, conditional collection.

## Your task — code first

1. Call errors.Join(nil,nil) and compare against nil.
2. Call errors.Join(nil,ErrOne) and inspect it.
3. Explain why callers can append conditional errors without custom nil filtering.

## Failure cases and acceptance tests

- [ ] All-nil returns nil.
- [ ] A single non-nil child remains matchable.

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

