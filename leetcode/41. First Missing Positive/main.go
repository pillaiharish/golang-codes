package main

import "fmt"

// firstMissingPositive rearranges nums in place.
// Put x in index x-1 for 1 <= x <= n, unless the target already holds x.
// Time O(n) amortized; auxiliary space O(1). Input is mutated.
func firstMissingPositive(nums []int) int {
    n:=len(nums)
    for i:=0;i<n;i++ {
        for nums[i]>=1 && nums[i]<=n && nums[nums[i]-1]!=nums[i] {
            target:=nums[i]-1
            nums[i],nums[target]=nums[target],nums[i]
        }
    }
    for i,v:=range nums {
        if v!=i+1 { return i+1 }
    }
    return n+1
}

func main() {
    nums:=[]int{3,4,-1,1}
    fmt.Println(firstMissingPositive(nums)) // 2
}
