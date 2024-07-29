/*
16 Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.



Example 1:

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
Example 2:

Input: nums = [0,0,0], target = 1
Output: 0
Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(arr []int, target int) int {
	sort.Ints(arr)
	closest := arr[0] + arr[1] + arr[2]
	//loop through all elements
	for i := 0; i < len(arr)-2; i++ {
		//creating two pointer array andle left and right
		left, right := i+1, len(arr)-1

		for left < right {
			// handle the sum condition
			sum := arr[i] + arr[left] + arr[right]
			if math.Abs(float64(sum-target)) < math.Abs(float64(closest-target)) {
				closest = sum
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return sum
			}
		}
	}
	// return closest value
	return closest
}
