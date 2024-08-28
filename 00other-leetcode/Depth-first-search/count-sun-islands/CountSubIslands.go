/*
You are given two m x n binary matrices grid1 and grid2 containing only 0's (representing water) and 1's (representing land). An island is a group of 1's connected 4-directionally (horizontal or vertical). Any cells outside of the grid are considered water cells.

An island in grid2 is considered a sub-island if there is an island in grid1 that contains all the cells that make up this island in grid2.

Return the number of islands in grid2 that are considered sub-islands.

Example 1:

Input: grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]], grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
Output: 3
Explanation: In the picture above, the grid on the left is grid1 and the grid on the right is grid2.
The 1s colored red in grid2 are those considered to be part of a sub-island. There are three sub-islands.
Example 2:

Input: grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]], grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
Output: 2
Explanation: In the picture above, the grid on the left is grid1 and the grid on the right is grid2.
The 1s colored red in grid2 are those considered to be part of a sub-island. There are two sub-islands.
*/
package main

import "fmt"

// DFS function to explore the island and check if it's a sub-island
func dfs(grid1, grid2 [][]int, i, j int, m, n int) bool {
	// Check if the current position is out of bounds or water
	if i < 0 || i >= m || j < 0 || j >= n || grid2[i][j] == 0 {
		return true
	}

	// If grid1 does not have land where grid2 does, it's not a sub-island
	if grid1[i][j] == 0 {
		return false
	}

	// Mark the cell in grid2 as visited
	grid2[i][j] = 0

	// Check all 4 directions
	up := dfs(grid1, grid2, i-1, j, m, n)
	down := dfs(grid1, grid2, i+1, j, m, n)
	left := dfs(grid1, grid2, i, j-1, m, n)
	right := dfs(grid1, grid2, i, j+1, m, n)

	// Return true if all parts of this island are valid sub-islands
	return up && down && left && right
}

func countSubIslands(grid1, grid2 [][]int) int {
	if len(grid1) == 0 || len(grid2) == 0 {
		return 0
	}

	m, n := len(grid1), len(grid1[0])
	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				// If this island is a sub-island, increment the count
				if dfs(grid1, grid2, i, j, m, n) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	grid1 := [][]int{
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1},
		{1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0},
		{1, 0, 1, 0, 1},
	}

	grid2 := [][]int{
		{1, 1, 1, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
		{0, 1, 1, 1, 0},
		{1, 0, 1, 0, 1},
	}

	result := countSubIslands(grid1, grid2)
	fmt.Println("Number of sub-islands:", result) // Output: Number of sub-islands: 3
}
