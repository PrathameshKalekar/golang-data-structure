/*
9 Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/
package main

import "fmt"

func main() {
	answer := palindrome(121)
	fmt.Printf("Is provided number is an palindrome : %t", answer)
}

func palindrome(num int) bool {
	// If the number is negative, it cannot be a palindrome
	if num < 0 {
		return false
	}

	// Store the original number in temp
	temp := num
	// This will hold the reversed number
	result := 0

	// Reverse the number
	for temp > 0 {
		// Get the last digit of temp
		digit := temp % 10
		// Remove the last digit from temp
		temp = temp / 10
		// Append the digit to the result
		result = result*10 + digit
	}

	// Check if the reversed number is equal to the original number
	return result == num
}
