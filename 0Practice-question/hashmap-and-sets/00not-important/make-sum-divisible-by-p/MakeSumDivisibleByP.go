/*
1590 Given an array of positive integers nums, remove the smallest subarray (possibly empty) such that the sum of the remaining elements is divisible by p. It is not allowed to remove the whole array.

Return the length of the smallest subarray that you need to remove, or -1 if it's impossible.

A subarray is defined as a contiguous block of elements in the array.

Example 1:

Input: nums = [3,1,4,2], p = 6
Output: 1
Explanation: The sum of the elements in nums is 10, which is not divisible by 6. We can remove the subarray [4], and the sum of the remaining elements is 6, which is divisible by 6.
Example 2:

Input: nums = [6,3,5,2], p = 9
Output: 2
Explanation: We cannot remove a single element to get a sum divisible by 9. The best way is to remove the subarray [5,2], leaving us with [6,3] with sum 9.
Example 3:

Input: nums = [1,2,3], p = 3
Output: 0
Explanation: Here the sum is 6. which is already divisible by 3. Thus we do not need to remove anything.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{3, 1, 4, 2}
	p1 := 6
	fmt.Println(minSubarray(nums1, p1)) // Output: 1

	nums2 := []int{6, 3, 5, 2}
	p2 := 9
	fmt.Println(minSubarray(nums2, p2)) // Output: 2

	nums3 := []int{1, 2, 3}
	p3 := 3
	fmt.Println(minSubarray(nums3, p3)) // Output: 0
}

func minSubarray(nums []int, p int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	d := sum % p
	if d%p == 0 {
		return 0
	}

	dict := make(map[int]int)
	dict[0] = -1

	res := math.MaxInt64

	sum = 0
	for i, v := range nums {
		sum += v
		a := sum % p
		tar := (a + p - d) % p
		if j, ok := dict[tar]; ok {
			res = min(res, i-j)
		}
		dict[sum%p] = i
	}
	if res == math.MaxInt64 || res == len(nums) {
		return -1
	}
	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
