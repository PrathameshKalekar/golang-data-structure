/*
2326 You are given two integers m and n, which represent the dimensions of a matrix.

You are also given the head of a linked list of integers.

Generate an m x n matrix that contains the integers in the linked list presented in spiral order (clockwise), starting from the top-left of the matrix. If there are remaining empty spaces, fill them with -1.

Return the generated matrix.

Example 1:

Input: m = 3, n = 5, head = [3,0,2,6,8,1,7,9,4,2,5,5,0]
Output: [[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]
Explanation: The diagram above shows how the values are printed in the matrix.
Note that the remaining spaces in the matrix are filled with -1.
Example 2:

Input: m = 1, n = 4, head = [0,1,2]
Output: [[0,1,2,-1]]
Explanation: The diagram above shows how the values are printed from left to right in the matrix.
The last space in the matrix is set to -1.
*/
package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix4(m int, n int, head *ListNode) [][]int {
	// Initialize the m x n matrix with -1
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}

	// Define boundaries
	top, bottom, left, right := 0, m-1, 0, n-1
	current := head

	for current != nil && top <= bottom && left <= right {
		// Move right
		for i := left; i <= right && current != nil; i++ {
			matrix[top][i] = current.Val
			current = current.Next
		}
		top++

		// Move down
		for i := top; i <= bottom && current != nil; i++ {
			matrix[i][right] = current.Val
			current = current.Next
		}
		right--

		// Move left
		for i := right; i >= left && current != nil; i-- {
			matrix[bottom][i] = current.Val
			current = current.Next
		}
		bottom--

		// Move up
		for i := bottom; i >= top && current != nil; i-- {
			matrix[i][left] = current.Val
			current = current.Next
		}
		left++
	}

	return matrix
}

func main() {
	// Example 1:
	head := &ListNode{3, &ListNode{0, &ListNode{2, &ListNode{6, &ListNode{8,
		&ListNode{1, &ListNode{7, &ListNode{9, &ListNode{4, &ListNode{2,
			&ListNode{5, &ListNode{5, &ListNode{0, nil}}}}}}}}}}}}}
	m, n := 3, 5
	result := spiralMatrix4(m, n, head)
	fmt.Println(result) // Output: [[3 0 2 6 8] [5 0 -1 -1 1] [5 2 4 9 7]]

	// Example 2:
	head2 := &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
	m2, n2 := 1, 4
	result2 := spiralMatrix4(m2, n2, head2)
	fmt.Println(result2) // Output: [[0 1 2 -1]]
}
