/*
40 Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.

Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]
*/
package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) // Sort candidates to handle duplicates
	var result [][]int
	var backtrack func(start, target int, combination []int)

	backtrack = func(start, target int, combination []int) {
		if target == 0 {
			// If the target is reached, add the combination to the result
			combinationCopy := make([]int, len(combination))
			copy(combinationCopy, combination)
			result = append(result, combinationCopy)
			return
		}

		for i := start; i < len(candidates); i++ {
			// Skip duplicates
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			if candidates[i] > target {
				// If the current candidate exceeds the target, stop further exploration
				break
			}

			// Include the candidate and move forward with the next element
			backtrack(i+1, target-candidates[i], append(combination, candidates[i]))
		}
	}

	backtrack(0, target, []int{}) // Initial call with start = 0
	return result
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
