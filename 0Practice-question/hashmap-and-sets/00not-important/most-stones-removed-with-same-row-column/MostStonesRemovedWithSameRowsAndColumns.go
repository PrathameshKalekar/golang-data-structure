/*
947 On a 2D plane, we place n stones at some integer coordinate points. Each coordinate point may have at most one stone.

A stone can be removed if it shares either the same row or the same column as another stone that has not been removed.

Given an array stones of length n where stones[i] = [xi, yi] represents the location of the ith stone, return the largest possible number of stones that can be removed.

Example 1:

Input: stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
Output: 5
Explanation: One way to remove 5 stones is as follows:
1. Remove stone [2,2] because it shares the same row as [2,1].
2. Remove stone [2,1] because it shares the same column as [0,1].
3. Remove stone [1,2] because it shares the same row as [1,0].
4. Remove stone [1,0] because it shares the same column as [0,0].
5. Remove stone [0,1] because it shares the same row as [0,0].
Stone [0,0] cannot be removed since it does not share a row/column with another stone still on the plane.
Example 2:

Input: stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
Output: 3
Explanation: One way to make 3 moves is as follows:
1. Remove stone [2,2] because it shares the same row as [2,0].
2. Remove stone [2,0] because it shares the same column as [0,0].
3. Remove stone [0,2] because it shares the same row as [0,0].
Stones [0,0] and [1,1] cannot be removed since they do not share a row/column with another stone still on the plane.
Example 3:

Input: stones = [[0,0]]
Output: 0
Explanation: [0,0] is the only stone on the plane, so you cannot remove it.
*/
package main

import "fmt"

func removeStones(stones [][]int) int {
	// Create adjacency lists
	rowMap := make(map[int][]int)
	colMap := make(map[int][]int)

	for i, stone := range stones {
		row := stone[0]
		col := stone[1]
		rowMap[row] = append(rowMap[row], i)
		colMap[col] = append(colMap[col], i)
	}

	visited := make([]bool, len(stones))
	var dfs func(int)

	// Depth-First Search function to explore all stones in the same component
	dfs = func(node int) {
		visited[node] = true
		row := stones[node][0]
		col := stones[node][1]

		for _, neighbor := range rowMap[row] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
		for _, neighbor := range colMap[col] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	components := 0

	// Iterate over all stones and perform DFS if not visited
	for i := range stones {
		if !visited[i] {
			dfs(i)
			components++
		}
	}

	// Total stones - number of connected components gives the maximum stones that can be removed
	return len(stones) - components
}

func main() {
	stones1 := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}
	stones2 := [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}

	fmt.Println(removeStones(stones1)) // Output: 5
	fmt.Println(removeStones(stones2)) // Output: 3
}
