/*
30 You are given a string s and an array of strings words. All the strings of words are of the same length.

A concatenated string is a string that exactly contains all the strings of any permutation of words concatenated.

For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated string because it is not the concatenation of any permutation of words.
Return an array of the starting indices of all the concatenated substrings in s. You can return the answer in any order.

Example 1:

Input: s = "barfoothefoobarman", words = ["foo","bar"]

Output: [0,9]

Explanation:

The substring starting at 0 is "barfoo". It is the concatenation of ["bar","foo"] which is a permutation of words.
The substring starting at 9 is "foobar". It is the concatenation of ["foo","bar"] which is a permutation of words.

Example 2:

Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]

Output: []

Explanation:

There is no concatenated substring.

Example 3:

Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]

Output: [6,9,12]

Explanation:

The substring starting at 6 is "foobarthe". It is the concatenation of ["foo","bar","the"].
The substring starting at 9 is "barthefoo". It is the concatenation of ["bar","the","foo"].
The substring starting at 12 is "thefoobar". It is the concatenation of ["the","foo","bar"].
*/
package main

import "fmt"

func main() {
	s := "barfoofoobarthefoobarman"
	words := []string{"bar", "foo", "the"}
	answer := substringWithConcatenationOfAllWords(s, words)
	fmt.Println(answer)
}

func substringWithConcatenationOfAllWords(s string, words []string) []int {
	wordLen := len(words[0])         // Length of each word (all words are of the same length)
	wordCount := len(words)          // Number of words in the list
	concatLen := wordLen * wordCount // Total length of concatenated string
	result := []int{}                // Stores the starting indices of the valid substrings

	// Edge case: If s is shorter than the total concatenation length, return empty
	if len(s) < concatLen {
		return result
	}

	// Create a frequency map for the words
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	// Check every possible starting point in s
	for i := 0; i <= len(s)-concatLen; i++ {
		seen := make(map[string]int)
		j := 0
		for j < wordCount {
			wordStart := i + j*wordLen
			word := s[wordStart : wordStart+wordLen]

			// If the word is not part of the words array, break
			if _, exists := wordFreq[word]; !exists {
				break
			}

			// Count the word in the current window
			seen[word]++
			if seen[word] > wordFreq[word] {
				break
			}

			j++
		}

		// If all words are used exactly once, add the starting index
		if j == wordCount {
			result = append(result, i)
		}
	}

	return result
}
