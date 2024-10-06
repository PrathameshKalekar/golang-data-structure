/*
1813 You are given two strings sentence1 and sentence2, each representing a sentence composed of words. A sentence is a list of words that are separated by a single space with no leading or trailing spaces. Each word consists of only uppercase and lowercase English characters.

Two sentences s1 and s2 are considered similar if it is possible to insert an arbitrary sentence (possibly empty) inside one of these sentences such that the two sentences become equal. Note that the inserted sentence must be separated from existing words by spaces.

For example,

s1 = "Hello Jane" and s2 = "Hello my name is Jane" can be made equal by inserting "my name is" between "Hello" and "Jane" in s1.
s1 = "Frog cool" and s2 = "Frogs are cool" are not similar, since although there is a sentence "s are" inserted into s1, it is not separated from "Frog" by a space.
Given two sentences sentence1 and sentence2, return true if sentence1 and sentence2 are similar. Otherwise, return false.

Example 1:

Input: sentence1 = "My name is Haley", sentence2 = "My Haley"

Output: true

Explanation:

sentence2 can be turned to sentence1 by inserting "name is" between "My" and "Haley".

Example 2:

Input: sentence1 = "of", sentence2 = "A lot of words"

Output: false

Explanation:

No single sentence can be inserted inside one of the sentences to make it equal to the other.

Example 3:

Input: sentence1 = "Eating right now", sentence2 = "Eating"

Output: true

Explanation:

sentence2 can be turned to sentence1 by inserting "right now" at the end of the sentence.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Hello Jane"
	s2 := "Hello my name is Jane"
	answer := sentenceSimilarity3(s1, s2)
	fmt.Println(answer)
}

func sentenceSimilarity3(s1, s2 string) bool {
	full := strings.Split(s1, " ")
	partial := strings.Split(s2, " ")

	// ensure full is the longer sentence
	if len(full) < len(partial) {
		full, partial = partial, full
	}

	// find matching prefix
	left := 0
	for left < len(partial) && full[left] == partial[left] {
		left++
	}

	// find matching suffix
	right := 0
	for right < len(partial)-left && full[len(full)-1-right] == partial[len(partial)-1-right] {
		right++
	}

	// check if the sum of matched prefix and suffix covers the partial sentence
	return left+right >= len(partial)
}
