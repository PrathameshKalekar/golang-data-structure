/*
202 Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.

Example 1:

Input: n = 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
Example 2:

Input: n = 2
Output: false
*/
package main

import (
	"fmt"
)

// Function to calculate the sum of the squares of digits of a number
func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

// Function to determine if a number is a happy number
func isHappy(n int) bool {
	seen := make(map[int]bool)
	for n != 1 && !seen[n] {
		seen[n] = true
		n = sumOfSquares(n)
	}
	return n == 1
}

func main() {
	n := 19
	fmt.Println("Input:", n)
	fmt.Println("Output:", isHappy(n)) // Output: true

	n = 2
	fmt.Println("Input:", n)
	fmt.Println("Output:", isHappy(n)) // Output: false
}
