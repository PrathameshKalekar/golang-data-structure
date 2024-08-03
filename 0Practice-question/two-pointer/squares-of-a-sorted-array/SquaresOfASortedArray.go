/*
977 Given an integer array arr sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

Example 1:

Input: arr = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].
Example 2:

Input: arr = [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Constraints:

1 <= arr.length <= 104
-104 <= arr[i] <= 104
arr is sorted in non-decreasing order.
*/
package main

import (
	"fmt"
)

func main() {
	arr := []int{-4, -1, 0, 3, 10}
	answer := SquareOfSortedArrays(arr)
	fmt.Println(answer)
}
func SquareOfSortedArrays(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	n := len(arr)

	//make a result length of an array
	result := make([]int, len(arr))
	// for i th value of result its a square of ith value of arr
	left, right := 0, n-1
	//looping
	for i := n - 1; i >= 0; i-- {
		leftSquare := arr[left] * arr[left]
		rightSquare := arr[right] * arr[right]
		// if left side bigger then add left sideat i th location else right
		if leftSquare > rightSquare {
			result[i] = leftSquare
			left++
		} else {
			result[i] = rightSquare
			right--
		}
	}
	// return the sorted array
	return result
}
