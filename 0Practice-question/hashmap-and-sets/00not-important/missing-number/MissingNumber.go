/*
268 Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

Example 1:

Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
Example 2:

Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.
Example 3:

Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.
*/
package main

import "fmt"

func main() {
	arr := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	answer := missingNumber(arr)
	fmt.Println(answer)
}
func missingNumber(nums []int) int {
	// Calculate the length of the input array
	n := len(nums)

	// Calculate the expected sum of the first n natural numbers
	expectedSum := n * (n + 1) / 2

	// Initialize a variable to hold the actual sum of the array elements
	actualSum := 0

	// Calculate the actual sum of the array elements
	for _, num := range nums {
		actualSum += num
	}

	// The missing number is the difference between the expected sum and the actual sum
	return expectedSum - actualSum
}
