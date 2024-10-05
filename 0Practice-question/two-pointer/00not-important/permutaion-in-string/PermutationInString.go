/*
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.

Example 1:

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input: s1 = "ab", s2 = "eidboaoo"
Output: false
*/
package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	answer := permutationInString(s1, s2)
	fmt.Println(answer)
}

func permutationInString(s1, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	var pattern, arr [26]int
	for i := 0; i < len(s1); i++ {
		pattern[int(s1[i]-'a')]++
		arr[int(s2[i]-'a')]++
	}

	i := 0
	for {
		if arr == pattern {
			return true
		}
		if i+len(s1) >= len(s2) {
			break
		}
		arr[int(s2[i]-'a')]--
		arr[int(s2[i+len(s1)]-'a')]++
		i++
	}
	return false
}
