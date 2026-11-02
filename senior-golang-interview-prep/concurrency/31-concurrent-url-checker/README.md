# Q31. Build a concurrent URL checker

[Concurrency index](../README.md) · [Q20: semaphore](../20-semaphore-limit-concurrency/README.md) · [Revision notes](../NOTES.md)

**Level 6 — interview-sized capstone.**

## Task

Given a list of URLs, check them concurrently, limit active HTTP requests to N, use a per-request timeout, collect one result per input, report errors, and support cancellation. Preserve input identity even when completions arrive out of order.

## Reference design

```go
type Result struct {
    Index int
    URL string
    StatusCode int
    Err error
}
```

1. Validate inputs and choose how to handle duplicate URLs.
2. Use a bounded worker pool or acquire a permit **before** launching each goroutine.
3. Construct requests with `http.NewRequestWithContext`; derive per-request deadlines from a parent context.
4. Reuse one configured `http.Client` instead of a new client per URL. Close every non-nil response body.
5. Write each worker outcome to a disjoint `results[index]` slot, join all workers, and return results in input order.
6. Define canceled/never-admitted versus attempted-and-failed results distinctly.

## Scenario Q&A

| Question | Answer |
|---|---|
| A URL hangs forever; what bounds it? | Request context/deadline and HTTP client transport timeouts; join after cooperative cancellation. |
| One request fails; cancel all? | Depends on contract: report all results or fail-fast via `errgroup`; state the choice. |
| Can shared `results[index]` writes race? | Disjoint indices are safe when each slot has one writer and callers read after joining. Shared `append` without a lock is not. |
| How does the caller know which URL produced a result? | Include index/URL in the result rather than relying on receive order. |
| What if URLs come from untrusted input? | Enforce SSRF protections and destination/redirect policy before issuing network requests. |

**Recall:** Test timeout, cancellation while waiting for admission, HTTP error responses, malformed URL, empty list, and max active requests.

References: [`http.NewRequestWithContext`](https://pkg.go.dev/net/http#NewRequestWithContext), [`context` package](https://pkg.go.dev/context).
