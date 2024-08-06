/*
35 Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: arr = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: arr = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: arr = [1,3,5,6], target = 7
Output: 4
*/
package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 6}
	answer := searchInsertPosition(arr, 7)
	fmt.Println(answer)
}

func searchInsertPosition(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
