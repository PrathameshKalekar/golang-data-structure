/*
39 Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
Example 3:

Input: candidates = [2], target = 1
Output: []
*/
package main

import "fmt"

func main() {
	candidates1 := []int{2, 3, 6, 7}
	target1 := 7
	fmt.Println(combinationSum(candidates1, target1)) // Output: [[2, 2, 3], [7]]

	candidates2 := []int{2, 3, 5}
	target2 := 8
	fmt.Println(combinationSum(candidates2, target2)) // Output: [[2, 2, 2, 2], [2, 3, 3], [3, 5]]

	candidates3 := []int{2}
	target3 := 1
	fmt.Println(combinationSum(candidates3, target3)) // Output: []
}

func combinationSum(arr []int, target int) [][]int {
	var result [][]int
	var combination []int
	backtrack(arr, target, 0, combination, &result)
	return result
}

func backtrack(candidates []int, target int, start int, combination []int, result *[][]int) {
	if target == 0 {
		// Found a valid combination
		temp := make([]int, len(combination))
		copy(temp, combination)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] <= target {
			combination = append(combination, candidates[i])
			backtrack(candidates, target-candidates[i], i, combination, result)
			combination = combination[:len(combination)-1]
		}
	}
}
