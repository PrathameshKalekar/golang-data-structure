/*
1508 You are given the array nums consisting of n positive integers. You computed the sum of all non-empty continuous subarrays from the array and then sorted them in non-decreasing order, creating a new array of n * (n + 1) / 2 numbers.

Return the sum of the numbers from index left to index right (indexed from 1), inclusive, in the new array. Since the answer can be a huge number return it modulo 109 + 7.



Example 1:

Input: nums = [1,2,3,4], n = 4, left = 1, right = 5
Output: 13
Explanation: All subarray sums are 1, 3, 6, 10, 2, 5, 9, 3, 7, 4. After sorting them in non-decreasing order we have the new array [1, 2, 3, 3, 4, 5, 6, 7, 9, 10]. The sum of the numbers from index le = 1 to ri = 5 is 1 + 2 + 3 + 3 + 4 = 13.
Example 2:

Input: nums = [1,2,3,4], n = 4, left = 3, right = 4
Output: 6
Explanation: The given array is the same as example 1. We have the new array [1, 2, 3, 3, 4, 5, 6, 7, 9, 10]. The sum of the numbers from index le = 3 to ri = 4 is 3 + 3 = 6.
Example 3:

Input: nums = [1,2,3,4], n = 4, left = 1, right = 10
Output: 50
*/package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4}
	answer := rangeOfSortedSubarraySums(arr, len(arr), 1, 10)
	fmt.Println(answer)
}

const MOD = 1_000_000_007

func rangeOfSortedSubarraySums(arr []int, n, left, right int) int {
	sums := []int{} //to keep track of all sums
	total := 0
	//to add the sums of the array in sums
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			sums = append(sums, sum)
		}
	}
	// sort the new array of sums
	sort.Ints(sums)
	//loop from left ot right to add the total and maek it non overflow by % MOD
	for i := left - 1; i < right; i++ {
		total = (total + sums[i]) % MOD
	}
	//return the total
	return total
}
