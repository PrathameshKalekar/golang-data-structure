/*
840 A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to 9 such that each row, column, and both diagonals all have the same sum.

Given a row x col grid of integers, how many 3 x 3 contiguous magic square subgrids are there?

Note: while a magic square can only contain numbers from 1 to 9, grid may contain numbers up to 15.

Example 1:

Input: grid = [[4,3,8,4],[9,5,1,9],[2,7,6,2]]
Output: 1
Explanation:
The following subgrid is a 3 x 3 magic square:

while this one is not:

In total, there is only one magic square inside the given grid.
Example 2:

Input: grid = [[8]]
Output: 0
*/
package main

import (
	"fmt"
)

func numMagicSquaresInside(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	// Function to check if a 3x3 grid is a magic square
	isMagic := func(r, c int) bool {
		// Ensure all values are distinct and in the range [1, 9]
		seen := make([]bool, 10)
		for i := r; i < r+3; i++ {
			for j := c; j < c+3; j++ {
				num := grid[i][j]
				if num < 1 || num > 9 || seen[num] {
					return false
				}
				seen[num] = true
			}
		}

		// Check rows, columns, and diagonals sum to 15
		return grid[r][c]+grid[r][c+1]+grid[r][c+2] == 15 &&
			grid[r+1][c]+grid[r+1][c+1]+grid[r+1][c+2] == 15 &&
			grid[r+2][c]+grid[r+2][c+1]+grid[r+2][c+2] == 15 &&
			grid[r][c]+grid[r+1][c]+grid[r+2][c] == 15 &&
			grid[r][c+1]+grid[r+1][c+1]+grid[r+2][c+1] == 15 &&
			grid[r][c+2]+grid[r+1][c+2]+grid[r+2][c+2] == 15 &&
			grid[r][c]+grid[r+1][c+1]+grid[r+2][c+2] == 15 &&
			grid[r][c+2]+grid[r+1][c+1]+grid[r+2][c] == 15
	}

	// Check all possible 3x3 subgrids
	for r := 0; r <= rows-3; r++ {
		for c := 0; c <= cols-3; c++ {
			if isMagic(r, c) {
				count++
			}
		}
	}

	return count
}

func main() {
	grid1 := [][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	}
	fmt.Println(numMagicSquaresInside(grid1)) // Output: 1

	grid2 := [][]int{
		{8},
	}
	fmt.Println(numMagicSquaresInside(grid2)) // Output: 0
}
