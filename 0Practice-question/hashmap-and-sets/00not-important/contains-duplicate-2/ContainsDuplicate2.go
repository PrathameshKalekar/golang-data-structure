/*
219 Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
*/
package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 1}
	k1 := 3
	fmt.Println("Input:", nums1, "k =", k1)
	fmt.Println("Output:", containsDuplicate2(nums1, k1)) // Output: true

	nums2 := []int{1, 0, 1, 1}
	k2 := 1
	fmt.Println("Input:", nums2, "k =", k2)
	fmt.Println("Output:", containsDuplicate2(nums2, k2)) // Output: true

	nums3 := []int{1, 2, 3, 1, 2, 3}
	k3 := 2
	fmt.Println("Input:", nums3, "k =", k3)
	fmt.Println("Output:", containsDuplicate2(nums3, k3)) // Output: false
}

func containsDuplicate2(arr []int, k int) bool {
	arrMap := map[int]int{}
	for i, num := range arr {
		if j, found := arrMap[num]; found && i-j <= k {
			return true
		}
		arrMap[num] = i
	}
	return false
}
