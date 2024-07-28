/*
Given an m x n matrix, return all elements of the matrix in spiral order.

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

*/

package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	answer := spiralMatrix(matrix)
	fmt.Println(answer)
}

func spiralMatrix(matrix [][]int) []int {
	//int lenghth of matrix == 0 return nil array
	if len(matrix) == 0 {
		return []int{}
	}
	var result []int
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	//iterate through all matrix value till al value is added
	for top <= bottom && left <= right {
		// Traverse from left to right along the top row
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse from top to bottom along the right column
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		if top <= bottom {
			// Traverse from right to left along the bottom row
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			// Traverse from bottom to top along the left column
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}
	//return result
	return result
}
