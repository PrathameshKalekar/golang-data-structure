/*
2419 You are given an integer array nums of size n.

Consider a non-empty subarray from nums that has the maximum possible bitwise AND.

In other words, let k be the maximum value of the bitwise AND of any subarray of nums. Then, only subarrays with a bitwise AND equal to k should be considered.
Return the length of the longest such subarray.

The bitwise AND of an array is the bitwise AND of all the numbers in it.

A subarray is a contiguous sequence of elements within an array.

Example 1:

Input: nums = [1,2,3,3,2,2]
Output: 2
Explanation:
The maximum possible bitwise AND of a subarray is 3.
The longest subarray with that value is [3,3], so we return 2.
Example 2:

Input: nums = [1,2,3,4]
Output: 1
Explanation:
The maximum possible bitwise AND of a subarray is 4.
The longest subarray with that value is [4], so we return 1.
*/
package main

import "fmt"

// Function to find the length of the longest subarray with maximum bitwise AND
func longestSubarray(nums []int) int {
	maxVal := nums[0]
	maxLength := 0
	currentLength := 0

	// Step 1: Find the maximum value in the array
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}

	// Step 2: Find the longest subarray where all elements are equal to maxVal
	for _, num := range nums {
		if num == maxVal {
			currentLength++
			if currentLength > maxLength {
				maxLength = currentLength
			}
		} else {
			currentLength = 0
		}
	}

	return maxLength
}

func main() {
	// Test case 1
	nums1 := []int{1, 2, 3, 3, 2, 2}
	fmt.Println("Output for [1, 2, 3, 3, 2, 2]:", longestSubarray(nums1)) // Output: 2

	// Test case 2
	nums2 := []int{1, 2, 3, 4}
	fmt.Println("Output for [1, 2, 3, 4]:", longestSubarray(nums2)) // Output: 1
}
