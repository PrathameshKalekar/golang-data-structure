/*
290 Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

Example 1:

Input: pattern = "abba", s = "dog cat cat dog"
Output: true
Example 2:

Input: pattern = "abba", s = "dog cat cat fish"
Output: false
Example 3:

Input: pattern = "aaaa", s = "dog cat cat dog"
Output: false
*/
package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	// Maps to store the mapping from pattern to words and words to pattern
	patternToWord := make(map[rune]string)
	wordToPattern := make(map[string]rune)

	for i, ch := range pattern {
		word := words[i]

		// Check if there is a conflict in the mapping pattern -> word
		if mappedWord, exists := patternToWord[ch]; exists {
			if mappedWord != word {
				return false
			}
		} else {
			patternToWord[ch] = word
		}

		// Check if there is a conflict in the mapping word -> pattern
		if mappedPattern, exists := wordToPattern[word]; exists {
			if mappedPattern != ch {
				return false
			}
		} else {
			wordToPattern[word] = ch
		}
	}

	return true
}

func main() {
	pattern1 := "abba"
	s1 := "dog cat cat dog"
	fmt.Println("Input: pattern =", pattern1, ", s =", s1)
	fmt.Println("Output:", wordPattern(pattern1, s1)) // Output: true

	pattern2 := "abba"
	s2 := "dog cat cat fish"
	fmt.Println("Input: pattern =", pattern2, ", s =", s2)
	fmt.Println("Output:", wordPattern(pattern2, s2)) // Output: false

	pattern3 := "aaaa"
	s3 := "dog cat cat dog"
	fmt.Println("Input: pattern =", pattern3, ", s =", s3)
	fmt.Println("Output:", wordPattern(pattern3, s3)) // Output: false
}
