# 17 — Complexity, amortization, and input contracts

[← Foundations index](../README.md) · [← LeetCode index](../../README.md)

## Mental model
Complexity claims should follow the actual work done by the code, not a memorized pattern label. Correctness also depends on input promises such as sortedness, valid k, distinct values and lowercase alphabets.

## Operations and reusable rules
- O(n²) pair loops differ from O(n) two pointers where indices move only forward/backward.
- A map lookup is typically expected O(1), not an absolute worst-case guarantee.
- A monotonic deque pushes and removes each index once: total O(n), despite an inner loop.
- Enumerating all n(n+1)/2 subarrays costs Ω(n²); sorting that list costs O(n² log n).
- A recursive tree traversal takes O(n) time; extra stack depends on height h, O(h).
- Memory includes results, map key/entry overhead, retained slice backing arrays, and recursion stack; distinguish auxiliary space from output space.
- Test empty inputs, one-element inputs, repeated values, negative numbers, Unicode and integer overflow only where allowed or relevant.

## Runnable example
```go
package main
import "fmt"
func countSubarrays(n int)int{
    count:=0
    for i:=0;i<n;i++{
        for j:=i;j<n;j++{count++}
    }
    return count
}
func main(){
    for n:=1;n<=4;n++{
        fmt.Println(n,countSubarrays(n),n*(n+1)/2)
    }
    // 1 1 1; 2 3 3; 3 6 6; 4 10 10
}
```

## Trace it on paper
For n=4, starts 0,1,2,3 have 4,3,2,1 subarrays respectively → 10 total; these are distinct intervals, not necessarily distinct sums.

## Pre-problem drill
State the complexities of #1, #239, #1508 and #94, including the memory not just time. Why is worst-case LRU map lookup not formally constant?

## Pitfalls to recall
- A slice operation `q=q[1:]` is O(1), but repeated use can retain memory and cause later growth.
- Complexity of sorting **N=n(n+1)/2 values** is O(N log N), not O(n log n).
- A Go `int` can overflow in expressions such as `mid*(mid+1)/2` with large values.

## Problems that depend on it
- [#1 Two Sum](../../1.%20Two%20Sum/main.go)
- [#41 First Missing Positive](../../41.%20First%20Missing%20Positive/main.go)
- [#94 Binary Tree Inorder Traversal](../../94.%20Binary%20Tree%20Inorder%20Traversal/main.go)
- [#1508 Range Sum of Sorted Subarray Sums](../../1508.%20Range%20Sum%20of%20Sorted%20Subarray%20Sums/main.go)
- [#239 Sliding Window Maximum](../../239.%20Sliding%20Window%20Maximum/main.go)
- [#441 Arranging Coins](../../441.%20Arranging%20Coins/main.go)

## Cold recall
Write the minimum state and invariant, then rewrite the example in Go and test one edge case.
