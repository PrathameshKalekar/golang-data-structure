/*
5 ven a string s, return the longest palindromic substring in s.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
*/
package main

import "fmt"

func main() {
	s := "cbbds"
	answer := longestPalindromeSubstring(s)
	fmt.Println(answer)
}
func longestPalindromeSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0 // to track the start and end of the longest palindrome
	for i := 0; i < len(s); i++ {
		// Check for odd-length palindromes centered at i
		l1, r1 := expandAroundCenter(s, i, i)
		// Check for even-length palindromes centered between i and i+1
		l2, r2 := expandAroundCenter(s, i, i+1)

		// Update the longest palindrome substring
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}

// expandAroundCenter expands the palindrome around the given center and returns the start and end indices of the longest palindrome
func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
