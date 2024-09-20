/*
You are given a string s. You can convert s to a palindrome by adding characters in front of it.

Return the shortest palindrome you can find by performing this transformation.

Example 1:

Input: s = "aacecaaa"
Output: "aaacecaaa"
Example 2:

Input: s = "abcd"
Output: "dcbabcd"
*/
package main

import (
	"fmt"
)

// Compute LPS (Longest Prefix which is also Suffix) array for KMP algorithm
func computeLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0
	i := 1

	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	// Create a new string by concatenating s and its reverse with a special character in between
	reverseS := reverse(s)
	newString := s + "#" + reverseS

	// Find the LPS array for the concatenated string
	lps := computeLPS(newString)

	// Characters to add in front are the non-palindromic suffix of the original string
	toAdd := reverseS[:len(s)-lps[len(lps)-1]]

	return toAdd + s
}

// Helper function to reverse a string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	s1 := "aacecaaa"
	s2 := "abcd"

	fmt.Println("Shortest palindrome for", s1, ":", shortestPalindrome(s1)) // Output: "aaacecaaa"
	fmt.Println("Shortest palindrome for", s2, ":", shortestPalindrome(s2)) // Output: "dcbabcd"
}
