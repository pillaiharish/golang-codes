# 15 — Subarrays, running sums, prefix sums and modulo

[← Foundations index](../README.md) · [← Problem index](../../README.md)

## Recognize the building block
A subarray is a contiguous range. Running sums build all subarray sums without re-summing from scratch, and prefix sums make arbitrary range sums a subtraction.

## Operations you must be able to write
- For each start `i`, set sum 0, advance end `j` and add `nums[j]` to emit every contiguous subarray sum.
- Prefix sum `prefix[i+1]=prefix[i]+nums[i]`; range `[l,r)` has sum `prefix[r]-prefix[l]`.
- There are exactly `n(n+1)/2` nonempty subarrays; enumerating all is Θ(n²) output size.
- Sort the sum list if the problem orders subarray sums, then sum the rank range `[left-1,right)` after checking bounds.
- For nonnegative values, modulo `1_000_000_007` reduces the final result; use sufficiently wide intermediate integers where needed.

## Complete small Go example
```go
package main
import ("fmt";"sort")
func main(){
    nums:=[]int{1,2,3}
    sums:=[]int{}
    for i:=0;i<len(nums);i++{
        sum:=0
        for j:=i;j<len(nums);j++{sum+=nums[j];sums=append(sums,sum)}
    }
    sort.Ints(sums)
    fmt.Println(sums) // [1 2 3 3 5 6]
    prefix:=[]int{0}
    for _,v:=range nums{prefix=append(prefix,prefix[len(prefix)-1]+v)}
    fmt.Println(prefix[3]-prefix[1]) // sum nums[1:3] = 5
}
```

## Dry run
For nums [1,2,3], subarrays: [1]=1,[1,2]=3,[1,2,3]=6,[2]=2,[2,3]=5,[3]=3. Sorted sums [1,2,3,3,5,6].

## Tiny practice task
Count how many contiguous subarrays length 4 has: 10. Generate sums for `[2,1]`: 2,3,1; sorted [1,2,3].

## Common mistakes
- Prefix sums and enumerating all subarrays solve different tasks: prefix sums don't sort the n(n+1)/2 sums.
- Beware integer overflow and modulo conventions when changing constraints.
- #1508's existing implementation is a direct enumeration+sort approach; advanced binary-search optimizations are not already implemented in the repository.

## Problem links from this repository
- [#1508 Range Sum of Sorted Subarray Sums](../../1508.%20Range%20Sum%20of%20Sorted%20Subarray%20Sums/main.go)
- [#441 Arranging Coins](../../441.%20Arranging%20Coins/main.go)

## Blank-screen check
Close this note. Restate the invariant, retype the operations, then solve the first linked problem without looking.
