/*
1371 Given the string s, return the size of the longest substring containing each vowel an even number of times. That is, 'a', 'e', 'i', 'o', and 'u' must appear an even number of times.

Example 1:

Input: s = "eleetminicoworoep"
Output: 13
Explanation: The longest substring is "leetminicowor" which contains two each of the vowels: e, i and o and zero of the vowels: a and u.
Example 2:

Input: s = "leetcodeisgreat"
Output: 5
Explanation: The longest substring is "leetc" which contains two e's.
Example 3:

Input: s = "bcbcbc"
Output: 6
Explanation: In this case, the given string "bcbcbc" is the longest because all vowels: a, e, i, o and u appear zero times.
*/
package main

import "fmt"

// Function to find the longest substring with vowels in even counts
func findTheLongestSubstring(s string) int {
	// Initialize the mask, where all vowels are even (00000 -> mask = 0)
	mask := 0
	// Map to store the first occurrence of each mask
	pos := map[int]int{0: -1} // The mask 0 is found at index -1 (before the start)
	maxLen := 0

	// Iterate over the string
	for i := 0; i < len(s); i++ {
		// Update the mask based on the current character
		switch s[i] {
		case 'a':
			mask ^= 1 << 0
		case 'e':
			mask ^= 1 << 1
		case 'i':
			mask ^= 1 << 2
		case 'o':
			mask ^= 1 << 3
		case 'u':
			mask ^= 1 << 4
		}

		// Check if we have seen this mask before
		if firstPos, found := pos[mask]; found {
			// Update maxLen if the current substring is longer
			maxLen = max(maxLen, i-firstPos)
		} else {
			// Store the first occurrence of this mask
			pos[mask] = i
		}
	}

	return maxLen
}

// Utility function to get the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test cases
	fmt.Println(findTheLongestSubstring("eleetminicoworoep")) // Output: 13
	fmt.Println(findTheLongestSubstring("leetcodeisgreat"))   // Output: 5
	fmt.Println(findTheLongestSubstring("bcbcbc"))            // Output: 6
}
