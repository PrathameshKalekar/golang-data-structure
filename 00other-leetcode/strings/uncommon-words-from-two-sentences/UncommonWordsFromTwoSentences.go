/*
884 A sentence is a string of single-space separated words where each word consists only of lowercase letters.

A word is uncommon if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Given two sentences s1 and s2, return a list of all the uncommon words. You may return the answer in any order.

Example 1:

Input: s1 = "this apple is sweet", s2 = "this apple is sour"

Output: ["sweet","sour"]

Explanation:

The word "sweet" appears only in s1, while the word "sour" appears only in s2.

Example 2:

Input: s1 = "apple apple", s2 = "banana"

Output: ["banana"]
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "this apple is sweet"
	s2 := "this apple is sour"
	answer := uncommonWordsFromTwoSentences(s1, s2)
	fmt.Println(answer)
}

func uncommonWordsFromTwoSentences(s1, s2 string) []string {
	//combine both string that will make the comparison easier
	str := s1 + " " + s2
	//make a map to keep track of reapeating words
	strMap := map[string]int{}
	// split the word with the space
	words := strings.Split(str, " ")
	// map each word with thier count
	for _, value := range words {
		strMap[value]++
	}
	var result []string
	for char, value := range strMap {
		//count is one means the word id came once so add it in the result
		if value == 1 {
			result = append(result, char)
		}
	}
	//return the result
	return result
}
