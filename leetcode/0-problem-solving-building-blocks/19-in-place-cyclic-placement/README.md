# 19 — In-place cyclic placement and the duplicate guard

[← Foundations index](../README.md) · [← LeetCode index](../../README.md)

## Mental model
When values have a natural index, the input slice itself can become a lookup table. This is how First Missing Positive achieves O(n) time and O(1) auxiliary space.

## Operations and reusable rules
- For length n, only 1..n can occupy positions in the candidate answer range.
- Correct destination of value x is index x-1.
- At position i: while 1<=nums[i]<=n and nums[nums[i]-1]!=nums[i], swap nums[i] into its target.
- After a swap, do not increment i until the new current item also has a chance to move.
- Finally scan i from 0; first `nums[i]!=i+1` returns i+1; else answer n+1.

## Runnable example
```go
package main
import "fmt"
func missingPositive(a []int)int{
    n:=len(a)
    for i:=0;i<n;i++{
        for a[i]>=1&&a[i]<=n&&a[a[i]-1]!=a[i]{
            j:=a[i]-1
            a[i],a[j]=a[j],a[i]
        }
    }
    for i,v:=range a{if v!=i+1{return i+1}}
    return n+1
}
func main(){a:=[]int{3,4,-1,1};fmt.Println(missingPositive(a),a)}
```

## Trace it on paper
`[3,4,-1,1]` → put 3 at index 2 → `[-1,4,3,1]` → put 4 at index 3 → `[-1,1,3,4]` → put 1 at index 0 → `[1,-1,3,4]` → missing index 1 means answer 2.

## Pre-problem drill
For `[1,1]`, explain why checking the destination value prevents swapping forever. For `[1,2,3]`, explain why answer is 4.

## Pitfalls to recall
- The algorithm intentionally mutates input.
- Without target-slot inequality, duplicate values can create an infinite loop.
- Each successful swap fixes at least one position, proving O(n) swaps overall.

## Problems that depend on it
- [#41 First Missing Positive](../../41.%20First%20Missing%20Positive/main.go)

## Cold recall
Write the minimum state and invariant, then rewrite the example in Go and test one edge case.
