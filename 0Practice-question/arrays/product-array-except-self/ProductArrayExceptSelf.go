/*
238  Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]/*
*/
package main

import "fmt"

func main() {
	arr := []int{-1, 1, 0, -3, 3}
	answer := productArrayExceptSelf(arr)
	fmt.Println(answer)
}

func productArrayExceptSelf(arr []int) []int {
	n := len(arr)
	//making array with same length as original array
	result := make([]int, n)

	//filling result array with 1
	for i := 0; i < n; i++ {
		result[i] = 1
	}

	//Making left to right pointer variavle that track all value
	leftToRightValue := 1
	//looping all element from left to right in the arr for multiplication
	for i := 0; i < n; i++ {
		result[i] = leftToRightValue
		leftToRightValue *= arr[i]
	}

	rightToLeftValue := 1
	//looping all element from right to left in the arr for multiplication
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightToLeftValue
		rightToLeftValue *= arr[i]
	}
	return result
}
