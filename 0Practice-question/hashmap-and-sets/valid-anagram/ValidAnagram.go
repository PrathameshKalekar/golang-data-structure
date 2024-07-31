/*
242 Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
*/

package main

import "fmt"

func main() {
	string1 := "anagram"
	string2 := "nagarama"
	answer := validAnagram(string1, string2)
	fmt.Printf("Is %s and %s is valid anagram : %t", string1, string2, answer)
}

func validAnagram(string1, string2 string) bool {
	string1Map, string2Map := make(map[rune]int), make(map[rune]int)
	//mapping string 1
	for _, value := range string1 {
		string1Map[value]++
	}
	//mapping string 2
	for _, value := range string2 {
		string2Map[value]++
	}
	//if len of both map is same means it may contains the anagram
	if len(string1Map) == len(string2Map) {
		for char, count := range string1Map {
			// if each char of string 1 and 2 maps contains same count then its a anagram else it not
			if string2Map[char] != count {
				// if not the return false
				return false
			}
		}
		//is all condition staisfies return true
		return true
	}
	//else return false
	return false
}
