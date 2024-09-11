/*
189 Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
*/
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Before rotation:", arr)
	rotateArray(arr, 2)
	fmt.Println("After rotation:", arr)
}

func rotateArray(arr []int, k int) {
	n := len(arr)
	k = k % n // Handle k larger than the array size
	if k == 0 {
		return
	}

	// Step 1: Reverse the entire array
	reverse(arr, 0, n-1)

	// Step 2: Reverse the first k elements
	reverse(arr, 0, k-1)

	// Step 3: Reverse the remaining elements
	reverse(arr, k, n-1)
}

// Helper function to reverse the slice in place
func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
