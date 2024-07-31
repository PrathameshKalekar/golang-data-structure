/*
383 Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.



Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true
*/

package main

import "fmt"

func main() {
	ransomNot := "aab"
	magazine := "baa"
	fmt.Println(ransomNote(ransomNot, magazine))

	ransomNot = "aab"
	magazine = "ab"
	fmt.Println(ransomNote(ransomNot, magazine))
}

func ransomNote(ransomNote, magazine string) bool {
	ransomMap, magazineMap := make(map[rune]int), make(map[rune]int)
	for _, value := range ransomNote {
		ransomMap[value]++
	}
	for _, value := range magazine {
		magazineMap[value]++
	}
	for char, count := range ransomMap {
		if magazineMap[char] < count {
			return false
		}
	}
	return true
}
