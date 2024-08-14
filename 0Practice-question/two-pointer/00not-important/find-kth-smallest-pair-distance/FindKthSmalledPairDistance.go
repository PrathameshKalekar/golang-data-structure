/*
719 The distance of a pair of integers a and b is defined as the absolute difference between a and b.

Given an integer array nums and an integer k, return the kth smallest distance among all the pairs nums[i] and nums[j] where 0 <= i < j < nums.length.

Example 1:

Input: nums = [1,3,1], k = 1
Output: 0
Explanation: Here are all the pairs:
(1,3) -> 2
(1,1) -> 0
(3,1) -> 2
Then the 1st smallest distance pair is (1,1), and its distance is 0.
Example 2:

Input: nums = [1,1,1], k = 2
Output: 0
Example 3:

Input: nums = [1,6,1], k = 3
Output: 5
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 0, 2, 1, 0, 1, 1, 0, 2, 2, 0, 2, 0, 1, 1, 1, 0, 1, 0, 1, 1, 2, 2, 2, 2, 0, 0, 2, 1, 2, 1, 2, 0, 0, 0, 1, 0, 0, 1, 0, 2, 1, 1, 1, 1, 0, 2, 2, 1, 0, 2, 0, 2, 2, 2, 1, 0, 2, 2, 2, 2, 0, 0, 1, 0, 1, 1, 2, 1, 2, 2, 1, 1, 0, 2, 0, 1, 0, 1, 1, 2, 0, 1, 1, 1, 1, 2, 0, 2, 2, 0, 0, 1, 1, 1, 1, 2, 1, 2, 2, 1, 2, 0, 1, 2, 2, 1, 1, 2, 1, 0, 1, 1, 1, 0, 0, 1, 2, 0, 2, 1, 0, 1, 2, 0, 2, 2, 1, 0, 0, 2, 0, 1, 1, 0, 1, 0, 1, 0, 1, 2, 2, 2, 0, 1, 1, 1, 1, 1, 0, 2, 2, 2, 1, 0, 1, 0, 1, 0, 0, 0, 0, 2, 0, 0, 1, 1, 2, 0, 2, 1, 2, 0, 0, 1, 0, 1, 2, 1, 0, 1, 1, 1, 1, 0, 0, 2, 2, 1, 1, 0, 0, 0, 0, 1, 2, 2, 1, 1, 0, 1, 2, 0, 0, 2, 1, 2, 1, 2, 0, 2, 1, 1, 2, 2, 2, 2, 2, 0, 1, 1, 0, 1, 2, 2, 0, 2, 2, 2, 0, 2, 0, 1, 1, 2, 2, 0, 2, 2, 2, 2, 2, 2, 1, 0, 2, 2, 2, 0, 2, 0, 2, 0, 2, 1, 0, 1, 0, 1, 1, 0, 2, 0, 1, 0, 0, 2, 0, 0, 0, 2, 0, 2, 2, 0, 2, 0, 0, 1, 1, 0, 2, 0, 1, 2, 2, 0, 1, 1, 2, 0, 0, 0, 2, 1, 0, 1, 0, 2, 1, 2, 0, 1, 2, 1, 1, 1, 0, 1, 2, 1, 2, 2, 1, 2, 0, 2, 1, 0, 1, 1, 1, 2, 0, 2, 2, 2, 2, 1, 0, 2, 2, 1, 1, 1, 1, 0, 1, 2, 2, 0, 1, 2, 2, 1, 0, 0, 1, 2, 1, 0, 2, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 2, 0, 2, 0, 1, 2, 2, 0, 2, 1, 2, 2, 0, 0, 0, 0, 2, 1, 1, 2, 0, 2, 0, 1, 0, 1, 0, 2, 0, 0, 0, 2, 1, 2, 0, 1, 2, 1, 2, 1, 0, 1, 0, 1, 0, 0, 1, 2, 1, 1, 2, 1, 1, 2, 2, 1, 0, 1, 1, 2, 2, 1, 2, 2, 2, 1, 0, 1, 1, 1, 0, 0, 0, 2, 1, 1, 1, 2, 2, 1, 2, 0, 2, 1, 2, 0, 2, 2, 2, 1, 1, 2, 2, 0, 0, 2, 2, 2, 1, 2, 2, 1, 0, 2, 0, 2, 0, 2, 1, 2, 2, 1, 1, 1, 1, 1, 0, 0, 1, 0, 2, 2, 0, 0, 2, 1, 0, 1, 0, 2, 1, 0, 0, 0, 0, 1, 2, 1, 2, 0, 0, 1, 1, 2, 2, 2, 2, 0, 2, 1, 0, 0, 0, 2, 0, 1, 1, 0, 2, 1, 1, 1, 2, 1, 0, 0, 1, 0, 1, 0, 1, 2, 0, 2, 1, 0, 0, 1, 2, 1, 1, 0, 0, 0, 2, 2, 2, 1, 1, 2, 2, 0, 1, 0, 2, 2, 2, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 2, 1, 2, 0, 2, 0, 2, 1, 2, 2, 2, 0, 1, 0, 1, 2, 0, 2, 2, 1, 2, 0, 1, 2, 2, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 2, 1, 0, 0, 2, 0, 1, 0, 2, 2, 0, 1, 0, 0, 0, 1, 0, 0, 0, 2, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 2, 2, 0, 2, 2, 0, 0, 1, 2, 0, 1, 1, 1, 2, 0, 0, 0, 2, 0, 2, 2, 2, 2, 1, 0, 2, 2, 0, 0, 1, 1, 2, 2, 2, 1, 1, 2, 0, 1, 0, 0, 1, 0, 1, 2, 0, 1, 2, 0, 1, 1, 1, 1, 2, 0, 1, 0, 1, 2, 2, 0, 0, 2, 0, 1, 2, 1, 2, 1, 0, 0, 1, 0, 0, 1, 2, 0, 1, 0, 0, 0, 0, 2, 0, 1, 0, 1, 2, 0, 1, 2, 0, 0, 0, 0, 1, 1, 2, 0, 0, 0, 2, 1, 1, 0, 0, 2, 2, 1, 0, 0, 1, 0, 1, 0, 1, 1, 0, 2, 0, 1, 1, 2, 2, 1, 1, 0, 2, 0, 0, 1, 0, 1, 2, 2, 1, 2, 2, 2, 2, 2, 1, 2, 0, 0, 0, 1, 1, 0, 0, 2, 1, 0, 1, 0, 2, 2, 0, 1, 2, 2, 2, 0, 0, 0, 2, 2, 1, 2, 1, 0, 0, 0, 1, 1, 2, 1, 2, 0, 1, 2, 1, 0, 1, 2, 0, 1, 2, 1, 2, 1, 1, 1, 2, 2, 1, 0, 1, 0, 2, 1, 2, 0, 2, 0, 0, 0, 1, 2, 1, 0, 0, 0, 1, 2, 0, 2, 2, 2, 1, 0, 2, 2, 1, 1, 2, 0, 0, 1, 2, 2, 2, 1, 2, 1, 2, 1, 0, 2, 2, 0, 0, 0, 2, 0, 1, 0, 1, 0, 2, 2, 2, 2, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 2, 0, 1, 0, 0, 2, 0, 1, 1, 1, 1, 2, 0, 1, 1, 2, 0, 2, 1, 2, 1, 1, 0, 1, 1, 2, 1, 2, 0, 1, 0, 0, 0, 1, 0, 2, 0, 0, 2, 0, 0, 1, 0, 1, 2, 2, 1, 2, 0, 2, 2, 2, 0, 2, 0, 2, 1, 0, 0, 0, 0, 0, 0, 1, 0, 2, 0, 1, 1, 0, 2, 1, 2, 1, 1, 0, 2, 1, 2, 0, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 2, 2, 1, 1, 0, 1, 0, 0, 0, 2, 2, 2, 0, 1, 2, 1, 2, 0, 2, 1, 0, 0, 1, 2, 2, 1, 0, 1, 2, 0, 0, 2, 1, 1, 2, 0, 0, 1, 0, 2, 2, 2, 2, 0, 1, 0, 0, 0, 1, 1, 0, 2, 0, 2, 0, 2, 2, 2, 0, 2, 0, 0, 2, 1, 0, 2, 2, 2, 1, 2, 0, 2, 2, 0, 2, 1, 2, 2, 0, 1, 0, 2, 0, 1, 1, 2, 2, 2, 2, 0, 0, 0, 0, 2, 2, 2, 2, 1, 0, 2, 0, 2, 0, 1, 0, 1, 1, 2, 2, 1, 2, 1, 2, 0, 1, 2, 2, 2, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 2, 2, 0, 2, 2, 2, 0, 1, 0, 2, 1, 1, 1, 0, 0, 0, 0, 0, 2, 0, 1, 0, 2, 0, 0, 2, 2, 2, 2, 1, 2, 2, 1, 2, 0, 1, 0, 1, 1, 0, 0, 2, 2, 1, 2, 2, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 2, 0, 1, 0, 2, 2, 2, 0, 1, 0, 2, 1, 2, 2, 0, 2, 1, 1, 1, 1, 1, 1, 2, 2, 1, 0, 1, 2, 1, 1, 0, 0, 1, 0, 2, 0, 2, 2, 2, 2, 1, 2, 2, 2, 2, 1, 2, 2, 1, 1, 1, 0, 1, 2, 2, 0, 0, 2, 1, 0, 2, 1, 0, 2, 2, 0, 2, 1, 0, 2, 1, 0, 0, 0, 0, 2, 0, 1, 0, 2, 0, 1, 2, 2, 0, 1, 0, 1, 1, 1, 2, 1, 0, 0, 0, 0, 0, 2, 0, 1, 1, 1, 0, 1, 2, 2, 1, 2, 2, 0, 2, 2, 0, 0, 0, 2, 2, 2, 1, 1, 2, 1, 2, 1, 2, 1, 1, 0, 0, 2, 0, 2, 0, 1, 1, 0, 0, 2, 1, 0, 0, 1, 2, 0, 1, 0, 2, 1, 1, 1, 1, 2, 2, 0, 2, 2, 0, 1, 1, 2, 2, 0, 2, 0, 1, 0, 2, 0, 2, 2, 2, 1, 2, 2, 0, 1, 1, 2, 2, 1, 2, 2, 2, 2, 2, 2, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 2, 1, 2, 1, 2, 2, 1, 1, 0, 2, 1, 0, 0, 1, 2, 2, 1, 1, 1, 1, 0, 0, 0, 2, 0, 0, 2, 1, 1, 2, 2, 2, 1, 2, 1, 0, 1, 2, 1, 2, 2, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 2, 2, 1, 2, 1, 0, 2, 0, 1, 0, 0, 2, 0, 0, 0, 2, 1, 0, 2, 0, 2, 1, 1, 2, 0, 1, 1, 0, 2, 0, 1, 2, 2, 0, 0, 2, 0, 1, 1, 0, 2, 2, 1, 0, 0, 0, 2, 0, 1, 0, 0, 0, 2, 0, 2, 2, 1, 0, 1, 2, 2, 2, 2, 1, 1, 2, 2, 2, 0, 1, 2, 0, 2, 1, 2, 1, 0, 2, 0, 1, 0, 0, 2, 1, 0, 0, 0, 2, 0, 0, 1, 2, 2, 0, 0, 1, 0, 1, 1, 2, 2, 2, 0, 2, 1, 0, 2, 0, 1, 2, 1, 0, 1, 1, 1, 2, 1, 0, 0, 0, 2, 1, 0, 2, 2, 0, 1, 0, 2, 0, 2, 2, 2, 0, 1, 2, 2, 2, 1, 0, 0, 1, 0, 0, 0, 1, 1, 0, 1, 2, 2, 2, 0, 1, 0, 0, 1, 2, 2, 2, 0, 1, 1, 1, 2, 0, 2, 1, 1, 2, 2, 1, 0, 1, 2, 0, 2, 1, 1, 0, 1, 1, 2, 0, 1, 0, 2, 1, 2, 0, 2, 2, 2, 2, 0, 2, 0, 2, 1, 0, 1, 0, 2, 1, 0, 0, 0, 2, 0, 0, 0, 1, 0, 2, 1, 0, 0, 0, 0, 0, 2, 1, 0, 0, 0, 2, 1, 2, 0, 0, 0, 2, 2, 2, 2, 1, 1, 1, 0, 1, 2, 2, 0, 0, 1, 1, 0, 2, 0, 2, 1, 2, 0, 0, 2, 1, 0, 0, 0, 2, 2, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 2, 1, 2, 1, 2, 0, 1, 1, 2, 0, 0, 2, 1, 0, 0, 2, 0, 1, 2, 1, 1, 2, 2, 2, 1, 2, 2, 0, 1, 2, 0, 1, 2, 1, 2, 0, 2, 1, 0, 2, 1, 1, 2, 0, 2, 0, 1, 2, 2, 1, 2, 1, 1, 1, 1, 2, 1, 2, 1, 2, 0, 1, 1, 1, 2, 0, 1, 2, 1, 1, 0, 2, 1, 0, 2, 0, 0, 0, 2, 1, 0, 1, 2, 0, 1, 0, 1, 1, 0, 0, 1, 0, 2, 0, 0, 0, 0, 2, 2, 0, 1, 1, 2, 1, 2, 2, 1, 0, 1, 1, 2, 2, 0, 1, 0, 1, 0, 2, 1, 0, 2, 0, 2, 2, 1, 1, 1, 1, 2, 2, 2, 2, 0, 1, 1, 0, 2, 0, 1, 2, 1, 2, 0, 0, 2, 2, 1, 2, 2, 1, 1, 0, 0, 0, 0, 2, 1, 0, 2, 2, 0, 1, 2, 2, 2, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 1, 2, 2, 0, 2, 0, 1, 2, 2, 1, 0, 2, 2, 0, 0, 0, 2, 0, 0, 0, 1, 2, 1, 0, 2, 2, 1, 2, 0, 0, 2, 2, 2, 1, 1, 1, 2, 0, 0, 2, 0, 0, 2, 1, 2, 2, 2, 1, 0, 0, 1, 1, 1, 0, 0, 2, 0, 1, 1, 0, 0, 0, 1, 0, 2, 0, 0, 1, 1, 2, 0, 2, 0, 2, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 2, 1, 0, 1, 0, 1, 2, 0, 0, 1, 2, 0, 2, 1, 0, 1, 1, 2, 0, 2, 2, 1, 1, 2, 2, 1, 1, 0, 2, 2, 2, 2, 1, 2, 0, 2, 0, 2, 2, 1, 1, 0, 1, 0, 0, 1, 2, 2, 1, 1, 1, 0, 1, 0, 1, 2, 0, 2, 1, 2, 2, 0, 2, 2, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 2, 0, 2, 2, 0, 1, 2, 2, 1, 1, 2, 0, 1, 1, 1, 2, 0, 2, 1, 1, 2, 1, 1, 1, 1, 2, 1, 0, 1, 2, 0, 1, 1, 1, 0, 2, 1, 2, 0, 0, 2, 1, 0, 1, 0, 2, 2, 2, 1, 1, 1, 2, 2, 0, 2, 1, 0, 2, 2, 2, 2, 2, 2, 0, 0, 1, 2, 2, 2, 0, 0, 2, 0, 2, 2, 2, 2, 2, 2, 2, 0, 2, 1, 1, 2, 0, 2, 2, 0, 2, 1, 0, 0, 1, 1, 1, 1, 2, 1, 2, 0, 2, 0, 0, 0, 2, 2, 0, 0, 2, 1, 2, 1, 0, 1, 0, 2, 0, 1, 2, 0, 1, 2, 0, 1, 1, 1, 1, 0, 2, 2, 1, 0, 2, 1, 0, 2, 0, 1, 0, 0, 2, 1, 0, 1, 0, 2, 2, 2, 1, 2, 2, 0, 2, 1, 0, 2, 0, 2, 1, 1, 2, 0, 1, 2, 0, 0, 1, 1, 1, 0, 2, 0, 0, 2, 2, 0, 0, 2, 2, 1, 0, 1, 2, 1, 1, 2, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 2, 2, 2, 1, 1, 2, 0, 1, 2, 1, 2, 1, 1, 1, 1, 1, 0, 2, 1, 1, 1, 0, 2, 0, 0, 0, 2, 1, 0, 1, 2, 0, 1, 1, 2, 1, 1, 2, 2, 0, 2, 0, 2, 0, 2, 0, 2, 2, 0, 0, 1, 0, 2, 0, 0, 0, 0, 0, 1, 0, 2, 1, 1, 0, 1, 2, 2, 1, 0, 2, 2, 1, 2, 0, 1, 2, 2, 2, 0, 1, 0, 0, 1, 0, 2, 1, 0, 1, 1, 1, 1, 2, 1, 2, 0, 1, 2, 2, 2, 2, 2, 0, 1, 0, 2, 2, 0, 0, 0, 0, 1, 2, 1, 1, 1, 1, 0, 1, 0, 0, 0, 0, 2, 2, 2, 0, 0, 0, 1, 2, 2, 2, 2, 2, 2, 2, 1, 2, 1, 1, 1, 2, 0, 1, 1, 2, 0, 0, 1, 2, 0, 2, 0, 2, 2, 1, 2, 2, 2, 2, 2, 1, 1, 2, 1, 1, 1, 2, 0, 1, 2, 2, 1, 1, 0, 2, 1, 2, 1, 1, 1, 0, 2, 1, 0, 0, 2, 2, 2, 1, 0, 1, 2, 1, 2, 1, 1, 0, 0, 0, 2, 1, 0, 1, 0, 0, 1, 2, 0, 0, 2, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 2, 2, 0, 0, 0, 1, 2, 0, 1, 1, 1, 1, 2, 1, 1, 0, 1, 1, 1, 0, 1, 1, 2, 1, 2, 2, 1, 2, 2, 2, 1, 1, 0, 2, 1, 0, 2, 0, 1, 1, 0, 1, 2, 1, 1, 0, 2, 0, 0, 1, 0, 0, 1, 0, 0, 1, 2, 2, 0, 1, 2, 1, 2, 2, 2, 2, 0, 2, 1, 1, 0, 0, 0, 1, 1, 2, 2, 0, 0, 0, 1, 2, 0, 1, 2, 0, 1, 2, 1, 1, 2, 1, 2, 2, 0, 1, 2, 2, 1, 2, 0, 1, 1, 0, 0, 0, 0, 1, 2, 2, 2, 0, 1, 0, 1, 0, 2, 1, 1, 0, 2, 1, 2, 1, 2, 0, 2, 0, 0, 1, 2, 0, 1, 0, 0, 0, 1, 1, 1, 2, 1, 1, 1, 0, 1, 2, 0, 0, 0, 1, 0, 2, 0, 2, 0, 1, 1, 1, 0, 0, 0, 2, 0, 1, 2, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 2, 1, 1, 2, 2, 1, 1, 1, 0, 1, 0, 1, 1, 2, 1, 2, 2, 1, 1, 2, 2, 2, 1, 1, 1, 1, 1, 1, 2, 1, 1, 0, 2, 1, 0, 0, 1, 1, 1, 0, 2, 2, 0, 0, 1, 1, 2, 2, 1, 2, 2, 2, 0, 0, 2, 2, 2, 0, 1, 0, 2, 1, 0, 0, 2, 0, 0, 1, 2, 1, 2, 1, 2, 0, 0, 1, 0, 2, 0, 0, 2, 0, 1, 1, 2, 2, 0, 1, 0, 1, 1, 1, 2, 2, 1, 1, 1, 2, 0, 1, 2, 0, 1, 1, 1, 0, 1, 0, 2, 2, 0, 2, 1, 0, 0, 1, 2, 0, 0, 1, 0, 1, 0, 2, 2, 1, 0, 1, 2, 2, 0, 2, 2, 0, 0, 1, 1, 1, 2, 1, 2, 2, 2, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 2, 1, 0, 2, 1, 2, 0, 2, 0, 2, 2, 0, 2, 2, 0, 1, 0, 0, 2, 0, 2, 2, 1, 0, 2, 0, 2, 2, 1, 1, 2, 2, 2, 1, 0, 0, 1, 2, 1, 0, 1, 0, 2, 2, 1, 2, 1, 2, 0, 1, 0, 1, 2, 2, 1, 0, 2, 0, 2, 1, 0, 1, 2, 0, 0, 2, 2, 1, 1, 0, 0, 1, 2, 0, 2, 2, 1, 2, 1, 0, 1, 1, 0, 0, 2, 2, 2, 2, 2, 0, 1, 0, 1, 1, 1, 0, 1, 2, 1, 0, 2, 2, 0, 1, 1, 2, 0, 2, 2, 0, 2, 1, 1, 2, 2, 2, 1, 1, 0, 1, 0, 2, 2, 2, 2, 0, 2, 2, 1, 0, 0, 1, 1, 0, 1, 2, 1, 0, 0, 0, 2, 2, 0, 2, 0, 1, 2, 2, 1, 1, 1, 2, 0, 2, 0, 0, 1, 2, 0, 1, 1, 0, 1, 1, 1, 2, 0, 0, 0, 2, 1, 0, 2, 1, 2, 1, 0, 1, 1, 2, 2, 0, 1, 0, 2, 2, 2, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 0, 1, 1, 1, 1, 0, 2, 1, 0, 0, 2, 0, 2, 0, 1, 1, 0, 0, 1, 2, 0, 2, 2, 2, 1, 1, 0, 0, 0, 0, 1, 0, 1, 2, 2, 2, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 1, 2, 2, 0, 0, 2, 2, 0, 0, 0, 2, 2, 2, 2, 1, 1, 2, 2, 2, 0, 0, 2, 0, 1, 2, 2, 2, 2, 0, 0, 2, 0, 2, 2, 2, 0, 2, 1, 1, 1, 1, 0, 2, 0, 2, 0, 0, 1, 2, 2, 1, 0, 2, 0, 1, 2, 2, 1, 2, 0, 0, 2, 0, 0, 1, 1, 1, 0, 2, 0, 2, 2, 2, 2, 1, 2, 2, 0, 1, 2, 2, 2, 0, 0, 0, 1, 0, 1, 2, 0, 0, 0, 0, 1, 2, 0, 0, 2, 2, 0, 1, 1, 2, 2, 2, 2, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 0, 0, 2, 2, 0, 0, 1, 1, 0, 2, 2, 1, 1, 0, 2, 0, 2, 1, 0, 1, 2, 1, 1, 2, 1, 0, 0, 1, 2, 1, 1, 2, 2, 0, 1, 2, 0, 0, 0, 1, 1, 2, 0, 2, 1, 2, 0, 1, 0, 0, 0, 2, 0, 1, 1, 2, 2, 2, 0, 0, 0, 2, 2, 1, 0, 2, 1, 0, 0, 2, 1, 1, 2, 0, 2, 0, 2, 0, 1, 2, 2, 0, 0, 0, 2, 2, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 2, 1, 1, 0, 1, 0, 2, 0, 1, 0, 2, 0, 2, 1, 0, 1, 2, 0, 0, 2, 1, 1, 0, 0, 1, 0, 2, 1, 1, 1, 2, 1, 2, 2, 0, 2, 2, 1, 2, 2, 2, 1, 0, 1, 1, 1, 2, 1, 1, 1, 1, 0, 0, 1, 1, 0, 2, 2, 0, 1, 1, 1, 1, 2, 1, 1, 0, 0, 2, 2, 0, 2, 1, 1, 1, 1, 2, 0, 2, 2, 2, 0, 1, 0, 1, 0, 0, 1, 0, 2, 0, 0, 0, 2, 1, 1, 2, 2, 0, 1, 1, 0, 0, 1, 0, 2, 0, 2, 2, 0, 2, 2, 0, 0, 1, 2, 0, 1, 0, 0, 0, 0, 1, 2, 0, 1, 1, 0, 0, 2, 0, 0, 2, 0, 1, 1, 2, 2, 1, 0, 1, 0, 0, 2, 2, 2, 0, 0, 1, 0, 2, 0, 1, 2, 1, 0, 1, 1, 2, 2, 2, 2, 0, 0, 0, 2, 1, 1, 2, 0, 1, 1, 1, 0, 0, 0, 1, 0, 2, 0, 2, 0, 2, 0, 0, 1, 0, 0, 2, 0, 2, 2, 2, 2, 2, 1, 0, 2, 0, 0, 0, 0, 0, 0, 2, 0, 2, 0, 1, 0, 2, 2, 2, 1, 1, 1, 0, 1, 0, 1, 1, 0, 2, 0, 0, 1, 2, 2, 1, 2, 0, 0, 0, 2, 2, 0, 1, 1, 0, 2, 0, 1, 2, 1, 0, 0, 0, 1, 1, 1, 2, 0, 1, 1, 1, 2, 1, 0, 2, 1, 0, 1, 0, 1, 2, 0, 1, 2, 0, 2, 0, 0, 0, 2, 0, 2, 2, 2, 0, 1, 1, 2, 0, 0, 1, 1, 0, 1, 1, 0, 2, 0, 1, 1, 2, 2, 1, 0, 2, 1, 2, 0, 0, 0, 2, 2, 2, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 2, 0, 2, 2, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 2, 0, 0, 2, 0, 2, 2, 2, 2, 1, 1, 1, 2, 0, 2, 1, 1, 0, 0, 1, 2, 1, 2, 2, 2, 2, 0, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 2, 0, 2, 2, 2, 0, 1, 0, 1, 2, 2, 0, 0, 2, 0, 0, 1, 0, 0, 1, 2, 1, 0, 1, 2, 2, 1, 2, 2, 1, 1, 0, 2, 1, 0, 1, 1, 2, 1, 2, 1, 0, 2, 0, 1, 1, 0, 0, 1, 0, 2, 1, 1, 1, 2, 1, 0, 2, 2, 1, 2, 2, 1, 2, 2, 2, 2, 1, 1, 2, 1, 0, 0, 2, 2, 1, 1, 2, 1, 1, 1, 0, 0, 0, 0, 0, 0, 2, 2, 1, 0, 1, 0, 1, 1, 2, 2, 0, 1, 1, 2, 2, 1, 0, 2, 1, 1, 1, 2, 1, 2, 0, 0, 2, 0, 2, 0, 0, 0, 0, 2, 2, 2, 0, 0, 0, 2, 0, 2, 1, 0, 2, 1, 1, 0, 1, 0, 0, 2, 2, 1, 1, 2, 0, 0, 0, 1, 1, 1, 1, 1, 0, 2, 2, 2, 2, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 2, 2, 1, 0, 1, 1, 0, 2, 0, 2, 2, 1, 1, 1, 1, 0, 0, 2, 2, 1, 2, 0, 0, 2, 1, 0, 2, 1, 2, 1, 2, 1, 0, 2, 2, 2, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 2, 2, 2, 1, 0, 2, 0, 2, 2, 2, 1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 2, 1, 2, 0, 0, 0, 1, 2, 0, 1, 2, 0, 1, 0, 0, 0, 1, 0, 1, 2, 2, 1, 2, 0, 1, 0, 2, 1, 0, 1, 2, 1, 1, 1, 2, 1, 0, 0, 2, 1, 1, 1, 0, 2, 1, 2, 2, 2, 2, 2, 1, 1, 2, 1, 2, 1, 0, 1, 1, 0, 1, 2, 0, 1, 2, 2, 2, 2, 2, 1, 0, 2, 2, 2, 2, 0, 1, 2, 1, 2, 0, 0, 2, 0, 0, 0, 0, 0, 0, 2, 0, 0, 2, 0, 2, 1, 1, 2, 0, 1, 1, 2, 0, 0, 2, 2, 0, 0, 2, 0, 2, 2, 1, 0, 1, 0, 2, 0, 0, 0, 0, 2, 0, 1, 2, 1, 0, 1, 2, 0, 2, 0, 2, 1, 0, 0, 0, 1, 2, 0, 1, 1, 2, 1, 1, 1, 0, 0, 0, 0, 0, 2, 2, 0, 2, 2, 2, 1, 2, 1, 1, 1, 2, 0, 2, 1, 0, 0, 0, 2, 2, 1, 0, 2, 0, 2, 0, 1, 0, 1, 0, 2, 0, 1, 1, 1, 1, 0, 2, 0, 1, 2, 1, 2, 0, 2, 1, 2, 1, 2, 2, 2, 2, 1, 2, 0, 0, 2, 0, 2, 2, 2, 0, 1, 2, 1, 0, 2, 0, 2, 2, 2, 0, 1, 1, 0, 2, 2, 2, 2, 0, 1, 2, 1, 2, 2, 1, 2, 0, 1, 2, 0, 0, 0, 2, 2, 1, 1, 0, 0, 2, 1, 1, 1, 0, 1, 2, 0, 1, 2, 0, 0, 1, 2, 0, 2, 0, 0, 2, 0, 2, 0, 1, 1, 1, 1, 1, 0, 2, 1, 1, 2, 2, 0, 0, 1, 2, 1, 2, 2, 1, 0, 2, 2, 2, 2, 2, 0, 1, 0, 0, 1, 1, 2, 0, 1, 2, 1, 2, 0, 0, 0, 2, 2, 2, 1, 2, 1, 1, 2, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 2, 0, 1, 1, 2, 2, 0, 0, 1, 2, 0, 2, 2, 1, 2, 0, 0, 1, 0, 1, 2, 2, 0, 1, 2, 1, 2, 1, 2, 1, 2, 0, 1, 1, 1, 1, 1, 2, 0, 2, 2, 1, 2, 2, 1, 1, 0, 0, 1, 0, 2, 0, 1, 1, 1, 1, 2, 2, 0, 0, 1, 0, 1, 0, 2, 2, 2, 1, 0, 2, 2, 2, 0, 0, 2, 2, 1, 2, 2, 0, 2, 0, 2, 1, 0, 2, 0, 0, 1, 1, 1, 0, 2, 0, 0, 0, 1, 0, 2, 1, 2, 1, 2, 1, 0, 1, 2, 0, 2, 0, 1, 2, 2, 1, 0, 1, 2, 1, 2, 1, 0, 2, 2, 2, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 0, 0, 2, 0, 1, 2, 2, 0, 1, 2, 2, 2, 2, 0, 2, 1, 1, 2, 2, 2, 2, 1, 0, 1, 0, 2, 2, 0, 1, 0, 2, 2, 0, 0, 0, 0, 2, 0, 2, 1, 0, 1, 1, 2, 2, 1, 1, 2, 1, 2, 2, 0, 1, 1, 2, 0, 0, 0, 0, 2, 0, 0, 1, 2, 1, 2, 2, 2, 1, 0, 1, 1, 1, 1, 0, 0, 2, 0, 2, 1, 0, 2, 2, 1, 0, 0, 2, 1, 2, 2, 2, 1, 0, 2, 1, 0, 1, 2, 2, 0, 1, 2, 2, 0, 0, 2, 1, 2, 0, 1, 0, 1, 2, 1, 1, 0, 1, 1, 0, 0, 1, 2, 0, 1, 2, 2, 0, 1, 0, 0, 2, 2, 2, 2, 2, 0, 2, 0, 1, 1, 2, 2, 0, 0, 2, 0, 0, 2, 0, 1, 1, 0, 1, 0, 1, 2, 1, 1, 2, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 1, 0, 1, 2, 2, 0, 1, 2, 2, 1, 1, 1, 2, 1, 1, 1, 2, 2, 2, 0, 0, 0, 0, 1, 1, 0, 0, 2, 2, 1, 2, 0, 2, 2, 0, 0, 1, 1, 0, 2, 1, 1, 1, 0, 2, 0, 0, 0, 0, 0, 1, 2, 1, 0, 0, 0, 0, 2, 2, 1, 0, 2, 2, 0, 1, 0, 0, 1, 2, 0, 0, 0, 1, 0, 2, 0, 2, 0, 0, 2, 2, 1, 1, 0, 0, 2, 0, 0, 2, 1, 0, 2, 1, 0, 1, 2, 2, 1, 0, 2, 0, 0, 0, 0, 2, 2, 2, 2, 0, 1, 1, 0, 1, 2, 0, 0, 0, 2, 0, 1, 1, 1, 1, 0, 0, 1, 2, 1, 1, 2, 2, 1, 0, 1, 1, 1, 1, 0, 2, 0, 1, 2, 2, 0, 2, 0, 2, 1, 1, 0, 1, 0, 2, 2, 2, 2, 1, 1, 2, 1, 1, 2, 0, 0, 2, 0, 0, 1, 0, 0, 2, 2, 0, 2, 1, 2, 2, 2, 2, 1, 0, 1, 2, 2, 1, 1, 1, 1, 1, 0, 1, 2, 0, 0, 2, 0, 0, 2, 2, 1, 1, 1, 0, 0, 0, 2, 0, 2, 0, 1, 2, 2, 0, 2, 2, 2, 0, 1, 0, 1, 2, 2, 2, 1, 0, 2, 1, 1, 2, 0, 1, 2, 2, 1, 0, 2, 2, 1, 2, 0, 1, 2, 1, 0, 2, 0, 0, 2, 2, 1, 1, 0, 2, 0, 0, 2, 2, 0, 1, 0, 1, 2, 0, 1, 0, 1, 2, 0, 0, 1, 1, 1, 0, 1, 1, 0, 2, 2, 2, 0, 1, 0, 1, 1, 0, 2, 1, 0, 1, 0, 0, 2, 0, 0, 1, 2, 0, 2, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 2, 0, 2, 0, 0, 1, 1, 0, 0, 2, 0, 0, 2, 1, 2, 1, 2, 2, 1, 2, 1, 0, 1, 2, 0, 2, 0, 1, 2, 2, 0, 2, 2, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 2, 1, 0, 1, 0, 1, 0, 2, 2, 2, 2, 2, 1, 2, 0, 1, 2, 2, 1, 0, 1, 2, 0, 1, 1, 0, 1, 2, 0, 1, 2, 0, 0, 2, 0, 0, 2, 2, 0, 0, 1, 2, 2, 0, 0, 1, 2, 2, 1, 1, 1, 2, 0, 2, 0, 2, 2, 1, 2, 1, 1, 2, 1, 1, 2, 2, 2, 2, 2, 0, 0, 1, 2, 1, 2, 0, 0, 2, 0, 0, 2, 1, 1, 1, 2, 0, 1, 2, 0, 1, 1, 0, 1, 1, 2, 0, 2, 2, 0, 1, 2, 2, 2, 0, 1, 2, 0, 2, 2, 0, 0, 1, 0, 0, 0, 0, 2, 0, 1, 0, 0, 2, 0, 0, 0, 2, 1, 0, 2, 0, 1, 1, 1, 0, 2, 0, 0, 1, 0, 2, 2, 1, 0, 0, 1, 1, 0, 1, 2, 1, 1, 2, 2, 2, 0, 1, 0, 2, 2, 2, 1, 1, 2, 2, 0, 1, 2, 0, 0, 1, 1, 2, 0, 1, 2, 1, 2, 1, 2, 2, 1, 2, 2, 0, 1, 1, 2, 0, 1, 0, 0, 2, 0, 0, 1, 0, 0, 0, 2, 1, 2, 2, 2, 0, 0, 1, 2, 0, 0, 2, 1, 1, 2, 0, 0, 2, 1, 1, 0, 0, 1, 2, 2, 0, 1, 2, 0, 1, 0, 2, 2, 1, 1, 1, 0, 0, 2, 2, 2, 2, 0, 1, 0, 2, 2, 1, 2, 1, 1, 1, 0, 2, 2, 1, 1, 2, 2, 0, 2, 1, 2, 2, 2, 1, 1, 2, 1, 0, 0, 1, 1, 0, 2, 1, 0, 2, 2, 2, 2, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 2, 1, 0, 2, 1, 0, 0, 2, 0, 2, 2, 1, 2, 0, 2, 1, 2, 2, 1, 2, 0, 0, 2, 2, 0, 0, 1, 0, 1, 2, 1, 1, 0, 2, 1, 0, 1, 2, 2, 2, 1, 1, 2, 0, 1, 1, 2, 1, 2, 0, 2, 1, 0, 2, 1, 2, 0, 0, 0, 0, 0, 2, 0, 2, 1, 0, 0, 0, 1, 2, 0, 0, 2, 1, 2, 0, 2, 1, 2, 0, 1, 1, 2, 0, 0, 1, 1, 0, 0, 2, 0, 1, 2, 1, 1, 0, 1, 0, 2, 2, 2, 2, 1, 2, 0, 2, 1, 2, 2, 1, 1, 1, 0, 1, 2, 2, 1, 2, 0, 0, 0, 2, 1, 1, 1, 0, 2, 0, 1, 0, 2, 0, 0, 2, 0, 2, 1, 1, 1, 2, 0, 1, 0, 0, 2, 2, 1, 0, 0, 0, 2, 2, 0, 2, 1, 1, 0, 1, 1, 2, 1, 2, 2, 0, 0, 1, 2, 2, 2, 2, 2, 1, 2, 0, 1, 2, 1, 0, 2, 2, 2, 2, 2, 1, 2, 1, 2, 0, 0, 1, 1, 2, 2, 1, 0, 1, 2, 0, 2, 0, 1, 2, 2, 1, 2, 1, 0, 0, 1, 2, 0, 1, 1, 0, 2, 2, 2, 0, 1, 2, 2, 1, 0, 1, 2, 1, 0, 2, 2, 2, 2, 0, 1, 0, 2, 0, 0, 0, 1, 1, 2, 2, 2, 0, 2, 2, 2, 0, 2, 1, 0, 2, 1, 2, 1, 2, 2, 1, 0, 1, 1, 2, 2, 0, 1, 2, 1, 2, 1, 2, 2, 0, 2, 0, 1, 1, 0, 2, 2, 2, 2, 1, 0, 1, 0, 0, 1, 1, 2, 1, 2, 1, 1, 2, 1, 1, 0, 2, 2, 1, 1, 2, 2, 1, 1, 2, 2, 1, 2, 2, 1, 2, 1, 2, 0, 1, 2, 1, 0, 2, 0, 2, 0, 0, 0, 2, 0, 0, 0, 1, 2, 2, 2, 2, 1, 0, 2, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1, 2, 2, 0, 2, 1, 2, 1, 2, 2, 1, 0, 1, 0, 2, 2, 0, 2, 2, 1, 1, 0, 0, 1, 1, 0, 2, 2, 2, 0, 1, 0, 1, 0, 0, 1, 0, 2, 1, 2, 2, 0, 0, 2, 0, 0, 2, 0, 2, 1, 1, 1, 1, 0, 0, 1, 0, 1, 0, 1, 0, 2, 2, 1, 1, 2, 1, 2, 0, 2, 1, 0, 2, 2, 0, 2, 2, 0, 2, 2, 2, 0, 1, 2, 0, 2, 2, 2, 0, 1, 0, 0, 2, 2, 1, 1, 2, 1, 2, 2, 0, 2, 2, 2, 2, 0, 0, 0, 1, 0, 2, 0, 0, 2, 1, 2, 0, 2, 2, 2, 2, 2, 2, 1, 2, 0, 2, 2, 2, 2, 1, 1, 0, 0, 1, 2, 0, 2, 0, 2, 0, 2, 2, 0, 1, 2, 1, 2, 0, 2, 2, 1, 2, 1, 0, 0, 1, 2, 1, 0, 2, 2, 1, 1, 1, 2, 1, 1, 2, 1, 0, 2, 2, 0, 0, 0, 2, 1, 2, 2, 0, 1, 0, 0, 2, 1, 1, 1, 0, 1, 1, 2, 2, 1, 1, 0, 0, 1, 1, 0, 2, 0, 2, 1, 0, 0, 1, 2, 2, 1, 0, 2, 2, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 2, 1, 0, 2, 2, 0, 2, 1, 1, 0, 2, 1, 1, 1, 2, 2, 2, 2, 2, 2, 1, 1, 1, 2, 0, 2, 2, 2, 1, 1, 1, 0, 2, 1, 2, 1, 0, 1, 1, 0, 0, 0, 2, 1, 2, 2, 1, 1, 2, 0, 1, 0, 2, 0, 0, 2, 2, 1, 0, 2, 1, 2, 0, 2, 2, 2, 2, 0, 0, 0, 1, 0, 2, 1, 1, 0, 2, 1, 2, 2, 1, 2, 2, 1, 0, 1, 2, 2, 0, 1, 0, 1, 2, 2, 1, 1, 1, 0, 0, 0, 1, 2, 2, 1, 2, 2, 0, 2, 2, 1, 0, 0, 2, 0, 1, 1, 2, 1, 0, 0, 0, 1, 2, 0, 0, 0, 0, 1, 1, 1, 2, 1, 2, 2, 1, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 0, 2, 0, 2, 1, 2, 0, 1, 1, 1, 0, 0, 2, 2, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 2, 0, 2, 0, 1, 2, 2, 1, 2, 1, 1, 1, 0, 1, 0, 0, 2, 0, 1, 1, 2, 2, 2, 1, 1, 0, 1, 2, 2, 2, 0, 2, 2, 0, 1, 1, 1, 1, 1, 2, 0, 0, 2, 2, 2, 0, 0, 0, 1, 1, 2, 0, 2, 1, 0, 1, 1, 2, 1, 1, 2, 2, 2, 2, 2, 2, 0, 0, 2, 1, 2, 2, 2, 1, 2, 0, 0, 2, 1, 1, 0, 1, 0, 2, 0, 1, 2, 2, 0, 0, 0, 2, 0, 1, 0, 1, 2, 1, 0, 1, 2, 2, 1, 0, 0, 1, 1, 0, 2, 2, 0, 1, 2, 1, 1, 0, 1, 1, 2, 1, 2, 2, 1, 0, 2, 0, 2, 1, 0, 0, 2, 2, 1, 1, 0, 0, 1, 1, 0, 0, 2, 0, 0, 2, 0, 2, 0, 2, 1, 1, 2, 1, 0, 0, 0, 0, 2, 0, 0, 1, 1, 2, 0, 0, 2, 1, 1, 2, 0, 2, 1, 1, 1, 2, 1, 2, 1, 2, 0, 1, 2, 2, 2, 1, 2, 2, 1, 1, 2, 0, 2, 1, 0, 1, 2, 1, 1, 0, 2, 1, 2, 0, 1, 1, 0, 0, 0, 2, 1, 0, 1, 0, 0, 0, 1, 2, 0, 0, 1, 0, 0, 2, 2, 0, 2, 1, 0, 0, 1, 1, 2, 2, 2, 0, 1, 1, 0, 1, 0, 2, 0, 0, 0, 2, 0, 1, 0, 0, 0, 1, 0, 2, 2, 0, 1, 1, 1, 1, 0, 0, 2, 1, 2, 0, 0, 0, 1, 1, 2, 2, 0, 0, 0, 1, 2, 1, 0, 1, 1, 2, 1, 1, 0, 0, 2, 2, 2, 0, 1, 0, 0, 2, 1, 0, 1, 0, 0, 0, 2, 1, 2, 1, 0, 2, 0, 0, 1, 2, 2, 1, 2, 1, 2, 0, 0, 1, 1, 0, 1, 1, 1, 1, 2, 2, 1, 2, 0, 2, 1, 0, 0, 0, 2, 2, 2, 0, 1, 2, 1, 1, 0, 1, 2, 1, 0, 0, 0, 2, 0, 0, 1, 2, 2, 0, 0, 1, 0, 0, 1, 0, 2, 1, 0, 2, 2, 2, 0, 2, 1, 1, 2, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 2, 0, 2, 2, 2, 2, 2, 0, 2, 1, 1, 0, 1, 0, 0, 1, 0, 0, 0, 2, 1, 2, 2, 2, 1, 0, 1, 0, 2, 2, 0, 0, 0, 1, 2, 0, 2, 2, 1, 1, 1, 0, 2, 2, 0, 2, 2, 0, 1, 2, 2, 0, 1, 1, 1, 1, 2, 2, 0, 2, 2, 2, 2, 1, 1, 0, 2, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 2, 1, 2, 0, 1, 0, 2, 1, 2, 2, 2, 2, 0, 2, 1, 0, 2, 0, 2, 0, 2, 2, 0, 0, 2, 0, 2, 1, 2, 2, 0, 0, 1, 2, 0, 1, 1, 2, 0, 0, 2, 0, 1, 0, 0, 2, 0, 1, 2, 2, 1, 0, 2, 2, 1, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 0, 2, 0, 2, 2, 2, 2, 1, 0, 2, 2, 0, 2, 0, 0, 0, 2, 0, 0, 2, 0, 2, 1, 1, 2, 1, 1, 0, 1, 0, 1, 1, 2, 0, 2, 1, 0, 0, 0, 2, 0, 2, 0, 0, 2, 1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 2, 0, 1, 2, 1, 0, 1, 0, 2, 2, 2, 2, 0, 2, 0, 0, 0, 1, 1, 0, 1, 1, 0, 2, 0, 2, 2, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 2, 1, 0, 1, 2, 1, 2, 2, 0, 1, 0, 1, 2, 0, 1, 1, 2, 0, 0, 2, 0, 2, 0, 0, 1, 2, 2, 2, 0, 0, 2, 0, 1, 1, 2, 2, 0, 0, 2, 1, 1, 0, 1, 1, 0, 2, 0, 0, 2, 2, 0, 0, 2, 0, 1, 2, 1, 2, 0, 2, 1, 1, 1, 0, 2, 0, 0, 2, 2, 2, 2, 2, 1, 2, 0, 2, 1, 1, 2, 2, 2, 0, 0, 1, 0, 0, 2, 2, 2, 1, 1, 1, 2, 1, 2, 2, 2, 0, 1, 0, 0, 1, 1, 0, 1, 1, 2, 0, 2, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 2, 0, 0, 0, 2, 2, 0, 0, 0, 0, 2, 1, 2, 1, 1, 1, 2, 1, 1, 0, 0, 0, 0, 2, 1, 2, 1, 0, 1, 2, 2, 0, 0, 0, 2, 2, 0, 0, 1, 0, 0, 1, 1, 2, 1, 2, 1, 0, 1, 1, 1, 1, 2, 1, 1, 2, 0, 1, 2, 1, 0, 2, 2, 2, 0, 0, 2, 2, 0, 0, 2, 1, 0, 0, 2, 2, 1, 2, 0, 0, 0, 0, 2, 1, 0, 1, 1, 0, 2, 1, 2, 1, 2, 1, 2, 1, 2, 2, 1, 2, 0, 2, 2, 2, 1, 1, 1, 2, 1, 2, 2, 2, 1, 1, 2, 1, 1, 2, 1, 2, 1, 2, 2, 0, 2, 2, 1, 2, 1, 2, 2, 2, 0, 0, 0, 1, 1, 0, 2, 1, 0, 1, 1, 2, 1, 0, 2, 1, 0, 1, 0, 0, 2, 2, 1, 2, 1, 0, 1, 1, 0, 1, 1, 2, 2, 0, 0, 2, 0, 0, 0, 2, 0, 2, 1, 1, 2, 2, 1, 0, 0, 2, 1, 0, 2, 2, 1, 1, 2, 1, 2, 2, 2, 2, 0, 1, 1, 1, 2, 0, 2, 0, 0, 1, 2, 0, 2, 2, 0, 2, 2, 0, 0, 2, 1, 2, 1, 1, 0, 2, 2, 2, 1, 2, 2, 0, 0, 0, 0, 2, 0, 1, 1, 2, 1, 1, 1, 2, 2, 1, 0, 1, 1, 1, 1, 0, 1, 1, 2, 1, 0, 2, 0, 2, 0, 1, 2, 2, 0, 1, 1, 2, 1, 2, 0, 2, 1, 0, 1, 0, 1, 2, 1, 2, 0, 2, 1, 2, 1, 0, 1, 1, 1, 0, 1, 0, 0, 2, 0, 0, 2, 2, 0, 0, 2, 0, 2, 2, 0, 0, 2, 2, 0, 2, 1, 2, 0, 1, 0, 2, 1, 1, 2, 1, 2, 0, 2, 1, 2, 1, 1, 2, 2, 1, 0, 1, 0, 1, 1, 1, 1, 0, 0, 0, 2, 0, 0, 0, 0, 0, 2, 1, 2, 0, 1, 1, 2, 0, 2, 2, 0, 2, 0, 1, 1, 1, 0, 0, 2, 1, 1, 2, 0, 1, 2, 0, 0, 2, 0, 0, 2, 0, 1, 2, 2, 2, 2, 2, 0, 1, 0, 2, 2, 1, 1, 2, 2, 2, 1, 0, 2, 2, 2, 2, 1, 0, 0, 2, 2, 1, 1, 1, 1, 1, 1, 2, 2, 0, 2, 2, 1, 1, 2, 0, 1, 0, 1, 2, 2, 0, 2, 2, 2, 0, 2, 1, 1, 1, 2, 0, 0, 0, 1, 2, 0, 1, 0, 1, 0, 1, 1, 2, 2, 1, 0, 0, 2, 0, 1, 2, 0, 0, 1, 2, 1, 1, 2, 2, 1, 0, 0, 0, 0, 2, 2, 0, 2, 0, 0, 1, 0, 2, 0, 2, 1, 1, 1, 0, 2, 2, 0, 0, 0, 1, 0, 1, 0, 1, 0, 2, 1, 2, 0, 1, 2, 2, 0, 0, 0, 0, 0, 0, 2, 0, 2, 1, 0, 2, 1, 1, 1, 2, 0, 1, 1, 2, 1, 0, 1, 2, 0, 2, 1, 2, 0, 0, 1, 1, 0, 1, 1, 2, 0, 2, 2, 0, 2, 1, 0, 0, 2, 2, 0, 2, 1, 2, 1, 1, 1, 1, 2, 2, 0, 1, 2, 2, 0, 2, 1, 1, 1, 2, 0, 0, 1, 0, 0, 0, 2, 1, 1, 0, 0, 0, 1, 0, 0, 2, 2, 0, 1, 1, 0, 2, 1, 2, 2, 0, 0, 2, 2, 0, 0, 2, 2, 0, 0, 1, 1, 1, 0, 0, 0, 1, 2, 0, 0, 0, 1, 0, 1, 2, 2, 2, 1, 1, 1, 1, 0, 2, 2, 1, 1, 1, 0, 2, 0, 2, 0, 2, 1, 2, 2, 1, 2, 1, 0, 1, 1, 2, 1, 2, 2, 1, 0, 1, 1, 0, 1, 2, 1, 0, 0, 0, 0, 0, 2, 1, 2, 0, 2, 0, 1, 0, 2, 1, 0, 0, 2, 0, 1, 2, 0, 2, 2, 1, 1, 1, 0, 0, 2, 0, 1, 1, 1, 1, 2, 0, 1, 1, 0, 0, 2, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 2, 1, 2, 1, 1, 2, 2, 2, 1, 1, 0, 2, 2, 0, 2, 2, 2, 0, 2, 0, 1, 1, 0, 1, 2, 1, 2, 1, 1, 2, 0, 0, 1, 2, 1, 1, 0, 1, 2, 1, 2, 2, 1, 2, 1, 1, 0, 2, 2, 1, 0, 2, 0, 1, 2, 1, 1, 2, 0, 1, 2, 2, 1, 2, 0, 1, 2, 1, 1, 2, 0, 2, 1, 1, 0, 2, 2, 2, 0, 2, 0, 1, 1, 0, 0, 1, 0, 1, 2, 2, 1, 1, 0, 1, 1, 1, 1, 1, 2, 2, 0, 2, 1, 1, 1, 1, 0, 1, 0, 1, 1, 0, 0, 1, 0, 1, 2, 2, 1, 2, 0, 1, 2, 1, 0, 0, 0, 1, 0, 0, 1, 2, 1, 1, 0, 2, 0, 0, 0, 0, 0, 1, 0, 1, 2, 1, 2, 2, 2, 0, 2, 1, 1, 2, 2, 1, 0, 1, 0, 0, 1, 1, 2, 1, 2, 0, 1, 2, 2, 2, 0, 2, 2, 0, 0, 2, 1, 2, 1, 2, 1, 0, 2, 2, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 2, 2, 2, 0, 1, 2, 0, 1, 1, 2, 1, 0, 1, 2, 0, 2, 2, 2, 1, 2, 2, 1, 1, 2, 0, 2, 1, 1, 0, 2, 1, 1, 1, 2, 1, 2, 1, 1, 2, 1, 2, 2, 1, 0, 0, 2, 0, 0, 1, 0, 2, 0, 2, 0, 2, 1, 0, 1, 2, 1, 0, 1, 2, 1, 1, 0, 2, 2, 1, 2, 1, 1, 0, 2, 2, 2, 2, 0, 2, 1, 0, 2, 1, 2, 2, 0, 1, 1, 1, 0, 1, 2, 1, 2, 0, 1, 2, 2, 0, 2, 2, 2, 0, 1, 1, 0, 2, 1, 0, 0, 2, 1, 0, 1, 0, 0, 1, 1, 1, 1, 2, 0, 1, 0, 1, 1, 2, 2, 2, 0, 0, 2, 1, 1, 1, 2, 2, 0, 0, 0, 0, 2, 0, 1, 1, 1, 2, 0, 0, 1, 1, 1, 2, 0, 0, 1, 0, 2, 0, 0, 2, 2, 1, 2, 1, 1, 2, 2, 2, 2, 1, 2, 2, 1, 2, 1, 1, 0, 1, 1, 1, 1, 2, 1, 1, 2, 2, 2, 1, 1, 1, 0, 0, 0, 0, 0, 1, 1, 2, 1, 1, 0, 1, 1, 2, 2, 0, 1, 2, 2, 1, 2, 1, 0, 1, 2, 1, 0, 2, 1, 1, 1, 2, 2, 1, 1, 0, 0, 0, 1, 1, 2, 1, 0, 2, 1, 0, 1, 2, 1, 1, 1, 1, 2, 2, 0, 2, 1, 1, 1, 1, 2, 0, 0, 1, 1, 2, 1, 2, 0, 2, 1, 0, 0, 0, 0, 1, 2, 0, 1, 2, 0, 0, 1, 1, 1, 2, 2, 2, 2, 2, 1, 2, 0, 0, 1, 0, 0, 2, 0, 1, 0, 2, 0, 1, 2, 1, 1, 1, 1, 2, 0, 2, 0, 2, 0, 0, 1, 2, 2, 1, 1, 1, 1, 1, 2, 1, 2, 0, 0, 2, 1, 2, 1, 1, 1, 2, 2, 2, 1, 2, 0, 2, 1, 2, 1, 2, 1, 2, 1, 2, 2, 2, 2, 1, 1, 2, 0, 0, 0, 2, 0, 2, 2, 0, 1, 0, 1, 2, 1, 1, 0, 0, 1, 1, 2, 2, 0, 1, 1, 2, 2, 1, 2, 0, 1, 2, 1, 0, 2, 2, 1, 0, 1, 0, 1, 1, 0, 2, 2, 2, 1, 2, 0, 2, 2, 2, 2, 2, 1, 0, 1, 1, 2, 1, 2, 1, 2, 0, 0, 0, 0, 2, 0, 2, 2, 1, 1, 1, 1, 2, 2, 0, 0, 0, 2, 0, 2, 1, 0, 0, 1, 0, 1, 2, 1, 0, 1, 0, 2, 0, 2, 0, 0, 2, 2, 0, 0, 1, 2, 1, 0, 1, 0, 2, 1, 0, 2, 2, 2, 2, 2, 2, 0, 0, 1, 2, 2, 2, 2, 0, 1, 0, 2, 1, 2, 2, 2, 0, 2, 2, 2, 1, 1, 2, 2, 1, 0, 0, 1, 1, 2, 1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 1, 0, 2, 0, 2, 2, 1, 2, 2, 2, 0, 0, 2, 1, 1, 2, 1, 2, 2, 0, 0, 2, 2, 2, 2, 0, 0, 0, 0, 0, 0, 2, 0, 2, 2, 1, 0, 0, 1, 0, 2, 2, 1, 0, 2, 2, 0, 1, 0, 2, 1, 0, 0, 2, 0, 0, 2, 1, 2, 2, 2, 0, 0, 1, 0, 0, 2, 0, 2, 1, 2, 2, 2, 1, 1, 1, 2, 1, 2, 2, 2, 2, 0, 1, 0, 2, 1, 2, 0, 2, 1, 0, 2, 0, 1, 1, 0, 1, 1, 2, 2, 0, 0, 2, 1, 0, 2, 0, 2, 1, 2, 2, 1, 2, 1, 2, 1, 2, 2, 2, 0, 2, 1, 2, 2, 1, 0, 2, 0, 0, 1, 0, 2, 0, 1, 1, 1, 2, 2, 2, 2, 2, 0, 1, 1, 1, 1, 2, 1, 0, 0, 2, 2, 1, 2, 2, 2, 0, 1, 2, 1, 1, 1, 2, 2, 0, 1, 1, 2, 2, 0, 0, 0, 0, 0, 2, 1, 0, 1, 0, 0, 0, 0, 1, 2, 1, 1, 0, 2, 1, 2, 1, 0, 0, 0, 0, 1, 2, 1, 0, 0, 0, 0, 1, 2, 0, 0, 1, 2, 2, 1, 2, 2, 2, 2, 0, 0, 0, 2, 0, 0, 2, 0, 2, 1, 1, 1, 0, 0, 0, 2, 1, 0, 2, 1, 0, 2, 0, 2, 2, 0, 0, 0, 0, 2, 1, 1, 1, 0, 1, 2, 2, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 2, 1, 2, 1, 1, 2, 2, 2, 2, 2, 2, 1, 0, 2, 2, 2, 1, 2, 1, 2, 1, 0, 2, 1, 0, 1, 1, 1, 1, 0, 0, 2, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 0, 0, 2, 0, 1, 1, 0, 1, 1, 1, 2, 1, 0, 0, 1, 2, 0, 1, 1, 0, 2, 1, 1, 2, 1, 0, 1, 2, 0, 1, 1, 0, 1, 1, 0, 2, 2, 1, 0, 1, 1, 0, 0, 0, 2, 2, 2, 1, 0, 2, 0, 1, 0, 2, 2, 1, 2, 0, 0, 1, 2, 1, 0, 1, 0, 1, 2, 2, 2, 2, 1, 2, 1, 0, 2, 1, 2, 2, 2, 2, 2, 0, 2, 2, 0, 1, 1, 2, 2, 2, 0, 2, 2, 2, 0, 0, 2, 1, 1, 1, 0, 0, 1, 0, 0, 2, 1, 1, 1, 0, 1, 0, 0, 0, 2, 0, 0, 1, 1, 1, 2, 0, 1, 1, 0, 1, 2, 0, 1, 2, 0, 2, 2, 2, 1, 0, 2, 2, 2, 1, 0, 0, 2, 2, 2, 0, 2, 0, 2, 1, 1, 0, 1, 0, 0, 1, 2, 1, 1, 2, 1, 1, 0, 2, 0, 0, 1, 1, 0, 2, 1, 2, 1, 1, 1, 2, 0, 2, 1, 0, 1, 0, 1, 0, 1, 0, 1, 2, 0, 1, 0, 0, 2, 2, 1, 2, 0, 0, 2, 1, 2, 1, 2, 2, 0, 1, 1, 0, 1, 1, 2, 0, 0, 1, 2, 0, 0, 0, 2, 0, 2, 2, 0, 0, 2, 1, 2, 2, 2, 0, 0, 1, 1, 2, 2, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 2, 2, 1, 0, 1, 1, 1, 1, 2, 0, 0, 1, 0, 2, 1, 0, 2, 1, 1, 0, 1, 1, 0, 1, 2, 2, 0, 0, 1, 1, 2, 0, 2, 1, 0, 2, 2, 0, 2, 1, 0, 0, 1, 0, 0, 1, 0, 2, 0, 2, 0, 0, 1, 0, 0, 1, 2, 0, 1, 1, 1, 2, 0, 0, 0, 0, 2, 2, 2, 0, 1, 0, 1, 1, 0, 0, 2, 2, 0, 2, 0, 1, 2, 1, 1, 1, 0, 1, 0, 2, 0, 1, 2, 0, 2, 2, 0, 1, 1, 1, 2, 1, 1, 0, 1, 2, 2, 0, 1, 0, 1, 0, 1, 2, 1, 1, 0, 0, 2, 1, 2, 2, 2, 0, 1, 2, 0, 2, 1, 1, 0, 2, 1, 2, 1, 2, 0, 2, 0, 0, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 1, 2, 0, 1, 0, 0, 0, 1, 0, 2, 2, 1, 0, 1, 1, 0, 2, 2, 1, 0, 2, 2, 1, 1, 0, 0, 0, 1, 2, 2, 1, 2, 0, 2, 0, 1, 0, 2, 2, 0, 0, 2, 0, 1, 2, 2, 1, 1, 2, 2, 1, 1, 1, 2, 0, 0, 0, 0, 2, 0, 0, 1, 1, 2, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0, 2, 0, 0, 1, 1, 2, 2, 2, 0, 0, 1, 0, 0, 2, 1, 2, 1, 2, 0, 1, 2, 2, 0, 2, 1, 1, 0, 0, 0, 2, 0, 2, 0, 1, 2, 0, 1, 2, 1, 0, 1, 1, 0, 0, 0, 1, 2, 1, 0, 2, 0, 2, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 2, 1, 0, 2, 1, 1, 0, 1, 1, 2, 1, 1, 1, 0, 2, 2, 2, 1, 0, 1, 0, 1, 2, 1, 2, 1, 2, 0, 0, 2, 2, 1, 0, 2, 2, 2, 2, 2, 0, 1, 0, 1, 1, 2, 1, 1, 2, 1, 1, 1, 2, 1, 1, 2, 1, 0, 0, 2, 2, 0, 0, 1, 1, 2, 1, 1, 0, 0, 0, 2, 0, 2, 2, 1, 0, 2, 0, 0, 1, 2, 1, 2, 2, 1, 2, 0, 0, 2, 0, 0, 0, 2, 2, 1, 0, 1, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 0, 2, 0, 0, 0, 0, 1, 0, 0, 1, 2, 1, 2, 2, 2, 0, 2, 2, 0, 1, 2, 0, 1, 0, 2, 0, 2, 1, 1, 1, 0, 2, 1, 2, 1, 2, 0, 1, 0, 2, 2, 2, 0, 1, 1, 1, 2, 2, 2, 0, 1, 0, 2, 2, 2, 2, 0, 0, 0, 2, 1, 2, 2, 1, 2, 0, 1, 1, 1, 0, 0, 1, 1, 1, 2, 1, 1, 2, 2, 2, 0, 0, 0, 2, 2, 1, 1, 0, 2, 1, 2, 1, 1, 2, 1, 1, 2, 0, 1, 1, 1, 1, 2, 2, 2, 2, 2, 1, 1, 0, 0, 2, 1, 2, 0, 2, 2, 1, 1, 1, 2, 2, 0, 1, 2, 0, 0, 0, 0, 0, 2, 0, 2, 2, 1, 2, 0, 2, 1, 0, 1, 2, 2, 2, 0, 2, 2, 2, 2, 1, 2, 1, 2, 2, 0, 0, 2, 2, 0, 1, 2, 0, 2, 1, 0, 2, 1, 1, 2, 1, 2, 2, 1, 2, 2, 2, 0, 1, 1, 0, 0, 1, 0, 2, 2, 0, 1, 2, 1, 0, 2, 2, 0, 1, 0, 2, 2, 1, 1, 0, 2, 1, 1, 2, 2, 1, 2, 2, 1, 0, 1, 2, 0, 1, 1, 1, 0, 2, 0, 0, 2, 2, 0, 2, 1, 0, 2, 0, 1, 0, 0, 1, 0, 1, 0, 2, 2, 1, 1, 0, 1, 0, 1, 0, 1, 2, 2, 2, 2, 1, 2, 1, 0, 0, 2, 1, 0, 0, 0, 1, 2, 2, 0, 1, 1, 2, 2, 0, 1, 0, 0, 1, 0, 1, 1, 1, 2, 0, 2, 0, 0, 1, 1, 1, 0, 1, 1, 2, 2, 2, 1, 0, 0, 2, 2, 0, 2, 2, 1, 2, 0, 0, 2, 2, 0, 0, 2, 2, 0, 0, 1, 0, 2, 1, 0, 1, 1, 0, 2, 0, 0, 0, 0, 2, 1, 1, 2, 0, 0, 2, 1, 0, 0, 1, 2, 1, 1, 0, 0, 0, 2, 2, 2, 1, 2, 1, 0, 1, 1, 1, 0, 1, 0, 0, 2, 0, 0, 1, 0, 2, 2, 2, 2, 2, 1, 1, 0, 0, 0, 1, 1, 0, 2, 1, 1, 1, 2, 1, 0, 0, 1, 0, 2, 2, 2, 0, 1, 1, 0, 1, 1, 2, 0, 1, 2, 1, 0, 2, 1, 0, 0, 2, 1, 0, 2, 1, 2, 2, 1, 2, 0, 2, 2, 0, 0, 1, 0, 2, 2, 1, 2, 1, 2, 2, 1, 1, 0, 0, 2, 2, 1, 1, 2, 2, 1, 2, 0, 0, 0, 2, 2, 2, 0, 1, 0, 2, 1, 2, 2, 2, 2, 2, 1, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 2, 1, 2, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 2, 0, 0, 1, 1, 2, 0, 2, 1, 0, 0, 2, 0, 1, 2, 0, 0, 0, 1, 1, 0, 2, 0, 2, 0, 2, 0, 0, 1, 1, 0, 0, 0, 2, 0, 2, 1, 0, 1, 0, 1, 0, 1, 2, 1, 1, 0, 0, 1, 2, 1, 0, 2, 0, 1, 1, 1, 2, 2, 2, 2, 0, 1, 0, 0, 1, 2, 2, 2, 2, 1, 0, 0, 0, 2, 2, 1, 1, 1, 2, 0, 1, 0, 0, 1, 1, 1, 1, 1, 2, 0, 2, 1, 1, 2, 2, 1, 1, 2, 2, 2, 0, 2, 2, 2, 1, 1, 2, 1, 1, 0, 1, 2, 0, 1, 2, 0, 2, 0, 2, 2, 0, 2, 1, 0, 2, 0, 1, 1, 0, 2, 1, 1, 2, 0, 1, 1, 1, 2, 1, 2, 1, 2, 2, 1, 1, 0, 0, 2, 2, 1, 2, 0, 1, 1, 2, 1, 1, 1, 1, 0, 1, 0, 2, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 2, 1, 2, 2, 1, 1, 2, 1, 0, 1, 2, 1, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 2, 0, 2, 1, 1, 0, 2, 2, 2, 0, 2, 1, 2, 2, 1, 1, 1, 1, 1, 0, 0, 0, 1, 2, 2, 2, 0, 0, 0, 0, 2, 1, 0, 1, 0, 2, 1, 0, 1, 0, 1, 2, 0, 1, 0, 2, 1, 1, 1, 0, 2, 1, 2, 0, 1, 1, 1, 0, 0, 0, 0, 2, 0, 0, 2, 0, 2, 1, 0, 1, 0, 0, 0, 1, 1, 0, 2, 1, 2, 1, 2, 2, 0, 1, 2, 0, 2, 1, 1, 0, 1, 1, 1, 2, 0, 0, 1, 2, 1, 2, 0, 1, 2, 1, 2, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1, 1, 1, 0, 1, 1, 2, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 2, 2, 0, 2, 2, 0, 1, 1, 2, 1, 2, 0, 2, 1, 0, 2, 1, 1, 1, 2, 0, 0, 2, 1, 2, 2, 2, 2, 2, 1, 0, 2, 2, 0, 1, 1, 0, 2, 2, 1, 1, 0, 2, 1, 2, 0, 0, 2, 1, 1, 2, 0, 0, 0, 0, 2, 2, 0, 1, 1, 0, 2, 2, 1, 2, 1, 1, 0, 2, 2, 0, 1, 2, 1, 2, 1, 0, 1, 1, 0, 0, 2, 2, 2, 0, 0, 2, 2, 0, 0, 0, 0, 0, 0, 1, 1, 0, 2, 0, 0, 2, 0, 2, 0, 0, 0, 2, 1, 2, 2, 1, 0, 2, 2, 1, 2, 2, 0, 2, 0, 1, 0, 0, 1, 2, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1, 1, 2, 2, 2, 1, 0, 2, 2, 2, 0, 2, 1, 2, 1, 2, 0, 1, 1, 0, 0, 0, 2, 2, 1, 1, 1, 2, 0, 2, 1, 1, 1, 0, 0, 1, 2, 2, 0, 0, 0, 1, 1, 2, 1, 2, 0, 2, 1, 2, 2}
	answer := findKthSmallestPairDistance(arr, 25000)
	fmt.Println(answer)
}

func findKthSmallestPairDistance(nums []int, k int) int {
	sort.Ints(nums) // Sorting the array
	n := len(nums)

	// Binary search on the answer
	left, right := 0, nums[n-1]-nums[0]
	for left < right {
		mid := (left + right) / 2
		count := 0
		j := 0

		for i := 0; i < n; i++ {
			for j < n && nums[j]-nums[i] <= mid {
				j++
			}
			count += j - i - 1
		}

		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}