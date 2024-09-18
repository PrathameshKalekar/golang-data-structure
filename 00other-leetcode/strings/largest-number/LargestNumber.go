/*
Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.

Since the result may be very large, so you need to return a string instead of an integer.

Example 1:

Input: nums = [10,2]
Output: "210"
Example 2:

Input: nums = [3,30,34,5,9]
Output: "9534330"
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr := []int{3, 30, 34, 5, 9}
	answer := largestNumber(arr)
	fmt.Println(answer)
}

func largestNumber(arr []int) string {
	// Convert each integer in the array to a string
	nums := make([]string, len(arr)) // Create a slice of strings to store the number strings
	for i, value := range arr {      // Loop through the input array
		nums[i] = strconv.Itoa(value) // Convert each integer to string and store in 'nums'
	}

	// Implement a custom sorting algorithm (bubble sort in this case)
	// Sort based on the concatenated order of strings
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			// Combine two strings in both possible order
			num1, _ := strconv.Atoi(nums[j] + nums[j+1]) // Concatenate as one possible number
			num2, _ := strconv.Atoi(nums[j+1] + nums[j]) // Concatenate as the other possible number

			// Compare the two possible numbers
			// If num1 is less than num2, swap nums[j] and nums[j+1]
			// This ensures that the larger number (when concatenated) comes first
			if num1 < num2 {
				nums[j], nums[j+1] = nums[j+1], nums[j] // Swap the two numbers
			}
		}
	}

	// If the largest number is "0", the whole number should be "0"
	// This prevents cases like [0, 0] returning "00"
	if nums[0] == "0" {
		return "0"
	}

	// Join the sorted strings into one large concatenated string and return it
	return strings.Join(nums, "")
}
