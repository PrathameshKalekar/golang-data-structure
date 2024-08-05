/*
28 Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.
Example 2:

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
*/
package main

import "fmt"

func main() {
	s1 := "sabutsa"
	s2 := "sad"
	fmt.Println(findTheIndexOfTheFirstOccurenceInAString(s1, s2))
}
func findTheIndexOfTheFirstOccurenceInAString(s1, s2 string) int {
	// Iterate over each character in s1 using range.
	for s1Index, char := range s1 {
		// Check if the current character matches the first character of s2.
		if char == rune(s2[0]) {
			// Ensure the remaining length of s1 from s1Index is at least as long as s2.
			if len(s1) >= s1Index+len(s2) && s1[s1Index:s1Index+len(s2)] == s2 {
				// If the substring from s1Index matches s2, return the current index.
				return s1Index
			}
		}
	}
	// If no match is found, return -1.
	return -1
}
