/*
564 Given a string n representing an integer, return the closest integer (not including itself), which is a palindrome. If there is a tie, return the smaller one.

The closest is defined as the absolute difference minimized between two integers.

Example 1:

Input: n = "123"
Output: "121"
Example 2:

Input: n = "1"
Output: "0"
Explanation: 0 and 2 are the closest palindromes but we return the smallest which is 0.
*/
package main

import (
	"fmt"
	"math"
	"strconv"
)

func closestPalindrome(n string) string {
	length := len(n)
	num, _ := strconv.Atoi(n)
	candidates := make(map[int]struct{})

	// Edge cases: numbers with all 9s or single digits
	candidates[int(math.Pow10(length))+1] = struct{}{}   // 100...001
	candidates[int(math.Pow10(length-1))-1] = struct{}{} // 999...999 or 0

	// Use the first half to generate the palindrome candidates
	prefix, _ := strconv.Atoi(n[:(length+1)/2])
	for i := -1; i <= 1; i++ {
		newPrefix := prefix + i
		palindrome := generatePalindrome(newPrefix, length%2 == 0)
		candidates[palindrome] = struct{}{}
	}

	// Remove the original number itself from candidates
	delete(candidates, num)

	// Find the closest palindrome
	closest := -1
	for candidate := range candidates {
		if closest == -1 || abs(candidate-num) < abs(closest-num) || (abs(candidate-num) == abs(closest-num) && candidate < closest) {
			closest = candidate
		}
	}
	return strconv.Itoa(closest)
}

// Helper function to generate palindrome from the prefix
func generatePalindrome(prefix int, evenLength bool) int {
	prefixStr := strconv.Itoa(prefix)
	rev := reverseString(prefixStr)
	if evenLength {
		return toInt(prefixStr + rev)
	}
	return toInt(prefixStr + rev[1:])
}

// Helper function to reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Helper function to convert string to integer
func toInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

// Helper function to calculate absolute value
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(closestPalindrome("123")) // Output: "121"
	fmt.Println(closestPalindrome("1"))   // Output: "0"
}
