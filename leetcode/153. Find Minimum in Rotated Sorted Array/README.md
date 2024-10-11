# 153. Find Minimum in Rotated Sorted Array

We can solve this problem using a modified binary search. The minimum element in a rotated sorted array must be in the unsorted part of the array, so we can keep narrowing down the search based on the mid element.

Binary Search: We use a binary search technique to minimize the search space. 

### Comparison with Right Element:
We compare the middle element with the rightmost element. If nums[mid] > nums[right], the minimum must be in the right half. Otherwise, the minimum is in the left half (including mid).

## Stopping Condition: 
The loop stops when left == right, which gives the index of the minimum element.

## Time Complexity:
O(log n): The array is divided into half at each step, leading to logarithmic time complexity.

## Space Complexity:
O(1): Constant space is used, as the algorithm only requires a few variables regardless of the input size.