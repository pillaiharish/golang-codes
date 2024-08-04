
# Two Pointers Algorithm
https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

## How It Works

### Initialization
Two pointers, `left` and `right`, are initialized to the beginning and end of the array, respectively.

### Iteration
1. The sum of the elements at the `left` and `right` pointers is calculated.
2. If the sum equals the target, the indices are returned (1-indexed).
3. If the sum is less than the target, increment the `left` pointer to increase the sum.
4. If the sum is greater than the target, decrement the `right` pointer to decrease the sum.

### Termination
The loop continues until the `left` pointer is no longer less than the `right` pointer.