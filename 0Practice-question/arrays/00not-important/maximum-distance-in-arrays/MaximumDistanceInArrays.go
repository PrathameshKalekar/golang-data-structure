/*
You are given m arrays, where each array is sorted in ascending order.

You can pick up two integers from two different arrays (each array picks one) and calculate the distance. We define the distance between two integers a and b to be their absolute difference |a - b|.

Return the maximum distance.

Example 1:

Input: arrays = [[1,2,3],[4,5],[1,2,3]]
Output: 4
Explanation: One way to reach the maximum distance 4 is to pick 1 in the first or third array and pick 5 in the second array.
Example 2:

Input: arrays = [[1],[1]]
Output: 0
*/
package main

import (
	"fmt"
)

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}
	answer := maximumDistanceInArrays(arr)
	fmt.Println(answer)
}

func maximumDistanceInArrays(arrs [][]int) int {
	if len(arrs) == 0 || len(arrs) == 1 {
		return 0
	}

	// Initialize min and max based on the first array
	minValue := arrs[0][0]
	maxValue := arrs[0][len(arrs[0])-1]

	// Track the maximum distance
	maxDistance := 0

	for i := 1; i < len(arrs); i++ {
		// Update the max distance by comparing the current array's elements
		// with min and max from previous arrays
		maxDistance = max(maxDistance, arrs[i][len(arrs[i])-1]-minValue) // Compare max of current array with global min
		maxDistance = max(maxDistance, maxValue-arrs[i][0])              // Compare min of current array with global max

		// Update minValue and maxValue for future comparisons
		minValue = min(minValue, arrs[i][0])
		maxValue = max(maxValue, arrs[i][len(arrs[i])-1])
	}

	return maxDistance
}

// Helper function to find the maximum of two values
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to find the minimum of two values
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
