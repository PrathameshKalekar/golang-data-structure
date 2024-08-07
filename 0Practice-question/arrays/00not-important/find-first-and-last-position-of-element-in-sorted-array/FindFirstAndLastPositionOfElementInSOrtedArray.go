/*
34 Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:

Input: nums = [], target = 0
Output: [-1,-1]
*/
package main

import (
	"fmt"
)

func main() {
	arr := []int{}
	answer := findFirstAndLastPositionOfElementInSortedArray(arr, 1)
	fmt.Println(answer)
}

func findFirstAndLastPositionOfElementInSortedArray(arr []int, target int) []int {

	left, right := 0, len(arr)-1
	for left <= right {
		//first check the left and right value same return the them else manage the left right postition
		if arr[left] == target && arr[right] == target {
			return []int{left, right}
		}
		if arr[left] < target {
			left++
		}
		if arr[right] > target {
			right--
		}
	}
	// if all element done scan still the left right element not found return -1 -1
	return []int{-1, -1}
}
