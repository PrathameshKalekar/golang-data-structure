/*
7 Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21
*/
package main

import (
	"math"
)

func main() {
	reverseInteger(123)
}
func reverseInteger(x int) int {
	result := 0

	// Loop through each digit in x
	for x != 0 {
		digit := x % 10 // Extract the last digit of x
		x /= 10         // Remove the last digit from x

		// Check for overflow/underflow before updating result
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			return 0 // Overflow case
		}
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && digit < -8) {
			return 0 // Underflow case
		}

		result = result*10 + digit // Append the digit to the reversed number
	}

	return result
}
