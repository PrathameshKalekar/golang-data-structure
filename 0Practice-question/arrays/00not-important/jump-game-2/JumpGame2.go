/*
45 You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:

0 <= j <= nums[i] and
i + j < n
Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].

Example 1:

Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [2,3,0,1,4]
Output: 2
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(jumpGame2([]int{2, 3, 1, 1, 4})) // Output: 2
	fmt.Println(jumpGame2([]int{2, 3, 0, 1, 4})) // Output: 2
}

func jumpGame2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])

		// If we have reached the end of the current jump,
		// increment the number of jumps and update the end to the farthest position we can reach.
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}

		// If the farthest we can reach is beyond or at the last index, we can break early.
		if currentEnd >= len(nums)-1 {
			break
		}
	}

	return jumps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
