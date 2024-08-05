/*
33 There is an integer array arr sorted in ascending order (with distinct values).

Prior to being passed to your function, arr is possibly rotated at an unknown pivot index k (1 <= k < arr.length) such that the resulting array is [arr[k], arr[k+1], ..., arr[n-1], arr[0], arr[1], ..., arr[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array arr after the possible rotation and an integer target, return the index of target if it is in arr, or -1 if it is not in arr.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: arr = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: arr = [4,5,6,7,0,1,2], target = 3
Output: -1
Example 3:

Input: arr = [1], target = 0
Output: -1
*/
package main

import "fmt"

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	index := searchInRotatedSortedArray(arr, 0)
	fmt.Println("Target found at index : ", index)
}
func searchInRotatedSortedArray(arr []int, target int) int {
	// Initialize pointers for binary search
	left, right := 0, len(arr)-1

	for left <= right {
		// Calculate the middle index
		mid := left + (right-left)/2

		// Check if the middle element is the target
		if arr[mid] == target {
			return mid
		}

		// Determine which side of the array is sorted
		if arr[left] <= arr[mid] { // Left side is sorted
			// Check if the target lies within the sorted left side
			if arr[left] <= target && target < arr[mid] {
				right = mid - 1 // Search in the left side
			} else {
				left = mid + 1 // Search in the right side
			}
		} else { // Right side is sorted
			// Check if the target lies within the sorted right side
			if arr[mid] < target && target <= arr[right] {
				left = mid + 1 // Search in the right side
			} else {
				right = mid - 1 // Search in the left side
			}
		}
	}

	// Return -1 if the target is not found
	return -1
}

// func searchInRotatedSortedArray(arr []int, target int) int {

// 	for i, val := range arr {
// 		if val == target {
// 			return i
// 		}
// 	}
//  return target - arr[0]
// }
