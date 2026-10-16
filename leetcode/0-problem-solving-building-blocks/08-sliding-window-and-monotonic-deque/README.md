# 08 — Sliding windows and monotonic deques

[← Foundations index](../README.md) · [← Problem index](../../README.md)

## Recognition and intuition
For successive windows, reuse candidates instead of rescanning every k elements. A monotonic deque remembers surviving maxima.

## Operations to know first
- Store **indices** in the deque; indices distinguish values that expire.
- Before adding index `i`, discard a front index `<i-k+1` (equiv. `<= i-k`).
- Pop from the back while old values are `<=nums[i]`; newer equal/larger values dominate.
- Append i; once i>=k-1, output nums[front].
- Every index is inserted and deleted at most once: amortized O(n) operations.

## Runnable Go micro-example
```go
package main
import "fmt"
func maxWindows(a []int,k int) []int {
    if k<=0||k>len(a){return []int{}}
    q,result:=[]int{},[]int{}
    for i,v:=range a {
        if len(q)>0&&q[0]<=i-k {q=q[1:]}
        for len(q)>0&&a[q[len(q)-1]]<=v{q=q[:len(q)-1]}
        q=append(q,i)
        if i>=k-1 {result=append(result,a[q[0]])}
    }
    return result
}
func main(){fmt.Println(maxWindows([]int{1,3,-1,-3,5},3))}
```

## Hand trace
For `[1,3,-1,-3,5]`, k=3: outputs are 3 (window [1,3,-1]), 3 ([3,-1,-3]), 5 ([-1,-3,5]).

## Recall exercise (no notes)
Why can't you store only the values? Show when index 1 expires as i moves to 4 for k=3. Why does the nested while loop remain O(n)?

## Common mistakes
- Deque needs both an expiry rule and a dominance rule.
- The simple `q=q[1:]` Go slice version may retain or grow its backing array; O(k) *logical entries* is not a strict heap-memory guarantee.
- Use a circular deque for hard bounded storage, and handle invalid k before allocating result.

## Stronger problems using this block
- [#239 Sliding Window Maximum](../../239.%20Sliding%20Window%20Maximum/main.go)
- [#641 Design Circular Deque](../../641.%20Design%20Circular%20Deque/main.go)

## Review in 60 seconds
Name the invariant, rewrite the code without looking, then explain one mistake the tests should catch.
