/*
962 A ramp in an integer array nums is a pair (i, j) for which i < j and nums[i] <= nums[j]. The width of such a ramp is j - i.

Given an integer array nums, return the maximum width of a ramp in nums. If there is no ramp in nums, return 0.

Example 1:

Input: nums = [6,0,8,2,1,5]
Output: 4
Explanation: The maximum width ramp is achieved at (i, j) = (1, 5): nums[1] = 0 and nums[5] = 5.
Example 2:

Input: nums = [9,8,1,0,1,9,4,0,4,1]
Output: 7
Explanation: The maximum width ramp is achieved at (i, j) = (2, 9): nums[2] = 1 and nums[9] = 1.
*/
package main

import (
	"fmt"
)

func maximumWidthRamp(nums []int) int {
	n := len(nums)
	stack := []int{}

	// Step 1: Populate the stack with indices of nums in increasing order
	for i := 0; i < n; i++ {
		if len(stack) == 0 || nums[stack[len(stack)-1]] > nums[i] {
			stack = append(stack, i)
		}
	}

	maxWidth := 0

	// Step 2: Iterate from the end of nums to find the maximum width ramp
	for j := n - 1; j >= 0; j-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[j] {
			i := stack[len(stack)-1]
			maxWidth = max(maxWidth, j-i)
			stack = stack[:len(stack)-1] // Remove the used index
		}
	}

	return maxWidth
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	example1 := []int{6, 0, 8, 2, 1, 5}
	example2 := []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}

	fmt.Println(maximumWidthRamp(example1)) // Output: 4
	fmt.Println(maximumWidthRamp(example2)) // Output: 7
}
